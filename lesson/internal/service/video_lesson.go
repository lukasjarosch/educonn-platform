package service

import (
	"context"

	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/errors"
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/mongodb"
	pbLesson "github.com/lukasjarosch/educonn-platform/lesson/proto"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	merr "github.com/micro/go-micro/errors"
	"github.com/rs/zerolog/log"
	"gopkg.in/mgo.v2/bson"
)

type VideoLessonService struct {
	videoClient     pbVideo.VideoClient
	videoLessonRepo *mongodb.VideoLessonRepository
}

func NewVideoLessonService(videoClient pbVideo.VideoClient, repo *mongodb.VideoLessonRepository) *VideoLessonService {
	return &VideoLessonService{
		videoClient:     videoClient,
		videoLessonRepo: repo,
	}
}

// Create a new video lesson. This method is called from LessonService.Create as middleware
func (v *VideoLessonService) Create(ctx context.Context, req *pbLesson.CreateVideoLessonRequest, res *pbLesson.CreateVideoLessonResponse) error {

	// videoId must be set
	videoId := req.Lesson.GetVideoId()
	if videoId == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.MissingVideoId.Error())
	}

	if !bson.IsObjectIdHex(videoId) {
		return merr.BadRequest(config.ServiceName, "%s", errors.MalformedId)
	}

	// check if videoClient actually exists
	_, err := v.videoClient.GetById(ctx, &pbVideo.GetVideoRequest{Id: videoId})
	if err != nil {
		log.Info().Err(err).Str("videoClient", videoId).Msg("unable to call Video.GetById")
		return err
	}

	// create VideoLesson in mongodb
	lessId := bson.NewObjectId()
	lesson, err := v.videoLessonRepo.CreateVideoLesson(&mongodb.VideoLesson{
		ID:      lessId,
		VideoId: bson.ObjectIdHex(req.Lesson.VideoId),
	})
	if err != nil {
		log.Info().Err(err).Msg("unable to create VideoLesson")
	}

	log.Debug().Str("video_lesson", lesson.ID.Hex()).Msg("created VideoLesson")

	// return
	res.Lesson = &pbLesson.VideoLesson{
		VideoId: lesson.VideoId.Hex(),
		Id:      lesson.ID.Hex(),
	}
	return nil
}

func (v *VideoLessonService) GetById(ctx context.Context, req *pbLesson.GetVideoLessonByIdRequest, res *pbLesson.GetVideoLessonByIdResponse) error {

	if req.LessonId == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.MissingLessonId)
	}

	// Fetch lesson
	lesson, err := v.videoLessonRepo.GetById(req.LessonId)
	if err != nil {
		log.Debug().Str("lesson", req.LessonId).Err(err).Msg("unable to fetch lesson")
		return merr.InternalServerError(config.ServiceName, "%s", err.Error())
	}
	log.Debug().Str("lesson", lesson.ID.Hex()).Msg("fetched video-lesson")

	res.Lesson = &pbLesson.VideoLesson{
		Id:      lesson.ID.Hex(),
		VideoId: lesson.VideoId.Hex(),
	}

	return nil
}
