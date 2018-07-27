package service

import (
	"context"
	"github.com/lukasjarosch/educonn-master-thesis/video/proto"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/amazon"
	log "github.com/sirupsen/logrus"
)

type videoService struct {
	videoCreatedPublisher videoCreatedPublisher
	sqsConsumer *amazon.SqsS3EventConsumer
	sqsS3Context context.Context
}

type videoCreatedPublisher interface {
	PublishVideoCreated(event *educonn_video.VideoCreatedEvent) (err error)
}

func NewVideoService(vidCreatedPub videoCreatedPublisher, sqsS3EventConsumer *amazon.SqsS3EventConsumer, sqsS3Context context.Context) educonn_video.VideoHandler {
	svc := &videoService{
		videoCreatedPublisher: vidCreatedPub,
		sqsConsumer:sqsS3EventConsumer,
		sqsS3Context:sqsS3Context,
	}

	log.Info("Start consuming SQS S3 events")
	go svc.sqsConsumer.Consume()
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
	for s3event := range v.sqsConsumer.VideoUploadedChannel {
		log.Infof("HANDLING event %s: %s", s3event.EventName, s3event.S3.Object.Key)

		// Find video entry by object-key in our db
		// Update entry
		// Check video
		// Pub event: VideoUploadedEvent
	}
}
