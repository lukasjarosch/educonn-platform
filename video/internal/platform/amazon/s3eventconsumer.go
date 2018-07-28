package amazon

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/evalphobia/aws-sdk-go-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-wrapper/sqs"
	log "github.com/sirupsen/logrus"
)

const TranscodeStatusCompleted = "COMPLETED"
const TranscodeStatusWarning = "WARNING"
const TranscodeStatusError = "ERROR"

type SQSTranscodeEventConsumer struct {
	sqs                  *sqs.SQS
	queue                *sqs.Queue
	VideoUploadedChannel chan *events.S3EventRecord
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

// ElasticTranscodeEventHandler defines the interface for all Elastic Transcode handlers
type ElasticTranscodeEventHandler interface {
	OnCompleted(message *ElasticTranscoderMessage) (error)
	OnWarning(message *ElasticTranscoderMessage) (error)
	OnError(message *ElasticTranscoderMessage) (error)
}

// NewSQSTranscodeEventConsumer creates a new consumer object
func NewSQSTranscodeEventConsumer(videoUploadedChannel chan *events.S3EventRecord, accessKey string, secretKey string, region string, queueName string) (*SQSTranscodeEventConsumer, error) {
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
		sqs:                  svc,
		queue:                queue,
		VideoUploadedChannel: videoUploadedChannel,
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
func (s *SQSTranscodeEventConsumer) Consume(eventHandler ElasticTranscodeEventHandler) error {
	for {
		message, err := s.FetchMessage()
		if err != nil {
			log.Info(err)
			continue
		}
		if message == nil {
			continue
		}

		transcodeMessage, err := s.ExtractElasticTranscoderMessage(message)
		if err != nil {
		    return err
		}

		// TODO: These if's could be replaced by a cool callback map

		// COMPLETED
		if transcodeMessage.State == TranscodeStatusCompleted {
			err := eventHandler.OnCompleted(transcodeMessage)
			if err != nil {
				log.Info(err)
				continue
			}
		}

		// WARNING
		if transcodeMessage.State == TranscodeStatusWarning {
			err := eventHandler.OnWarning(transcodeMessage)
			if err != nil {
				log.Info(err)
				continue
			}
		}

		// ERROR
		if transcodeMessage.State == TranscodeStatusError {
			err := eventHandler.OnError(transcodeMessage)
			if err != nil {
				log.Info(err)
				continue
			}
		}
	}
}
