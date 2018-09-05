package main

import (
	"context"
	"time"

	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/service"
	pbTranscode "github.com/lukasjarosch/educonn-platform/transcode/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/amazon"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/broker"
	"github.com/micro/go-micro/server"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/mongodb"
	"os"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
)

func main() {
	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// setup consumer
	videoCreatedChannel := make(chan *pbVideo.VideoCreatedEvent)
	videoCreatedSubscriber := broker.NewVideoCreatedSubscriber(videoCreatedChannel)

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
		log.Print("Broker Init error: %v", err)
	}
	if err := rabbitBroker.Connect(); err != nil {
		log.Print("Broker Connect error: %v", err)
	}
	micro.Broker(rabbitBroker)


	// setup mongodb
	transcodeRepository,err := mongodb.NewTranscodeRepository(config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName)
	if err != nil {
	    log.Fatal().Interface("error", err).Msg("unable to connect to database")
	}

	// setup SQS
	elasticTranscoderChan := make(chan *amazon.ElasticTranscoderMessage)
	sqsConsumer, err := amazon.NewSQSTranscodeEventConsumer(elasticTranscoderChan, config.AwsAccessKey, config.AwsSecretKey, config.AwsRegion, config.AwsSqsVideoQueueName)
	if err != nil {
		log.Fatal().Err( err).Str("queue", config.AwsSqsVideoQueueName).Msg("unable to connect to SQS")
	}
	sqsContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Info().Str("queue", config.AwsSqsVideoQueueName).Msg("attached to SQS queue")

	// setup Elastic Transcoder
	transcoder, err := amazon.NewElasticTranscoderClient(config.AwsAccessKey, config.AwsSecretKey, config.AwsRegion)
	if err != nil {
		log.Fatal().Interface("error", err).Msg("unable to create ElasticTranscoderClient")
	}
	log.Info().Msg("attached ElasticTranscoder")

	// register video.events.created subscriber
	err = micro.RegisterSubscriber(
		broker.VideoCreatedTopic,
		svc.Server(),
		videoCreatedSubscriber,
		server.SubscriberQueue(broker.VideoCreatedQueue),
	)

	// register publishers
	transcodingCompletedPublisher := micro.NewPublisher(broker.VideoTranscodingCompleted, svc.Client())
	transcodingFailedPublisher := micro.NewPublisher(broker.VideoTranscodingFailed, svc.Client())

	// register service handler
	pbTranscode.RegisterTranscodeHandler(
		svc.Server(),
		service.NewTranscodeService(
			sqsConsumer,
			sqsContext,
			transcoder,
			videoCreatedChannel,
			transcodingCompletedPublisher,
			transcodingFailedPublisher,
			transcodeRepository,
		),
	)

	// fire..
	if err := svc.Run(); err != nil {
		panic(err)
	}
}
