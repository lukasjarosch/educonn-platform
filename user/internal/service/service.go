package service

import (
	"context"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/errors"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/mongodb"
	pb "github.com/lukasjarosch/educonn-platform/user/proto"
	merr "github.com/micro/go-micro/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"github.com/micro/go-micro/metadata"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"
)

type userService struct {
	repo                      *mongodb.UserRepository
	tokenService              *jwt_handler.JwtTokenHandler
	userCreatedEventPublisher userCreatedEventPublisher
	userDeletedEventPublisher userDeletedEventPublisher
}

type userCreatedEventPublisher interface {
	PublishUserCreated(event *pb.UserCreatedEvent) (err error)
}

type userDeletedEventPublisher interface {
	PublishUserDeleted(event *pb.UserDeletedEvent) (err error)
}

// NewUserService creates a new userService
func NewUserService(repo *mongodb.UserRepository, createdPublisher userCreatedEventPublisher, deletedPublisher userDeletedEventPublisher, tokenService *jwt_handler.JwtTokenHandler) pb.UserHandler {
	return &userService{
		repo: repo,
		userCreatedEventPublisher: createdPublisher,
		userDeletedEventPublisher: deletedPublisher,
		tokenService:              tokenService,
	}
}

// Create a new user
func (s *userService) Create(ctx context.Context, req *pb.UserDetails, res *pb.UserResponse) error {

	// create user type
	hashedPass, _ := HashAndSalt([]byte(req.Password))
	user := &mongodb.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPass,
	}

	// check if email is already taken
	existingUser, err := s.repo.FindByEmail(user.Email)
	if existingUser != nil {
		log.Warn().Msg(errors.EmailExists.Error())
		return errors.EmailExists
	}

	// create user
	user, err = s.repo.CreateUser(user)
	if err != nil {
		log.Info().Interface("error", err).Msg("unable to create user")
		return err
	}
	req.Id = user.ID.Hex()
	log.Debug().Str("user", user.ID.Hex()).Msg("successfully created new user")

	// build response
	res.User = &pb.UserDetails{
		Id:        user.ID.Hex(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Email:     user.Email,
	}

	// publish user.events.created
	s.userCreatedEventPublisher.PublishUserCreated(&pb.UserCreatedEvent{
		User: res.User,
	})

	return nil
}

// Get a specific user, either by ID or by email depending on which field is set
func (s *userService) Get(ctx context.Context, req *pb.UserDetails, res *pb.UserResponse) (err error) {
	var user *mongodb.User

	if req.GetId() != "" {
		if bson.IsObjectIdHex(req.GetId()) == false {
			return merr.BadRequest(config.ServiceName, "%s", errors.MalformedUserId)
		}
		user, err = s.repo.FindById(req.GetId())
		if err != nil {
			return err
		}
	}
	if req.GetEmail() != "" && req.GetId() == "" {
		user, err = s.repo.FindByEmail(req.GetEmail())
		if err != nil {
			return err
		}
	}
	if req.GetId() == "" && req.GetEmail() == "" {
		return errors.MalformedUserId
	}

	res.User = &pb.UserDetails{
		Id:        user.ID.Hex(),
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return nil
}

// GetAll fetch all users
func (s *userService) GetAll(ctx context.Context, req *pb.Request, res *pb.UserResponse) error {
	users, err := s.repo.GetAll()
	if err != nil {
		return merr.InternalServerError(config.ServiceName, "%s", err)
	}

	for _, user := range users {
		pbUser := &pb.UserDetails{
			Id:        user.ID.Hex(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  user.Password,
		}
		res.Users = append(res.Users, pbUser)
	}

	return nil
}

// Delete deletes a user
func (s *userService) Delete(ctx context.Context, req *pb.DeleteRequest, res *pb.DeleteResponse) error {
	if req.User.GetId() == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.MissingUserId)
	}

	md, _ := metadata.FromContext(ctx)
	token, err := s.tokenService.GetBearerToken(md)
	if err != nil {
	    return merr.Unauthorized(config.ServiceName, "%s", jwt_handler.Unauthorized)
	}
	claims, _ := s.tokenService.Decode(token)

	// TODO: check user permissions

	// Currently only the user can delete itself
	if claims.User.Id != req.User.Id {
		return merr.Unauthorized(config.ServiceName, "%s", jwt_handler.Unauthorized)
	}

	user, _ := s.repo.FindById(req.User.Id)

	err = s.repo.DeleteUser(req.User.GetId())
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return merr.NotFound(config.ServiceName, "%s", errors.UserNotFound)
		}
		if err.Error() == errors.MalformedUserId.Error() {
			return merr.BadRequest(config.ServiceName, "%s", err.Error())
		}
		return merr.InternalServerError(config.ServiceName, "%s", err)
	}

	log.Info().Str("user_id", req.User.Id).Msg("deleted user from database")

	s.userDeletedEventPublisher.PublishUserDeleted(&pb.UserDeletedEvent{
		User: &pb.UserDetails{
			Id:        user.ID.Hex(),
			Email:     user.Email,
			LastName:  user.LastName,
			FirstName: user.FirstName,
			Password:  user.Password,
		},
	})
	return nil
}

// Auth creates a new RS256 JWT token using the 4096bit-RSA auth keys of the project including CustomClaims
func (s *userService) Auth(ctx context.Context, req *pb.UserDetails, res *pb.Token) error {

	// validate input
	if req.GetEmail() == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.MissingEmail)
	}
	if req.GetPassword() == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.MissingPassword)
	}

	// fetch user
	user, err := s.repo.FindByEmail(req.GetEmail())
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return merr.NotFound(config.ServiceName, "%s", errors.UserNotFound)
		}
		return merr.InternalServerError(config.ServiceName, "%s", err)
	}

	// compare password hashes
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword())); err != nil {
		log.Warn().Interface("error", errors.InvalidCredentials)
		return merr.Unauthorized(config.ServiceName, "%s", errors.InvalidCredentials)
	}

	// create a new token
	token, err := s.tokenService.Encode(&pb.UserDetails{
		Id:        user.ID.Hex(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
	if err != nil {
		log.Error().Interface("error", err).Msg("ParseInt was unable to parse the JWT_EXPIRE_SECONDS to Int64")
		return merr.InternalServerError(config.ServiceName, "%s", "unable to create token")
	}

	res.Token = token

	return nil
}

// Decode decodes and validates the token with the public key
func (s *userService) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	res.Valid = false

	_, err := s.tokenService.Decode(req.Token)
	if err != nil {
		log.Info().Interface("error", err).Msg(errors.JwtDecodingFailed.Error())
		return errors.JwtDecodingFailed
	}

	res.Valid = true
	return nil
}
