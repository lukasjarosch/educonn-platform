package service

import (
	"context"
	"github.com/lukasjarosch/educonn-master-thesis/video/proto"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/amazon"
	log "github.com/sirupsen/logrus"
)

type videoService struct {
	videoCreatedPublisher videoCreatedPublisher
	sqsConsumer *amazon.SQSTranscodeEventConsumer
	sqsS3Context context.Context
}

type videoCreatedPublisher interface {
	PublishVideoCreated(event *educonn_video.VideoCreatedEvent) (err error)
}

func NewVideoService(vidCreatedPub videoCreatedPublisher, sqsS3EventConsumer *amazon.SQSTranscodeEventConsumer, sqsS3Context context.Context) educonn_video.VideoHandler {
	svc := &videoService{
		videoCreatedPublisher: vidCreatedPub,
		sqsConsumer:sqsS3EventConsumer,
		sqsS3Context:sqsS3Context,
	}

	log.Info("Start consuming SQS S3 events")
	go svc.sqsConsumer.Consume(NewTranscodeHandler())
	go svc.awaitSqsS3Event()

	return svc
}

func (v *videoService) Create(ctx context.Context, req *educonn_video.CreateVideoRequest, res *educonn_video.CreateVideoResponse) error {
	v.videoCreatedPublisher.PublishVideoCreated(&educonn_video.VideoCreatedEvent{
		Video:  res.Video,
		UserId: "asdfasdf",
	})
	return nil
}

func (v *videoService) awaitSqsS3Event() {
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
