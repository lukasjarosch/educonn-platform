// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lesson/proto/lesson.proto

package educonn_lesson

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

type Type int32

const (
	Type_VIDEO Type = 0
	Type_TEXT  Type = 1
)

var Type_name = map[int32]string{
	0: "VIDEO",
	1: "TEXT",
}
var Type_value = map[string]int32{
	"VIDEO": 0,
	"TEXT":  1,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_lesson_8fc30f429e5170d7, []int{0}
}

// ----------------------------
// LESSON
// ----------------------------
type Lesson struct {
	Base                 *LessonBase       `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Stats                *LessonStatistics `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
	Video                *VideoLesson      `protobuf:"bytes,3,opt,name=video,proto3" json:"video,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Lesson) Reset()         { *m = Lesson{} }
func (m *Lesson) String() string { return proto.CompactTextString(m) }
func (*Lesson) ProtoMessage()    {}
func (*Lesson) Descriptor() ([]byte, []int) {
	return fileDescriptor_lesson_8fc30f429e5170d7, []int{0}
}
func (m *Lesson) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lesson.Unmarshal(m, b)
}
func (m *Lesson) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lesson.Marshal(b, m, deterministic)
}
func (dst *Lesson) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lesson.Merge(dst, src)
}
func (m *Lesson) XXX_Size() int {
	return xxx_messageInfo_Lesson.Size(m)
}
func (m *Lesson) XXX_DiscardUnknown() {
	xxx_messageInfo_Lesson.DiscardUnknown(m)
}

var xxx_messageInfo_Lesson proto.InternalMessageInfo

func (m *Lesson) GetBase() *LessonBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *Lesson) GetStats() *LessonStatistics {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *Lesson) GetVideo() *VideoLesson {
	if m != nil {
		return m.Video
	}
	return nil
}

type LessonBase struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type                 Type     `protobuf:"varint,4,opt,name=type,proto3,enum=educonn.lesson.Type" json:"type,omitempty"`
	UserId               string   `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LessonBase) Reset()         { *m = LessonBase{} }
func (m *LessonBase) String() string { return proto.CompactTextString(m) }
func (*LessonBase) ProtoMessage()    {}
func (*LessonBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_lesson_8fc30f429e5170d7, []int{1}
}
func (m *LessonBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LessonBase.Unmarshal(m, b)
}
func (m *LessonBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LessonBase.Marshal(b, m, deterministic)
}
func (dst *LessonBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LessonBase.Merge(dst, src)
}
func (m *LessonBase) XXX_Size() int {
	return xxx_messageInfo_LessonBase.Size(m)
}
func (m *LessonBase) XXX_DiscardUnknown() {
	xxx_messageInfo_LessonBase.DiscardUnknown(m)
}

var xxx_messageInfo_LessonBase proto.InternalMessageInfo

func (m *LessonBase) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LessonBase) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LessonBase) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LessonBase) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_VIDEO
}

