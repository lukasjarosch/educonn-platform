package main

import (
	"github.com/rs/zerolog/log"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"

	"github.com/lukasjarosch/educonn-platform/mail/internal/platform/broker"
	"github.com/lukasjarosch/educonn-platform/mail/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/mail/internal/platform/mail"
	"github.com/lukasjarosch/educonn-platform/mail/internal/service"
	"github.com/lukasjarosch/educonn-platform/mail/proto"
	"github.com/lukasjarosch/educonn-platform/user/proto"
	"github.com/micro/go-micro/server"
	_ "github.com/joho/godotenv/autoload"
	"fmt"
	"os"
	"github.com/rs/zerolog"
)

func main() {

	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// setup the consumer
	userCreatedChannel := make(chan *educonn_user.UserCreatedEvent)
	userDeletedChannel := make(chan *educonn_user.UserDeletedEvent)
	userCreatedSubscriber := broker.NewUserCreatedSubscriber(userCreatedChannel)
	userDeletedSubscriber := broker.NewUserDeletedSubscriber(userDeletedChannel)

	// setup micro service
	svc := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	svc.Init()

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

	// service handler
	educonn_mail.RegisterEmailHandler(
		svc.Server(),
		service.NewMailService(
			userCreatedChannel,
			userDeletedChannel,
			mailer,
		))

	// fire
	if err := svc.Run(); err != nil {
		panic(err)
	}
}
