package service

import (
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/amazon"
	"github.com/micro/go-micro"
	"github.com/prometheus/common/log"
	"github.com/lukasjarosch/educonn-platform/transcode/proto"
	"context"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/broker"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/mongodb"
	"time"
	"github.com/pkg/errors"
	"fmt"
)

// transcodeHandler implements the ElasticTranscodeEvendHandler interface
type transcodeHandler struct {
	completedPublisher micro.Publisher
	failedPublisher    micro.Publisher
	transcodeRepository *mongodb.TranscodeRepository
}

func NewTranscodeHandler(completedPublisher micro.Publisher, failedPublisher micro.Publisher, transcodeRepo *mongodb.TranscodeRepository) *transcodeHandler {
	return &transcodeHandler{
		completedPublisher: completedPublisher,
		failedPublisher: failedPublisher,
		transcodeRepository:transcodeRepo,
	}
}

func (t *transcodeHandler) OnCompleted(message *amazon.ElasticTranscoderMessage) error {
	log.Infof("[ElasticTranscoder] COMPLETED job '%s' on pipeline '%s': %s", message.JobId, message.PipelineId, message.Outputs[0].Key)

	// Update transcoding job
	job, err := t.transcodeRepository.FindByJobId(message.JobId)
	if err != nil {
	    log.Warnf("[DB] received completed event for UNKNOWN JOB '%s': %s", message.JobId, err)
	    return err
	}
	job.Status.Started = false
	job.Status.Completed = true
	job.EndedAt = time.Now()
	job, err = t.transcodeRepository.UpdateJob(job)
	if err != nil {
	    err = errors.New(fmt.Sprintf("[DB] unable to update job '%s': %s",job.ID.Hex(), err))
	    return err
	}
	log.Infof("[DB] updated job '%s'", job.ID.Hex())

	// Prepare event
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
		VideoId: job.VideoId.Hex(),
	}

	// Publish
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

	job, err := t.transcodeRepository.FindByJobId(message.JobId)
	if err != nil {
		err = errors.New(fmt.Sprintf("[DB] received error event for UNKNOWN JOB '%s': %s", message.JobId, err))
		return err
	}

	// TODO: save to DB
	log.Infof("[DB] updated job '%s'", job.ID.Hex())

	evt := &educonn_transcode.TranscodingFailedEvent{
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
		VideoId: job.VideoId.Hex(),
	}

	t.failedPublisher.Publish(context.Background(), evt)
	log.Infof("[PUB] pubbed '%s'", broker.VideoTranscodingFailed)
	return nil
}
