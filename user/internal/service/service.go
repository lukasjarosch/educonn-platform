package service

import (
	"context"
	"github.com/lukasjarosch/educonn-master-thesis/user/proto"
)

type userService struct {
	repo                      userRepository
	userCreatedEventPublisher userCreatedEventPublisher
	userDeletedEventPublisher userDeletedEventPublisher
}

type userRepository interface {
}

type userCreatedEventPublisher interface {
	PublishUserCreated(event *educonn_user.UserCreatedEvent) (err error)
}

type userDeletedEventPublisher interface {
	PublishUserDeleted(event *educonn_user.UserDeletedEvent) (err error)
}

// NewUserService creates a new userService
func NewUserService(repo userRepository, createdPublisher userCreatedEventPublisher, deletedPublisher userDeletedEventPublisher) educonn_user.UserHandler {
	return &userService{
		repo: repo,
		userCreatedEventPublisher: createdPublisher,
		userDeletedEventPublisher: deletedPublisher,
	}
}

func (s *userService) Create(ctx context.Context, req *educonn_user.UserDetails, res *educonn_user.UserResponse) error {
	s.userCreatedEventPublisher.PublishUserCreated(&educonn_user.UserCreatedEvent{
		User: req,
	})
	return nil
}

func (s *userService) Get(ctx context.Context, req *educonn_user.UserDetails, res *educonn_user.UserResponse) error {
	return nil
}

func (s *userService) GetAll(ctx context.Context, req *educonn_user.Request, res *educonn_user.UserResponse) error {
	return nil
}

func (s *userService) Auth(ctx context.Context, req *educonn_user.UserDetails, res *educonn_user.Token) error {
	return nil
}

func (s *userService) ValidateToken(ctx context.Context, req *educonn_user.Token, res *educonn_user.Token) error {
	return nil
}

func (s *userService) Delete(ctx context.Context, req *educonn_user.DeleteRequest, res *educonn_user.DeleteResponse) error {
	s.userDeletedEventPublisher.PublishUserDeleted(&educonn_user.UserDeletedEvent{
		User: req.User,
	})
	return nil
}
