package service

import (
	"context"
	pb "github.com/lukasjarosch/educonn-platform/user-api/proto"
	pbUser "github.com/lukasjarosch/educonn-platform/user/proto"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	merr "github.com/micro/go-micro/errors"
	"github.com/rs/zerolog/log"
	"net/mail"
	"github.com/micro/go-micro/metadata"
	"github.com/lukasjarosch/educonn-platform/user-api/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/user-api/internal/platform/errors"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"
)

type User struct {
	userClient pbUser.UserClient
	videoClient pbVideo.VideoClient
	jwtService *jwt_handler.JwtTokenHandler
}

func NewUserApi(userClient pbUser.UserClient, videoClient pbVideo.VideoClient, jwtService *jwt_handler.JwtTokenHandler) *User {
	return &User{userClient: userClient, videoClient: videoClient, jwtService:jwtService}
}

func (u *User) Create(ctx context.Context, req *pb.CreateRequest, res *pb.CreateResponse) error {


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

	user, err := u.userClient.Create(ctx, req.User)
	if err != nil {
		log.Warn().Interface("error", err).Msg("unable to create user")
		return merr.InternalServerError(config.ServiceName, "%s: %s", "Unable to create user", err.Error())
	}

	res.User = user.User

	return nil
}

func (u *User) Delete(ctx context.Context, req *pb.DeleteRequest, res *pb.DeleteResponse) error {

	md, _ := metadata.FromContext(ctx)

	// check for token
	token, err := u.jwtService.GetBearerToken(md)
	if err != nil {
	    log.Error().Interface("err", err).Msg("asdf")
	    return merr.Unauthorized(config.ServiceName, "%s", err.Error())
	}

	// validate jwt
	_, err = u.jwtService.Decode(token)
	if err != nil {
	    log.Error().Interface("error", err).Msg(errors.InvalidJWTToken.Error())
	    return merr.Unauthorized(config.ServiceName, "%s", errors.InvalidJWTToken)
	}


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

	return nil
}

func (u *User) Login(ctx context.Context, req *pb.LoginRequest, res *pb.LoginResponse) error {

	if req.Email == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.EmailMissing.Error())
	}

	if req.Password == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.PasswordMissing.Error())
	}

	user := &pbUser.UserDetails{
		Email: req.Email,
		Password: req.Password,
	}

	// login
	response, err := u.userClient.Auth(ctx, user)
	if err != nil {
	    log.Debug().Interface("error", err).Str("email", req.Email).Msgf("user login failed")
	    return err
	}

	res.Token = response.Token

	return nil
}

func (u *User) Videos(ctx context.Context, req *pb.VideoRequest, res *pb.VideoResponse) error {
	md, _ := metadata.FromContext(ctx)

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

	// fetch videos
	videos, err := u.videoClient.GetByUserId(ctx, userRequest)
	if err != nil {
		log.Debug().Interface("error", err).Str("user", claims.User.Id).Msg("unable to fetch videos for user")
	    return err
	}

	res.Videos = videos.Videos

	return nil
}
