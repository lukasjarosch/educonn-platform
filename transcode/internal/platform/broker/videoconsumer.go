package broker

import (
	"github.com/lukasjarosch/educonn-platform/video/proto"
	"context"
)

const (
	VideoCreatedTopic = "video.events.created"
	VideoCreatedQueue = "video-created-queue"
)

type VideoCreatedSubscriber struct {
	videoCreatedChan chan *educonn_video.VideoCreatedEvent
}

func NewVideoCreatedSubscriber(videoCreatedChan chan *educonn_video.VideoCreatedEvent) *VideoCreatedSubscriber {
	return &VideoCreatedSubscriber{
		videoCreatedChan:videoCreatedChan,
	}
}

func (v *VideoCreatedSubscriber) Process(ctx context.Context, event *educonn_video.VideoCreatedEvent) error {
	v.videoCreatedChan <- event
	return nil
}