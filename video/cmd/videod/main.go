package main

import (
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/broker"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/video/internal/service"
	"github.com/lukasjarosch/educonn-platform/video/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/prometheus/common/log"
	"time"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/amazon"
	"github.com/lukasjarosch/educonn-platform/transcode/proto"
	"github.com/micro/go-micro/server"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/mongodb"
)

func main() {

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
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := rabbitBroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	micro.Broker(rabbitBroker)

	// Setup S3
	bucket, err := amazon.NewS3Bucket(config.AwsS3VideoBucket, config.AwsRegion, config.AwsAccessKey, config.AwsSecretKey)
	if err != nil {
	    log.Warn(err)
	    return
	}
	log.Infof("[S3] attached to bucket: %s", bucket.Bucket)

	// Create repository
	videoRepository, err := mongodb.NewVideoRepository(config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName)
	if err != nil {
	    log.Fatalf("[DB] unable to connect to database: %s", err)
	}
	log.Infof("[DB] connected to %s:%s/%s", config.DbHost, config.DbPort, config.DbName)

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
