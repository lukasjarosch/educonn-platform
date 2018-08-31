package main

import (
	"github.com/rs/zerolog/log"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"

	"github.com/lukasjarosch/educonn-platform/mail/internal/platform/broker"
	"github.com/lukasjarosch/educonn-platform/mail/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/mail/internal/platform/mail"
	"github.com/lukasjarosch/educonn-platform/mail/internal/service"
	pbMail "github.com/lukasjarosch/educonn-platform/mail/proto"
	pbUser "github.com/lukasjarosch/educonn-platform/user/proto"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	"github.com/micro/go-micro/server"
	_ "github.com/joho/godotenv/autoload"
	"fmt"
	"os"
	"github.com/rs/zerolog"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"github.com/opentracing/opentracing-go"
	opentrace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/lukasjarosch/educonn-platform/mail/internal/middleware"
	"github.com/lukasjarosch/educonn-platform/mail/internal/platform/errors"
)

func main() {

	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	tracer, err := initTracing()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to init tracer")
	}

	// setup the consumer
	userCreatedChannel := make(chan broker.UserCreatedEvent)
	userCreatedSubscriber := broker.NewUserCreatedSubscriber(userCreatedChannel)
	userDeletedChannel := make(chan broker.UserDeletedEvent)
	userDeletedSubscriber := broker.NewUserDeletedSubscriber(userDeletedChannel)
	videoProcessedChannel := make(chan broker.VideoProcessedEvent)
	videoProcessedSubscriber := broker.NewVideoProcessedSubscriber(videoProcessedChannel)

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

	// setup rabbitmq
	rabbitBroker := svc.Server().Options().Broker
	if err := rabbitBroker.Init(rabbitmq.Exchange(config.ExchangeName)); err != nil {
		log.Fatal().Interface("error", err).Msg("broker init error")
	}
	if err := rabbitBroker.Connect(); err != nil {
		log.Fatal().Interface("error", err).Msg("broker connect error")
	}
	micro.Broker(rabbitBroker)

	// setup mail
	mailer, err := mail.NewSmtpMail(
		config.SmtpHostname,
		config.SmtpPort,
		config.SmtpUsername,
		config.SmtpPassword,
	)
	if err != nil {
		log.Fatal().Interface("error", err).Msg("unable to setup SmtpMail")
	}

	// UserCreatedSubscriber
	err = micro.RegisterSubscriber(
		broker.UserCreatedTopic,
		svc.Server(),
		userCreatedSubscriber,
		server.SubscriberQueue(broker.UserCreatedQueue),
	)
	if err != nil {
		panic(err)
	}
	log.Info().Msg(fmt.Sprintf("subscribed to %s", broker.UserCreatedTopic))

	// UserDeletedSubscriber
	err = micro.RegisterSubscriber(
		broker.UserDeletedTopic,
		svc.Server(),
		userDeletedSubscriber,
		server.SubscriberQueue(broker.UserDeletedQueue),
	)
	if err != nil {
		panic(err)
	}
	log.Info().Msg(fmt.Sprintf("subscribed to %s", broker.UserDeletedTopic))

	// VideoProcessedSubscriber
	err = micro.RegisterSubscriber(
		broker.VideoProcessedTopic,
		svc.Server(),
		videoProcessedSubscriber,
		server.SubscriberQueue(broker.VideoProcessedQueue),
	)
	if err != nil {
		panic(err)
	}
	log.Info().Msg(fmt.Sprintf("subscribed to %s", broker.VideoProcessedTopic))

	// rpc clients
	userClient := pbUser.NewUserClient("educonn.srv.user", svc.Client())
	videoClient := pbVideo.NewVideoClient("educonn.srv.video", svc.Client())

	// service handler
	pbMail.RegisterEmailHandler(
		svc.Server(),
		service.NewMailService(
			userCreatedChannel,
			userDeletedChannel,
			videoProcessedChannel,
			mailer,
			userClient,
			videoClient,
		))

	// fire
	if err := svc.Run(); err != nil {
		panic(err)
	}
}

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
	opentracing.InitGlobalTracer(tracer)

	return tracer, nil
}
