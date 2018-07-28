package service

import (
	"context"
	"github.com/lukasjarosch/educonn-master-thesis/video/proto"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/amazon"
	log "github.com/sirupsen/logrus"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/config"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/errors"
)

type videoService struct {
	videoCreatedPublisher videoCreatedPublisher
	sqsConsumer *amazon.SQSTranscodeEventConsumer
	sqsS3Context context.Context
	s3Bucket *amazon.S3Bucket
}

type videoCreatedPublisher interface {
	PublishVideoCreated(event *educonn_video.VideoCreatedEvent) (err error)
}

func NewVideoService(vidCreatedPub videoCreatedPublisher, sqsS3EventConsumer *amazon.SQSTranscodeEventConsumer, sqsS3Context context.Context, bucket *amazon.S3Bucket) educonn_video.VideoHandler {
	svc := &videoService{
		videoCreatedPublisher: vidCreatedPub,
		sqsConsumer:sqsS3EventConsumer,
		sqsS3Context:sqsS3Context,
		s3Bucket:bucket,
	}

	log.Infof("[SQS] start consuming from: %s", config.AwsSqsVideoQueueName)
	go svc.sqsConsumer.Consume(NewTranscodeHandler())
	go svc.awaitSQSEvent()

	return svc
}

func (v *videoService) Create(ctx context.Context, req *educonn_video.CreateVideoRequest, res *educonn_video.CreateVideoResponse) error {
	v.videoCreatedPublisher.PublishVideoCreated(&educonn_video.VideoCreatedEvent{
		Video:  res.Video,
		UserId: "asdfasdf",
	})

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

	return nil
}

func (v *videoService) awaitSQSEvent() {
	handler := NewTranscodeHandler()
	for msg := range v.sqsConsumer.ElasticTranscoderChannel {
		log.Infof("Handling: %s", msg.MessageDetails)

		// COMPLETED
		if msg.State == amazon.TranscodeStatusCompleted {
			err := handler.OnCompleted(msg)
			if err != nil {
				log.Info(err)
				continue
			}
		}

		// WARNING
		if msg.State == amazon.TranscodeStatusWarning{
			err := handler.OnWarning(msg)
			if err != nil {
				log.Info(err)
				continue
			}
		}

		// ERROR
		if msg.State == amazon.TranscodeStatusError{
			err := handler.OnError(msg)
			if err != nil {
				log.Info(err)
				continue
			}
		}
	}
}
