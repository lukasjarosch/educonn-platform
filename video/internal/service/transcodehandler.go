package service

import (
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/amazon"
	log "github.com/sirupsen/logrus"
)

// transcodeHandler implements the ElasticTranscodeEvendHandler interface
type transcodeHandler struct {
}

func NewTranscodeHandler() *transcodeHandler {
	return &transcodeHandler{}
}

func (t *transcodeHandler) OnCompleted(message *amazon.ElasticTranscoderMessage) error {
	log.Info("OnCompleted")
	return nil
}

func (t *transcodeHandler) OnWarning(message *amazon.ElasticTranscoderMessage) error {
	log.Info("OnWarning")
	return nil
}

func (t *transcodeHandler) OnError(message *amazon.ElasticTranscoderMessage) error {
	log.Info("OnError")
	return nil
}
