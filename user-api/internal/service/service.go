package service

import (
	"context"
	"net/mail"

	"github.com/lukasjarosch/educonn-platform/user-api/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/user-api/internal/platform/errors"
	pb "github.com/lukasjarosch/educonn-platform/user-api/proto"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"
	pbUser "github.com/lukasjarosch/educonn-platform/user/proto"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	merr "github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
)

type User struct {
	userClient  pbUser.UserClient
	videoClient pbVideo.VideoClient
	jwtService  *jwt_handler.JwtTokenHandler
}

func NewUserApi(userClient pbUser.UserClient, videoClient pbVideo.VideoClient, jwtService *jwt_handler.JwtTokenHandler) *User {
	return &User{userClient: userClient, videoClient: videoClient, jwtService: jwtService}
}

func (u *User) Create(ctx context.Context, req *pb.CreateRequest, res *pb.CreateResponse) error {

	span := opentracing.SpanFromContext(ctx)
	span.SetOperationName("Create")

	if req.User == nil {
		return merr.BadRequest(config.ServiceName, "%s", "Request contains no data")
	}

	// validate request
	if req.User.FirstName == "" || req.User.LastName == "" {
		return merr.BadRequest(config.ServiceName, "%s", "Please specify your first and last name")
	}

	// validate email
	email := req.User.Email
	if email == "" {
		return merr.BadRequest(config.ServiceName, "%s", "Please specify an email address")
	}
	e, err := mail.ParseAddress(req.User.Email)
	if err != nil {
		log.Debug().Msg(err.Error())
		return merr.BadRequest(config.ServiceName, "%s", "Please provide a valid email address")
	}
	req.User.Email = e.Address // override the user input with parsed email

	// validate password
	pass := req.User.Password
	if pass == "" {
		return merr.BadRequest(config.ServiceName, "%s", "Please specify a password")
	}

	span.LogKV("event", "call userClient.Create")

	user, err := u.userClient.Create(ctx, req.User)
	if err != nil {
		log.Warn().Interface("error", err).Msg("unable to create user")
		span.LogKV("event", "userClient.Create failed")
		return merr.InternalServerError(config.ServiceName, "%s: %s", "Unable to create user", err.Error())
	}

	span.LogKV("event", "userClient.Create finished")

	res.User = user.User

	return nil
}

func (u *User) Delete(ctx context.Context, req *pb.DeleteRequest, res *pb.DeleteResponse) error {
	span := opentracing.SpanFromContext(ctx)
	span.SetOperationName("delete")

	md, _ := metadata.FromContext(ctx)

	// check for token
	token, err := u.jwtService.GetBearerToken(md)
	if err != nil {
		span.LogKV("event", "missing JWT token")
		return merr.Unauthorized(config.ServiceName, "%s", err.Error())
	}

	// validate jwt
	_, err = u.jwtService.Decode(token)
	if err != nil {
		log.Error().Interface("error", err).Msg(errors.InvalidJWTToken.Error())
		span.LogKV("event", "invalid jwt token")
		return merr.Unauthorized(config.ServiceName, "%s", errors.InvalidJWTToken)
	}

	span.LogKV("event", "call userClient.Delete")

	// delete user
	_, err = u.userClient.Delete(ctx, &pbUser.DeleteRequest{
		User: &pbUser.UserDetails{
			Id: req.UserId,
		},
	})
	if err != nil {
		return err
	}
	log.Debug().Str("user", req.UserId).Msg("deleted user from database")

	span.LogKV("event", "userClient.Delete finished")

	return nil
}

func (u *User) Login(ctx context.Context, req *pb.LoginRequest, res *pb.LoginResponse) error {
	span := opentracing.SpanFromContext(ctx)
	span.SetOperationName("login")

	if req.Email == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.EmailMissing.Error())
	}

	if req.Password == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.PasswordMissing.Error())
	}

	user := &pbUser.UserDetails{
		Email:    req.Email,
		Password: req.Password,
	}

	span.LogKV("event", "call userClient.Auth")

	// login
	response, err := u.userClient.Auth(ctx, user)
	if err != nil {
		log.Debug().Interface("error", err).Str("email", req.Email).Msgf("user login failed")
		span.LogKV("event", "auth failed")
		return err
	}
	span.LogKV("event", "auth suceeded")

	res.Token = response.Token

	return nil
}

func (u *User) Videos(ctx context.Context, req *pb.VideoRequest, res *pb.VideoResponse) error {
	md, _ := metadata.FromContext(ctx)

	span := opentracing.SpanFromContext(ctx)
	span.SetOperationName("get videos")

	// check for token
	token, err := u.jwtService.GetBearerToken(md)
	if err != nil {
		log.Error().Interface("err", err).Msg("asdf")
		return merr.Unauthorized(config.ServiceName, "%s", err.Error())
	}

	// validate jwt
	claims, err := u.jwtService.Decode(token)
	if err != nil {
		log.Error().Interface("error", err).Msg(errors.InvalidJWTToken.Error())
		return merr.Unauthorized(config.ServiceName, "%s", errors.InvalidJWTToken)
	}

	userRequest := &pbVideo.GetByUserIdRequest{
		UserId: claims.User.Id,
	}

	span.LogKV("event", "call userClient.GetByUserId")
	// fetch videos
	videos, err := u.videoClient.GetByUserId(ctx, userRequest)
	if err != nil {
		span.LogKV("event", "videos could not be fetched")
		log.Debug().Interface("error", err).Str("user", claims.User.Id).Msg("unable to fetch videos for user")
		return err
	}

	span.LogKV("event", "videos fetched")
	res.Videos = videos.Videos

	return nil
}
