package main

import (
	"time"

	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/service"
	pbTranscode "github.com/lukasjarosch/educonn-platform/transcode/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/amazon"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/broker"
	"github.com/micro/go-micro/server"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/mongodb"
	"os"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	opentrace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/middleware"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/errors"
	"github.com/opentracing/opentracing-go"
)

func main() {
	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	tracer, err := initTracing()
	if err != nil {
	    log.Fatal().Err(err).Msg("unable to init tracing")
	}

	// setup micro service
	svc := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),


		micro.WrapHandler(opentrace.NewHandlerWrapper(tracer)),
		micro.WrapSubscriber(opentrace.NewSubscriberWrapper(tracer)),
		micro.WrapCall(opentrace.NewCallWrapper(tracer)),
		micro.WrapClient(opentrace.NewClientWrapper(tracer)),

		micro.WrapHandler(middleware.LogHandlerWrapper),
		micro.WrapSubscriber(middleware.LogSubscriberWrapper),
	)
	svc.Init()

	// setup consumer
	videoCreatedChannel := make(chan broker.VideoCreatedEvent)
	videoCreatedSubscriber := broker.NewVideoCreatedSubscriber(videoCreatedChannel)

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
		log.Fatal().Interface("error", err).Str("queue", config.AwsSqsVideoQueueName).Msg("unable to connect to SQS")
	}
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

// initialize zipkin and opentracing
func initTracing() (opentracing.Tracer, error) {
	collector, err := zipkin.NewHTTPCollector(config.ZipkinCollectorUrl)
	if err != nil {
		return nil, errors.Error("unable to create zipkin collector")
	}
	tracer, err := zipkin.NewTracer(
		zipkin.NewRecorder(collector, true, "9411", config.ServiceName),
	)
	if err != nil {
		return nil, errors.Error("unable to create new zipkin tracer")
	}
	opentracing.SetGlobalTracer(tracer)

	return tracer, nil
}
