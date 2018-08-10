package mongodb

import (
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/errors"
)

type VideoLessonRepository struct {
	session *mgo.Session
}

// NewLessonRepository creates a new lesson repository
func NewVideoLessonRepository(session *mgo.Session) *VideoLessonRepository {
	return &VideoLessonRepository{
		session: session,
	}
}

func (v *VideoLessonRepository) CreateVideoLesson(videoLesson *VideoLesson) (*VideoLesson, error) {
	session := v.session.Clone()
	defer session.Close()

	videoLesson.ID = bson.NewObjectId()

	err := session.DB(config.DbName).C(config.VideoLessonCollection).Insert(videoLesson)
	if err != nil {
		return nil, err
	}
	return videoLesson, nil
}

func (v *VideoLessonRepository) GetById(id string)  (*VideoLesson, error) {
	session := v.session.Clone()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		return nil, errors.MalformedId
	}

	lesson := &VideoLesson{}

	err := session.DB(config.DbName).C(config.VideoLessonCollection).FindId(bson.ObjectIdHex(id)).One(lesson)
	if err != nil {
		return nil, err
	}
	return lesson, nil
}
