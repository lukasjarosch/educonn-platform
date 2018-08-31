package main

import (
	"github.com/micro/go-micro"
	"os"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
	"github.com/lukasjarosch/educonn-platform/api/user/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"
	pbUser"github.com/lukasjarosch/educonn-platform/user/proto"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	api "github.com/lukasjarosch/educonn-platform/api/user/internal/service"
	_ "github.com/joho/godotenv/autoload"
	"time"
	"github.com/lukasjarosch/educonn-platform/api/user/internal/middleware"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	opentrace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/lukasjarosch/educonn-platform/api/user/internal/platform/errors"
)

func main() {
	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	tracer, err := initTracing()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to init tracing")
	}

	// setup service
	service := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),

		micro.WrapHandler(middleware.AuthHandlerWrapper),

		micro.WrapHandler(opentrace.NewHandlerWrapper(tracer)),
		micro.WrapSubscriber(opentrace.NewSubscriberWrapper(tracer)),
		micro.WrapCall(opentrace.NewCallWrapper(tracer)),
		micro.WrapClient(opentrace.NewClientWrapper(tracer)),

		micro.WrapHandler(middleware.LogHandlerWrapper),
	)
	service.Init()

	// jwt handler
	jwtService, err := jwt_handler.NewJwtTokenHandler(config.PublicKeyPath, "")
	if err != nil {
	    log.Fatal().Err(err).Str("AUTH_PUBLIC_KEY_PATH", config.PublicKeyPath).Msg("unable to create JwtTokenHandler")
	}

	// rpc clients
	user := pbUser.NewUserClient("educonn.srv.user", service.Client())
	video := pbVideo.NewVideoClient("educonn.srv.video", service.Client())

	// api handler
	micro.RegisterHandler(service.Server(), api.NewUserApi(user, video, jwtService))

	if err := service.Run(); err != nil {
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
	opentracing.SetGlobalTracer(tracer)

	return tracer, nil
}
