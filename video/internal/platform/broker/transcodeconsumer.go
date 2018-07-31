package broker

import (
	"github.com/lukasjarosch/educonn-platform/transcode/proto"
	"context"
	"github.com/prometheus/common/log"
)

const (
	TranscodeCompletedTopic = "transcode.events.completed"
	TranscodeFailedTopic = "transcode.events.failed"
	TranscodeCompletedQueue = "transcode-completed"
	TranscodeFailedQueue = "transcode-failed"
)

type TranscodeCompletedSubscriber struct {
	transcodeCompletedChan chan *educonn_transcode.TranscodingCompletedEvent
}

func NewTranscodeCompletedSubscriber(transcodeCompletedChan chan *educonn_transcode.TranscodingCompletedEvent) *TranscodeCompletedSubscriber{
	return &TranscodeCompletedSubscriber{transcodeCompletedChan:transcodeCompletedChan}
}

func (t *TranscodeCompletedSubscriber) Process(ctx context.Context, event *educonn_transcode.TranscodingCompletedEvent) error {
	t.transcodeCompletedChan <- event
	log.Infof("[SUB] received '%s' event for transcode-job-id '%s'", TranscodeCompletedTopic, event.Transcode.JobId)
	return nil
}

// ------------------------------------ //

type TranscodeFailedSubscriber struct {
	transcodeFailedChan chan *educonn_transcode.TranscodingFailedEvent
}

func NewTranscodeFailedSubscriber(transcodeFailedChan chan *educonn_transcode.TranscodingFailedEvent) *TranscodeFailedSubscriber{
	return &TranscodeFailedSubscriber{transcodeFailedChan:transcodeFailedChan}
}

func (t *TranscodeFailedSubscriber) Process(ctx context.Context, event *educonn_transcode.TranscodingFailedEvent) error {
	t.transcodeFailedChan<- event
	log.Infof("[SUB] received '%s' event for transcode-job-id '%s'", TranscodeFailedTopic, event.Transcode.JobId)
	return nil
}
