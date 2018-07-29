package broker

import (
	"context"
	"github.com/lukasjarosch/educonn-master-thesis/user/proto"
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
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

func (p *EventPublisher) PublishUserCreated(event *educonn_user.UserCreatedEvent) (err error) {
	if err = p.userCreatedPublisher.Publish(context.Background(), event); err != nil {
		log.Warnf("[pub] Unable to publish to %s: %+v", UserCreatedTopic, err)
		return nil
	}
	log.Infof("[pub] published '%s' for user '%s'", UserCreatedTopic, event.User.Id)
	return nil
}

func (p *EventPublisher) PublishUserDeleted(event *educonn_user.UserDeletedEvent) (err error) {
	if err = p.userCreatedPublisher.Publish(context.Background(), event); err != nil {
		log.Warnf("[pub] Unable to publish to %s: %+v", UserDeletedTopic, err)
		return nil
	}
	log.Infof("[pub] published to '%s' '%s'", UserDeletedTopic, event.User.Id)
	return nil
}
