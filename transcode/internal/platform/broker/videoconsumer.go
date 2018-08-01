package broker

import (
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	"context"
)

const (
	VideoCreatedTopic = "video.events.created"
	VideoCreatedQueue = "video-created-queue"
)

type VideoCreatedSubscriber struct {
	videoCreatedChan chan *pbVideo.VideoCreatedEvent
}

func NewVideoCreatedSubscriber(videoCreatedChan chan *pbVideo.VideoCreatedEvent) *VideoCreatedSubscriber {
	return &VideoCreatedSubscriber{
		videoCreatedChan:videoCreatedChan,
	}
}

func (v *VideoCreatedSubscriber) Process(ctx context.Context, event *pbVideo.VideoCreatedEvent) error {
	v.videoCreatedChan <- event
	return nil
}