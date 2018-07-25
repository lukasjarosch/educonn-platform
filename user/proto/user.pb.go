// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package educonn_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserDetails struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDetails) Reset()         { *m = UserDetails{} }
func (m *UserDetails) String() string { return proto.CompactTextString(m) }
func (*UserDetails) ProtoMessage()    {}
func (*UserDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_f3b1ebad4557c80e, []int{0}
}
func (m *UserDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDetails.Unmarshal(m, b)
}
func (m *UserDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDetails.Marshal(b, m, deterministic)
}
func (dst *UserDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDetails.Merge(dst, src)
}
func (m *UserDetails) XXX_Size() int {
	return xxx_messageInfo_UserDetails.Size(m)
}
func (m *UserDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDetails.DiscardUnknown(m)
}

var xxx_messageInfo_UserDetails proto.InternalMessageInfo

func (m *UserDetails) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserDetails) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserDetails) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserDetails) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserDetails) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_f3b1ebad4557c80e, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type UserResponse struct {
	User                 *UserDetails   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*UserDetails `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error       `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_f3b1ebad4557c80e, []int{2}
}
func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (dst *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(dst, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUser() *UserDetails {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserResponse) GetUsers() []*UserDetails {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *UserResponse) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_f3b1ebad4557c80e, []int{3}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (dst *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(dst, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_f3b1ebad4557c80e, []int{4}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type DeleteRequest struct {
	User                 *UserDetails `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_f3b1ebad4557c80e, []int{5}
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

func (m *DeleteRequest) GetUser() *UserDetails {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteResponse struct {
	Errors               []*Error `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_f3b1ebad4557c80e, []int{6}
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

func (m *DeleteResponse) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type UserCreatedEvent struct {
	User                 *UserDetails `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UserCreatedEvent) Reset()         { *m = UserCreatedEvent{} }
func (m *UserCreatedEvent) String() string { return proto.CompactTextString(m) }
func (*UserCreatedEvent) ProtoMessage()    {}
func (*UserCreatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_f3b1ebad4557c80e, []int{7}
}
func (m *UserCreatedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCreatedEvent.Unmarshal(m, b)
}
func (m *UserCreatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCreatedEvent.Marshal(b, m, deterministic)
}
func (dst *UserCreatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreatedEvent.Merge(dst, src)
}
func (m *UserCreatedEvent) XXX_Size() int {
	return xxx_messageInfo_UserCreatedEvent.Size(m)
}
func (m *UserCreatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreatedEvent proto.InternalMessageInfo

func (m *UserCreatedEvent) GetUser() *UserDetails {
	if m != nil {
		return m.User
	}
	return nil
}

type UserDeletedEvent struct {
	User                 *UserDetails `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UserDeletedEvent) Reset()         { *m = UserDeletedEvent{} }
func (m *UserDeletedEvent) String() string { return proto.CompactTextString(m) }
func (*UserDeletedEvent) ProtoMessage()    {}
func (*UserDeletedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_f3b1ebad4557c80e, []int{8}
}
func (m *UserDeletedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeletedEvent.Unmarshal(m, b)
}
func (m *UserDeletedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeletedEvent.Marshal(b, m, deterministic)
}
func (dst *UserDeletedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeletedEvent.Merge(dst, src)
}
func (m *UserDeletedEvent) XXX_Size() int {
	return xxx_messageInfo_UserDeletedEvent.Size(m)
}
func (m *UserDeletedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeletedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeletedEvent proto.InternalMessageInfo

func (m *UserDeletedEvent) GetUser() *UserDetails {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*UserDetails)(nil), "educonn.user.UserDetails")
	proto.RegisterType((*Request)(nil), "educonn.user.Request")
	proto.RegisterType((*UserResponse)(nil), "educonn.user.UserResponse")
	proto.RegisterType((*Token)(nil), "educonn.user.Token")
	proto.RegisterType((*Error)(nil), "educonn.user.Error")
	proto.RegisterType((*DeleteRequest)(nil), "educonn.user.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "educonn.user.DeleteResponse")
	proto.RegisterType((*UserCreatedEvent)(nil), "educonn.user.UserCreatedEvent")
	proto.RegisterType((*UserDeletedEvent)(nil), "educonn.user.UserDeletedEvent")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_f3b1ebad4557c80e) }

var fileDescriptor_user_f3b1ebad4557c80e = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6a, 0x14, 0x41,
	0x14, 0x4d, 0x4f, 0x3f, 0xcc, 0xdc, 0x49, 0x82, 0x5c, 0x15, 0xda, 0x31, 0x8b, 0xa1, 0x56, 0x82,
	0x38, 0x42, 0xdc, 0x08, 0x32, 0x81, 0xc1, 0x0c, 0xd9, 0xb9, 0x68, 0xd4, 0xb5, 0xe5, 0xd4, 0x15,
	0x0b, 0x3b, 0x5d, 0x6d, 0x55, 0x75, 0xfc, 0x03, 0xc1, 0x6f, 0xf0, 0x67, 0xa5, 0x1e, 0x6d, 0xa6,
	0xcd, 0x8b, 0x99, 0x55, 0xd7, 0x3d, 0xf7, 0x9c, 0xaa, 0x73, 0x1f, 0x34, 0x40, 0x67, 0x48, 0xcf,
	0x5b, 0xad, 0xac, 0xc2, 0x03, 0x12, 0xdd, 0x5a, 0x35, 0xcd, 0xdc, 0x61, 0xec, 0x77, 0x02, 0x93,
	0x8f, 0x86, 0xf4, 0x19, 0x59, 0x2e, 0x6b, 0x83, 0x47, 0x30, 0x92, 0xa2, 0x4c, 0x66, 0xc9, 0xf3,
	0x71, 0x35, 0x92, 0x02, 0x8f, 0x61, 0xfc, 0x55, 0x6a, 0x63, 0xdf, 0xf3, 0x0b, 0x2a, 0x47, 0x1e,
	0xbe, 0x02, 0x70, 0x0a, 0xfb, 0x35, 0x8f, 0xc9, 0xd4, 0x27, 0xff, 0xc5, 0xf8, 0x18, 0x72, 0xba,
	0xe0, 0xb2, 0x2e, 0x33, 0x9f, 0x08, 0x81, 0x53, 0xb4, 0xdc, 0x98, 0x9f, 0x4a, 0x8b, 0x32, 0x0f,
	0x8a, 0x3e, 0x66, 0x63, 0x78, 0x50, 0xd1, 0x8f, 0x8e, 0x8c, 0x65, 0x7f, 0x12, 0x38, 0x70, 0xb6,
	0x2a, 0x32, 0xad, 0x6a, 0x0c, 0xe1, 0x4b, 0xc8, 0x9c, 0x5f, 0xef, 0x6c, 0x72, 0xf2, 0x74, 0xbe,
	0x59, 0xc4, 0x7c, 0xa3, 0x80, 0xca, 0xd3, 0xf0, 0x15, 0xe4, 0xee, 0x6b, 0xca, 0xd1, 0x2c, 0xbd,
	0x9b, 0x1f, 0x78, 0xf8, 0x02, 0x0a, 0xd2, 0x5a, 0x69, 0x53, 0xa6, 0x5e, 0xf1, 0x68, 0xa8, 0x58,
	0xb9, 0x5c, 0x15, 0x29, 0xec, 0x33, 0xe4, 0x1f, 0xd4, 0x77, 0x6a, 0x5c, 0x8d, 0xd6, 0x1d, 0x62,
	0xc3, 0x42, 0xe0, 0xd0, 0x4b, 0x5e, 0x4b, 0xe1, 0xfb, 0xb5, 0x5f, 0x85, 0x60, 0xbb, 0x17, 0x16,
	0x90, 0x7b, 0x00, 0x11, 0xb2, 0xb5, 0x12, 0xe4, 0x1f, 0xc8, 0x2b, 0x7f, 0xc6, 0x19, 0x4c, 0x04,
	0x99, 0xb5, 0x96, 0xad, 0x95, 0xaa, 0x89, 0x53, 0xd9, 0x84, 0xd8, 0x29, 0x1c, 0x9e, 0x51, 0x4d,
	0x96, 0x62, 0x3f, 0xb7, 0x6c, 0x1f, 0x5b, 0xc0, 0x51, 0xaf, 0x8f, 0xfd, 0xbf, 0x72, 0x9f, 0xdc,
	0xef, 0x7e, 0x09, 0x0f, 0xdd, 0x9d, 0xef, 0x34, 0x71, 0x4b, 0x62, 0x75, 0x49, 0xcd, 0xd6, 0x0e,
	0xe2, 0x15, 0xc1, 0xc5, 0x4e, 0x57, 0x9c, 0xfc, 0x4a, 0x21, 0x73, 0x28, 0x2e, 0xa1, 0x08, 0x56,
	0xf0, 0x76, 0xcd, 0x74, 0x7a, 0x3d, 0xd5, 0x17, 0xcf, 0xf6, 0xf0, 0x14, 0xd2, 0x73, 0xb2, 0xbb,
	0xeb, 0x17, 0x50, 0x9c, 0x93, 0x5d, 0xd6, 0x35, 0x3e, 0x19, 0xf2, 0xe2, 0x80, 0xee, 0x91, 0xbf,
	0x81, 0x6c, 0xd9, 0xd9, 0x6f, 0x77, 0xbd, 0xff, 0xdf, 0x40, 0xfc, 0x7e, 0xb2, 0x3d, 0x7c, 0x0b,
	0x87, 0x9f, 0xdc, 0xfa, 0x71, 0x4b, 0x61, 0x65, 0x6f, 0xe2, 0xdd, 0x26, 0x5e, 0x41, 0x11, 0x06,
	0x80, 0xcf, 0x86, 0x84, 0xc1, 0x72, 0x4d, 0x8f, 0x6f, 0x4e, 0xf6, 0xee, 0xbf, 0x14, 0xfe, 0xc7,
	0xf3, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0xa0, 0x3c, 0x8c, 0x86, 0x04, 0x00, 0x00,
}
