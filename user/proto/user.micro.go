// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

/*
Package educonn_user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	UserDetails
	Request
	UserResponse
	Token
	Error
	DeleteRequest
	DeleteResponse
	UserCreatedEvent
	UserDeletedEvent
*/
package educonn_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for User service

type UserClient interface {
	Create(ctx context.Context, in *UserDetails, opts ...client.CallOption) (*UserResponse, error)
	Get(ctx context.Context, in *UserDetails, opts ...client.CallOption) (*UserResponse, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*UserResponse, error)
	Auth(ctx context.Context, in *UserDetails, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type userClient struct {
	c           client.Client
	serviceName string
}

func NewUserClient(serviceName string, c client.Client) UserClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "educonn.user"
	}
	return &userClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userClient) Create(ctx context.Context, in *UserDetails, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "User.Create", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Get(ctx context.Context, in *UserDetails, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "User.Get", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "User.GetAll", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Auth(ctx context.Context, in *UserDetails, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "User.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "User.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "User.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Create(context.Context, *UserDetails, *UserResponse) error
	Get(context.Context, *UserDetails, *UserResponse) error
	GetAll(context.Context, *Request, *UserResponse) error
	Auth(context.Context, *UserDetails, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&User{hdlr}, opts...))
}

type User struct {
	UserHandler
}

func (h *User) Create(ctx context.Context, in *UserDetails, out *UserResponse) error {
	return h.UserHandler.Create(ctx, in, out)
}

func (h *User) Get(ctx context.Context, in *UserDetails, out *UserResponse) error {
	return h.UserHandler.Get(ctx, in, out)
}

func (h *User) GetAll(ctx context.Context, in *Request, out *UserResponse) error {
	return h.UserHandler.GetAll(ctx, in, out)
}

func (h *User) Auth(ctx context.Context, in *UserDetails, out *Token) error {
	return h.UserHandler.Auth(ctx, in, out)
}

func (h *User) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserHandler.ValidateToken(ctx, in, out)
}

func (h *User) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.UserHandler.Delete(ctx, in, out)
}
