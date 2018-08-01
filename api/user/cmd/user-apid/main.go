package main

import (
	"github.com/micro/go-micro"
	service2 "github.com/lukasjarosch/educonn-platform/api/user/internal/service"
)

func main() {
	service := micro.NewService(
		micro.Name("educonn.api.user"),
	)
	service.Init()

	micro.RegisterHandler(service.Server(), service2.NewUserApiService())

	service.Run()
}
