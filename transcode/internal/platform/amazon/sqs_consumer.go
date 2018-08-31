package amazon

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/evalphobia/aws-sdk-go-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-wrapper/sqs"
	"context"
	"github.com/rs/zerolog/log"
)

const TranscodeStatusCompleted = "COMPLETED"
const TranscodeStatusWarning = "WARNING"
const TranscodeStatusError = "ERROR"

type SQSTranscodeEventConsumer struct {
	sqs                      *sqs.SQS
	queue                    *sqs.Queue
	ElasticTranscoderChannel chan *ElasticTranscoderMessage
}

// ElasticTranscoderMessage
type ElasticTranscoderMessage struct {
	State string `json:"state"`
	ErrorCode int64 `json:"errorCode"`
	MessageDetails string `json:"messageDetails"`
	Version string `json:"version"`
	JobId string `json:"jobId"`
	PipelineId string `json:"pipelineId"`
	Input struct {
		Key string `json:"key"`
	}
	InputCount int64 `json:"inputCount"`
	OutputKeyPrefix string `json:"outputKeyPrefix"`
	Outputs []struct {
		Id string `json:"id"`
		PresetId string `json:"presetId"`
		Key string `json:"key"`
		ThumbnailPattern string `json:"thumbnailPattern"`
		Rotate string `json:"rotate"`
		Status string `json:"status"`
		StatusDetail string `json:"statusDetail"`
		ErrorCode int64 `json:"errorCode"`
	}

}

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *ElasticTranscoderMessage) error {
	log.Error().Msg("Handle")
	return nil
}

func Handler(ctx context.Context, msg *ElasticTranscoderMessage) error {
	log.Error().Msg("Handler")
	return nil
}

// ElasticTranscodeEventHandler defines the interface for all Elastic Transcode handlers
type ElasticTranscodeEventHandler interface {
	OnCompleted(message *ElasticTranscoderMessage) (error)
	OnWarning(message *ElasticTranscoderMessage) (error)
	OnError(message *ElasticTranscoderMessage) (error)
}

// NewSQSTranscodeEventConsumer creates a new consumer object
func NewSQSTranscodeEventConsumer(videoUploadedChannel chan *ElasticTranscoderMessage, accessKey string, secretKey string, region string, queueName string) (*SQSTranscodeEventConsumer, error) {
	svc, err := sqs.New(config.Config{
		AccessKey: accessKey,
		SecretKey: secretKey,
		Region:    region,
	})
	if err != nil {
		return nil, err
	}

	queue, err := svc.GetQueue(queueName)
	if err != nil {
		return nil, err
	}

	return &SQSTranscodeEventConsumer{
		sqs:                      svc,
		queue:                    queue,
		ElasticTranscoderChannel: videoUploadedChannel,
	}, nil
}

// FetchMessage tries to retrieve one message from SQS. If a message exists, it is unmarshalled and returned as S3Event
func (s *SQSTranscodeEventConsumer) FetchMessage() (map[string]interface{}, error) {
	message, err := s.queue.FetchOne()
	if err != nil {
		return nil, err
	}

	if message == nil {
		return nil, nil // this is no error
	}

	// Unmarshal message body
	var data map[string]interface{}
	err = json.Unmarshal([]byte(message.Body()), &data)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error unmarshalling SQS message: %v", err))
	}

	err = s.queue.DeleteMessage(message)

	return data, err
}

// ExtractElasticTranscoderMessage tries to unmarshal a map[string]interface{} into the ElasticTranscoderMessage type
func (s *SQSTranscodeEventConsumer) ExtractElasticTranscoderMessage(message map[string]interface{}) (*ElasticTranscoderMessage, error) {
	var transcodeEvent ElasticTranscoderMessage
	err := json.Unmarshal([]byte(message["Message"].(string)), &transcodeEvent)
	if err != nil {
		return nil, err
	}
	return &transcodeEvent, nil
}

// Consume is the actual polling method
func (s *SQSTranscodeEventConsumer) Consume() error {
	for {
		message, err := s.FetchMessage()
		if err != nil {
			continue
		}
		if message == nil {
			continue
		}

		transcodeMessage, err := s.ExtractElasticTranscoderMessage(message)
		if err != nil {
		    return err
		}

		s.ElasticTranscoderChannel <- transcodeMessage

	}
}
