package broker

import (
	"context"
	pb "github.com/lukasjarosch/educonn-platform/user/proto"
	"github.com/micro/go-micro"
	"github.com/rs/zerolog/log"
)

const (
	UserCreatedTopic = "user.events.created"
	UserDeletedTopic = "user.events.deleted"
)

// EventPublisher is an even publisher for the go-micro broker
type EventPublisher struct {
	userCreatedPublisher micro.Publisher
	userDeletedPublisher micro.Publisher
}

// NewEventPublisher creates a new broker event publisher
func NewEventPublisher(userCreatedPublisher micro.Publisher) *EventPublisher {
	return &EventPublisher{userCreatedPublisher: userCreatedPublisher}
}

func (p *EventPublisher) PublishUserCreated(event *pb.UserCreatedEvent) (err error) {
	if err = p.userCreatedPublisher.Publish(context.Background(), event); err != nil {
		log.Warn().
			Str("topic", UserCreatedTopic).
			Interface("event", event).
			Msg("unable to publish event")
		return nil
	}
	log.Debug().
		Str("topic", UserCreatedTopic).
		Interface("event", event).
		Msg("published event")
	return nil
}

func (p *EventPublisher) PublishUserDeleted(event *pb.UserDeletedEvent) (err error) {
	if err = p.userCreatedPublisher.Publish(context.Background(), event); err != nil {
		log.Warn().Str("topic", UserDeletedTopic).Interface("event", event).Msg("unable to publish event")
		return nil
	}
	log.Debug().Str("topic", UserDeletedTopic).Interface("event", event).Msg("published event")
	return nil
}
