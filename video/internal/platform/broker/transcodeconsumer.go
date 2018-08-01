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

type TranscodeCompletedSubscriber struct {
	transcodeCompletedChan chan *pbTranscode.TranscodingCompletedEvent
}

func NewTranscodeCompletedSubscriber(transcodeCompletedChan chan *pbTranscode.TranscodingCompletedEvent) *TranscodeCompletedSubscriber{
	return &TranscodeCompletedSubscriber{transcodeCompletedChan:transcodeCompletedChan}
}

func (t *TranscodeCompletedSubscriber) Process(ctx context.Context, event *pbTranscode.TranscodingCompletedEvent) error {
	t.transcodeCompletedChan <- event
	log.Info().Str("topic", TranscodeCompletedTopic).Str("job", event.Transcode.JobId).Msg("received TranscodingCompletedEvent")
	return nil
}

// ------------------------------------ //

type TranscodeFailedSubscriber struct {
	transcodeFailedChan chan *pbTranscode.TranscodingFailedEvent
}

func NewTranscodeFailedSubscriber(transcodeFailedChan chan *pbTranscode.TranscodingFailedEvent) *TranscodeFailedSubscriber{
	return &TranscodeFailedSubscriber{transcodeFailedChan:transcodeFailedChan}
}

func (t *TranscodeFailedSubscriber) Process(ctx context.Context, event *pbTranscode.TranscodingFailedEvent) error {
	t.transcodeFailedChan<- event
	log.Info().Str("topic", TranscodeFailedTopic).Str("job", event.Transcode.JobId).Msg("received TranscodingCompletedEvent")
	return nil
}
