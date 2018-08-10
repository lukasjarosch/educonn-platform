package service

import "context"
import (
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/errors"
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/mongodb"
	pbLesson "github.com/lukasjarosch/educonn-platform/lesson/proto"
	merr "github.com/micro/go-micro/errors"
	"github.com/rs/zerolog/log"
	"gopkg.in/mgo.v2/bson"
	"github.com/micro/go-micro/metadata"
)

const (
	VideoServiceName = "educonn.srv.videoClient"
)

type LessonService struct {
	videoLessonService *VideoLessonService
	lessonRepo         *mongodb.LessonRepository
}

func NewLessonService(videoLesson *VideoLessonService, repo *mongodb.LessonRepository) *LessonService {
	return &LessonService{
		videoLessonService: videoLesson,
		lessonRepo:         repo,
	}
}

func (l *LessonService) Create(ctx context.Context, req *pbLesson.CreateLessonRequest, res *pbLesson.CreateLessonRequest) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		log.Info().Msg("no metadata in request")
	}
	log.Info().Interface("md", md).Msg("metadata")

	// LessonBase
	if req.Lesson.Base.Name == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.MissingLessonName.Error())
	}

	// TODO: fetch userId from context
	userId := bson.NewObjectId()

	// determine lesson type
	switch req.Lesson.Base.GetType() {

	case pbLesson.Type_VIDEO:

		// Create video lesson
		videoLessonResponse := &pbLesson.CreateVideoLessonResponse{}
		err := l.videoLessonService.Create(ctx, &pbLesson.CreateVideoLessonRequest{Lesson:req.Lesson.Video}, videoLessonResponse)
		if err != nil {
			return err
		}

		// Create BaseLesson and attach newly created VideoLesson
		lesson, err := l.lessonRepo.CreateBaseLesson(&mongodb.BaseLesson{
			Title:       req.Lesson.Base.Name,
			Description: req.Lesson.Base.Description,
			Type:        req.Lesson.Base.Type.String(),
			TypeID:      bson.ObjectIdHex(videoLessonResponse.Lesson.Id),
			UserID:      userId,
		})
		if err != nil {
			log.Info().Err(err).Msg(errors.MongoCreateFailed.Error())
			return merr.BadRequest(config.ServiceName, "%s", errors.MongoCreateFailed)
		}

		log.Debug().Str("lesson", lesson.ID.Hex()).Str("type", req.Lesson.Base.Type.String()).Str("type_id", videoLessonResponse.Lesson.Id).Msg("created Lesson")

		// TODO: Pub lesson.events.created
		break

	case pbLesson.Type_TEXT:
		log.Debug().Msg("TEXT")
		break

	default:
		return merr.BadRequest(config.ServiceName, "%s", errors.UnknownLessonType.Error())
		break
	}

	res.Lesson = req.Lesson
	return nil
}
