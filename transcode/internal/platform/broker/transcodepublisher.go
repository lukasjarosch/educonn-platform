package broker


import (
	"github.com/micro/go-micro"
	pbTranscode "github.com/lukasjarosch/educonn-platform/transcode/proto"
	"context"
	"github.com/rs/zerolog/log"
)

const (
	VideoTranscodingCompleted = "transcode.events.completed"
	VideoTranscodingFailed    = "transcode.events.failed"
)

type TranscodeEventPublisher struct {
	transcodingCompletedPublisher micro.Publisher
	transcodingFailedPublisher    micro.Publisher
}

func NewTranscodeEventPublisher(completedPublisher micro.Publisher, failedPublisher micro.Publisher) *TranscodeEventPublisher {
	return &TranscodeEventPublisher{
		transcodingCompletedPublisher: completedPublisher,
		transcodingFailedPublisher:    failedPublisher,
	}
}

func (t *TranscodeEventPublisher) PublishTranscodingCompleted(event pbTranscode.TranscodingCompletedEvent) (err error) {
	if err = t.transcodingCompletedPublisher.Publish(context.Background(), event); err != nil {
		log.Warn().Str("topic", VideoTranscodingCompleted).Interface("error", err).Interface("event", event).Msg("unable to publish event")
		return nil
	}
	log.Info().Str("topic", VideoTranscodingCompleted).Interface("event", event).Msg("published event")
	return nil
}

func (t *TranscodeEventPublisher) PublishTranscodingFailed(event pbTranscode.TranscodingFailedEvent) (err error) {
	if err = t.transcodingFailedPublisher.Publish(context.Background(), event); err != nil {
		log.Warn().Str("topic", VideoTranscodingFailed).Interface("error", err).Interface("event", event).Msg("unable to publish event")
		return nil
	}
	log.Info().Str("topic", VideoTranscodingFailed).Interface("event", event).Msg("published event")
	return nil
}
