package broker

import (
	"context"
	"github.com/lukasjarosch/educonn-platform/user/proto"
	"github.com/micro/go-micro/metadata"
	log "github.com/sirupsen/logrus"
)

const (
	UserCreatedTopic = "user.events.created"
	UserCreatedQueue = "user-created-queue"
	UserDeletedTopic = "user.events.deleted"
	UserDeletedQueue = "user-deleted-queue"
)

type UserCreatedSubscriber struct {
	userCreatedChan chan *educonn_user.UserCreatedEvent
}

func NewUserCreatedSubscriber(userCreatedChannel chan *educonn_user.UserCreatedEvent) *UserCreatedSubscriber {
	return &UserCreatedSubscriber{
		userCreatedChan: userCreatedChannel,
	}
}

func (s *UserCreatedSubscriber) Process(ctx context.Context, event *educonn_user.UserCreatedEvent) error {
	s.userCreatedChan <- event
	event.User.Password = "" // unset the password or we would log the plaintext password
	log.Infof("[sub] received event '%s' for user '%s'", UserCreatedTopic, event.User.Id)
	return nil
}

// ---------------------------

type UserDeletedSubscriber struct {
	userDeletedChan chan *educonn_user.UserDeletedEvent
}

func NewUserDeletedSubscriber(userDeletedChannel chan *educonn_user.UserDeletedEvent) *UserDeletedSubscriber {
	return &UserDeletedSubscriber{
		userDeletedChan: userDeletedChannel,
	}
}

func (s *UserDeletedSubscriber) Process(ctx context.Context, event *educonn_user.UserDeletedEvent) error {
	md, _ := metadata.FromContext(ctx)
	log.Infof("%+v", md)
	s.userDeletedChan <- event
	event.User.Password = "" // unset the password or we would log the plaintext password
	log.Infof("[sub] received event '%s' for user '%s' ", UserDeletedTopic, event.User.Id)
	return nil
}
