// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: video.proto

/*
Package educonn_video is a generated protocol buffer package.

It is generated from these files:
	video.proto

It has these top-level messages:
	VideoDetails
	VideoStorage
	VideoStatus
	VideoStatistics
	VideoThumbnail
	VideoCreatedEvent
	CreateVideoRequest
	CreateVideoResponse
	Error
*/
package educonn_video

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Video service

type VideoClient interface {
	Create(ctx context.Context, in *CreateVideoRequest, opts ...client.CallOption) (*CreateVideoResponse, error)
}

type videoClient struct {
	c           client.Client
	serviceName string
}

func NewVideoClient(serviceName string, c client.Client) VideoClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "educonn.video"
	}
	return &videoClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *videoClient) Create(ctx context.Context, in *CreateVideoRequest, opts ...client.CallOption) (*CreateVideoResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Video.Create", in)
	out := new(CreateVideoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Video service

type VideoHandler interface {
	Create(context.Context, *CreateVideoRequest, *CreateVideoResponse) error
}

func RegisterVideoHandler(s server.Server, hdlr VideoHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Video{hdlr}, opts...))
}

type Video struct {
	VideoHandler
}

func (h *Video) Create(ctx context.Context, in *CreateVideoRequest, out *CreateVideoResponse) error {
	return h.VideoHandler.Create(ctx, in, out)
}
