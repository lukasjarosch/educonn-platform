package main

import (
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/broker"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/user/internal/service"
	pb "github.com/lukasjarosch/educonn-platform/user/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/rs/zerolog/log"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/joho/godotenv/autoload"
	"time"
	"os"
	"github.com/rs/zerolog"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/mongodb"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"
)

func main() {

	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

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
