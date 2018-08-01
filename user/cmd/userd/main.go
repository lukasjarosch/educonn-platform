package main

import (
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/broker"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/user/internal/service"
	"github.com/lukasjarosch/educonn-platform/user/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/rs/zerolog/log"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/joho/godotenv/autoload"
	"time"
	"os"
	"github.com/rs/zerolog"
)

func main() {

	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// TODO: mysqlrepo
	repo := false

	svc := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	svc.Init()

	// setup rabbitmq
	rabbitBroker := svc.Server().Options().Broker
	if err := rabbitBroker.Init(rabbitmq.Exchange("educonn")); err != nil {
		log.Print("Broker Init error: %v", err)
	}
	if err := rabbitBroker.Connect(); err != nil {
		log.Print("Broker Connect error: %v", err)
	}
	micro.Broker(rabbitBroker)

	userCreatedPublisher := micro.NewPublisher(broker.UserCreatedTopic, svc.Client())
	userDeletedPublisher := micro.NewPublisher(broker.UserDeletedTopic, svc.Client())

	educonn_user.RegisterUserHandler(
		svc.Server(),
		service.NewUserService(
			repo,
			broker.NewEventPublisher(userCreatedPublisher),
			broker.NewEventPublisher(userDeletedPublisher),
		),
	)

	if err := svc.Run(); err != nil {
		panic(err)
	}
}
