package service

import (
	"context"
	"github.com/lukasjarosch/educonn-master-thesis/video/proto"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/amazon"
	"github.com/prometheus/common/log"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/config"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/errors"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/proto"
)

type videoService struct {
	videoCreatedPublisher videoCreatedPublisher
	s3Bucket *amazon.S3Bucket
	transcodeCompletedChan chan *educonn_transcode.TranscodingCompletedEvent
	transcodeFailedChan chan *educonn_transcode.TranscodingFailedEvent
}

type videoCreatedPublisher interface {
	PublishVideoCreated(event *educonn_video.VideoCreatedEvent) (err error)
}

func NewVideoService(vidCreatedPub videoCreatedPublisher,
	bucket *amazon.S3Bucket,
	transcodeCompletedChan chan *educonn_transcode.TranscodingCompletedEvent,
	transcodeFailedChan chan *educonn_transcode.TranscodingFailedEvent) educonn_video.VideoHandler {
	svc := &videoService{
		videoCreatedPublisher: vidCreatedPub,
		s3Bucket:bucket,
		transcodeCompletedChan:transcodeCompletedChan,
		transcodeFailedChan:transcodeFailedChan,
	}

	go svc.awaitTranscodeCompletedEvent()
	go svc.awaitTranscodeFailedEvent()

	return svc
}

func (v *videoService) Create(ctx context.Context, req *educonn_video.CreateVideoRequest, res *educonn_video.CreateVideoResponse) error {
	// Check if file actually exists
	fileKey := req.Video.Storage.RawKey
	err := v.s3Bucket.CheckFileExists(fileKey, config.AwsS3VideoBucket)
	if err != nil {
		log.Infof("[S3] key '%s' does not exist in bucket", fileKey)
		res.Errors = append(res.Errors, &educonn_video.Error{
			Description: errors.RawVideoFileS3NotFound.Error(),
			Code: 404,
		})
	    return err
	}
	log.Infof("[S3] key '%s' exists in bucket", fileKey)

	// TODO: insert into DB

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

		// TODO: persistence stuff
	}
}

func (v *videoService) awaitTranscodeFailedEvent() {
	for transcoded := range v.transcodeFailedChan{
		log.Infof("Job '%s' for video '%s' failed", transcoded.Transcode.JobId, "TODO")
		for _, err := range transcoded.Transcode.Status.ErrorMessages {
			log.Warn(err)
		}

		// TODO: persistence stuff
	}
}
