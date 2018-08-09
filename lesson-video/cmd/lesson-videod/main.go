package main

import (
	"os"
	"github.com/rs/zerolog/log"
	"github.com/micro/go-micro"
	"github.com/lukasjarosch/educonn-platform/lesson-video/internal/platform/config"
	"time"
	"github.com/rs/zerolog"
)

func main() {

	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Init service
	service := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()


	// register handler

	// fire...
	if err := service.Run(); err != nil {
		panic(err)
	}
}
