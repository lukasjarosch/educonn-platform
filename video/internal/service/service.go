package service

import (
	"context"
	"fmt"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/amazon"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/broker"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/errors"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/mongodb"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	merr "github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"github.com/rs/zerolog/log"
	"gopkg.in/mgo.v2/bson"
	"github.com/opentracing/opentracing-go"
)

type videoService struct {
	videoPublisher         videoPublisher
	s3Bucket               *amazon.S3Bucket
	transcodeCompletedChan chan broker.TranscodingCompletedEvent
	transcodeFailedChan    chan broker.TranscodingFailedEvent
	videoRepository        *mongodb.VideoRepository
	jwtHandler             *jwt_handler.JwtTokenHandler
}

type videoPublisher interface {
	PublishVideoCreated(ctx context.Context, event *pbVideo.VideoCreatedEvent) (err error)
	PublishVideoProcessed(ctx context.Context, event *pbVideo.VideoProcessedEvent) (err error)
}

func NewVideoService(vidCreatedPub videoPublisher,
	bucket *amazon.S3Bucket,
	transcodeCompletedChan chan broker.TranscodingCompletedEvent,
	transcodeFailedChan chan broker.TranscodingFailedEvent,
	videoRepo *mongodb.VideoRepository,
	jwtHandler *jwt_handler.JwtTokenHandler) pbVideo.VideoHandler {
	svc := &videoService{
		videoPublisher:         vidCreatedPub,
		s3Bucket:               bucket,
		transcodeCompletedChan: transcodeCompletedChan,
		transcodeFailedChan:    transcodeFailedChan,
		videoRepository:        videoRepo,
		jwtHandler:             jwtHandler,
	}

	go svc.awaitTranscodeCompletedEvent()
	go svc.awaitTranscodeFailedEvent()

	return svc
}

// Create a new Video and publish a VideoCreatedEvent
func (v *videoService) Create(ctx context.Context, req *pbVideo.CreateVideoRequest, res *pbVideo.CreateVideoResponse) error {

	// Fetch token
	md, _ := metadata.FromContext(ctx)
	token, err := v.jwtHandler.GetBearerToken(md)
	if err != nil {
		return merr.Unauthorized(config.ServiceName, "%s", err.Error())
	}
	// Decode jwt and extract user
	claims, err := v.jwtHandler.Decode(token)
	if err != nil {
		log.Debug().Interface("error", err).Msg("unable to decode jwt token")
		return merr.Unauthorized(config.ServiceName, "%s", err.Error())
	}
	user := claims.User

	// TODO: Authorization: create:video

	// Check if raw file does already exist
	existingVideo, _ := v.videoRepository.FindByRawStorageKey(ctx, req.Video.Storage.RawKey)
	if existingVideo != nil {
		log.Warn().Str("file_key", existingVideo.Storage.RawKey).Msg("a source video with that key already exists")
		return errors.RawVideoAlreadyExists
	}

	// Check if file actually exists in S3
	fileKey := req.Video.Storage.RawKey
	err = v.s3Bucket.CheckFileExists(ctx, fileKey, config.AwsS3VideoBucket)
	if err != nil {
		log.Warn().Str("key", fileKey).Msg("key does not exist in bucket")
		res.Errors = append(res.Errors, &pbVideo.Error{
			Description: errors.RawVideoFileS3NotFound.Error(),
			Code:        404,
		})
		return merr.NotFound(config.ServiceName, "%s", errors.RawVideoFileS3NotFound)
	}
	log.Info().Str("key", fileKey).Str("bucket", config.AwsS3VideoBucket).Msg("key found in bucket")

	// Insert video into DB
	video, err := v.videoRepository.CreateVideo(ctx, mongodb.UnmarshalProtobuf(req.Video, user.Id))
	if err != nil {
		log.Warn().Interface("error", err).Msg("unable to create video in database")
		return err
	}
	log.Info().Str("video_id", video.ID.Hex()).Msg("created video")
	req.Video.Id = video.ID.Hex()

	// Publish event
	v.videoPublisher.PublishVideoCreated(ctx, &pbVideo.VideoCreatedEvent{
		Video:  req.Video,
		UserId: user.Id,
	})

	res.Video = req.Video

	return nil
}

