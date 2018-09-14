package service

import (
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	pb "github.com/lukasjarosch/educonn-platform/video-api/proto"
	merr "github.com/micro/go-micro/errors"
	"context"
	"github.com/lukasjarosch/educonn-platform/video-api/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/video-api/internal/platform/errors"
	"github.com/micro/go-micro/metadata"
	"github.com/rs/zerolog/log"
)

type Video struct {
	videoClient pbVideo.VideoClient
	jwtService *jwt_handler.JwtTokenHandler
}

func NewVideoApi(videoClient pbVideo.VideoClient, jwtService *jwt_handler.JwtTokenHandler) *Video {
	return  &Video{videoClient:videoClient, jwtService:jwtService}
}

func (v *Video) Create(ctx context.Context, req *pb.CreateRequest, res *pb.CreateResponse) error {
	md, _ := metadata.FromContext(ctx)

	// check for token
	token, err := v.jwtService.GetBearerToken(md)
	if err != nil {
		log.Error().Interface("err", err).Msg("asdf")
		return merr.Unauthorized(config.ServiceName, "%s", err.Error())
	}

	// validate jwt
	_, err = v.jwtService.Decode(token)
	if err != nil {
		log.Error().Interface("error", err).Msg(errors.InvalidJWTToken.Error())
		return merr.Unauthorized(config.ServiceName, "%s", errors.InvalidJWTToken)
	}

	videoRequest := &pbVideo.CreateVideoRequest{
		Video: req.Video,
	}
	videoResponse, err := v.videoClient.Create(ctx, videoRequest)
	if err != nil {
	    return err
	}

	res.Video = videoResponse.Video

	return nil
}

func (v *Video) Delete(ctx context.Context, req *pb.DeleteRequest, res *pb.DeleteResponse) error {
	md, _ := metadata.FromContext(ctx)

	// check for token
	token, err := v.jwtService.GetBearerToken(md)
	if err != nil {
		log.Error().Interface("err", err).Msg("asdf")
		return merr.Unauthorized(config.ServiceName, "%s", err.Error())
	}

	// validate jwt
	_, err = v.jwtService.Decode(token)
	if err != nil {
		log.Error().Interface("error", err).Msg(errors.InvalidJWTToken.Error())
		return merr.Unauthorized(config.ServiceName, "%s", errors.InvalidJWTToken)
	}

	// check for video id
	if req.VideoId == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.MissingVideoId)
	}

	err = errors.Error("'educonn.srv.video' does not yet implement this call")
	log.Error().Err(err)
	return merr.InternalServerError(config.ServiceName, "%s", err.Error())

	return nil
}

// Get retrieves a video by ID and returns the video as well as a SignedURL vor the file
//
// The signed Url is only valid for a few minutes. That's because the client only needs to embed it.
// Else we would need to make all videos public by default in the S3 bucket.
func (v *Video) Get(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {

	if req.VideoId == "" {
		return merr.BadRequest(config.ServiceName, "%s", errors.MissingVideoId)
	}

	getRequest := &pbVideo.GetVideoRequest{
		Id: req.VideoId,
	}

	video, err := v.videoClient.GetById(ctx, getRequest)
	if err != nil {
	    return err
	}

	res.Video = video.Video
	res.SignedURL = video.SignedUrl

	return nil
}
