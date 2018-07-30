package mongodb

import (
	"fmt"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/config"
	"github.com/lukasjarosch/educonn-master-thesis/video/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type VideoRepository struct {
	session *mgo.Session
}

// NewVideoRepository creates a new video repository
func NewVideoRepository(host string, port string, user string, pass string, dbName string) (*VideoRepository, error) {
	connString := fmt.Sprintf("%s:%s/%s", host, port, dbName)
	session, err := mgo.Dial(connString)
	if err != nil {
		return nil, err
	}
	return &VideoRepository{
		session: session,
	}, nil
}

func UnmarshalProtobuf(video *educonn_video.VideoDetails) *Video {
	return &Video{
		Title:       video.Title,
		Description: video.Description,
		Tags:        video.Tags,
		Storage: Storage{
			RawKey: video.Storage.RawKey,
		},
	}
}

// CreateVideo will insert a new video entry
func (v *VideoRepository) CreateVideo(video *Video) (*Video, error) {
	session := v.session.Clone()
	defer session.Close()

	video.ID = bson.NewObjectId()

	err := session.DB(config.DbName).C(config.VideoCollectionName).Insert(video)
	if err != nil {
		return nil, nil
	}
	return video, nil
}

func (v *VideoRepository) UpdateVideo(video *Video) error {
	session := v.session.Clone()
	defer session.Close()

	err := session.DB(config.DbName).C(config.VideoCollectionName).UpdateId(video.ID, video)
	if err != nil {
		return err
	}
	return nil
}

func (v *VideoRepository) FindByRawStorageKey(key string) (*Video, error) {
	session := v.session.Clone()
	defer session.Close()

	var video = &Video{}
	err := session.DB(config.DbName).C(config.VideoCollectionName).Find(bson.M{"storage.raw_key": key}).One(video)
	if err != nil {
	    return nil, err
	}
	return video, nil
}