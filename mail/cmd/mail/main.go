package main

import (
	"time"

	"github.com/lukasjarosch/educonn-platform/mail/internal/wrapper"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"

	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/lukasjarosch/educonn-platform/mail/internal/platform/broker"
	"github.com/lukasjarosch/educonn-platform/mail/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/mail/internal/platform/mail"
	"github.com/lukasjarosch/educonn-platform/mail/internal/service"
	pbMail "github.com/lukasjarosch/educonn-platform/mail/proto"
	pbUser "github.com/lukasjarosch/educonn-platform/user/proto"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	"github.com/micro/go-micro/server"
	"github.com/rs/zerolog"
)

func main() {
	InitTracer(config.ZipkinConnectionString, "9411", config.ServiceName)

	// setup micro service
	svc := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	svc.Init(
		micro.WrapHandler(wrapper.NewTraceHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapCall(wrapper.NewTraceCallWrapper(opentracing.GlobalTracer())),
		micro.WrapSubscriber(wrapper.NewTraceSubscriberWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(wrapper.RequestIdWrapper),
		micro.WrapHandler(wrapper.NewLogWrapper),
	)

	InitLogging(svc.Server().Options().Id)

	// setup the consumer
	userCreatedChannel := make(chan *pbUser.UserCreatedEvent)
	userCreatedSubscriber := broker.NewUserCreatedSubscriber(userCreatedChannel)
	userDeletedChannel := make(chan *pbUser.UserDeletedEvent)
	userDeletedSubscriber := broker.NewUserDeletedSubscriber(userDeletedChannel)
	videoProcessedChannel := make(chan *pbVideo.VideoProcessedEvent)
	videoProcessedSubscriber := broker.NewVideoProcessedSubscriber(videoProcessedChannel)

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
func InitLogging(instanceId string) {
	log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	log.Logger = log.Logger.With().Str("instance_id", instanceId).Logger()
}

func InitTracer(zipkinURL string, hostPort string, serviceName string) {
	log.Debug().Msg("Initialize tracing")
	collector, err := zipkin.NewHTTPCollector(zipkinURL)
	if err != nil {
		log.Error().Msgf("unable to create Zipkin HTTP collector: %v", err)
		return
	}
	tracer, err := zipkin.NewTracer(
		zipkin.NewRecorder(collector, true, hostPort, serviceName),
	)
	if err != nil {
		log.Error().Msgf("unable to create Zipkin tracer: %v", err)
		return
	}
	opentracing.InitGlobalTracer(tracer)
	return
}
