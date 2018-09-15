package main

import (
	"os"
	"time"

	"github.com/micro/go-micro"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/lukasjarosch/educonn-platform/user-api/internal/middleware"
	"github.com/lukasjarosch/educonn-platform/user-api/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"

	pbUser"github.com/lukasjarosch/educonn-platform/user/proto"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	api "github.com/lukasjarosch/educonn-platform/user-api/internal/service"
)

func main() {
	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// setup service
	service := micro.NewService(
		micro.Name(config.ServiceName),
		micro.WrapHandler(middleware.TraceHandlerWrapper),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	// jwt handler
	jwtService, err := jwt_handler.NewJwtTokenHandler(config.PublicKeyPath, "")
	if err != nil {
	    log.Fatal().Err(err).Str("AUTH_PUBLIC_KEY_PATH", config.PublicKeyPath).Msg("unable to create JwtTokenHandler")
	}

	// rpc clients
	user := pbUser.NewUserClient("educonn.srv.user", service.Client())
	video := pbVideo.NewVideoClient("educonn.srv.video", service.Client())

	// api handler
	micro.RegisterHandler(service.Server(), api.NewUserApi(user, video, jwtService))

	if err := service.Run(); err != nil {
		panic(err)
	}
}
