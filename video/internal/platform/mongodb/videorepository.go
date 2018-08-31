package mongodb

import (
	"fmt"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/video/internal/platform/errors"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"context"
	"github.com/opentracing/opentracing-go"
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

func UnmarshalProtobuf(video *pbVideo.VideoDetails, userId string) *Video {
	return &Video{
		Title:       video.Title,
		UserID:      bson.ObjectIdHex(userId),
		Description: video.Description,
		Tags:        video.Tags,
		Storage: Storage{
			RawKey: video.Storage.RawKey,
		},
	}
}

// CreateVideo will insert a new video entry
func (v *VideoRepository) CreateVideo(ctx context.Context, video *Video) (*Video, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "VideoRepository.CreateVideo")
	defer span.Finish()

	session := v.session.Clone()
	defer session.Close()

	video.ID = bson.NewObjectId()

	err := session.DB(config.DbName).C(config.VideoCollectionName).Insert(video)
	if err != nil {
		return nil, nil
	}
	return video, nil
}

func (v *VideoRepository) UpdateVideo(ctx context.Context, video *Video) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "VideoRepository.UpdateVideo")
	defer span.Finish()

	session := v.session.Clone()
	defer session.Close()

	err := session.DB(config.DbName).C(config.VideoCollectionName).UpdateId(video.ID, video)
	if err != nil {
		return err
	}
	return nil
}

func (v *VideoRepository) FindById(ctx context.Context, id string) (*Video, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "VideoRepository.FindById")
	defer span.Finish()

	session := v.session.Clone()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		return nil, errors.InvalidVideoId
	}

	var video = &Video{}
	err := session.DB(config.DbName).C(config.VideoCollectionName).FindId(bson.ObjectIdHex(id)).One(video)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (v *VideoRepository) FindByRawStorageKey(ctx context.Context, key string) (*Video, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "VideoRepository.FindByRawStorageKey")
	defer span.Finish()

	session := v.session.Clone()
	defer session.Close()

	var video = &Video{}
	err := session.DB(config.DbName).C(config.VideoCollectionName).Find(bson.M{"storage.raw_key": key}).One(video)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (v *VideoRepository) FindByUserId(ctx context.Context, userId string) ([]*Video, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "VideoRepository.FindByUserId")
	defer span.Finish()

	session := v.session.Clone()
	defer session.Close()

	if !bson.IsObjectIdHex(userId) {
		return nil, errors.Error("Malformed user id")
	}

	var videos = []*Video{nil}
	err := session.DB(config.DbName).C(config.VideoCollectionName).
		Find(bson.M{
			"user_id": bson.ObjectIdHex(userId),
		}).
		All(&videos)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (v *VideoRepository) IncrementViews(ctx context.Context, videoId bson.ObjectId) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "VideoRepository.IncrementViews")
	defer span.Finish()

	session := v.session.Clone()
	defer session.Close()

	video, err := v.FindById(ctx, videoId.Hex())
	if err != nil {
	    return err
	}

	video.Statistics.ViewCount = video.Statistics.ViewCount + 1

	err = v.UpdateVideo(ctx, video)
	if err != nil {
	    return err
	}

	return nil	
}
