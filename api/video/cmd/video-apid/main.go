package main

import (
	"os"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
	"github.com/micro/go-micro"
	"github.com/lukasjarosch/educonn-platform/api/video/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	api "github.com/lukasjarosch/educonn-platform/api/video/internal/service"
)

func main() {

	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Init service
	service := micro.NewService(
		micro.Name(config.ServiceName),
	)
	service.Init()

	// create JWT handler
	jwtService, err := jwt_handler.NewJwtTokenHandler(config.PublicKeyPath, "")
	if err != nil {
		log.Fatal().Interface("error", err).Msg("unable to create JwtTokenHandler")
	}

	// create rpc clients
	video := pbVideo.NewVideoClient("educonn.srv.video", service.Client())

	// register handler
	micro.RegisterHandler(service.Server(), api.NewVideoApi(
		video,
		jwtService,
	))

	// fire...
	if err := service.Run(); err != nil {
		panic(err)
	}
}
