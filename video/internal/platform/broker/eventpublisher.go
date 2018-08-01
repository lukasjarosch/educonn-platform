package broker

import (
	"context"
	"github.com/lukasjarosch/educonn-platform/video/proto"
	"github.com/micro/go-micro"
	"github.com/rs/zerolog/log"
)

const (
	VideoCreatedTopic = "video.events.created"
)

type EventPublisher struct {
	videoCreatedPublisher micro.Publisher
}

func NewEventPublisher(videoCreatedPublisher micro.Publisher) *EventPublisher {
	return &EventPublisher{videoCreatedPublisher: videoCreatedPublisher}
}

func (p *EventPublisher) PublishVideoCreated(event *educonn_video.VideoCreatedEvent) (err error) {
	if err = p.videoCreatedPublisher.Publish(context.Background(), event); err != nil {
		log.Warn().Str("topic", VideoCreatedTopic).Interface("error", err).Str("event", "VideoCreatedEvent").Msg("unable to publish event")
		return nil
	}
	log.Info().Str("topic", VideoCreatedTopic).Interface("event", event).Msg("published event")
	return nil
}
