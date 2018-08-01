package service

import (
	"context"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	pb "github.com/lukasjarosch/educonn-platform/user/proto"
)

type userService struct {
	repo                      userRepository
	userCreatedEventPublisher userCreatedEventPublisher
	userDeletedEventPublisher userDeletedEventPublisher
}

type userRepository interface {
}

type userCreatedEventPublisher interface {
	PublishUserCreated(event *pb.UserCreatedEvent) (err error)
}

type userDeletedEventPublisher interface {
	PublishUserDeleted(event *pb.UserDeletedEvent) (err error)
}

// NewUserService creates a new userService
func NewUserService(repo userRepository, createdPublisher userCreatedEventPublisher, deletedPublisher userDeletedEventPublisher) pb.UserHandler {
	return &userService{
		repo: repo,
		userCreatedEventPublisher: createdPublisher,
		userDeletedEventPublisher: deletedPublisher,
	}
}

func (s *userService) Create(ctx context.Context, req *pb.UserDetails, res *pb.UserResponse) error {
	userId := xid.New()

	// TODO Database stuff
	log.Info().Str("user_id", userId.String()).Msg("created new user")

	req.Id = userId.String()
	s.userCreatedEventPublisher.PublishUserCreated(&pb.UserCreatedEvent{
		User: req,
	})
	return nil
}

func (s *userService) Get(ctx context.Context, req *pb.UserDetails, res *pb.UserResponse) error {
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
