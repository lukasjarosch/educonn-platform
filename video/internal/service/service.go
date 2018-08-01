package service

import (
	"context"
	"fmt"
	"github.com/lukasjarosch/educonn-platform/transcode/proto"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/amazon"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/errors"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/mongodb"
	"github.com/rs/zerolog/log"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/broker"
	"github.com/lukasjarosch/educonn-platform/video/proto"
)

type videoService struct {
	videoCreatedPublisher  videoCreatedPublisher
	s3Bucket               *amazon.S3Bucket
	transcodeCompletedChan chan *educonn_transcode.TranscodingCompletedEvent
	transcodeFailedChan    chan *educonn_transcode.TranscodingFailedEvent
	videoRepository        *mongodb.VideoRepository
}

type videoCreatedPublisher interface {
	PublishVideoCreated(event *educonn_video.VideoCreatedEvent) (err error)
}

func NewVideoService(vidCreatedPub videoCreatedPublisher,
	bucket *amazon.S3Bucket,
	transcodeCompletedChan chan *educonn_transcode.TranscodingCompletedEvent,
	transcodeFailedChan chan *educonn_transcode.TranscodingFailedEvent,
	videoRepo *mongodb.VideoRepository) educonn_video.VideoHandler {
	svc := &videoService{
		videoCreatedPublisher:  vidCreatedPub,
		s3Bucket:               bucket,
		transcodeCompletedChan: transcodeCompletedChan,
		transcodeFailedChan:    transcodeFailedChan,
		videoRepository:        videoRepo,
	}

	go svc.awaitTranscodeCompletedEvent()
	go svc.awaitTranscodeFailedEvent()

	return svc
}

func (v *videoService) Create(ctx context.Context, req *educonn_video.CreateVideoRequest, res *educonn_video.CreateVideoResponse) error {

	// Check if raw file does already exist
	existingVideo, _ := v.videoRepository.FindByRawStorageKey(req.Video.Storage.RawKey)
	if existingVideo != nil {
		log.Warn().Str("file_key", existingVideo.Storage.RawKey).Msg("a source video with that key already exists")
		return errors.RawVideoAlreadyExists
	}

	// Check if file actually exists in S3
	fileKey := req.Video.Storage.RawKey
	err := v.s3Bucket.CheckFileExists(fileKey, config.AwsS3VideoBucket)
	if err != nil {
		log.Warn().Str("key", fileKey).Msg("key does not exist in bucket")
		res.Errors = append(res.Errors, &educonn_video.Error{
			Description: errors.RawVideoFileS3NotFound.Error(),
			Code:        404,
		})
		return err
	}
	log.Info().Str("key", fileKey).Str("bucket", config.AwsS3VideoBucket).Msg("key found in bucket")

	// Insert video into DB
	video, err := v.videoRepository.CreateVideo(mongodb.UnmarshalProtobuf(req.Video))
	if err != nil {
		log.Warn().Interface("error", err).Msg("unable to create video in database")
		return err
	}
	log.Info().Str("video_id", video.ID.Hex()).Msg("created video")
	req.Video.Id = video.ID.Hex()

	// Publish event
	v.videoCreatedPublisher.PublishVideoCreated(&educonn_video.VideoCreatedEvent{
		Video:  req.Video,
		UserId: "TODO",
	})

	res.Video = req.Video

	return nil
}

func (v *videoService) GetById(ctx context.Context, req *educonn_video.GetVideoRequest, res *educonn_video.GetVideoResponse) error {
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
	res.Video = &educonn_video.VideoDetails{
		Id: video.ID.Hex(),
		Title: video.Title,
		Description: video.Description,
		Tags: video.Tags,
		Storage: &educonn_video.VideoStorage{
			RawKey: video.Storage.RawKey,
			TranscodedKey: video.Storage.OutputKey,
		},
		Statistics: &educonn_video.VideoStatistics{
			DislikeCound: video.Statistics.DislikeCount,
			LikeCount: video.Statistics.LikeCount,
			ViewCount: video.Statistics.ViewCount,
		},
		Status: &educonn_video.VideoStatus{
			Completed: video.Transcode.Completed,
			Error: trError,
			ErrorMessages: video.Transcode.Errors,
		},
	}

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