// GetById fetches a video by it's ID. This method does increment the view counter of the video.
func (v *videoService) GetById(ctx context.Context, req *pbVideo.GetVideoRequest, res *pbVideo.GetVideoResponse) error {
	if req.Id == "" {
		return errors.MissingVideoId
	}

	video, err := v.videoRepository.FindById(ctx, req.Id)
	if err != nil {
		log.Info().Str("video", req.Id).Msg("unable to find video")
		return errors.VideoNotFound
	}

	errorCount := len(video.Transcode.Errors)
	trError := false

	if errorCount >= 1 {
		trError = true
	} else {
		trError = false
	}
	res.SignedUrl, err = v.s3Bucket.GetSignedResourceURL(video.Storage.OutputKey)
	if err != nil {
		log.Warn().Err(err).Str("video", req.Id).Msg("unable to fetch signed url for video")
	}

	// count the view
	err = v.videoRepository.IncrementViews(ctx, video.ID)
	if err != nil {
		log.Warn().Err(err).Str("video", video.ID.Hex()).Msg("unable to increment view counter")
	}
	// TODO: publish video.events.view

	res.Video = &pbVideo.VideoDetails{
		Id:          video.ID.Hex(),
		Title:       video.Title,
		Description: video.Description,
		Tags:        video.Tags,
		Storage: &pbVideo.VideoStorage{
			RawKey:        video.Storage.RawKey,
			TranscodedKey: video.Storage.OutputKey,
		},
		Statistics: &pbVideo.VideoStatistics{
			ViewCount: video.Statistics.ViewCount,
		},
		Status: &pbVideo.VideoStatus{
			Completed:     video.Transcode.Completed,
			Error:         trError,
			ErrorMessages: video.Transcode.Errors,
		},
	}

	return nil
}

// GetByUserId fetches all videos uploaded a given user
func (v *videoService) GetByUserId(ctx context.Context, req *pbVideo.GetByUserIdRequest, res *pbVideo.GetByUserIdResponse) error {

	if req.UserId == "" {
		return errors.MissingUserId
	}

	videos, err := v.videoRepository.FindByUserId(ctx, req.UserId)
	if err != nil {
		return err
	}

	for _, video := range videos {
		res.Videos = append(res.Videos, &pbVideo.VideoDetails{
			Id:          video.ID.Hex(),
			Title:       video.Title,
			Description: video.Description,
			Tags:        video.Tags,
			Statistics: &pbVideo.VideoStatistics{
				ViewCount: video.Statistics.ViewCount,
			},
		})
	}

	log.Debug().Str("user", req.UserId).Msgf("fetched %d videos of user", len(videos))

	return nil
}

func (v *videoService) Update(ctx context.Context, req *pbVideo.UpdateVideoRequest, res *pbVideo.UpdateVideoResponse) error {

	if req.Video.Id == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.MissingVideoId)
	}
	// TODO: make sure userId is valid and user is authorized
	if req.Video.UserId == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.MissingUserId)
	}
	if !bson.IsObjectIdHex(req.Video.Id) || !bson.IsObjectIdHex(req.Video.UserId) {
		return merr.BadRequest(config.ServiceName, "%s", errors.MalformedId)
	}

	video := &mongodb.Video{
		ID:          bson.ObjectIdHex(req.Video.Id),
		Title:       req.Video.Title,
		Description: req.Video.Description,
		UserID:      bson.ObjectIdHex(req.Video.UserId),
		Transcode: mongodb.Transcode{
			Completed: true,
		},
		Tags: req.Video.Tags,
		Storage: mongodb.Storage{
			OutputKey: req.Video.Storage.TranscodedKey,
			RawKey:    req.Video.Storage.RawKey,
		},
	}

	err := v.videoRepository.UpdateVideo(ctx, video)
	if err != nil {
		log.Warn().Err(err).Msg("unable to update video")
		return merr.InternalServerError(config.ServiceName, "%s", "unable to update video")
	}
	return nil
}

