package service

import (
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/amazon"
	"github.com/micro/go-micro"
	"github.com/rs/zerolog/log"
	pbTranscode "github.com/lukasjarosch/educonn-platform/transcode/proto"
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
	log.Info().
		Str("job", message.JobId).
		Str("pipeline", message.PipelineId).
		Str("file_key", message.Outputs[0].Key).
		Msg("ElasticTranscoder job completed")

	// Update transcoding job
	job, err := t.transcodeRepository.FindByJobId(message.JobId)
	if err != nil {
		err = errors.New("ElasticTranscoder job completed but the job is unknown and was not started by me")
		log.Warn().Str("job", message.JobId).Interface("error", err).Msg(err.Error())
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

	// Prepare event
	evt := &pbTranscode.TranscodingCompletedEvent{
		Transcode: &pbTranscode.TranscodeDetails{
			JobId: message.JobId,
			InputKey: message.Input.Key,
			PipelineId: message.PipelineId,
			Status: &pbTranscode.TranscodeStatus{
				Completed: true,
			},
			OutputKey:message.Outputs[0].Key,
			OutputKeyPrefix:message.OutputKeyPrefix,
		},
		VideoId: job.VideoId.Hex(),
	}

	// Publish
	t.completedPublisher.Publish(context.Background(), evt)
	log.Info().Str("topic", broker.VideoTranscodingCompleted).Interface("event", evt)

	return nil
}

func (t *transcodeHandler) OnWarning(message *amazon.ElasticTranscoderMessage) error {
	log.Warn().Msg(message.MessageDetails)
	return nil
}

func (t *transcodeHandler) OnError(message *amazon.ElasticTranscoderMessage) error {
	log.Info().
		Str("job", message.JobId).
		Str("pipeline", message.PipelineId).
		Str("file_key", message.Outputs[0].Key).
		Msg("ElasticTranscoder job failed")

	job, err := t.transcodeRepository.FindByJobId(message.JobId)
	if err != nil {
		err = errors.New("ElasticTranscoder job completed but the job is unknown and was not started by me")
		log.Warn().Str("job", message.JobId).Interface("error", err).Msg(err.Error())
		return err
	}

	// TODO: save to DB

	evt := &pbTranscode.TranscodingFailedEvent{
		Transcode: &pbTranscode.TranscodeDetails{
			JobId: message.JobId,
			InputKey: message.Input.Key,
			PipelineId: message.PipelineId,
			Status: &pbTranscode.TranscodeStatus{
				Error:true,
				ErrorMessages: []string{message.MessageDetails},
			},
			OutputKey:message.Outputs[0].Key,
			OutputKeyPrefix:message.OutputKeyPrefix,
		},
		VideoId: job.VideoId.Hex(),
	}

	t.failedPublisher.Publish(context.Background(), evt)
	log.Info().Str("topic", broker.VideoTranscodingFailed).Interface("event", evt)
	return nil
}
