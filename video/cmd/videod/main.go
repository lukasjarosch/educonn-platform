package main

import (
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/broker"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/video/internal/service"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/rs/zerolog/log"
	"time"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/amazon"
	pbTranscode "github.com/lukasjarosch/educonn-platform/transcode/proto"
	"github.com/micro/go-micro/server"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/mongodb"
	"os"
	"github.com/rs/zerolog"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"
)

func main() {

	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	transcodeCompletedChannel := make(chan *pbTranscode.TranscodingCompletedEvent)
	transcodeCompletedSubscriber := broker.NewTranscodeCompletedSubscriber(transcodeCompletedChannel)
	transcodeFailedChannel := make(chan *pbTranscode.TranscodingFailedEvent)
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
	videoProcessedPublisher := micro.NewPublisher(broker.VideoProcessedTopic, svc.Client())

	// JWT handler (without private key, we only want to validate)
	jwtHandler, err := jwt_handler.NewJwtTokenHandler(config.AuthPublicKeyPath, "")
	if err != nil {
	    log.Fatal().Err(err).Msg("unable to create new JwtTokenHandler")
	}

	// Create subscribers
	err = micro.RegisterSubscriber(
		broker.TranscodeCompletedTopic,
		svc.Server(),
		transcodeCompletedSubscriber,
		server.SubscriberQueue(broker.TranscodeCompletedQueue),
	)
	if err != nil {
	    log.Fatal().Err(err)
	}
	err = micro.RegisterSubscriber(
		broker.TranscodeFailedTopic,
		svc.Server(),
		transcodeFailedSubscriber,
		server.SubscriberQueue(broker.TranscodeFailedQueue),
	)
	if err != nil {
		log.Fatal().Err(err)
	}

	// Attach handler
	pbVideo.RegisterVideoHandler(
		svc.Server(),
		service.NewVideoService(
			broker.NewEventPublisher(videoCreatedPublisher, videoProcessedPublisher),
			bucket,
			transcodeCompletedChannel,
			transcodeFailedChannel,
			videoRepository,
			jwtHandler,
		),
	)

	if err := svc.Run(); err != nil {
		panic(err)
	}
}
