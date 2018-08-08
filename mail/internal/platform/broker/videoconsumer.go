package broker

import (
	"context"

	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
)

type VideoProcessedSubscriber struct {
	videoProcessedChan chan *pbVideo.VideoProcessedEvent
}

func NewVideoProcessedSubscriber(videoProcessedChan chan *pbVideo.VideoProcessedEvent) *VideoProcessedSubscriber {
	return &VideoProcessedSubscriber{
		videoProcessedChan: videoProcessedChan,
	}
}

func (v *VideoProcessedSubscriber) Process(ctx context.Context, event *pbVideo.VideoProcessedEvent) error {
	v.videoProcessedChan <- event
	return nil
}
