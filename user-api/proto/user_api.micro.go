// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user-api/proto/user_api.proto

/*
Package educonn_user_api is a generated protocol buffer package.

It is generated from these files:
	user-api/proto/user_api.proto

It has these top-level messages:
	CreateRequest
	CreateResponse
	DeleteRequest
	DeleteResponse
	LoginRequest
	LoginResponse
	VideoRequest
	VideoResponse
*/
package educonn_user_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lukasjarosch/educonn-platform/user/proto"
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

// Client API for UserApi service

type UserApiClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	// Videos returns all videos of the user
	Videos(ctx context.Context, in *VideoRequest, opts ...client.CallOption) (*VideoResponse, error)
}

type userApiClient struct {
	c           client.Client
	serviceName string
}

func NewUserApiClient(serviceName string, c client.Client) UserApiClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "educonn.user_api"
	}
	return &userApiClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userApiClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserApi.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserApi.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiClient) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserApi.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiClient) Videos(ctx context.Context, in *VideoRequest, opts ...client.CallOption) (*VideoResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserApi.Videos", in)
	out := new(VideoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserApi service

type UserApiHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Login(context.Context, *LoginRequest, *LoginResponse) error
	// Videos returns all videos of the user
	Videos(context.Context, *VideoRequest, *VideoResponse) error
}

func RegisterUserApiHandler(s server.Server, hdlr UserApiHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserApi{hdlr}, opts...))
}

type UserApi struct {
	UserApiHandler
}

func (h *UserApi) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.UserApiHandler.Create(ctx, in, out)
}

func (h *UserApi) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.UserApiHandler.Delete(ctx, in, out)
}

func (h *UserApi) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.UserApiHandler.Login(ctx, in, out)
}

func (h *UserApi) Videos(ctx context.Context, in *VideoRequest, out *VideoResponse) error {
	return h.UserApiHandler.Videos(ctx, in, out)
}
