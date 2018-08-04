// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api/video/proto/video_api.proto

/*
Package educonn_api_user is a generated protocol buffer package.

It is generated from these files:
	api/video/proto/video_api.proto

It has these top-level messages:
	CreateRequest
	CreateResponse
	DeleteRequest
	DeleteResponse
	GetRequest
	GetResponse
*/
package educonn_api_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lukasjarosch/educonn-platform/video/proto"

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

// Client API for VideoApi service

type VideoApiClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
}

type videoApiClient struct {
	c           client.Client
	serviceName string
}

func NewVideoApiClient(serviceName string, c client.Client) VideoApiClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "educonn.api.user"
	}
	return &videoApiClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *videoApiClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "VideoApi.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoApiClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "VideoApi.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoApiClient) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.serviceName, "VideoApi.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VideoApi service

type VideoApiHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Get(context.Context, *GetRequest, *GetResponse) error
}

func RegisterVideoApiHandler(s server.Server, hdlr VideoApiHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VideoApi{hdlr}, opts...))
}

type VideoApi struct {
	VideoApiHandler
}

func (h *VideoApi) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.VideoApiHandler.Create(ctx, in, out)
}

func (h *VideoApi) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.VideoApiHandler.Delete(ctx, in, out)
}

func (h *VideoApi) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.VideoApiHandler.Get(ctx, in, out)
}
