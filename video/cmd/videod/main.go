package main

import (
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/broker"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/video/internal/service"
	"github.com/lukasjarosch/educonn-platform/video/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/rs/zerolog/log"
	"time"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/amazon"
	"github.com/lukasjarosch/educonn-platform/transcode/proto"
	"github.com/micro/go-micro/server"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/mongodb"
	"os"
	"github.com/rs/zerolog"
)

func main() {

	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	transcodeCompletedChannel := make(chan *educonn_transcode.TranscodingCompletedEvent)
	transcodeCompletedSubscriber := broker.NewTranscodeCompletedSubscriber(transcodeCompletedChannel)
	transcodeFailedChannel := make(chan *educonn_transcode.TranscodingFailedEvent)
	transcodeFailedSubscriber := broker.NewTranscodeFailedSubscriber(transcodeFailedChannel)

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
		log.Fatal().Interface("error", err).Msg("broker init error")
	}
	if err := rabbitBroker.Connect(); err != nil {
		log.Fatal().Interface("error", err).Msg("broker connect error")
	}
	micro.Broker(rabbitBroker)

	// Setup S3
	bucket, err := amazon.NewS3Bucket(config.AwsS3VideoBucket, config.AwsRegion, config.AwsAccessKey, config.AwsSecretKey)
	if err != nil {
		log.Warn().Interface("error", err)
	    return
	}
	log.Info().Str("bucket", bucket.Bucket).Msg( "attached to S3 bucket")

	// Create repository
	videoRepository, err := mongodb.NewVideoRepository(config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName)
	if err != nil {
	    log.Fatal().Interface("error", err).Msg("unable to connect to database")
	}
	log.Info().Str("host", config.DbHost).Str("db_name", config.DbName).Msg("connected to database")

	// Create publishers
	videoCreatedPublisher := micro.NewPublisher(broker.VideoCreatedTopic, svc.Client())

	// Create subscribers
	err = micro.RegisterSubscriber(
		broker.TranscodeCompletedTopic,
		svc.Server(),
		transcodeCompletedSubscriber,
		server.SubscriberQueue(broker.TranscodeCompletedQueue),
	)
	err = micro.RegisterSubscriber(
		broker.TranscodeFailedTopic,
		svc.Server(),
		transcodeFailedSubscriber,
		server.SubscriberQueue(broker.TranscodeFailedQueue),
	)

	// Attach handler
	educonn_video.RegisterVideoHandler(
		svc.Server(),
		service.NewVideoService(
			broker.NewEventPublisher(videoCreatedPublisher),
			bucket,
			transcodeCompletedChannel,
			transcodeFailedChannel,
			videoRepository,
		),
	)

	if err := svc.Run(); err != nil {
		panic(err)
	}
}
