package service

import (
	"context"
	"github.com/lukasjarosch/educonn-master-thesis/video/proto"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/amazon"
	"github.com/prometheus/common/log"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/config"
	"github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/errors"
)

type videoService struct {
	videoCreatedPublisher videoCreatedPublisher
	s3Bucket *amazon.S3Bucket
}

type videoCreatedPublisher interface {
	PublishVideoCreated(event *educonn_video.VideoCreatedEvent) (err error)
}

func NewVideoService(vidCreatedPub videoCreatedPublisher, bucket *amazon.S3Bucket) educonn_video.VideoHandler {
	svc := &videoService{
		videoCreatedPublisher: vidCreatedPub,
		s3Bucket:bucket,
	}
	return svc
}

func (v *videoService) Create(ctx context.Context, req *educonn_video.CreateVideoRequest, res *educonn_video.CreateVideoResponse) error {
	// Check if file actually exists
	fileKey := req.Video.Storage.RawKey
	err := v.s3Bucket.CheckFileExists(fileKey, config.AwsS3VideoBucket)
	if err != nil {
		log.Infof("[S3] key '%s' does not exist in bucket", fileKey)
		res.Errors = append(res.Errors, &educonn_video.Error{
			Description: errors.RawVideoFileS3NotFound.Error(),
			Code: 404,
		})
	    return err
	}
	log.Infof("[S3] key '%s' exists in bucket", fileKey)

	// TODO: insert into DB

	v.videoCreatedPublisher.PublishVideoCreated(&educonn_video.VideoCreatedEvent{
		Video:  req.Video,
		UserId: "TODO",
	})

	res.Video = &educonn_video.VideoDetails{
		Title: "test",
	}

	return nil
}

