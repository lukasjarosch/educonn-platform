package service

import (
	"context"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/amazon"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/config"
	pbTranscode "github.com/lukasjarosch/educonn-platform/transcode/proto"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	"github.com/rs/zerolog/log"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/broker"
	"github.com/micro/go-micro"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/mongodb"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type transcodeService struct {
	sqsConsumer      *amazon.SQSTranscodeEventConsumer
	sqsContext       context.Context
	transcoderClient *amazon.ElasticTranscoderClient
	videoCreatedChan chan *pbVideo.VideoCreatedEvent
	transcodingCompletedPublisher micro.Publisher
	transcodingFailedPublisher micro.Publisher
	transcodeRepository *mongodb.TranscodeRepository
}

func NewTranscodeService(sqsConsumer *amazon.SQSTranscodeEventConsumer,
	sqsContext context.Context,
	transcoderClient *amazon.ElasticTranscoderClient,
	videoCreatedChan chan *pbVideo.VideoCreatedEvent,
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

	log.Info().Str("topic", broker.VideoCreatedTopic).Str("queue", broker.VideoCreatedQueue).Msg("start consuming VideoCreatedEvents")
	go svc.awaitVideoCreatedEvent()

	log.Info().Str("queue", config.AwsSqsVideoQueueName).Msg("start consuming from SQS queue")
	go svc.sqsConsumer.Consume()
	go svc.awaitSQSEvent()

	return svc
}

func (t *transcodeService) CreateJob(ctx context.Context, request *pbTranscode.CreateJobRequest, response *pbTranscode.CreateJobResponse) error {

	res, err := t.transcoderClient.CreateJob(request.Job.InputKey)
	if err != nil {
		log.Warn().Interface("error", err).Msg("unable to create new transcoding job")
		return err
	}

	response.Job = &pbTranscode.TranscodeDetails{
		JobId:           *res.Job.Id,
		InputKey:        request.Job.InputKey,
		PipelineId:      *res.Job.PipelineId,
		OutputKeyPrefix: *res.Job.OutputKeyPrefix,
		OutputKey:       *res.Job.Output.Key,
	}

	jobStatus := *res.Job.Status
	status := &pbTranscode.TranscodeStatus{
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
		log.Warn().Interface("error", err).Msg("failed to create transcoding job in database")
	}

	if jobStatus == "Submitted" {
		status.Started = true
	}
	if jobStatus == "Error" {
		status.Error = true
	}

	response.Job.Status = status

	log.Info().Str("pipeline", response.Job.PipelineId).Str("job", transcodeJob.JobId).Msg("created new ElasticTranscoder job")

	return nil
}

func (t *transcodeService) awaitSQSEvent() {
	handler := NewTranscodeHandler(t.transcodingCompletedPublisher, t.transcodingFailedPublisher, t.transcodeRepository)
	for msg := range t.sqsConsumer.ElasticTranscoderChannel {
		// COMPLETED
		if msg.State == amazon.TranscodeStatusCompleted {
			err := handler.OnCompleted(msg)
			if err != nil {
				log.Warn().Interface("error", err).Msg("SQS handler failed in OnCompleted")
				continue
			}
		}

		// WARNING
		if msg.State == amazon.TranscodeStatusWarning {
			err := handler.OnWarning(msg)
			if err != nil {
				log.Warn().Interface("error", err).Msg("SQS handler failed in OnWarning")
				continue
			}
		}

		// ERROR
		if msg.State == amazon.TranscodeStatusError {
			err := handler.OnError(msg)
			if err != nil {
				log.Warn().Interface("error", err).Msg("SQS handler failed in OnError")
				continue
			}
		}
	}
}

func (t *transcodeService) awaitVideoCreatedEvent() {
	for videoCreated := range t.videoCreatedChan {
		req := &pbTranscode.CreateJobRequest{
			Job: &pbTranscode.TranscodeDetails{
				PipelineId: config.AwsTranscodePipelineId,
				InputKey: videoCreated.Video.Storage.RawKey,
			},
			VideoId: videoCreated.Video.Id,
		}
		err := t.CreateJob(context.Background(), req, &pbTranscode.CreateJobResponse{})
		if err != nil {
			log.Warn().Interface("error", err).Msg("unable to call CreateJob")
		}
	}
}
