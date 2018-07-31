package mongodb

import (
	"fmt"
	"github.com/lukasjarosch/educonn-master-thesis/transcode/internal/platform/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TranscodeRepository struct {
	session *mgo.Session
}

// NewTranscodeRepository creates a new transcode repository
func NewTranscodeRepository(host string, port string, user string, pass string, dbName string) (*TranscodeRepository, error) {
	connString := fmt.Sprintf("%s:%s/%s", host, port, dbName)
	session, err := mgo.Dial(connString)
	if err != nil {
		return nil, err
	}
	return &TranscodeRepository{
		session: session,
	}, nil
}

// CreateTranscodingJob creates a new transcoding job in the DB
func (t *TranscodeRepository) CreateTranscodingJob(job *TranscodingJob) (*TranscodingJob, error) {
	session := t.session.Clone()
	defer session.Close()

	job.ID = bson.NewObjectId()
	err := session.DB(config.DbName).C(config.TranscodingJobCollection).Insert(job)
	if err != nil {
	    return nil, err
	}
	return job, err
}

// FindByJobId tries to find a job by JobID
func (t *TranscodeRepository) FindByJobId(id string) (*TranscodingJob, error) {
	session := t.session.Clone()
	defer session.Close()
	var job TranscodingJob
	err := session.DB(config.DbName).C(config.TranscodingJobCollection).Find(bson.M{"job_id": id}).One(&job)
	if err != nil {
		return nil, err
	}
	return &job, err
}

// UpdateJob by it's id
func (t *TranscodeRepository) UpdateJob(job *TranscodingJob) (*TranscodingJob, error) {
	session := t.session.Clone()
	defer session.Close()

	err := session.DB(config.DbName).C(config.TranscodingJobCollection).UpdateId(job.ID, job)

	if err != nil {
		return nil, err
	}
	return job, err
}
