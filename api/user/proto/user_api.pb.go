// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/user/proto/user_api.proto

package educonn_api_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "github.com/lukasjarosch/educonn-platform/user/proto"

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
	User                 *proto1.UserDetails `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_api_410802639415ec6b, []int{0}
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

func (m *CreateRequest) GetUser() *proto1.UserDetails {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateResponse struct {
	User                 *proto1.UserDetails `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_api_410802639415ec6b, []int{1}
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

func (m *CreateResponse) GetUser() *proto1.UserDetails {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_api_410802639415ec6b, []int{2}
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

func (m *DeleteRequest) GetUserId() string {
	if m != nil {
		return m.UserId
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
	return fileDescriptor_user_api_410802639415ec6b, []int{3}
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

func init() {
	proto.RegisterType((*CreateRequest)(nil), "educonn.api.user.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "educonn.api.user.CreateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "educonn.api.user.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "educonn.api.user.DeleteResponse")
}

func init() {
	proto.RegisterFile("api/user/proto/user_api.proto", fileDescriptor_user_api_410802639415ec6b)
}

var fileDescriptor_user_api_410802639415ec6b = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2c, 0xc8, 0xd4,
	0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x07, 0x33, 0xe3, 0x13, 0x0b, 0x32,
	0xf5, 0xc0, 0x5c, 0x21, 0x81, 0xd4, 0x94, 0xd2, 0xe4, 0xfc, 0xbc, 0x3c, 0x3d, 0x90, 0x10, 0x48,
	0x4e, 0x4a, 0x14, 0x4d, 0x31, 0x44, 0xa1, 0x92, 0x1d, 0x17, 0xaf, 0x73, 0x51, 0x6a, 0x62, 0x49,
	0x6a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x2e, 0x17, 0x0b, 0x48, 0x5a, 0x82, 0x51,
	0x81, 0x51, 0x83, 0xdb, 0x48, 0x52, 0x0f, 0x66, 0x10, 0x58, 0x4f, 0x68, 0x71, 0x6a, 0x91, 0x4b,
	0x6a, 0x49, 0x62, 0x66, 0x4e, 0x71, 0x10, 0x58, 0x99, 0x92, 0x3d, 0x17, 0x1f, 0x4c, 0x7f, 0x71,
	0x41, 0x7e, 0x5e, 0x71, 0x2a, 0xa9, 0x06, 0xa8, 0x73, 0xf1, 0xba, 0xa4, 0xe6, 0xa4, 0x22, 0x1c,
	0x20, 0xc6, 0xc5, 0x06, 0x92, 0xf0, 0x4c, 0x01, 0x9b, 0xc0, 0x19, 0x04, 0xe5, 0x29, 0x09, 0x70,
	0xf1, 0xc1, 0x14, 0x42, 0x6c, 0x32, 0x5a, 0xce, 0xc8, 0xc5, 0x0e, 0x32, 0xd0, 0xb1, 0x20, 0x53,
	0xc8, 0x97, 0x8b, 0x0d, 0xe2, 0x0e, 0x21, 0x79, 0x3d, 0x74, 0xbf, 0xeb, 0xa1, 0xf8, 0x50, 0x4a,
	0x01, 0xb7, 0x02, 0x88, 0xc1, 0x4a, 0x0c, 0x20, 0xe3, 0x20, 0x96, 0x61, 0x33, 0x0e, 0xc5, 0xbd,
	0xd8, 0x8c, 0x43, 0x75, 0xa7, 0x12, 0x43, 0x12, 0x1b, 0x38, 0xb0, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x6e, 0x01, 0xaa, 0x4f, 0xb6, 0x01, 0x00, 0x00,
}
