package service

import (
	"context"
	pb "github.com/lukasjarosch/educonn-platform/api/user/proto"
	pbUser "github.com/lukasjarosch/educonn-platform/user/proto"
	merr "github.com/micro/go-micro/errors"
	"github.com/rs/zerolog/log"
	"net/mail"
	"github.com/micro/go-micro/metadata"
	"github.com/lukasjarosch/educonn-platform/api/user/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/api/user/internal/platform/errors"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"
)

type User struct {
	client pbUser.UserClient
	jwtService *jwt_handler.JwtTokenHandler
}

func NewUserApi(userClient pbUser.UserClient, jwtService *jwt_handler.JwtTokenHandler) *User {
	return &User{client: userClient, jwtService:jwtService}
}

func (u *User) Create(ctx context.Context, req *pb.CreateRequest, res *pb.CreateResponse) error {

	log.Debug().Interface("u", req.User)

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

	user, err := u.client.Create(ctx, req.User)
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
	_, err = u.client.Delete(ctx, &pbUser.DeleteRequest{
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
