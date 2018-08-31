package broker

import (
	pbTranscode "github.com/lukasjarosch/educonn-platform/transcode/proto"
	"context"
	"github.com/rs/zerolog/log"
)

const (
	TranscodeCompletedTopic = "transcode.events.completed"
	TranscodeFailedTopic = "transcode.events.failed"
	TranscodeCompletedQueue = "transcode-completed"
	TranscodeFailedQueue = "transcode-failed"
)

type TranscodingCompletedEvent struct {
	Event   *pbTranscode.TranscodingCompletedEvent
	Context context.Context
}

type TranscodeCompletedSubscriber struct {
	transcodeCompletedChan chan TranscodingCompletedEvent
}

func NewTranscodeCompletedSubscriber(transcodeCompletedChan chan TranscodingCompletedEvent) *TranscodeCompletedSubscriber{
	return &TranscodeCompletedSubscriber{transcodeCompletedChan:transcodeCompletedChan}
}

func (t *TranscodeCompletedSubscriber) Process(ctx context.Context, event *pbTranscode.TranscodingCompletedEvent) error {
	t.transcodeCompletedChan <- TranscodingCompletedEvent{Context:ctx, Event:event}
	log.Info().Str("topic", TranscodeCompletedTopic).Str("job", event.Transcode.JobId).Msg("received TranscodingCompletedEvent")
	return nil
}

// ------------------------------------ //

type TranscodingFailedEvent struct {
	Event *pbTranscode.TranscodingFailedEvent
	Context context.Context
}

type TranscodeFailedSubscriber struct {
	transcodeFailedChan chan TranscodingFailedEvent
}

func NewTranscodeFailedSubscriber(transcodeFailedChan chan TranscodingFailedEvent) *TranscodeFailedSubscriber{
	return &TranscodeFailedSubscriber{transcodeFailedChan:transcodeFailedChan}
}

func (t *TranscodeFailedSubscriber) Process(ctx context.Context, event *pbTranscode.TranscodingFailedEvent) error {
	t.transcodeFailedChan<- TranscodingFailedEvent{Context:ctx, Event:event}
	log.Info().Str("topic", TranscodeFailedTopic).Str("job", event.Transcode.JobId).Msg("received TranscodingCompletedEvent")
	return nil
}
