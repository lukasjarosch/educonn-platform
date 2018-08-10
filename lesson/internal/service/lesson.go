package service

import "context"
import (
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/errors"
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/mongodb"
	pbLesson "github.com/lukasjarosch/educonn-platform/lesson/proto"
	merr "github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"github.com/rs/zerolog/log"
	"gopkg.in/mgo.v2/bson"
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

// Create a new Lesson
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
		err := l.videoLessonService.Create(ctx, &pbLesson.CreateVideoLessonRequest{Lesson: req.Lesson.Video}, videoLessonResponse)
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
		log.Debug().Str("lesson", lesson.ID.Hex()).
			Str("type", req.Lesson.Base.Type.String()).
			Str("type_id", videoLessonResponse.Lesson.Id).
			Msg("created Lesson")

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

func (l *LessonService) GetById(ctx context.Context, req *pbLesson.GetLessonByIdRequest, res *pbLesson.GetLessonByIdResponse) error {

	// id cannot be empty
	if req.LessonId == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.MissingLessonId)
	}

	// Fetch BaseLession
	lesson, err := l.lessonRepo.GetById(req.LessonId)
	if err != nil {
		log.Debug().Str("lesson", req.LessonId).Err(err).Msg("unable to fetch lesson")
	    return merr.InternalServerError(config.ServiceName, "%s", err.Error())
	}
	log.Debug().Str("lesson", lesson.ID.Hex()).Msg("fetched lesson")


	switch lesson.Type {
	case pbLesson.Type(pbLesson.Type_VIDEO).String():

		videoResponse := &pbLesson.GetVideoLessonByIdResponse{}
		err := l.videoLessonService.GetById(ctx, &pbLesson.GetVideoLessonByIdRequest{LessonId:lesson.TypeID.Hex()}, videoResponse)
		if err != nil {
		    log.Info().Str("video_lesson", lesson.TypeID.Hex()).Err(err).Msg("unable to fetch video-lession")
		}

		res.Lesson = &pbLesson.Lesson{
			Base: &pbLesson.LessonBase{
				Id: lesson.ID.Hex(),
				Type: pbLesson.Type_VIDEO,
				Description: lesson.Description,
				Name: lesson.Title,
				UserId: lesson.UserID.Hex(),
			},
			Stats: &pbLesson.LessonStatistics{
				Views: lesson.Statistics.ViewCount,
				Likes: lesson.Statistics.Likes,
				Dislikes: lesson.Statistics.Dislikes,
			},
			Video: &pbLesson.VideoLesson{
				Id: videoResponse.Lesson.Id,
				VideoId: videoResponse.Lesson.VideoId,
			},
		}

	break
	case pbLesson.Type(pbLesson.Type_TEXT).String():
		log.Debug().Msg("GetById():Type_TEXT")
		break
	default:
		log.Info().Str("lesson", lesson.ID.Hex()).Str("type", lesson.Type).Msg("lesson has unknown type")

	}

	return nil
}
