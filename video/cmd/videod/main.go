package main

import (
	"github.com/micro/go-micro"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/config"
	"time"
	"github.com/micro/go-plugins/broker/rabbitmq"
	log "github.com/sirupsen/logrus"
)

func main() {

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
	if err := rabbitBroker.Init(rabbitmq.Exchange("educonn")); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := rabbitBroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	micro.Broker(rabbitBroker)

	if err := svc.Run(); err != nil {
		panic(err)
	}
}