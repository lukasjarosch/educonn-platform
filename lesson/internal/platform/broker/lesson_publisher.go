package broker

import "github.com/micro/go-micro"
import (
	pbLesson "github.com/lukasjarosch/educonn-platform/lesson/proto"
	"context"
	"github.com/rs/zerolog/log"
)

const (
	LessonCreatedTopic = "lesson.events.created"
)

type LessonEventPublisher struct {
	lessonCreatedPublisher micro.Publisher
}

func NewLessonEventPublisher (lessonCreatedPublisher micro.Publisher) * LessonEventPublisher {
	return &LessonEventPublisher{lessonCreatedPublisher:lessonCreatedPublisher}
}

func (p *LessonEventPublisher) PublishLessonCreated(ctx context.Context, event *pbLesson.LessonCreatedEvent) (err error) {
	if err = p.lessonCreatedPublisher.Publish(ctx, event); err != nil {
		log.Warn().Str("topic", LessonCreatedTopic).Interface("event", event).Msg("unable to publish event")
		return err
	}
	log.Debug().Str("topic", LessonCreatedTopic).Interface("event", event).Msg("published event")
	return nil
}
