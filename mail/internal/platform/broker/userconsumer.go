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
	VideoProcessedTopic = "video.events.processed"
	VideoProcessedQueue = "video-processed-queue"
)

type UserCreatedEvent struct {
	Event *pbUser.UserCreatedEvent
	Context context.Context
}

type UserCreatedSubscriber struct {
	userCreatedChan chan UserCreatedEvent
}

func NewUserCreatedSubscriber(userCreatedChannel chan UserCreatedEvent) *UserCreatedSubscriber {
	return &UserCreatedSubscriber{
		userCreatedChan: userCreatedChannel,
	}
}

func (s *UserCreatedSubscriber) Process(ctx context.Context, event *pbUser.UserCreatedEvent) error {
	s.userCreatedChan <- UserCreatedEvent{Event:event, Context:ctx}
	event.User.Password = "" // unset the password or we would log the plaintext password
	log.Info().Str("topic", UserCreatedTopic).Str("user", event.User.Id).Msg("received UserCreatedEvent")
	return nil
}

// ---------------------------

type UserDeletedEvent struct {
	Event *pbUser.UserDeletedEvent
	Context context.Context
}

type UserDeletedSubscriber struct {
	userDeletedChan chan UserDeletedEvent
}

func NewUserDeletedSubscriber(userDeletedChannel chan UserDeletedEvent) *UserDeletedSubscriber {
	return &UserDeletedSubscriber{
		userDeletedChan: userDeletedChannel,
	}
}

func (s *UserDeletedSubscriber) Process(ctx context.Context, event *pbUser.UserDeletedEvent) error {
	s.userDeletedChan <- UserDeletedEvent{Event:event, Context:ctx}
	event.User.Password = "" // unset the password or we would log the plaintext password
	log.Info().Str("topic", UserDeletedTopic).Str("user", event.User.Id).Msg("received UserDeletedEvent")
	return nil
}
