package service

import (
	"context"
	"fmt"
	pbTranscode "github.com/lukasjarosch/educonn-platform/transcode/proto"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/amazon"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/broker"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/errors"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/mongodb"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	"github.com/rs/zerolog/log"
	"github.com/micro/go-micro/metadata"
	merr "github.com/micro/go-micro/errors"
)

type videoService struct {
	videoCreatedPublisher  videoCreatedPublisher
	s3Bucket               *amazon.S3Bucket
	transcodeCompletedChan chan *pbTranscode.TranscodingCompletedEvent
	transcodeFailedChan    chan *pbTranscode.TranscodingFailedEvent
	videoRepository        *mongodb.VideoRepository
	jwtHandler             *jwt_handler.JwtTokenHandler
}

type videoCreatedPublisher interface {
	PublishVideoCreated(event *pbVideo.VideoCreatedEvent) (err error)
}

func NewVideoService(vidCreatedPub videoCreatedPublisher,
	bucket *amazon.S3Bucket,
	transcodeCompletedChan chan *pbTranscode.TranscodingCompletedEvent,
	transcodeFailedChan chan *pbTranscode.TranscodingFailedEvent,
	videoRepo *mongodb.VideoRepository,
	jwtHandler *jwt_handler.JwtTokenHandler) pbVideo.VideoHandler {
	svc := &videoService{
		videoCreatedPublisher:  vidCreatedPub,
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
	existingVideo, _ := v.videoRepository.FindByRawStorageKey(req.Video.Storage.RawKey)
	if existingVideo != nil {
		log.Warn().Str("file_key", existingVideo.Storage.RawKey).Msg("a source video with that key already exists")
		return errors.RawVideoAlreadyExists
	}

	// Check if file actually exists in S3
	fileKey := req.Video.Storage.RawKey
	err = v.s3Bucket.CheckFileExists(fileKey, config.AwsS3VideoBucket)
	if err != nil {
		log.Warn().Str("key", fileKey).Msg("key does not exist in bucket")
		res.Errors = append(res.Errors, &pbVideo.Error{
			Description: errors.RawVideoFileS3NotFound.Error(),
			Code:        404,
		})
		return err
	}
	log.Info().Str("key", fileKey).Str("bucket", config.AwsS3VideoBucket).Msg("key found in bucket")

	// Insert video into DB
	log.Debug().Msg(user.Id)
	video, err := v.videoRepository.CreateVideo(mongodb.UnmarshalProtobuf(req.Video, user.Id))
	if err != nil {
		log.Warn().Interface("error", err).Msg("unable to create video in database")
		return err
	}
	log.Info().Str("video_id", video.ID.Hex()).Msg("created video")
	req.Video.Id = video.ID.Hex()

	// Publish event
	v.videoCreatedPublisher.PublishVideoCreated(&pbVideo.VideoCreatedEvent{
		Video:  req.Video,
		UserId: user.Id,
	})

	res.Video = req.Video

	return nil
}

// GetById fetches a video by it's ID
func (v *videoService) GetById(ctx context.Context, req *pbVideo.GetVideoRequest, res *pbVideo.GetVideoResponse) error {
	if req.Id == "" {
		return errors.MissingVideoId
	}

	video, err := v.videoRepository.FindById(req.Id)
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
	res.SignedUrl, _ = v.s3Bucket.GetSignedResourceURL(video.Storage.OutputKey)
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
			DislikeCound: video.Statistics.DislikeCount,
			LikeCount:    video.Statistics.LikeCount,
			ViewCount:    video.Statistics.ViewCount,
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

	videos, err := v.videoRepository.FindByUserId(req.UserId)
	if err != nil {
	    return err
	}

	for _, video := range videos {
		res.Videos = append(res.Videos, &pbVideo.VideoDetails{
			Id: video.ID.Hex(),
			Title: video.Title,
			Description: video.Description,
			Tags: video.Tags,
			Statistics: &pbVideo.VideoStatistics{
				DislikeCound: video.Statistics.DislikeCount,
				LikeCount:    video.Statistics.LikeCount,
				ViewCount:    video.Statistics.ViewCount,
			},
		})
	}

	log.Debug().Str("user", req.UserId).Msgf("fetched %d videos of user", len(videos))

	return nil
}

func (v *videoService) awaitTranscodeCompletedEvent() {
	for transcoded := range v.transcodeCompletedChan {
		log.Info().Str("job", transcoded.Transcode.JobId).Str("video", transcoded.VideoId).Msg("transcoding job completed")

		// TODO: change to FindById
		video, err := v.videoRepository.FindByRawStorageKey(transcoded.Transcode.InputKey)
		if err != nil {
			log.Warn().
				Str("topic", broker.TranscodeCompletedTopic).
				Str("job", transcoded.Transcode.JobId).
				Str("input_key", transcoded.Transcode.InputKey).
				Interface("error", err).
				Msg("received job from non-existing video-input key")
			continue
		}

		// update video
		video.Storage.OutputKey = fmt.Sprintf("%s%s", transcoded.Transcode.OutputKeyPrefix, transcoded.Transcode.OutputKey)
		video.Transcode.Completed = true

		err = v.videoRepository.UpdateVideo(video)
		if err != nil {
			log.Warn().Interface("error", err).Msg("unable to update video")
			continue
		}
		log.Info().Str("video", video.ID.Hex()).Msg("video transcoding completed, updated successfully")
	}
}

func (v *videoService) awaitTranscodeFailedEvent() {
	for transcoded := range v.transcodeFailedChan {
		log.Info().Str("job", transcoded.Transcode.JobId).Str("video", transcoded.VideoId).Msg("transcoding job failed")

		// TODO: change to FindById
		video, err := v.videoRepository.FindByRawStorageKey(transcoded.Transcode.InputKey)
		if err != nil {
			log.Warn().
				Str("topic", broker.TranscodeFailedTopic).
				Str("job", transcoded.Transcode.JobId).
				Str("input_key", transcoded.Transcode.InputKey).
				Interface("error", err).
				Msg("received job from non-existing video-input key")
			continue
		}

		// append errors
		for _, err := range transcoded.Transcode.Status.ErrorMessages {
			video.Transcode.Errors = append(video.Transcode.Errors, err)
		}

		// update video
		err = v.videoRepository.UpdateVideo(video)
		if err != nil {
			log.Warn().Interface("error", err).Msg("unable to update video")
			continue
		}
		log.Info().Str("video", video.ID.Hex()).Msg("video transcoding failed, updated successfully")
	}
}
