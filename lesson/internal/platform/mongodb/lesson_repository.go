package mongodb

import (
	"fmt"
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/lukasjarosch/educonn-platform/lesson/internal/platform/errors"
)

type LessonRepository struct {
	session *mgo.Session
}

func Dial(host string, port string, user string, pass string, dbName string) (session *mgo.Session, err error) {
	connString := fmt.Sprintf("%s:%s/%s", host, port, dbName)
	session, err = mgo.Dial(connString)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// NewLessonRepository creates a new lesson repository
func NewLessonRepository(session *mgo.Session) *LessonRepository {
	return &LessonRepository{
		session: session,
	}
}


// CreateVideo will insert a new video entry
func (v *LessonRepository) CreateBaseLesson(baseLesson *BaseLesson) (*BaseLesson, error) {
	session := v.session.Clone()
	defer session.Close()

	baseLesson.ID = bson.NewObjectId()

	err := session.DB(config.DbName).C(config.BaseLessonCollection).Insert(baseLesson)
	if err != nil {
		return nil, err
	}
	return baseLesson, nil
}

func (v *LessonRepository) GetById(id string)  (*BaseLesson, error) {
	session := v.session.Clone()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		return nil, errors.MalformedId
	}

	lesson := &BaseLesson{}

	err := session.DB(config.DbName).C(config.BaseLessonCollection).FindId(bson.ObjectIdHex(id)).One(lesson)
	if err != nil {
		return nil, err
	}
	return lesson, nil
}

