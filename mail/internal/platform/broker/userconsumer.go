package broker

import (
	"context"
	pbUser "github.com/lukasjarosch/educonn-platform/user/proto"
	"github.com/rs/zerolog/log"
)

const (
	UserCreatedTopic = "user.events.created"
	UserCreatedQueue = "user-created-queue"
	UserDeletedTopic = "user.events.deleted"
	UserDeletedQueue = "user-deleted-queue"
)

type UserCreatedSubscriber struct {
	userCreatedChan chan *pbUser.UserCreatedEvent
}

func NewUserCreatedSubscriber(userCreatedChannel chan *pbUser.UserCreatedEvent) *UserCreatedSubscriber {
	return &UserCreatedSubscriber{
		userCreatedChan: userCreatedChannel,
	}
}

func (s *UserCreatedSubscriber) Process(ctx context.Context, event *pbUser.UserCreatedEvent) error {
	s.userCreatedChan <- event
	event.User.Password = "" // unset the password or we would log the plaintext password
	log.Info().Str("topic", UserCreatedTopic).Str("user", event.User.Id).Msg("received UserCreatedEvent")
	return nil
}

// ---------------------------

type UserDeletedSubscriber struct {
	userDeletedChan chan *pbUser.UserDeletedEvent
}

func NewUserDeletedSubscriber(userDeletedChannel chan *pbUser.UserDeletedEvent) *UserDeletedSubscriber {
	return &UserDeletedSubscriber{
		userDeletedChan: userDeletedChannel,
	}
}

func (s *UserDeletedSubscriber) Process(ctx context.Context, event *pbUser.UserDeletedEvent) error {
	s.userDeletedChan <- event
	event.User.Password = "" // unset the password or we would log the plaintext password
	log.Info().Str("topic", UserDeletedTopic).Str("user", event.User.Id).Msg("received UserDeletedEvent")
	return nil
}
