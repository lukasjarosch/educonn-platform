package service

import (
	"github.com/lukasjarosch/educonn-master-thesis/transcode/internal/platform/amazon"
	"github.com/micro/go-micro"
	"github.com/prometheus/common/log"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/proto"
	"context"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/internal/platform/broker"
)

// transcodeHandler implements the ElasticTranscodeEvendHandler interface
type transcodeHandler struct {
	completedPublisher micro.Publisher
	failedPublisher    micro.Publisher
}

func NewTranscodeHandler(completedPublisher micro.Publisher, failedPublisher micro.Publisher) *transcodeHandler {
	return &transcodeHandler{
		completedPublisher: completedPublisher,
		failedPublisher: failedPublisher,
	}
}

func (t *transcodeHandler) OnCompleted(message *amazon.ElasticTranscoderMessage) error {
	log.Infof("[ElasticTranscoder] COMPLETED job '%s' on pipeline '%s': %s", message.JobId, message.PipelineId, message.Outputs[0].Key)

	evt := &educonn_transcode.TranscodingCompletedEvent{
		Transcode: &educonn_transcode.TranscodeDetails{
			JobId: message.JobId,
			InputKey: message.Input.Key,
			PipelineId: message.PipelineId,
			Status: &educonn_transcode.TranscodeStatus{
				Completed: true,
			},
			OutputKey:message.Outputs[0].Key,
			OutputKeyPrefix:message.OutputKeyPrefix,
		},
	}

	t.completedPublisher.Publish(context.Background(), evt)
	log.Infof("[PUB] pubbed '%s'", broker.VideoTranscodingCompleted)

	return nil
}

func (t *transcodeHandler) OnWarning(message *amazon.ElasticTranscoderMessage) error {
	log.Warn(message.MessageDetails)
	return nil
}

func (t *transcodeHandler) OnError(message *amazon.ElasticTranscoderMessage) error {
	log.Infof("[ElasticTranscoder] FAILED job '%s' on pipeline '%s': %s", message.JobId, message.PipelineId, message.Outputs[0].Key)

	evt := &educonn_transcode.TranscodingCompletedEvent{
		Transcode: &educonn_transcode.TranscodeDetails{
			JobId: message.JobId,
			InputKey: message.Input.Key,
			PipelineId: message.PipelineId,
			Status: &educonn_transcode.TranscodeStatus{
				Error:true,
				ErrorMessages: []string{message.MessageDetails},
			},
			OutputKey:message.Outputs[0].Key,
			OutputKeyPrefix:message.OutputKeyPrefix,
		},
	}

	t.failedPublisher.Publish(context.Background(), evt)
	log.Infof("[PUB] pubbed '%s'", broker.VideoTranscodingFailed)
	return nil
}
