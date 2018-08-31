package main

import (
	"os"
	"time"

	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/mongodb"
	lesson "github.com/lukasjarosch/educonn-platform/lesson/internal/service"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	"github.com/micro/go-micro"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/broker"
)

func main() {
	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	// Init service
	service := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	// setup rabbitmq
	rabbitBroker := service.Server().Options().Broker
	if err := rabbitBroker.Init(rabbitmq.Exchange("educonn")); err != nil {
		log.Print("Broker Init error: %v", err)
		log.Fatal().Interface("error", err).Msg("broker Init error")
	}
	if err := rabbitBroker.Connect(); err != nil {
		log.Fatal().Interface("error", err).Msg("broker Connect error")
	}
	micro.Broker(rabbitBroker)

	// setup repositories
	mgoSession, err := mongodb.Dial(config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName)
	if err != nil {
		log.Fatal().Str("host", config.DbHost).Str("database", config.DbName).Interface("error", err).Msg("unable to connect to database")
	}
	lessonRepo := mongodb.NewLessonRepository(mgoSession)
	videoLessonRepo := mongodb.NewVideoLessonRepository(mgoSession)

	// rpc clients
	videoClient := pbVideo.NewVideoClient("educonn.srv.video", service.Client())

	// VideoLessonService handler
	videoLessonService := lesson.NewVideoLessonService(videoClient, videoLessonRepo)

	// publisher
	lessonCreatedPublisher := micro.NewPublisher(broker.LessonCreatedTopic, service.Client())
	lessonEventPublisher := broker.NewLessonEventPublisher(lessonCreatedPublisher)

	// register handler
	micro.RegisterHandler(service.Server(), videoLessonService)
	micro.RegisterHandler(service.Server(), lesson.NewLessonService(
		videoLessonService,
		lessonRepo,
		lessonEventPublisher,
	))

	// fire...
	if err := service.Run(); err != nil {
		panic(err)
	}
}
