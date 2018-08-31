package main

import (
	"os"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
	"github.com/micro/go-micro"
	"github.com/lukasjarosch/educonn-platform/api/video/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	api "github.com/lukasjarosch/educonn-platform/api/video/internal/service"
	"time"
	"github.com/lukasjarosch/educonn-platform/api/video/internal/middleware"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	opentrace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/lukasjarosch/educonn-platform/api/video/internal/platform/errors"
)

func main() {

	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	tracer, err := initTracing()
	if err != nil {
	    log.Fatal().Err(err).Msg("unable to init tracing")
	}


	// Init service
	service := micro.NewService(
		micro.Name(config.ServiceName),
		micro.WrapHandler(middleware.AuthHandlerWrapper),

		micro.WrapHandler(opentrace.NewHandlerWrapper(tracer)),
		micro.WrapSubscriber(opentrace.NewSubscriberWrapper(tracer)),
		micro.WrapCall(opentrace.NewCallWrapper(tracer)),
		micro.WrapClient(opentrace.NewClientWrapper(tracer)),

		micro.WrapHandler(middleware.LogHandlerWrapper),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	// create JWT handler
	jwtService, err := jwt_handler.NewJwtTokenHandler(config.PublicKeyPath, "")
	if err != nil {
		log.Fatal().Interface("error", err).Msg("unable to create JwtTokenHandler")
	}

	// create rpc clients
	video := pbVideo.NewVideoClient("educonn.srv.video", service.Client())

	// register handler
	micro.RegisterHandler(service.Server(), api.NewVideoApi(
		video,
		jwtService,
	))

	// fire...
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