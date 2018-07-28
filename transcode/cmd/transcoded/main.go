package main

import (
	"github.com/micro/go-micro"
	"time"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/internal/platform/config"
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



	if err := svc.Run(); err != nil {
		panic(err)
	}
}
