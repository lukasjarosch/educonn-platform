package main

import (
	"github.com/micro/go-micro"
	"os"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
	api "github.com/lukasjarosch/educonn-platform/api/user/internal/service"
	"github.com/lukasjarosch/educonn-platform/user/proto"
	"github.com/lukasjarosch/educonn-platform/api/user/internal/middleware"
	"github.com/lukasjarosch/educonn-platform/api/user/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"
)

func main() {
	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	service := micro.NewService(
		micro.Name(config.ServiceName),
		micro.WrapHandler(middleware.TraceHandlerWrapper),
	)
	service.Init()

	jwtService, err := jwt_handler.NewJwtTokenHandler(config.PublicKeyPath, "")
	if err != nil {
	    log.Fatal().Interface("error", err).Msg("unable to create JwtTokenHandler")
	}

	user := proto.NewUserClient("educonn.srv.user", service.Client())

	micro.RegisterHandler(service.Server(), api.NewUserApi(user, jwtService))

	if err := service.Run(); err != nil {
		panic(err)
	}
}
