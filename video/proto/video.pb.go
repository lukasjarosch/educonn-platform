// Code generated by protoc-gen-go. DO NOT EDIT.
// source: video.proto

package educonn_video

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

// ----------------------------
// VIDEO RESOURCE
// ----------------------------
type VideoDetails struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string            `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Tags                 []string          `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	Status               *VideoStatus      `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Statistics           *VideoStatistics  `protobuf:"bytes,6,opt,name=statistics,proto3" json:"statistics,omitempty"`
	Thumbnails           []*VideoThumbnail `protobuf:"bytes,7,rep,name=thumbnails,proto3" json:"thumbnails,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *VideoDetails) Reset()         { *m = VideoDetails{} }
func (m *VideoDetails) String() string { return proto.CompactTextString(m) }
func (*VideoDetails) ProtoMessage()    {}
func (*VideoDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_52938d9d859f5c4d, []int{0}
}
func (m *VideoDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoDetails.Unmarshal(m, b)
}
func (m *VideoDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoDetails.Marshal(b, m, deterministic)
}
func (dst *VideoDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoDetails.Merge(dst, src)
}
func (m *VideoDetails) XXX_Size() int {
	return xxx_messageInfo_VideoDetails.Size(m)
}
func (m *VideoDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoDetails.DiscardUnknown(m)
}

var xxx_messageInfo_VideoDetails proto.InternalMessageInfo

func (m *VideoDetails) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VideoDetails) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *VideoDetails) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *VideoDetails) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *VideoDetails) GetStatus() *VideoStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VideoDetails) GetStatistics() *VideoStatistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *VideoDetails) GetThumbnails() []*VideoThumbnail {
	if m != nil {
		return m.Thumbnails
	}
	return nil
}

