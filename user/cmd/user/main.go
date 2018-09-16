package main

import (
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/broker"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/mongodb"
	"github.com/lukasjarosch/educonn-platform/user/internal/service"
	"github.com/lukasjarosch/educonn-platform/user/internal/wrapper"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"
	pb "github.com/lukasjarosch/educonn-platform/user/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	ot "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	InitTracer("http://localhost:9411/api/v1/spans", "9411", config.ServiceName)

	svc := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	svc.Init(
		micro.WrapHandler(ot.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(ot.NewClientWrapper(opentracing.GlobalTracer())),
		micro.WrapCall(ot.NewCallWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(wrapper.RequestIdWrapper),
		micro.WrapHandler(wrapper.NewLogWrapper),
	)

	InitLogging(svc.Server().Options().Id)

	// setup rabbitmq
	rabbitBroker := svc.Server().Options().Broker
	if err := rabbitBroker.Init(rabbitmq.Exchange("educonn")); err != nil {
		log.Print("Broker Init error: %v", err)
		log.Fatal().Interface("error", err).Msg("broker Init error")
	}
	if err := rabbitBroker.Connect(); err != nil {
		log.Fatal().Interface("error", err).Msg("broker Connect error")
	}
	micro.Broker(rabbitBroker)

	// setup database
	repo, err := mongodb.NewUserRepository(config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName)
	if err != nil {
		log.Fatal().
			Str("host", config.DbHost).
			Str("database", config.DbName).
			Interface("error", err).
			Msg("unable to connect to database")
	}

	// Setup auth token service
	tokenService, err := jwt_handler.NewJwtTokenHandler(config.PublicKeyPath, config.PrivateKeyPath)
	if err != nil {
		log.Fatal().Interface("error", err).Msg("unable to create TokenService")
	}

	userCreatedPublisher := micro.NewPublisher(broker.UserCreatedTopic, svc.Client())
	userDeletedPublisher := micro.NewPublisher(broker.UserDeletedTopic, svc.Client())

	pb.RegisterUserHandler(
		svc.Server(),
		service.NewUserService(
			repo,
			broker.NewEventPublisher(userCreatedPublisher),
			broker.NewEventPublisher(userDeletedPublisher),
			tokenService,
		),
	)

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
