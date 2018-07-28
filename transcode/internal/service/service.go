package service

import (
	"context"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/internal/platform/amazon"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/proto"
	"github.com/prometheus/common/log"
)

type transcodeService struct {
	sqsConsumer      *amazon.SQSTranscodeEventConsumer
	sqsContext       context.Context
	transcoderClient *amazon.ElasticTranscoderClient
}

func NewTranscodeService(sqsConsumer *amazon.SQSTranscodeEventConsumer, sqsContext context.Context, transcoderClient *amazon.ElasticTranscoderClient) *transcodeService {
	return &transcodeService{
		sqsConsumer:      sqsConsumer,
		sqsContext:       sqsContext,
		transcoderClient: transcoderClient,
	}
}

func (t *transcodeService) CreateJob(ctx context.Context, request *educonn_transcode.CreateJobRequest, response *educonn_transcode.CreateJobResponse) error {

	res, err := t.transcoderClient.CreateJob(request.Job.InputKey)
	if err != nil {
		log.Warnf("transcoding failed: %v", err)
		return err
	}

	response.Job = &educonn_transcode.TranscodeDetails{
		JobId:           *res.Job.Id,
		InputKey:        request.Job.InputKey,
		PipelineId:      *res.Job.PipelineId,
		OutputKeyPrefix: *res.Job.OutputKeyPrefix,
		OutputKey:       *res.Job.Output.Key,
	}

	jobStatus := *res.Job.Status
	status := &educonn_transcode.TranscodeStatus{
		Completed: false,
		Error: false,
		Started: false,
	}

	if jobStatus == "Submitted" {
		status.Started = true
	}
	if jobStatus == "Error" {
		status.Error = true
	}

	response.Job.Status = status

	log.Infof("[ElasticTranscoder] started new job '%s' on pipeline '%s'", response.Job.JobId, response.Job.PipelineId)

	return nil
}