type VideoStatus struct {
	UploadStatus         string   `protobuf:"bytes,1,opt,name=uploadStatus,proto3" json:"uploadStatus,omitempty"`
	FailureReason        string   `protobuf:"bytes,2,opt,name=failureReason,proto3" json:"failureReason,omitempty"`
	RejectionReason      string   `protobuf:"bytes,3,opt,name=rejectionReason,proto3" json:"rejectionReason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoStatus) Reset()         { *m = VideoStatus{} }
func (m *VideoStatus) String() string { return proto.CompactTextString(m) }
func (*VideoStatus) ProtoMessage()    {}
func (*VideoStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_52938d9d859f5c4d, []int{1}
}
func (m *VideoStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoStatus.Unmarshal(m, b)
}
func (m *VideoStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoStatus.Marshal(b, m, deterministic)
}
func (dst *VideoStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoStatus.Merge(dst, src)
}
func (m *VideoStatus) XXX_Size() int {
	return xxx_messageInfo_VideoStatus.Size(m)
}
func (m *VideoStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoStatus.DiscardUnknown(m)
}

var xxx_messageInfo_VideoStatus proto.InternalMessageInfo

func (m *VideoStatus) GetUploadStatus() string {
	if m != nil {
		return m.UploadStatus
	}
	return ""
}

func (m *VideoStatus) GetFailureReason() string {
	if m != nil {
		return m.FailureReason
	}
	return ""
}

func (m *VideoStatus) GetRejectionReason() string {
	if m != nil {
		return m.RejectionReason
	}
	return ""
}

type VideoStatistics struct {
	ViewCount            int64    `protobuf:"varint,1,opt,name=viewCount,proto3" json:"viewCount,omitempty"`
	LikeCount            int64    `protobuf:"varint,2,opt,name=likeCount,proto3" json:"likeCount,omitempty"`
	DislikeCound         int64    `protobuf:"varint,3,opt,name=dislikeCound,proto3" json:"dislikeCound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoStatistics) Reset()         { *m = VideoStatistics{} }
func (m *VideoStatistics) String() string { return proto.CompactTextString(m) }
func (*VideoStatistics) ProtoMessage()    {}
func (*VideoStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_52938d9d859f5c4d, []int{2}
}
func (m *VideoStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoStatistics.Unmarshal(m, b)
}
func (m *VideoStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoStatistics.Marshal(b, m, deterministic)
}
func (dst *VideoStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoStatistics.Merge(dst, src)
}
func (m *VideoStatistics) XXX_Size() int {
	return xxx_messageInfo_VideoStatistics.Size(m)
}
func (m *VideoStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_VideoStatistics proto.InternalMessageInfo

func (m *VideoStatistics) GetViewCount() int64 {
	if m != nil {
		return m.ViewCount
	}
	return 0
}

func (m *VideoStatistics) GetLikeCount() int64 {
	if m != nil {
		return m.LikeCount
	}
	return 0
}

func (m *VideoStatistics) GetDislikeCound() int64 {
	if m != nil {
		return m.DislikeCound
	}
	return 0
}

type VideoThumbnail struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Width                string   `protobuf:"bytes,2,opt,name=width,proto3" json:"width,omitempty"`
	Height               string   `protobuf:"bytes,3,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoThumbnail) Reset()         { *m = VideoThumbnail{} }
func (m *VideoThumbnail) String() string { return proto.CompactTextString(m) }
func (*VideoThumbnail) ProtoMessage()    {}
func (*VideoThumbnail) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_52938d9d859f5c4d, []int{3}
}
func (m *VideoThumbnail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoThumbnail.Unmarshal(m, b)
}
func (m *VideoThumbnail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoThumbnail.Marshal(b, m, deterministic)
}
func (dst *VideoThumbnail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoThumbnail.Merge(dst, src)
}
func (m *VideoThumbnail) XXX_Size() int {
	return xxx_messageInfo_VideoThumbnail.Size(m)
}
func (m *VideoThumbnail) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoThumbnail.DiscardUnknown(m)
}

var xxx_messageInfo_VideoThumbnail proto.InternalMessageInfo

func (m *VideoThumbnail) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *VideoThumbnail) GetWidth() string {
	if m != nil {
		return m.Width
	}
	return ""
}

func (m *VideoThumbnail) GetHeight() string {
	if m != nil {
		return m.Height
	}
	return ""
}

// ----------------------------
// EVENTS
// ----------------------------
type VideoCreatedEvent struct {
	Video                *VideoDetails `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
	UserId               string        `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VideoCreatedEvent) Reset()         { *m = VideoCreatedEvent{} }
func (m *VideoCreatedEvent) String() string { return proto.CompactTextString(m) }
func (*VideoCreatedEvent) ProtoMessage()    {}
func (*VideoCreatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_52938d9d859f5c4d, []int{4}
}
func (m *VideoCreatedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoCreatedEvent.Unmarshal(m, b)
}
func (m *VideoCreatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoCreatedEvent.Marshal(b, m, deterministic)
}
func (dst *VideoCreatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoCreatedEvent.Merge(dst, src)
}
func (m *VideoCreatedEvent) XXX_Size() int {
	return xxx_messageInfo_VideoCreatedEvent.Size(m)
}
func (m *VideoCreatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoCreatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_VideoCreatedEvent proto.InternalMessageInfo

func (m *VideoCreatedEvent) GetVideo() *VideoDetails {
	if m != nil {
		return m.Video
	}
	return nil
}

func (m *VideoCreatedEvent) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type CreateVideoRequest struct {
	Video                *VideoDetails `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateVideoRequest) Reset()         { *m = CreateVideoRequest{} }
func (m *CreateVideoRequest) String() string { return proto.CompactTextString(m) }
func (*CreateVideoRequest) ProtoMessage()    {}
func (*CreateVideoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_52938d9d859f5c4d, []int{5}
}
func (m *CreateVideoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVideoRequest.Unmarshal(m, b)
}
func (m *CreateVideoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVideoRequest.Marshal(b, m, deterministic)
}
func (dst *CreateVideoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVideoRequest.Merge(dst, src)
}
func (m *CreateVideoRequest) XXX_Size() int {
	return xxx_messageInfo_CreateVideoRequest.Size(m)
}
func (m *CreateVideoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVideoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVideoRequest proto.InternalMessageInfo

func (m *CreateVideoRequest) GetVideo() *VideoDetails {
	if m != nil {
		return m.Video
	}
	return nil
}

type CreateVideoResponse struct {
	Video                *VideoDetails `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
	Errors               []*Error      `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateVideoResponse) Reset()         { *m = CreateVideoResponse{} }
func (m *CreateVideoResponse) String() string { return proto.CompactTextString(m) }
func (*CreateVideoResponse) ProtoMessage()    {}
func (*CreateVideoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_52938d9d859f5c4d, []int{6}
}
func (m *CreateVideoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVideoResponse.Unmarshal(m, b)
}
func (m *CreateVideoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVideoResponse.Marshal(b, m, deterministic)
}
func (dst *CreateVideoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVideoResponse.Merge(dst, src)
}
func (m *CreateVideoResponse) XXX_Size() int {
	return xxx_messageInfo_CreateVideoResponse.Size(m)
}
func (m *CreateVideoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVideoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVideoResponse proto.InternalMessageInfo

func (m *CreateVideoResponse) GetVideo() *VideoDetails {
	if m != nil {
		return m.Video
	}
	return nil
}

func (m *CreateVideoResponse) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

// ----------------------------
// MISC
// ----------------------------
type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_52938d9d859f5c4d, []int{7}
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

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*VideoDetails)(nil), "educonn.video.VideoDetails")
	proto.RegisterType((*VideoStatus)(nil), "educonn.video.VideoStatus")
	proto.RegisterType((*VideoStatistics)(nil), "educonn.video.VideoStatistics")
	proto.RegisterType((*VideoThumbnail)(nil), "educonn.video.VideoThumbnail")
	proto.RegisterType((*VideoCreatedEvent)(nil), "educonn.video.VideoCreatedEvent")
	proto.RegisterType((*CreateVideoRequest)(nil), "educonn.video.CreateVideoRequest")
	proto.RegisterType((*CreateVideoResponse)(nil), "educonn.video.CreateVideoResponse")
	proto.RegisterType((*Error)(nil), "educonn.video.Error")
}

