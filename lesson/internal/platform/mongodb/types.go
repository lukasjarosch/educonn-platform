package mongodb

import (
	"gopkg.in/mgo.v2/bson"
)

// Lesson is the base lesson document which is handled by the LessonService
type BaseLesson struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	UserID      bson.ObjectId `bson:"user_id" json:"user_id"`
	TypeID      bson.ObjectId `bson:"type_id" json:"type_id"`
	Title       string        `bson:"title" json:"title"`
	Type        string        `bson:"type" json:"type"`
	Description string        `bson:"description" json:"description"`
	Statistics  Statistics    `bson:"statistics" json:"statistics"`
}

// Statistics of a lesson
type Statistics struct {
	ViewCount int64 `bson:"view_count" json:"view_count"`
	Likes     int64 `bson:"likes" json:"likes"`
	Dislikes  int64 `bson:"dislikes" json:"dislikes"`
}

// VideoLesson is the concrete lesson type for videos
type VideoLesson struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	VideoId bson.ObjectId `bson:"video_id" json:"video_id"`
}
