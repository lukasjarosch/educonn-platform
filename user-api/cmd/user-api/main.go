package main

import (
	"os"
	"time"

	"github.com/micro/go-micro"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	opentracing "github.com/opentracing/opentracing-go"
	ot "github.com/micro/go-plugins/wrapper/trace/opentracing"

	"github.com/lukasjarosch/educonn-platform/user-api/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"

	api "github.com/lukasjarosch/educonn-platform/user-api/internal/service"
	pbUser "github.com/lukasjarosch/educonn-platform/user/proto"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
)

func main() {
	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	InitTracer("http://localhost:9411/api/v1/spans", "9411", config.ServiceName)

	// setup service
	service := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.WrapHandler(ot.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(ot.NewClientWrapper(opentracing.GlobalTracer())),
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
