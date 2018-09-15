package service

import (
	_ "github.com/rs/zerolog"
	_ "github.com/rs/zerolog/log"
)

type lessonApiService struct {
}

func NewLessonApiService() *lessonApiService {
	return &lessonApiService{}
}
