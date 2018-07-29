package main

import (
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"

	"github.com/lukasjarosch/educonn-master-thesis/mail/internal/platform/broker"
	"github.com/lukasjarosch/educonn-master-thesis/mail/internal/platform/config"
	"github.com/lukasjarosch/educonn-master-thesis/mail/internal/platform/mail"
	"github.com/lukasjarosch/educonn-master-thesis/mail/internal/service"
	"github.com/lukasjarosch/educonn-master-thesis/mail/proto"
	"github.com/lukasjarosch/educonn-master-thesis/user/proto"
	"github.com/micro/go-micro/server"
)

func main() {

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
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := rabbitBroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
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
		log.Fatalf("Unable to setup STMP mailer: %v", err)
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
	log.Infof("Subscribed %s", broker.UserCreatedTopic)

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
	log.Infof("Subscribed %s", broker.UserDeletedTopic)

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
