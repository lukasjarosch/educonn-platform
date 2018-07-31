package mongodb

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type TranscodingJob struct {
	ID              bson.ObjectId `bson:"_id" json:"id"`
	JobId           string        `bson:"job_id" json:"job_id"`
	VideoId         bson.ObjectId `bson:"video_id" json:"video_id"`
	PipelineId      string        `bson:"pipeline_id" json:"pipeline_id"`
	InputKey        string        `bson:"input_key" json:"input_key"`
	OutputKey       string        `bson:"output_key" json:"output_key"`
	OutputKeyPrefix string        `bson:"output_key_prefix" json:"output_key_prefis"`
	Status          Status        `bson:"status" json:"status"`
	StartedAt       time.Time     `bson:"started_at" json:"started_at"`
	EndedAt         time.Time     `bson:"ended_at" json:"ended_at"`
}

type Status struct {
	Started      bool   `bson:"started" json:"started"`
	Completed    bool   `bson:"completed" json:"completed"`
	Error        bool   `bson:"error" json:"error"`
	ErrorMessage string `bson:"error_message" json:"error_message"`
}
