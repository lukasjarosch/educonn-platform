// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/video/proto/video_api.proto

package educonn_api_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "github.com/lukasjarosch/educonn-platform/video/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateRequest struct {
	Video                *proto1.VideoDetails `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_api_6d03db99d0aa50f2, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetVideo() *proto1.VideoDetails {
	if m != nil {
		return m.Video
	}
	return nil
}

type CreateResponse struct {
	Video                *proto1.VideoDetails `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_api_6d03db99d0aa50f2, []int{1}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (dst *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(dst, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetVideo() *proto1.VideoDetails {
	if m != nil {
		return m.Video
	}
	return nil
}

type DeleteRequest struct {
	VideoId              string   `protobuf:"bytes,1,opt,name=videoId,proto3" json:"videoId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_api_6d03db99d0aa50f2, []int{2}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(dst, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetVideoId() string {
	if m != nil {
		return m.VideoId
	}
	return ""
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_api_6d03db99d0aa50f2, []int{3}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(dst, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type GetRequest struct {
	VideoId              string   `protobuf:"bytes,1,opt,name=videoId,proto3" json:"videoId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_api_6d03db99d0aa50f2, []int{4}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetVideoId() string {
	if m != nil {
		return m.VideoId
	}
	return ""
}

type GetResponse struct {
	Video                *proto1.VideoDetails `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
	SignedURL            string               `protobuf:"bytes,2,opt,name=signedURL,proto3" json:"signedURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_api_6d03db99d0aa50f2, []int{5}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetVideo() *proto1.VideoDetails {
	if m != nil {
		return m.Video
	}
	return nil
}

func (m *GetResponse) GetSignedURL() string {
	if m != nil {
		return m.SignedURL
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "educonn.api.user.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "educonn.api.user.CreateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "educonn.api.user.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "educonn.api.user.DeleteResponse")
	proto.RegisterType((*GetRequest)(nil), "educonn.api.user.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "educonn.api.user.GetResponse")
}

func init() {
	proto.RegisterFile("api/video/proto/video_api.proto", fileDescriptor_video_api_6d03db99d0aa50f2)
}

var fileDescriptor_video_api_6d03db99d0aa50f2 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2c, 0xc8, 0xd4,
	0x2f, 0xcb, 0x4c, 0x49, 0xcd, 0xd7, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x87, 0xb0, 0xe3, 0x13, 0x0b,
	0x32, 0xf5, 0xc0, 0x7c, 0x21, 0x81, 0xd4, 0x94, 0xd2, 0xe4, 0xfc, 0xbc, 0x3c, 0x3d, 0x90, 0x50,
	0x69, 0x71, 0x6a, 0x91, 0x94, 0x38, 0x86, 0x72, 0x88, 0x52, 0x25, 0x27, 0x2e, 0x5e, 0xe7, 0xa2,
	0xd4, 0xc4, 0x92, 0xd4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x43, 0x2e, 0x56, 0xb0,
	0xbc, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xb4, 0x1e, 0xcc, 0x2c, 0x88, 0xae, 0x30, 0x10,
	0xe9, 0x92, 0x5a, 0x92, 0x98, 0x99, 0x53, 0x1c, 0x04, 0x51, 0xa9, 0xe4, 0xcc, 0xc5, 0x07, 0x33,
	0xa3, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x95, 0x1c, 0x43, 0x34, 0xb9, 0x78, 0x5d, 0x52, 0x73, 0x52,
	0x11, 0x0e, 0x91, 0xe0, 0x62, 0x07, 0xcb, 0x78, 0xa6, 0x80, 0x4d, 0xe1, 0x0c, 0x82, 0x71, 0x95,
	0x04, 0xb8, 0xf8, 0x60, 0x4a, 0x21, 0xf6, 0x29, 0xa9, 0x71, 0x71, 0xb9, 0xa7, 0x96, 0x10, 0xd6,
	0x19, 0xc7, 0xc5, 0x0d, 0x56, 0x47, 0xb6, 0x33, 0x85, 0x64, 0xb8, 0x38, 0x8b, 0x33, 0xd3, 0xf3,
	0x52, 0x53, 0x42, 0x83, 0x7c, 0x24, 0x98, 0xc0, 0xa6, 0x23, 0x04, 0x8c, 0xde, 0x31, 0x72, 0x71,
	0x80, 0x75, 0x39, 0x16, 0x64, 0x0a, 0xf9, 0x72, 0xb1, 0x41, 0x82, 0x45, 0x48, 0x5e, 0x0f, 0x3d,
	0x42, 0xf4, 0x50, 0x02, 0x5d, 0x4a, 0x01, 0xb7, 0x02, 0xa8, 0x0f, 0x19, 0x40, 0xc6, 0x41, 0x7c,
	0x8d, 0xcd, 0x38, 0x94, 0xa0, 0xc3, 0x66, 0x1c, 0x5a, 0x80, 0x31, 0x08, 0xb9, 0x70, 0x31, 0xbb,
	0xa7, 0x96, 0x08, 0xc9, 0x60, 0x2a, 0x45, 0x84, 0xa4, 0x94, 0x2c, 0x0e, 0x59, 0x98, 0x29, 0x49,
	0x6c, 0xe0, 0x54, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x2e, 0xf9, 0x63, 0x93, 0x02,
	0x00, 0x00,
}
