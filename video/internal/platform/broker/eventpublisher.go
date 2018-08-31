package broker

import (
	"context"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	"github.com/micro/go-micro"
	"github.com/rs/zerolog/log"
)

const (
	VideoCreatedTopic = "video.events.created"
	VideoProcessedTopic = "video.events.processed"
)

type EventPublisher struct {
	videoCreatedPublisher micro.Publisher
	videoProcessedPublisher micro.Publisher
}

func NewEventPublisher(videoCreatedPublisher micro.Publisher, videoProcessedPublisher micro.Publisher) *EventPublisher {
	return &EventPublisher{videoCreatedPublisher: videoCreatedPublisher, videoProcessedPublisher:videoProcessedPublisher}
}

func (p *EventPublisher) PublishVideoCreated(ctx context.Context, event *pbVideo.VideoCreatedEvent) (err error) {
	if err = p.videoCreatedPublisher.Publish(ctx, event); err != nil {
		log.Warn().Str("topic", VideoCreatedTopic).Interface("error", err).Msg("unable to publish event")
		return nil
	}
	log.Info().Str("topic", VideoCreatedTopic).Interface("event", event).Msg("published event")
	return nil
}

func (p EventPublisher) PublishVideoProcessed(ctx context.Context, event *pbVideo.VideoProcessedEvent) (err error) {
	if err = p.videoProcessedPublisher.Publish(ctx, event); err != nil {
		log.Warn().Str("topic", VideoProcessedTopic).Interface("error", err).Msg("unable to publish event")
		return nil
	}
	log.Info().Str("topic", VideoProcessedTopic).Interface("event", event).Msg("published event")
	return nil
}
