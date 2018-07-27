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

type SqsS3EventConsumer struct {
	sqs                  *sqs.SQS
	queue                *sqs.Queue
	VideoUploadedChannel chan *events.S3EventRecord
}

// NewSqsS3EventConsumer creates a new consumer object
func NewSqsS3EventConsumer(videoUploadedChannel chan *events.S3EventRecord, accessKey string, secretKey string, region string, queueName string) (*SqsS3EventConsumer, error) {
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

	return &SqsS3EventConsumer{
		sqs:                  svc,
		queue:                queue,
		VideoUploadedChannel: videoUploadedChannel,
	}, nil
}

// FetchMessage tries to retrieve one message from SQS. If a message exists, it is unmarshalled and returned as S3Event
func (s *SqsS3EventConsumer) FetchMessage() (*events.S3Event, error) {
	messages, err := s.queue.Fetch(1)
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, nil
	}

	log.Info(messages)

	// unmarshal
	message := messages[0]
	var body events.S3Event
	err = json.Unmarshal([]byte(message.Body()), &body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error unmarshalling S3Event: %v", err))
	}

	return &body, nil
}

func (s *SqsS3EventConsumer) Consume() error {
	for {
		messages, err := s.FetchMessage()
		if err != nil {
		    log.Warn(err)
		    continue
		}
		if messages == nil {
			continue
		}
		if len(messages.Records) > 0 {
			message := messages.Records[0]
			s.VideoUploadedChannel <- &message
			log.Infof("[sub] received event from SQS '%s'", message.EventName)
		}
	}
}
