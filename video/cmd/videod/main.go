package main

import (
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/broker"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/config"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/service"
	"github.com/lukasjarosch/educonn-master-thesis/video/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"
	log "github.com/sirupsen/logrus"
	"time"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/amazon"
	"context"
)

func main() {

	// setup micro service
	svc := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	svc.Init()

	// setup rabbitmq
	rabbitBroker := svc.Server().Options().Broker
	if err := rabbitBroker.Init(rabbitmq.Exchange("educonn")); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := rabbitBroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	micro.Broker(rabbitBroker)

	// setup SQS
	elasticTranscoderChan := make(chan *amazon.ElasticTranscoderMessage)
	sqsConsumer, err := amazon.NewSQSTranscodeEventConsumer(elasticTranscoderChan, config.AwsAccessKey, config.AwsSecretKey, config.AwsRegion, config.AwsSqsVideoQueueName)
	if err != nil {
	    log.Fatal(err)
	}
	sqsContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup S3
	bucket, err := amazon.NewS3Bucket(config.AwsS3VideoBucket, config.AwsRegion, config.AwsAccessKey, config.AwsSecretKey)
	log.Infof("[S3] attached to bucket: %s", bucket.Bucket)


	videoCreatedPublisher := micro.NewPublisher(broker.VideoCreatedTopic, svc.Client())

	// Attach handler
	educonn_video.RegisterVideoHandler(
		svc.Server(),
		service.NewVideoService(
			broker.NewEventPublisher(videoCreatedPublisher),
			sqsConsumer,
			sqsContext,
			bucket,
		),
	)

	if err := svc.Run(); err != nil {
		panic(err)
	}
}
