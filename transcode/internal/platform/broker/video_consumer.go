package broker

import (
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	"context"
)

const (
	VideoCreatedTopic = "video.events.created"
	VideoCreatedQueue = "video-created-queue"
)

type VideoCreatedEvent struct {
	Event   *pbVideo.VideoCreatedEvent
	Context context.Context
}

type VideoCreatedSubscriber struct {
	videoCreatedChan chan VideoCreatedEvent
}

func NewVideoCreatedSubscriber(videoCreatedEventChan chan VideoCreatedEvent) *VideoCreatedSubscriber {
	return &VideoCreatedSubscriber{
		videoCreatedChan: videoCreatedEventChan,
	}
}

func (v *VideoCreatedSubscriber) Process(ctx context.Context, event *pbVideo.VideoCreatedEvent) error {
	v.videoCreatedChan <- VideoCreatedEvent{Context:ctx, Event:event}
	return nil
}

