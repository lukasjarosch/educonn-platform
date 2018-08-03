package mongodb

import (
	"gopkg.in/mgo.v2/bson"
)

type Video struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	UserID      bson.ObjectId `bson:"user_id" json:"user_id"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	Tags        []string      `bson:"tags" json:"tags"`
	Statistics  Statistics    `bson:"statistics" json:"statistics"`
	Storage     Storage       `bson:"storage" json:"storage"`
	Transcode   Transcode     `bson:"transcode" json:"transcode"`
}

type Transcode struct {
	Completed bool     `bson:"completed" json:"completed"`
	Errors    []string `bson:"errors" json:"errors"`
}

type Statistics struct {
	ViewCount    int64 `bson:"view_count" json:"view_count"`
	LikeCount    int64 `bson:"like_count" json:"like_count"`
	DislikeCount int64 `bson:"dislike_count" json:"dislike_count"`
}

type Storage struct {
	RawKey    string `bson:"raw_key" json:"raw_key"`
	OutputKey string `bson:"output_key" json:"output_key"`
}
