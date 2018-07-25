package broker

import (
	"github.com/lukasjarosch/educonn-master-thesis/user/proto"
	"github.com/micro/go-micro/broker"
	"github.com/lunny/log"
)

const (
	topic = "user.events.created"
)

// CreateEventConsumer creates a broker that converts messages into item shipped events.
// These events are then put on the channel so it can be processed by other modules.
func CreateEventConsumer(userCreatedChannel chan *educonn_user.UserCreatedEvent) (err error) {
	_, err = broker.Subscribe(topic, func(p broker.Publication) error {
		log.Debugf("[sub] received message $+v", p.Message().Header)
		return nil
	})
	if err != nil {
	    return err
	}
	return nil
}