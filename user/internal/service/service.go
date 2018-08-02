package service

import (
	"context"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/errors"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/mongodb"
	pb "github.com/lukasjarosch/educonn-platform/user/proto"
	merr "github.com/micro/go-micro/errors"
	"github.com/rs/zerolog/log"
	"gopkg.in/mgo.v2/bson"
)

type userService struct {
	repo                      *mongodb.UserRepository
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
func NewUserService(repo *mongodb.UserRepository, createdPublisher userCreatedEventPublisher, deletedPublisher userDeletedEventPublisher) pb.UserHandler {
	return &userService{
		repo: repo,
		userCreatedEventPublisher: createdPublisher,
		userDeletedEventPublisher: deletedPublisher,
	}
}

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

func (s *userService) GetAll(ctx context.Context, req *pb.Request, res *pb.UserResponse) error {
	return nil
}

func (s *userService) Auth(ctx context.Context, req *pb.UserDetails, res *pb.Token) error {
	return nil
}

func (s *userService) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}

func (s *userService) Delete(ctx context.Context, req *pb.DeleteRequest, res *pb.DeleteResponse) error {
	log.Info().Str("user_id", req.User.Id).Msg("deleted user")

	s.userDeletedEventPublisher.PublishUserDeleted(&pb.UserDeletedEvent{
		User: req.User,
	})
	return nil
}
