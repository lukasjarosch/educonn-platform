package service

import (
	"context"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/internal/platform/amazon"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/internal/platform/config"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/proto"
	"github.com/lukasjarosch/educonn-master-thesis/video/proto"
	"github.com/prometheus/common/log"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/internal/platform/broker"
	"github.com/micro/go-micro"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/internal/platform/mongodb"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type transcodeService struct {
	sqsConsumer      *amazon.SQSTranscodeEventConsumer
	sqsContext       context.Context
	transcoderClient *amazon.ElasticTranscoderClient
	videoCreatedChan chan *educonn_video.VideoCreatedEvent
	transcodingCompletedPublisher micro.Publisher
	transcodingFailedPublisher micro.Publisher
	transcodeRepository *mongodb.TranscodeRepository
}

func NewTranscodeService(sqsConsumer *amazon.SQSTranscodeEventConsumer,
	sqsContext context.Context,
	transcoderClient *amazon.ElasticTranscoderClient,
	videoCreatedChan chan *educonn_video.VideoCreatedEvent,
	completedPublisher micro.Publisher,
	failedPublisher micro.Publisher,
	transcodeRepo *mongodb.TranscodeRepository) *transcodeService {

	svc := &transcodeService{
		sqsConsumer:      sqsConsumer,
		sqsContext:       sqsContext,
		transcoderClient: transcoderClient,
		videoCreatedChan: videoCreatedChan,
		transcodingCompletedPublisher:completedPublisher,
		transcodingFailedPublisher: failedPublisher,
		transcodeRepository:transcodeRepo,
	}

	log.Infof("[SUB] consuming '%s' from '%s'", broker.VideoCreatedTopic, broker.VideoCreatedQueue)
	go svc.awaitVideoCreatedEvent()

	log.Infof("[SQS] start consuming from: %s", config.AwsSqsVideoQueueName)
	go svc.sqsConsumer.Consume()
	go svc.awaitSQSEvent()

	return svc
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
		Error:     false,
		Started:   true,
	}

	transcodeJob, err := t.transcodeRepository.CreateTranscodingJob(&mongodb.TranscodingJob{
		Status: mongodb.Status{
			Completed:false,
			Error:false,
			Started:true,
		},
		VideoId: bson.ObjectIdHex(request.VideoId),
		JobId: *res.Job.Id,
		PipelineId: *res.Job.PipelineId,
		InputKey: request.Job.InputKey,
		OutputKeyPrefix: *res.Job.OutputKeyPrefix,
		OutputKey: *res.Job.Output.Key,
		StartedAt: time.Now(),
		EndedAt: time.Time{},
	})
	if err != nil {
	    log.Warn(err)
	}
	log.Infof("[DB] created TranscodingJob '%s'", transcodeJob.ID)

	if jobStatus == "Submitted" {
		status.Started = true
	}
	if jobStatus == "Error" {
		status.Error = true
	}

	response.Job.Status = status

	log.Infof("[ElasticTranscoder] CREATED job '%s' on pipeline '%s'", response.Job.JobId, response.Job.PipelineId)

	return nil
}

func (t *transcodeService) awaitSQSEvent() {
	handler := NewTranscodeHandler(t.transcodingCompletedPublisher, t.transcodingFailedPublisher, t.transcodeRepository)
	for msg := range t.sqsConsumer.ElasticTranscoderChannel {
		// COMPLETED
		if msg.State == amazon.TranscodeStatusCompleted {
			err := handler.OnCompleted(msg)
			if err != nil {
				log.Info(err)
				continue
			}
		}

		// WARNING
		if msg.State == amazon.TranscodeStatusWarning {
			err := handler.OnWarning(msg)
			if err != nil {
				log.Info(err)
				continue
			}
		}

		// ERROR
		if msg.State == amazon.TranscodeStatusError {
			err := handler.OnError(msg)
			if err != nil {
				log.Info(err)
				continue
			}
		}
	}
}

func (t *transcodeService) awaitVideoCreatedEvent() {
	for videoCreated := range t.videoCreatedChan {
		req := &educonn_transcode.CreateJobRequest{
			Job: &educonn_transcode.TranscodeDetails{
				PipelineId: config.AwsTranscodePipelineId,
				InputKey: videoCreated.Video.Storage.RawKey,
			},
			VideoId: videoCreated.Video.Id,
		}
		err := t.CreateJob(context.Background(), req, &educonn_transcode.CreateJobResponse{})
		if err != nil {
		    log.Error(err)
		}
	}
}
