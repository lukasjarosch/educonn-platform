package service

import "context"
import (
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/broker"
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/errors"
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/mongodb"
	pbLesson "github.com/lukasjarosch/educonn-platform/lesson/proto"
	merr "github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"github.com/rs/zerolog/log"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

const (
	VideoServiceName = "educonn.srv.videoClient"
)

type LessonService struct {
	videoLessonService *VideoLessonService
	lessonRepo         *mongodb.LessonRepository
	publisher          *broker.LessonEventPublisher
}

func NewLessonService(videoLesson *VideoLessonService, repo *mongodb.LessonRepository, publisher *broker.LessonEventPublisher) *LessonService {
	return &LessonService{
		videoLessonService: videoLesson,
		lessonRepo:         repo,
		publisher:          publisher,
	}
}

// Create a new Lesson
func (l *LessonService) Create(ctx context.Context, req *pbLesson.CreateLesson_Request, res *pbLesson.CreateLesson_Response) error {

	// make sure we have a userId
	md, ok := metadata.FromContext(ctx)
	if !ok {
		log.Info().Msg("no metadata in request")
	}
	userId := md["X-User-Id"]
	if string(userId) == "" {
		log.Error().Msg("x-user-id empty")
		return merr.BadRequest(config.ServiceName, "%s", errors.MissingUserIdHeader)
	}

	// pre-flight checks
	if req.Name == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.MissingLessonName.Error())
	}

	if req.Type == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.MissingType.Error())
	}

	// determine lesson type
	switch strings.ToLower(req.GetType()) {

	case strings.ToLower(pbLesson.Type(pbLesson.Type_VIDEO).String()):

		// Create video lesson
		videoLessonRequest := &pbLesson.CreateVideoLessonRequest{
			Lesson: req.Video,
		}
		videoLessonResponse := &pbLesson.CreateVideoLessonResponse{}
		err := l.videoLessonService.Create(ctx, videoLessonRequest, videoLessonResponse)
		if err != nil {
			return err
		}

		// Create BaseLesson and attach newly created VideoLesson
		lesson, err := l.lessonRepo.CreateBaseLesson(&mongodb.BaseLesson{
			Title:       req.Name,
			Description: req.Description,
			Type:        req.Type,
			TypeID:      bson.ObjectIdHex(videoLessonResponse.Lesson.Id),
			UserID:      bson.ObjectIdHex(userId),
		})
		if err != nil {
			log.Info().Err(err).Msg(errors.MongoCreateFailed.Error())
			return merr.BadRequest(config.ServiceName, "%s", errors.MongoCreateFailed)
		}
		log.Debug().Str("lesson", lesson.ID.Hex()).
			Str("type", req.Type).
			Str("type_id", videoLessonResponse.Lesson.Id).
			Msg("created Lesson")

		// pub: lesson.events.created
		l.publisher.PublishLessonCreated(ctx, &pbLesson.LessonCreatedEvent{
			Lesson: &pbLesson.Lesson{
				Base: &pbLesson.LessonBase{
					Id:          lesson.ID.Hex(),
					Name:        lesson.Title,
					Description: lesson.Description,
					UserId:      lesson.UserID.Hex(),
					Type:        pbLesson.Type_VIDEO,
				},
			},
		})

		// response
		res.Lesson = &pbLesson.Lesson{
			Base: &pbLesson.LessonBase{
				Id:          lesson.ID.Hex(),
				UserId:      lesson.UserID.Hex(),
				Type:        pbLesson.Type_VIDEO,
				Description: lesson.Description,
				Name:        lesson.Title,
			},
			Video: &pbLesson.VideoLesson{
				Id:      videoLessonResponse.Lesson.Id,
				VideoId: videoLessonResponse.Lesson.VideoId,
			},
			// Stats are nil
		}
		break

	default:
		return merr.BadRequest(config.ServiceName, "%s: %s", errors.UnknownLessonType.Error(), req.Type)
		break
	}
	return nil
}

// GetById fetches a Lesson by id. The Lesson message will be aggregated with the correct LessonType
func (l *LessonService) GetById(ctx context.Context, req *pbLesson.GetLesson_ById_Request, res *pbLesson.GetLesson_ById_Response) error {

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

	switch strings.ToLower(lesson.Type) {
	case strings.ToLower(pbLesson.Type(pbLesson.Type_VIDEO).String()):

		videoResponse := &pbLesson.GetVideoLessonByIdResponse{}
		err := l.videoLessonService.GetById(ctx, &pbLesson.GetVideoLessonByIdRequest{LessonId: lesson.TypeID.Hex()}, videoResponse)
		if err != nil {
			log.Info().Str("video_lesson", lesson.TypeID.Hex()).Err(err).Msg("unable to fetch video-lession")
		}

		res.Lesson = &pbLesson.Lesson{
			Base: &pbLesson.LessonBase{
				Id:          lesson.ID.Hex(),
				Type:        pbLesson.Type_VIDEO,
				Description: lesson.Description,
				Name:        lesson.Title,
				UserId:      lesson.UserID.Hex(),
			},
			Stats: &pbLesson.LessonStatistics{
				Views:    lesson.Statistics.ViewCount,
				Likes:    lesson.Statistics.Likes,
				Dislikes: lesson.Statistics.Dislikes,
			},
			Video: &pbLesson.VideoLesson{
				Id:      videoResponse.Lesson.Id,
				VideoId: videoResponse.Lesson.VideoId,
			},
		}

		break
	default:
		log.Info().Str("lesson", lesson.ID.Hex()).Str("type", lesson.Type).Msg("lesson has unknown type")

	}

	return nil
}
