package broker

import (
	"context"

	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
)

type VideoProcessedEvent struct {
	Event *pbVideo.VideoProcessedEvent
	Context context.Context
}

type VideoProcessedSubscriber struct {
	videoProcessedChan chan VideoProcessedEvent
}

func NewVideoProcessedSubscriber(videoProcessedChan chan VideoProcessedEvent) *VideoProcessedSubscriber {
	return &VideoProcessedSubscriber{
		videoProcessedChan: videoProcessedChan,
	}
}

func (v *VideoProcessedSubscriber) Process(ctx context.Context, event *pbVideo.VideoProcessedEvent) error {
	v.videoProcessedChan <- VideoProcessedEvent{Context:ctx, Event:event}
	return nil
}
