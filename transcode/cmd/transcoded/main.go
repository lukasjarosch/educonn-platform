package main

import (
	"context"
	"time"

	"github.com/lukasjarosch/educonn-master-thesis/transcode/internal/platform/config"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/internal/service"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/prometheus/common/log"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/internal/platform/amazon"
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
	log.Infof("[SQS] attached to queue '%s'", config.AwsSqsVideoQueueName)

	// setup Elastic Transcoder
	transcoder, err := amazon.NewElasticTranscoderClient(config.AwsAccessKey, config.AwsSecretKey, config.AwsRegion)
	if err != nil {
		log.Infof("Unable to create ElasticTranscoder client: %v", err)
	}
	log.Infof("[ElasticTranscoder] attached")

	// register service handler
	educonn_transcode.RegisterTranscodeHandler(
		svc.Server(),
		service.NewTranscodeService(
			sqsConsumer,
			sqsContext,
			transcoder,
		),
	)

	// fire..
	if err := svc.Run(); err != nil {
		panic(err)
	}
}
