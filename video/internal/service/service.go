package service

import (
	"context"
	"fmt"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/proto"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/amazon"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/config"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/errors"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/mongodb"
	"github.com/lukasjarosch/educonn-master-thesis/video/proto"
	"github.com/prometheus/common/log"
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
		log.Warnf("[DB] video with raw key %s already exists", existingVideo.Storage.RawKey)
		return errors.RawVideoAlreadyExists
	}

	// Check if file actually exists
	fileKey := req.Video.Storage.RawKey
	err := v.s3Bucket.CheckFileExists(fileKey, config.AwsS3VideoBucket)
	if err != nil {
		log.Infof("[S3] key '%s' does not exist in bucket", fileKey)
		res.Errors = append(res.Errors, &educonn_video.Error{
			Description: errors.RawVideoFileS3NotFound.Error(),
			Code:        404,
		})
		return err
	}
	log.Infof("[S3] key '%s' exists in bucket", fileKey)

	// TODO: insert into DB
	video, err := v.videoRepository.CreateVideo(mongodb.UnmarshalProtobuf(req.Video))
	if err != nil {
		log.Warnf("[DB] error creating video: %s", err)
		return err
	}
	log.Infof("[DB] created video '%s'", video.ID.Hex())

	v.videoCreatedPublisher.PublishVideoCreated(&educonn_video.VideoCreatedEvent{
		Video:  req.Video,
		UserId: "TODO",
	})

	res.Video = req.Video

	return nil
}

func (v *videoService) awaitTranscodeCompletedEvent() {
	for transcoded := range v.transcodeCompletedChan {
		log.Infof("Job '%s' for video '%s' completed. Video can now be streamed", transcoded.Transcode.JobId, "TODO")

		video, err := v.videoRepository.FindByRawStorageKey(transcoded.Transcode.InputKey)
		if err != nil {
			log.Info(err)
		}
		video.Status.Completed = true
		video.Storage.OutputKey = fmt.Sprintf("%s%s", transcoded.Transcode.OutputKeyPrefix, transcoded.Transcode.OutputKey)

		err = v.videoRepository.UpdateVideo(video)
		if err != nil {
		    log.Warn(err)
		    continue
		}
		log.Infof("[DB] updated video '%s'", video.ID.Hex())
	}
}

func (v *videoService) awaitTranscodeFailedEvent() {
	for transcoded := range v.transcodeFailedChan {
		log.Infof("Job '%s' for video '%s' failed", transcoded.Transcode.JobId, "TODO")

		video, err := v.videoRepository.FindByRawStorageKey(transcoded.Transcode.InputKey)
		if err != nil {
			log.Info(err)
		}
		video.Status.Error = true
		video.Status.Completed = false
		video.Status.Warning = false
		video.Status.ErrorMessages = []string{}
		for _, err := range transcoded.Transcode.Status.ErrorMessages {
			video.Status.ErrorMessages = append(video.Status.ErrorMessages, err)
		}

		err = v.videoRepository.UpdateVideo(video)
		if err != nil {
			log.Warn(err)
			continue
		}
		log.Infof("[DB] updated video '%s'", video.ID.Hex())
	}
}
