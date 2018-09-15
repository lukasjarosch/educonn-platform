package main

import (
	"os"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/micro/go-micro"
	"github.com/lukasjarosch/educonn-platform/lesson-api/internal/platform/config"
	pbLessonApi "github.com/lukasjarosch/educonn-platform/lesson-api/proto"
	"time"
	"github.com/lukasjarosch/educonn-platform/lesson-api/internal/service"
)

func main() {
	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// setup micro service
	svc := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	svc.Init()

	// service handler
	pbLessonApi.RegisterLessonApiHandler(
		svc.Server(),
		service.NewLessonApiService(),
	)

	// fire
	if err := svc.Run(); err != nil {
		panic(err)
	}
}