func init() { proto.RegisterFile("video.proto", fileDescriptor_video_52938d9d859f5c4d) }

var fileDescriptor_video_52938d9d859f5c4d = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0x26, 0x76, 0xed, 0xaa, 0xe3, 0x3e, 0x60, 0xa8, 0x90, 0x55, 0x1e, 0x0a, 0x2b, 0x0e, 0x39,
	0xa0, 0x48, 0x04, 0x71, 0x84, 0x4b, 0xa9, 0x10, 0x37, 0x58, 0x10, 0x07, 0x0e, 0x48, 0x6e, 0x76,
	0x48, 0x16, 0x5c, 0x6f, 0xea, 0x5d, 0xa7, 0x67, 0x7e, 0x08, 0xff, 0x15, 0xed, 0xec, 0x1a, 0xe2,
	0x50, 0x38, 0xf4, 0x36, 0xf3, 0xcd, 0xf7, 0xcd, 0xcb, 0xb3, 0x86, 0x62, 0xad, 0x15, 0x99, 0xe9,
	0xaa, 0x35, 0xce, 0xe0, 0x01, 0xa9, 0x6e, 0x6e, 0x9a, 0x66, 0xca, 0xa0, 0xf8, 0x99, 0xc0, 0xfe,
	0x27, 0x6f, 0xbd, 0x26, 0x57, 0xe9, 0xda, 0xe2, 0x21, 0x24, 0x5a, 0x95, 0xa3, 0xf1, 0x68, 0xb2,
	0x27, 0x13, 0xad, 0xf0, 0x18, 0x32, 0xa7, 0x5d, 0x4d, 0x65, 0xc2, 0x50, 0x70, 0x70, 0x0c, 0x85,
	0x22, 0x3b, 0x6f, 0xf5, 0xca, 0x69, 0xd3, 0x94, 0x29, 0xc7, 0x36, 0x21, 0x44, 0xd8, 0x71, 0xd5,
	0xc2, 0x96, 0x3b, 0xe3, 0x74, 0xb2, 0x27, 0xd9, 0xc6, 0x19, 0xe4, 0xd6, 0x55, 0xae, 0xb3, 0x65,
	0x36, 0x1e, 0x4d, 0x8a, 0xd9, 0xc9, 0x74, 0xd0, 0xcc, 0x94, 0x1b, 0xf9, 0xc0, 0x0c, 0x19, 0x99,
	0xf8, 0x0a, 0xc0, 0x5b, 0xda, 0x3a, 0x3d, 0xb7, 0x65, 0xce, 0xba, 0x47, 0xff, 0xd2, 0x05, 0x96,
	0xdc, 0x50, 0xe0, 0x4b, 0x00, 0xb7, 0xec, 0x2e, 0xce, 0x1b, 0x3f, 0x5d, 0xb9, 0x3b, 0x4e, 0x27,
	0xc5, 0xec, 0xe1, 0x75, 0xfa, 0x8f, 0x3d, 0x4b, 0x6e, 0x08, 0xc4, 0x8f, 0x11, 0x14, 0x1b, 0x6d,
	0xa1, 0x80, 0xfd, 0x6e, 0x55, 0x9b, 0x4a, 0x05, 0x3f, 0x2e, 0x6a, 0x80, 0xe1, 0x13, 0x38, 0xf8,
	0x5a, 0xe9, 0xba, 0x6b, 0x49, 0x52, 0x65, 0x4d, 0x13, 0x57, 0x37, 0x04, 0x71, 0x02, 0x47, 0x2d,
	0x7d, 0xa3, 0xb9, 0xdf, 0x56, 0xe4, 0x85, 0x35, 0x6e, 0xc3, 0xe2, 0x12, 0x8e, 0xb6, 0x26, 0xc4,
	0x07, 0xb0, 0xb7, 0xd6, 0x74, 0x75, 0x6a, 0xba, 0xc6, 0x71, 0x0f, 0xa9, 0xfc, 0x03, 0xf8, 0x68,
	0xad, 0xbf, 0x53, 0x88, 0x26, 0x21, 0xfa, 0x1b, 0xf0, 0x23, 0x28, 0x6d, 0x7b, 0x5f, 0x71, 0xd5,
	0x54, 0x0e, 0x30, 0xf1, 0x0e, 0x0e, 0x87, 0x4b, 0xc1, 0xdb, 0x90, 0x76, 0x6d, 0x1d, 0xe7, 0xf5,
	0xa6, 0xbf, 0x8c, 0x2b, 0xad, 0xdc, 0xb2, 0xbf, 0x0c, 0x76, 0xf0, 0x1e, 0xe4, 0x4b, 0xd2, 0x8b,
	0xa5, 0x8b, 0xd3, 0x44, 0x4f, 0x7c, 0x81, 0x3b, 0x9c, 0xf1, 0xb4, 0xa5, 0xca, 0x91, 0x3a, 0x5b,
	0x53, 0xe3, 0xf0, 0x19, 0x64, 0xfc, 0x05, 0x38, 0x6d, 0x31, 0xbb, 0x7f, 0xdd, 0x77, 0x89, 0x87,
	0x29, 0x03, 0xd3, 0xe7, 0xef, 0x2c, 0xb5, 0x6f, 0x55, 0x2c, 0x1b, 0x3d, 0xf1, 0x06, 0x30, 0xa4,
	0x66, 0x91, 0xa4, 0xcb, 0x8e, 0xec, 0x4d, 0x0a, 0x88, 0x35, 0xdc, 0x1d, 0x24, 0xb2, 0x2b, 0xd3,
	0x58, 0xba, 0x49, 0xab, 0x4f, 0x21, 0xa7, 0xb6, 0x35, 0xad, 0x2d, 0x13, 0x3e, 0xbb, 0xe3, 0x2d,
	0xcd, 0x99, 0x0f, 0xca, 0xc8, 0x11, 0x2f, 0x20, 0x63, 0xc0, 0xbf, 0x9c, 0xb9, 0x51, 0xc4, 0x85,
	0x32, 0xc9, 0x36, 0x96, 0xb0, 0x7b, 0x41, 0xd6, 0x56, 0x8b, 0xfe, 0x1d, 0xf6, 0xee, 0xec, 0x33,
	0x64, 0x5c, 0x1b, 0xdf, 0x43, 0x1e, 0xfa, 0xc6, 0xc7, 0x5b, 0x75, 0xfe, 0xde, 0xcb, 0x89, 0xf8,
	0x1f, 0x25, 0x4c, 0x2c, 0x6e, 0x9d, 0xe7, 0xfc, 0xcb, 0x78, 0xfe, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0x51, 0x8a, 0x3f, 0xaa, 0x41, 0x04, 0x00, 0x00,
}
