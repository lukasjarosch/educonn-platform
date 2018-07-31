package broker


import (
	"github.com/micro/go-micro"
	"github.com/lukasjarosch/educonn-platform/transcode/proto"
	"context"
	"github.com/prometheus/common/log"
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

func (t *TranscodeEventPublisher) PublishTranscodingCompleted(event educonn_transcode.TranscodingCompletedEvent) (err error) {
	if err = t.transcodingCompletedPublisher.Publish(context.Background(), event); err != nil {
		log.Warnf("[pub] Unable to publish to %s: %+v", VideoTranscodingCompleted, err)
		return nil
	}
	log.Infof("[pub] published '%s' for transcoding job '%s'", VideoTranscodingCompleted, event.Transcode.JobId)
	return nil
}

func (t *TranscodeEventPublisher) PublishTranscodingFailed(event educonn_transcode.TranscodingFailedEvent) (err error) {
	if err = t.transcodingFailedPublisher.Publish(context.Background(), event); err != nil {
		log.Warnf("[pub] Unable to publish to %s: %+v", VideoTranscodingFailed, err)
		return nil
	}
	log.Infof("[pub] published '%s' for transcoding job '%s'", VideoTranscodingFailed, event.Transcode.JobId)
	return nil
}
