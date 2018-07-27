package broker

import (
	"context"
	"github.com/lukasjarosch/educonn-master-thesis/video/proto"
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
)

const (
	VideoCreatedTopic = "video.events.created"
)

type EventPublisher struct {
	videoCreatedPublisher micro.Publisher
}

func NewEventPublisher(videoCreatedPublisher micro.Publisher) *EventPublisher {
	return &EventPublisher{videoCreatedPublisher: videoCreatedPublisher}
}

func (p *EventPublisher) PublishVideoCreated(event *educonn_video.VideoCreatedEvent) (err error) {
	if err = p.videoCreatedPublisher.Publish(context.Background(), event); err != nil {
		log.Warnf("[pub] failed pub to %s: %+v", VideoCreatedTopic, err)
		return nil
	}
	log.Infof("[pub] published to %s: %v", VideoCreatedTopic, event)
	return nil
}