func (v *videoService) awaitTranscodeCompletedEvent() {
	for transcoded := range v.transcodeCompletedChan {
		log.Info().Str("job", transcoded.Event.Transcode.JobId).Str("video", transcoded.Event.VideoId).Msg("transcoding job completed")

		span, ctx := opentracing.StartSpanFromContext(transcoded.Context, "VideoService.awaitTranscodeCompletedEvent")
		defer span.Finish()

		// TODO: change to FindById
		video, err := v.videoRepository.FindByRawStorageKey(transcoded.Context, transcoded.Event.Transcode.InputKey)
		if err != nil {
			log.Warn().
				Str("topic", broker.TranscodeCompletedTopic).
				Str("job", transcoded.Event.Transcode.JobId).
				Str("input_key", transcoded.Event.Transcode.InputKey).
				Interface("error", err).
				Msg("received job from non-existing video-input key")
			continue
		}

		// update video
		video.Storage.OutputKey = fmt.Sprintf("%s%s", transcoded.Event.Transcode.OutputKeyPrefix, transcoded.Event.Transcode.OutputKey)
		video.Transcode.Completed = true

		// publish: video.events.transcoded
		videoDetails := &pbVideo.VideoDetails{
			Id:          video.ID.Hex(),
			Title:       video.Title,
			Description: video.Description,
			UserId:      video.UserID.Hex(),
			Tags:        video.Tags,
			Storage: &pbVideo.VideoStorage{
				RawKey:        video.Storage.RawKey,
				TranscodedKey: video.Storage.OutputKey,
			},
			Statistics: &pbVideo.VideoStatistics{
				ViewCount: video.Statistics.ViewCount,
			},
			Status: &pbVideo.VideoStatus{
				Completed: video.Transcode.Completed,
			},
		}

		res := &pbVideo.UpdateVideoResponse{}
		err = v.Update(ctx, &pbVideo.UpdateVideoRequest{Video: videoDetails}, res)
		if err != nil {
			log.Error().Err(err).Msg("call to Video.Update() failed")
		}

		err = v.videoPublisher.PublishVideoProcessed(ctx, &pbVideo.VideoProcessedEvent{Video:res.Video})
		if err != nil {
			log.Warn().Err(err).Msgf("unable to publish '%s'", broker.VideoProcessedTopic)
		}
	}
}

func (v *videoService) awaitTranscodeFailedEvent() {
	for transcoded := range v.transcodeFailedChan {
		log.Info().Str("job", transcoded.Event.Transcode.JobId).Str("video", transcoded.Event.VideoId).Msg("transcoding job failed")

		// TODO: change to FindById
		video, err := v.videoRepository.FindByRawStorageKey(transcoded.Context, transcoded.Event.Transcode.InputKey)
		if err != nil {
			log.Warn().
				Str("topic", broker.TranscodeFailedTopic).
				Str("job", transcoded.Event.Transcode.JobId).
				Str("input_key", transcoded.Event.Transcode.InputKey).
				Interface("error", err).
				Msg("received job from non-existing video-input key")
			continue
		}

		// append errors
		for _, err := range transcoded.Event.Transcode.Status.ErrorMessages {
			video.Transcode.Errors = append(video.Transcode.Errors, err)
		}

		// update video
		err = v.videoRepository.UpdateVideo(transcoded.Context, video)
		if err != nil {
			log.Warn().Interface("error", err).Msg("unable to update video")
			continue
		}
		log.Info().Str("video", video.ID.Hex()).Msg("video transcoding failed, updated successfully")
	}
}
