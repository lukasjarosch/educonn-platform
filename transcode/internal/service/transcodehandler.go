package service

import (
	"github.com/lukasjarosch/educonn-master-thesis/transcode/internal/platform/amazon"
	"github.com/prometheus/common/log"
)

// transcodeHandler implements the ElasticTranscodeEvendHandler interface
type transcodeHandler struct {
}

func NewTranscodeHandler() *transcodeHandler {
	return &transcodeHandler{}
}

func (t *transcodeHandler) OnCompleted(message *amazon.ElasticTranscoderMessage) error {
	log.Infof("[ElasticTranscoder] COMPLETED job '%s' on pipeline '%s': %s", message.JobId, message.PipelineId, message.Outputs[0].Key)
	return nil
}

func (t *transcodeHandler) OnWarning(message *amazon.ElasticTranscoderMessage) error {
	log.Warn(message.MessageDetails)
	return nil
}

func (t *transcodeHandler) OnError(message *amazon.ElasticTranscoderMessage) error {
	log.Info("OnError")
	return nil
}
