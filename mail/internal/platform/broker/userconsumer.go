package broker

import (
	"context"
	"github.com/lukasjarosch/educonn-platform/user/proto"
	"github.com/rs/zerolog/log"
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
	log.Info().Str("topic", UserCreatedTopic).Str("user", event.User.Id).Msg("received UserCreatedEvent")
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
	s.userDeletedChan <- event
	event.User.Password = "" // unset the password or we would log the plaintext password
	log.Info().Str("topic", UserDeletedTopic).Str("user", event.User.Id).Msg("received UserDeletedEvent")
	return nil
}