func (m *LessonBase) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type LessonStatistics struct {
	Likes                int64    `protobuf:"varint,1,opt,name=likes,proto3" json:"likes,omitempty"`
	Dislikes             int64    `protobuf:"varint,2,opt,name=dislikes,proto3" json:"dislikes,omitempty"`
	Views                int64    `protobuf:"varint,3,opt,name=views,proto3" json:"views,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LessonStatistics) Reset()         { *m = LessonStatistics{} }
func (m *LessonStatistics) String() string { return proto.CompactTextString(m) }
func (*LessonStatistics) ProtoMessage()    {}
func (*LessonStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_lesson_8fc30f429e5170d7, []int{2}
}
func (m *LessonStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LessonStatistics.Unmarshal(m, b)
}
func (m *LessonStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LessonStatistics.Marshal(b, m, deterministic)
}
func (dst *LessonStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LessonStatistics.Merge(dst, src)
}
func (m *LessonStatistics) XXX_Size() int {
	return xxx_messageInfo_LessonStatistics.Size(m)
}
func (m *LessonStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_LessonStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_LessonStatistics proto.InternalMessageInfo

func (m *LessonStatistics) GetLikes() int64 {
	if m != nil {
		return m.Likes
	}
	return 0
}

func (m *LessonStatistics) GetDislikes() int64 {
	if m != nil {
		return m.Dislikes
	}
	return 0
}

func (m *LessonStatistics) GetViews() int64 {
	if m != nil {
		return m.Views
	}
	return 0
}

// Requests & Responses
type CreateLessonRequest struct {
	Lesson               *Lesson  `protobuf:"bytes,1,opt,name=lesson,proto3" json:"lesson,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLessonRequest) Reset()         { *m = CreateLessonRequest{} }
func (m *CreateLessonRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLessonRequest) ProtoMessage()    {}
func (*CreateLessonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lesson_8fc30f429e5170d7, []int{3}
}
func (m *CreateLessonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLessonRequest.Unmarshal(m, b)
}
func (m *CreateLessonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLessonRequest.Marshal(b, m, deterministic)
}
func (dst *CreateLessonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLessonRequest.Merge(dst, src)
}
func (m *CreateLessonRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLessonRequest.Size(m)
}
func (m *CreateLessonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLessonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLessonRequest proto.InternalMessageInfo

func (m *CreateLessonRequest) GetLesson() *Lesson {
	if m != nil {
		return m.Lesson
	}
	return nil
}

type CreateLessonResponse struct {
	Lesson               *Lesson  `protobuf:"bytes,1,opt,name=lesson,proto3" json:"lesson,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLessonResponse) Reset()         { *m = CreateLessonResponse{} }
func (m *CreateLessonResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLessonResponse) ProtoMessage()    {}
func (*CreateLessonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_lesson_8fc30f429e5170d7, []int{4}
}
func (m *CreateLessonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLessonResponse.Unmarshal(m, b)
}
func (m *CreateLessonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLessonResponse.Marshal(b, m, deterministic)
}
func (dst *CreateLessonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLessonResponse.Merge(dst, src)
}
func (m *CreateLessonResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLessonResponse.Size(m)
}
func (m *CreateLessonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLessonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLessonResponse proto.InternalMessageInfo

func (m *CreateLessonResponse) GetLesson() *Lesson {
	if m != nil {
		return m.Lesson
	}
	return nil
}

type GetLessonByIdRequest struct {
	LessonId             string   `protobuf:"bytes,1,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLessonByIdRequest) Reset()         { *m = GetLessonByIdRequest{} }
func (m *GetLessonByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetLessonByIdRequest) ProtoMessage()    {}
func (*GetLessonByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lesson_8fc30f429e5170d7, []int{5}
}
func (m *GetLessonByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLessonByIdRequest.Unmarshal(m, b)
}
func (m *GetLessonByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLessonByIdRequest.Marshal(b, m, deterministic)
}
func (dst *GetLessonByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLessonByIdRequest.Merge(dst, src)
}
func (m *GetLessonByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetLessonByIdRequest.Size(m)
}
func (m *GetLessonByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLessonByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLessonByIdRequest proto.InternalMessageInfo

func (m *GetLessonByIdRequest) GetLessonId() string {
	if m != nil {
		return m.LessonId
	}
	return ""
}

type GetLessonByIdResponse struct {
	Lesson               *Lesson  `protobuf:"bytes,1,opt,name=lesson,proto3" json:"lesson,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLessonByIdResponse) Reset()         { *m = GetLessonByIdResponse{} }
func (m *GetLessonByIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetLessonByIdResponse) ProtoMessage()    {}
func (*GetLessonByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_lesson_8fc30f429e5170d7, []int{6}
}
func (m *GetLessonByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLessonByIdResponse.Unmarshal(m, b)
}
func (m *GetLessonByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLessonByIdResponse.Marshal(b, m, deterministic)
}
func (dst *GetLessonByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLessonByIdResponse.Merge(dst, src)
}
func (m *GetLessonByIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetLessonByIdResponse.Size(m)
}
func (m *GetLessonByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLessonByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLessonByIdResponse proto.InternalMessageInfo

func (m *GetLessonByIdResponse) GetLesson() *Lesson {
	if m != nil {
		return m.Lesson
	}
	return nil
}

// EVENTS
type LessonCreatedEvent struct {
	Lesson               *Lesson  `protobuf:"bytes,1,opt,name=lesson,proto3" json:"lesson,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LessonCreatedEvent) Reset()         { *m = LessonCreatedEvent{} }
func (m *LessonCreatedEvent) String() string { return proto.CompactTextString(m) }
func (*LessonCreatedEvent) ProtoMessage()    {}
func (*LessonCreatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_lesson_8fc30f429e5170d7, []int{7}
}
func (m *LessonCreatedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LessonCreatedEvent.Unmarshal(m, b)
}
func (m *LessonCreatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LessonCreatedEvent.Marshal(b, m, deterministic)
}
func (dst *LessonCreatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LessonCreatedEvent.Merge(dst, src)
}
func (m *LessonCreatedEvent) XXX_Size() int {
	return xxx_messageInfo_LessonCreatedEvent.Size(m)
}
func (m *LessonCreatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LessonCreatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LessonCreatedEvent proto.InternalMessageInfo

func (m *LessonCreatedEvent) GetLesson() *Lesson {
	if m != nil {
		return m.Lesson
	}
	return nil
}

// ----------------------------
// VIDEO-LESSON (Type: 0)
// ----------------------------
type VideoLesson struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VideoId              string   `protobuf:"bytes,2,opt,name=videoId,proto3" json:"videoId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoLesson) Reset()         { *m = VideoLesson{} }
func (m *VideoLesson) String() string { return proto.CompactTextString(m) }
func (*VideoLesson) ProtoMessage()    {}
func (*VideoLesson) Descriptor() ([]byte, []int) {
	return fileDescriptor_lesson_8fc30f429e5170d7, []int{8}
}
func (m *VideoLesson) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoLesson.Unmarshal(m, b)
}
func (m *VideoLesson) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoLesson.Marshal(b, m, deterministic)
}
func (dst *VideoLesson) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoLesson.Merge(dst, src)
}
func (m *VideoLesson) XXX_Size() int {
	return xxx_messageInfo_VideoLesson.Size(m)
}
func (m *VideoLesson) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoLesson.DiscardUnknown(m)
}

var xxx_messageInfo_VideoLesson proto.InternalMessageInfo

func (m *VideoLesson) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VideoLesson) GetVideoId() string {
	if m != nil {
		return m.VideoId
	}
	return ""
}

// Requests & Responses
type CreateVideoLessonRequest struct {
	Lesson               *VideoLesson `protobuf:"bytes,1,opt,name=lesson,proto3" json:"lesson,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateVideoLessonRequest) Reset()         { *m = CreateVideoLessonRequest{} }
func (m *CreateVideoLessonRequest) String() string { return proto.CompactTextString(m) }
func (*CreateVideoLessonRequest) ProtoMessage()    {}
func (*CreateVideoLessonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lesson_8fc30f429e5170d7, []int{9}
}
func (m *CreateVideoLessonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVideoLessonRequest.Unmarshal(m, b)
}
func (m *CreateVideoLessonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVideoLessonRequest.Marshal(b, m, deterministic)
}
func (dst *CreateVideoLessonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVideoLessonRequest.Merge(dst, src)
}
func (m *CreateVideoLessonRequest) XXX_Size() int {
	return xxx_messageInfo_CreateVideoLessonRequest.Size(m)
}
func (m *CreateVideoLessonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVideoLessonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVideoLessonRequest proto.InternalMessageInfo

func (m *CreateVideoLessonRequest) GetLesson() *VideoLesson {
	if m != nil {
		return m.Lesson
	}
	return nil
}

type CreateVideoLessonResponse struct {
	Lesson               *VideoLesson `protobuf:"bytes,1,opt,name=lesson,proto3" json:"lesson,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateVideoLessonResponse) Reset()         { *m = CreateVideoLessonResponse{} }
func (m *CreateVideoLessonResponse) String() string { return proto.CompactTextString(m) }
func (*CreateVideoLessonResponse) ProtoMessage()    {}
func (*CreateVideoLessonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_lesson_8fc30f429e5170d7, []int{10}
}
func (m *CreateVideoLessonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVideoLessonResponse.Unmarshal(m, b)
}
func (m *CreateVideoLessonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVideoLessonResponse.Marshal(b, m, deterministic)
}
func (dst *CreateVideoLessonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVideoLessonResponse.Merge(dst, src)
}
func (m *CreateVideoLessonResponse) XXX_Size() int {
	return xxx_messageInfo_CreateVideoLessonResponse.Size(m)
}
func (m *CreateVideoLessonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVideoLessonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVideoLessonResponse proto.InternalMessageInfo

func (m *CreateVideoLessonResponse) GetLesson() *VideoLesson {
	if m != nil {
		return m.Lesson
	}
	return nil
}

type GetVideoLessonByIdRequest struct {
	LessonId             string   `protobuf:"bytes,1,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVideoLessonByIdRequest) Reset()         { *m = GetVideoLessonByIdRequest{} }
func (m *GetVideoLessonByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetVideoLessonByIdRequest) ProtoMessage()    {}
func (*GetVideoLessonByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lesson_8fc30f429e5170d7, []int{11}
}
func (m *GetVideoLessonByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVideoLessonByIdRequest.Unmarshal(m, b)
}
func (m *GetVideoLessonByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVideoLessonByIdRequest.Marshal(b, m, deterministic)
}
func (dst *GetVideoLessonByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVideoLessonByIdRequest.Merge(dst, src)
}
func (m *GetVideoLessonByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetVideoLessonByIdRequest.Size(m)
}
func (m *GetVideoLessonByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVideoLessonByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVideoLessonByIdRequest proto.InternalMessageInfo

func (m *GetVideoLessonByIdRequest) GetLessonId() string {
	if m != nil {
		return m.LessonId
	}
	return ""
}

type GetVideoLessonByIdResponse struct {
	Lesson               *VideoLesson `protobuf:"bytes,1,opt,name=lesson,proto3" json:"lesson,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetVideoLessonByIdResponse) Reset()         { *m = GetVideoLessonByIdResponse{} }
func (m *GetVideoLessonByIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetVideoLessonByIdResponse) ProtoMessage()    {}
func (*GetVideoLessonByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_lesson_8fc30f429e5170d7, []int{12}
}
func (m *GetVideoLessonByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVideoLessonByIdResponse.Unmarshal(m, b)
}
func (m *GetVideoLessonByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVideoLessonByIdResponse.Marshal(b, m, deterministic)
}
func (dst *GetVideoLessonByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVideoLessonByIdResponse.Merge(dst, src)
}
func (m *GetVideoLessonByIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetVideoLessonByIdResponse.Size(m)
}
func (m *GetVideoLessonByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVideoLessonByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVideoLessonByIdResponse proto.InternalMessageInfo

func (m *GetVideoLessonByIdResponse) GetLesson() *VideoLesson {
	if m != nil {
		return m.Lesson
	}
	return nil
}

// ----------------------------
// TEXT-LESSON (Type: 1)
// ----------------------------
type TextLesson struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextLesson) Reset()         { *m = TextLesson{} }
func (m *TextLesson) String() string { return proto.CompactTextString(m) }
func (*TextLesson) ProtoMessage()    {}
func (*TextLesson) Descriptor() ([]byte, []int) {
	return fileDescriptor_lesson_8fc30f429e5170d7, []int{13}
}
func (m *TextLesson) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextLesson.Unmarshal(m, b)
}
func (m *TextLesson) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextLesson.Marshal(b, m, deterministic)
}
func (dst *TextLesson) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextLesson.Merge(dst, src)
}
func (m *TextLesson) XXX_Size() int {
	return xxx_messageInfo_TextLesson.Size(m)
}
func (m *TextLesson) XXX_DiscardUnknown() {
	xxx_messageInfo_TextLesson.DiscardUnknown(m)
}

var xxx_messageInfo_TextLesson proto.InternalMessageInfo

func (m *TextLesson) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Lesson)(nil), "educonn.lesson.Lesson")
	proto.RegisterType((*LessonBase)(nil), "educonn.lesson.LessonBase")
	proto.RegisterType((*LessonStatistics)(nil), "educonn.lesson.LessonStatistics")
	proto.RegisterType((*CreateLessonRequest)(nil), "educonn.lesson.CreateLessonRequest")
	proto.RegisterType((*CreateLessonResponse)(nil), "educonn.lesson.CreateLessonResponse")
	proto.RegisterType((*GetLessonByIdRequest)(nil), "educonn.lesson.GetLessonByIdRequest")
	proto.RegisterType((*GetLessonByIdResponse)(nil), "educonn.lesson.GetLessonByIdResponse")
	proto.RegisterType((*LessonCreatedEvent)(nil), "educonn.lesson.LessonCreatedEvent")
	proto.RegisterType((*VideoLesson)(nil), "educonn.lesson.VideoLesson")
	proto.RegisterType((*CreateVideoLessonRequest)(nil), "educonn.lesson.CreateVideoLessonRequest")
	proto.RegisterType((*CreateVideoLessonResponse)(nil), "educonn.lesson.CreateVideoLessonResponse")
	proto.RegisterType((*GetVideoLessonByIdRequest)(nil), "educonn.lesson.GetVideoLessonByIdRequest")
	proto.RegisterType((*GetVideoLessonByIdResponse)(nil), "educonn.lesson.GetVideoLessonByIdResponse")
	proto.RegisterType((*TextLesson)(nil), "educonn.lesson.TextLesson")
	proto.RegisterEnum("educonn.lesson.Type", Type_name, Type_value)
}

func init() { proto.RegisterFile("lesson/proto/lesson.proto", fileDescriptor_lesson_8fc30f429e5170d7) }

var fileDescriptor_lesson_8fc30f429e5170d7 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x51, 0x6f, 0x12, 0x41,
	0x10, 0xee, 0xd1, 0x83, 0xc2, 0x10, 0x09, 0x19, 0xb1, 0x39, 0x0e, 0x1f, 0xc8, 0xa9, 0x09, 0xed,
	0xc3, 0x35, 0x42, 0xa2, 0x3e, 0x6b, 0x91, 0x90, 0x98, 0x54, 0x57, 0x6c, 0x1a, 0x5f, 0x9a, 0x83,
	0x9d, 0x87, 0x8d, 0xf5, 0x0e, 0xd9, 0x05, 0xe5, 0x7f, 0xf8, 0x0f, 0xfc, 0x23, 0xfe, 0x0d, 0xff,
	0x8d, 0x61, 0xf7, 0x8e, 0xd2, 0x63, 0x53, 0x69, 0x7d, 0xbb, 0xd9, 0xfd, 0xbe, 0x99, 0x6f, 0x66,
	0xbe, 0xcd, 0x41, 0xf3, 0x8a, 0xa4, 0x4c, 0xe2, 0x93, 0xe9, 0x2c, 0x51, 0xc9, 0x89, 0x09, 0x42,
	0x1d, 0x60, 0x8d, 0xf8, 0x7c, 0x92, 0xc4, 0x71, 0x68, 0x4e, 0x83, 0x5f, 0x0e, 0x94, 0xde, 0xe9,
	0x4f, 0x0c, 0xc1, 0x1d, 0x47, 0x92, 0x3c, 0xa7, 0xed, 0x74, 0xaa, 0x5d, 0x3f, 0xbc, 0x89, 0x0c,
	0x0d, 0xea, 0x75, 0x24, 0x89, 0x69, 0x1c, 0xbe, 0x80, 0xa2, 0x54, 0x91, 0x92, 0x5e, 0x41, 0x13,
	0xda, 0x76, 0xc2, 0x47, 0x15, 0x29, 0x21, 0x95, 0x98, 0x48, 0x66, 0xe0, 0xf8, 0x1c, 0x8a, 0x0b,
	0xc1, 0x29, 0xf1, 0xf6, 0x35, 0xaf, 0x95, 0xe7, 0x9d, 0xaf, 0x2e, 0x0d, 0x99, 0x19, 0x64, 0xf0,
	0xd3, 0x01, 0xb8, 0xae, 0x8f, 0x35, 0x28, 0x08, 0xae, 0x75, 0x56, 0x58, 0x41, 0x70, 0x44, 0x70,
	0xe3, 0xe8, 0x2b, 0x69, 0x21, 0x15, 0xa6, 0xbf, 0xb1, 0x0d, 0x55, 0x4e, 0x72, 0x32, 0x13, 0x53,
	0x25, 0x92, 0x58, 0xd7, 0xaa, 0xb0, 0xcd, 0x23, 0xec, 0x80, 0xab, 0x96, 0x53, 0xf2, 0xdc, 0xb6,
	0xd3, 0xa9, 0x75, 0x1b, 0x79, 0x19, 0xa3, 0xe5, 0x94, 0x98, 0x46, 0xe0, 0x21, 0x94, 0xe6, 0x92,
	0x66, 0x43, 0xee, 0x15, 0x75, 0x9a, 0x34, 0x0a, 0x3e, 0x43, 0x3d, 0xdf, 0x24, 0x36, 0xa0, 0x78,
	0x25, 0xbe, 0x90, 0xd4, 0xf2, 0xf6, 0x99, 0x09, 0xd0, 0x87, 0x32, 0x17, 0xd2, 0x5c, 0x14, 0xf4,
	0xc5, 0x3a, 0x5e, 0x31, 0x16, 0x82, 0xbe, 0x4b, 0xad, 0x71, 0x9f, 0x99, 0x20, 0xe8, 0xc3, 0xc3,
	0x37, 0x33, 0x8a, 0x14, 0xa5, 0x93, 0xa0, 0x6f, 0x73, 0x92, 0x0a, 0x43, 0x28, 0x19, 0x7d, 0xe9,
	0x9a, 0x0e, 0xed, 0x53, 0x67, 0x29, 0x2a, 0x78, 0x0b, 0x8d, 0x9b, 0x69, 0xe4, 0x34, 0x89, 0x25,
	0xdd, 0x39, 0x4f, 0x0f, 0x1a, 0x03, 0x52, 0xe9, 0x0e, 0x96, 0x43, 0x9e, 0xe9, 0x69, 0x41, 0xc5,
	0x20, 0x2e, 0xd7, 0x1b, 0x29, 0x9b, 0x83, 0x21, 0x0f, 0x06, 0xf0, 0x28, 0x47, 0xba, 0x67, 0xf5,
	0x53, 0x40, 0x73, 0x62, 0x7a, 0xe1, 0xfd, 0x05, 0xc5, 0x77, 0x9f, 0xc5, 0x4b, 0xa8, 0x6e, 0x78,
	0x6b, 0xcb, 0x45, 0x1e, 0x1c, 0x68, 0xb7, 0x0d, 0x79, 0x6a, 0xa4, 0x2c, 0x0c, 0xce, 0xc0, 0x33,
	0x85, 0x37, 0xad, 0x99, 0x0e, 0xa0, 0x97, 0x13, 0x71, 0xab, 0x9d, 0x33, 0x25, 0xef, 0xa1, 0x69,
	0x49, 0x98, 0x0e, 0xe7, 0x5e, 0x19, 0x5f, 0x41, 0x73, 0x40, 0x6a, 0xe3, 0x66, 0xe7, 0x25, 0x7d,
	0x00, 0xdf, 0xc6, 0xfc, 0x1f, 0x31, 0x8f, 0x01, 0x46, 0xf4, 0x43, 0xd9, 0xe7, 0x7c, 0xdc, 0x02,
	0x77, 0xf5, 0xb6, 0xb0, 0x02, 0xc5, 0xf3, 0xe1, 0x69, 0xff, 0xac, 0xbe, 0x87, 0x65, 0x70, 0x47,
	0xfd, 0x8b, 0x51, 0xdd, 0xe9, 0xfe, 0x76, 0xe0, 0x41, 0xfa, 0xa6, 0x68, 0xb6, 0x10, 0x13, 0xc2,
	0x4f, 0x50, 0x32, 0xb3, 0xc2, 0x27, 0xf9, 0xda, 0x96, 0x07, 0xe2, 0x3f, 0xbd, 0x1d, 0x64, 0xda,
	0x0a, 0xf6, 0xf0, 0x02, 0x0e, 0x06, 0xa4, 0x56, 0xbd, 0xe2, 0x16, 0xc5, 0xe6, 0x74, 0xff, 0xd9,
	0x3f, 0x50, 0x59, 0xe6, 0xee, 0x1f, 0x07, 0x70, 0x63, 0x2a, 0x59, 0x1f, 0x97, 0xeb, 0x3e, 0x3a,
	0x76, 0x89, 0xdb, 0xe6, 0xf2, 0x8f, 0x76, 0x40, 0xae, 0x3b, 0x1a, 0x5f, 0x77, 0x74, 0x64, 0xd1,
	0x6a, 0xf7, 0x86, 0x7f, 0xbc, 0x0b, 0x34, 0xab, 0x31, 0x2e, 0xe9, 0xbf, 0x48, 0xef, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x89, 0xfd, 0xb1, 0x12, 0x62, 0x06, 0x00, 0x00,
}
