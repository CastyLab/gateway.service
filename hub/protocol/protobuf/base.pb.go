// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package protobuf

import (
	enums "movie.night.ws.server/hub/protocol/protobuf/enums"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
	messages "gitlab.com/movienight1/grpc.proto/messages"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CMsgProtoBufHeader struct {
	UserHash             []byte   `protobuf:"bytes,1,opt,name=UserHash,proto3" json:"UserHash,omitempty"`
	IpAddr               []byte   `protobuf:"bytes,2,opt,name=IpAddr,proto3" json:"IpAddr,omitempty"`
	SessionClientId      int32    `protobuf:"varint,3,opt,name=SessionClientId,proto3" json:"SessionClientId,omitempty"`
	AuthToken            []byte   `protobuf:"bytes,4,opt,name=AuthToken,proto3" json:"AuthToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMsgProtoBufHeader) Reset()         { *m = CMsgProtoBufHeader{} }
func (m *CMsgProtoBufHeader) String() string { return proto.CompactTextString(m) }
func (*CMsgProtoBufHeader) ProtoMessage()    {}
func (*CMsgProtoBufHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}

func (m *CMsgProtoBufHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgProtoBufHeader.Unmarshal(m, b)
}
func (m *CMsgProtoBufHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgProtoBufHeader.Marshal(b, m, deterministic)
}
func (m *CMsgProtoBufHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgProtoBufHeader.Merge(m, src)
}
func (m *CMsgProtoBufHeader) XXX_Size() int {
	return xxx_messageInfo_CMsgProtoBufHeader.Size(m)
}
func (m *CMsgProtoBufHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgProtoBufHeader.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgProtoBufHeader proto.InternalMessageInfo

func (m *CMsgProtoBufHeader) GetUserHash() []byte {
	if m != nil {
		return m.UserHash
	}
	return nil
}

func (m *CMsgProtoBufHeader) GetIpAddr() []byte {
	if m != nil {
		return m.IpAddr
	}
	return nil
}

func (m *CMsgProtoBufHeader) GetSessionClientId() int32 {
	if m != nil {
		return m.SessionClientId
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetAuthToken() []byte {
	if m != nil {
		return m.AuthToken
	}
	return nil
}

type PersonalStateMsgEvent struct {
	User                 *messages.User                 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	State                enums.EMSG_PERSONAL_STATE      `protobuf:"varint,2,opt,name=state,proto3,enum=enums.EMSG_PERSONAL_STATE" json:"state,omitempty"`
	Activity             *PersonalStateActivityMsgEvent `protobuf:"bytes,3,opt,name=activity,proto3" json:"activity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *PersonalStateMsgEvent) Reset()         { *m = PersonalStateMsgEvent{} }
func (m *PersonalStateMsgEvent) String() string { return proto.CompactTextString(m) }
func (*PersonalStateMsgEvent) ProtoMessage()    {}
func (*PersonalStateMsgEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{1}
}

func (m *PersonalStateMsgEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonalStateMsgEvent.Unmarshal(m, b)
}
func (m *PersonalStateMsgEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonalStateMsgEvent.Marshal(b, m, deterministic)
}
func (m *PersonalStateMsgEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonalStateMsgEvent.Merge(m, src)
}
func (m *PersonalStateMsgEvent) XXX_Size() int {
	return xxx_messageInfo_PersonalStateMsgEvent.Size(m)
}
func (m *PersonalStateMsgEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonalStateMsgEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PersonalStateMsgEvent proto.InternalMessageInfo

func (m *PersonalStateMsgEvent) GetUser() *messages.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *PersonalStateMsgEvent) GetState() enums.EMSG_PERSONAL_STATE {
	if m != nil {
		return m.State
	}
	return enums.EMSG_PERSONAL_STATE_OFFLINE
}

func (m *PersonalStateMsgEvent) GetActivity() *PersonalStateActivityMsgEvent {
	if m != nil {
		return m.Activity
	}
	return nil
}

type ChatMsgEvent struct {
	Message              []byte               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	From                 string               `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string               `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Emojies              []uint64             `protobuf:"varint,4,rep,packed,name=emojies,proto3" json:"emojies,omitempty"`
	Mentions             []uint64             `protobuf:"varint,5,rep,packed,name=mentions,proto3" json:"mentions,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ChatMsgEvent) Reset()         { *m = ChatMsgEvent{} }
func (m *ChatMsgEvent) String() string { return proto.CompactTextString(m) }
func (*ChatMsgEvent) ProtoMessage()    {}
func (*ChatMsgEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{2}
}

func (m *ChatMsgEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMsgEvent.Unmarshal(m, b)
}
func (m *ChatMsgEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMsgEvent.Marshal(b, m, deterministic)
}
func (m *ChatMsgEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMsgEvent.Merge(m, src)
}
func (m *ChatMsgEvent) XXX_Size() int {
	return xxx_messageInfo_ChatMsgEvent.Size(m)
}
func (m *ChatMsgEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMsgEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMsgEvent proto.InternalMessageInfo

func (m *ChatMsgEvent) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *ChatMsgEvent) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ChatMsgEvent) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ChatMsgEvent) GetEmojies() []uint64 {
	if m != nil {
		return m.Emojies
	}
	return nil
}

func (m *ChatMsgEvent) GetMentions() []uint64 {
	if m != nil {
		return m.Mentions
	}
	return nil
}

func (m *ChatMsgEvent) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type PersonalStateActivityMsgEvent struct {
	ActivityId           int64    `protobuf:"varint,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	ActivityText         []byte   `protobuf:"bytes,2,opt,name=activity_text,json=activityText,proto3" json:"activity_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonalStateActivityMsgEvent) Reset()         { *m = PersonalStateActivityMsgEvent{} }
func (m *PersonalStateActivityMsgEvent) String() string { return proto.CompactTextString(m) }
func (*PersonalStateActivityMsgEvent) ProtoMessage()    {}
func (*PersonalStateActivityMsgEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{3}
}

func (m *PersonalStateActivityMsgEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonalStateActivityMsgEvent.Unmarshal(m, b)
}
func (m *PersonalStateActivityMsgEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonalStateActivityMsgEvent.Marshal(b, m, deterministic)
}
func (m *PersonalStateActivityMsgEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonalStateActivityMsgEvent.Merge(m, src)
}
func (m *PersonalStateActivityMsgEvent) XXX_Size() int {
	return xxx_messageInfo_PersonalStateActivityMsgEvent.Size(m)
}
func (m *PersonalStateActivityMsgEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonalStateActivityMsgEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PersonalStateActivityMsgEvent proto.InternalMessageInfo

func (m *PersonalStateActivityMsgEvent) GetActivityId() int64 {
	if m != nil {
		return m.ActivityId
	}
	return 0
}

func (m *PersonalStateActivityMsgEvent) GetActivityText() []byte {
	if m != nil {
		return m.ActivityText
	}
	return nil
}

type TheaterLogOnEvent struct {
	Room                 []byte   `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	Token                []byte   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TheaterLogOnEvent) Reset()         { *m = TheaterLogOnEvent{} }
func (m *TheaterLogOnEvent) String() string { return proto.CompactTextString(m) }
func (*TheaterLogOnEvent) ProtoMessage()    {}
func (*TheaterLogOnEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{4}
}

func (m *TheaterLogOnEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TheaterLogOnEvent.Unmarshal(m, b)
}
func (m *TheaterLogOnEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TheaterLogOnEvent.Marshal(b, m, deterministic)
}
func (m *TheaterLogOnEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TheaterLogOnEvent.Merge(m, src)
}
func (m *TheaterLogOnEvent) XXX_Size() int {
	return xxx_messageInfo_TheaterLogOnEvent.Size(m)
}
func (m *TheaterLogOnEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TheaterLogOnEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TheaterLogOnEvent proto.InternalMessageInfo

func (m *TheaterLogOnEvent) GetRoom() []byte {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *TheaterLogOnEvent) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

type LogOnEvent struct {
	Username             []byte   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             []byte   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Token                []byte   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogOnEvent) Reset()         { *m = LogOnEvent{} }
func (m *LogOnEvent) String() string { return proto.CompactTextString(m) }
func (*LogOnEvent) ProtoMessage()    {}
func (*LogOnEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{5}
}

func (m *LogOnEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogOnEvent.Unmarshal(m, b)
}
func (m *LogOnEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogOnEvent.Marshal(b, m, deterministic)
}
func (m *LogOnEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogOnEvent.Merge(m, src)
}
func (m *LogOnEvent) XXX_Size() int {
	return xxx_messageInfo_LogOnEvent.Size(m)
}
func (m *LogOnEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LogOnEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LogOnEvent proto.InternalMessageInfo

func (m *LogOnEvent) GetUsername() []byte {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *LogOnEvent) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *LogOnEvent) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

type MsgEvent struct {
	Type                 enums.EMSG           `protobuf:"varint,1,opt,name=type,proto3,enum=enums.EMSG" json:"type,omitempty"`
	Data                 []byte               `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MsgEvent) Reset()         { *m = MsgEvent{} }
func (m *MsgEvent) String() string { return proto.CompactTextString(m) }
func (*MsgEvent) ProtoMessage()    {}
func (*MsgEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{6}
}

func (m *MsgEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgEvent.Unmarshal(m, b)
}
func (m *MsgEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgEvent.Marshal(b, m, deterministic)
}
func (m *MsgEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEvent.Merge(m, src)
}
func (m *MsgEvent) XXX_Size() int {
	return xxx_messageInfo_MsgEvent.Size(m)
}
func (m *MsgEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEvent proto.InternalMessageInfo

func (m *MsgEvent) GetType() enums.EMSG {
	if m != nil {
		return m.Type
	}
	return enums.EMSG_INVALID
}

func (m *MsgEvent) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MsgEvent) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*CMsgProtoBufHeader)(nil), "protobuf.CMsgProtoBufHeader")
	proto.RegisterType((*PersonalStateMsgEvent)(nil), "protobuf.PersonalStateMsgEvent")
	proto.RegisterType((*ChatMsgEvent)(nil), "protobuf.ChatMsgEvent")
	proto.RegisterType((*PersonalStateActivityMsgEvent)(nil), "protobuf.PersonalStateActivityMsgEvent")
	proto.RegisterType((*TheaterLogOnEvent)(nil), "protobuf.TheaterLogOnEvent")
	proto.RegisterType((*LogOnEvent)(nil), "protobuf.LogOnEvent")
	proto.RegisterType((*MsgEvent)(nil), "protobuf.MsgEvent")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150) }

var fileDescriptor_db1b6b0986796150 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x6e, 0xda, 0x4c,
	0x14, 0x95, 0x83, 0xe1, 0x83, 0x0b, 0x1f, 0x55, 0xa6, 0x3f, 0xb2, 0xac, 0x56, 0x41, 0xee, 0xa2,
	0xac, 0x4c, 0x45, 0x57, 0x5d, 0x74, 0xe1, 0x22, 0xd4, 0x20, 0x85, 0x06, 0x0d, 0xee, 0xa6, 0x1b,
	0x34, 0xc4, 0x17, 0xe3, 0x16, 0x7b, 0x90, 0xe7, 0x92, 0x26, 0x7d, 0x8a, 0x3e, 0x4b, 0xf7, 0x7d,
	0xb7, 0xca, 0x63, 0x8f, 0xa1, 0x59, 0x54, 0xea, 0xc6, 0x9e, 0x73, 0x7f, 0xce, 0x9c, 0x7b, 0xe6,
	0x02, 0xac, 0x85, 0x42, 0x7f, 0x9f, 0x4b, 0x92, 0xac, 0xad, 0x7f, 0xeb, 0xc3, 0xc6, 0xbd, 0x88,
	0xa5, 0x8c, 0x77, 0x38, 0x32, 0x81, 0x11, 0x25, 0x29, 0x2a, 0x12, 0xe9, 0xbe, 0x2c, 0x75, 0xcf,
	0x31, 0x3b, 0xa4, 0x6a, 0xa4, 0xbf, 0x55, 0xe8, 0x71, 0x8a, 0x4a, 0x89, 0x18, 0xd5, 0xe8, 0xa0,
	0x30, 0x2f, 0x83, 0xde, 0x0f, 0x0b, 0xd8, 0x64, 0xae, 0xe2, 0x45, 0x81, 0xde, 0x1f, 0x36, 0x97,
	0x28, 0x22, 0xcc, 0x99, 0x0b, 0xed, 0x4f, 0x0a, 0xf3, 0x4b, 0xa1, 0xb6, 0x8e, 0x35, 0xb0, 0x86,
	0x3d, 0x5e, 0x63, 0xf6, 0x0c, 0x5a, 0xb3, 0x7d, 0x10, 0x45, 0xb9, 0x73, 0xa6, 0x33, 0x15, 0x62,
	0x43, 0x78, 0xb4, 0x44, 0xa5, 0x12, 0x99, 0x4d, 0x76, 0x09, 0x66, 0x34, 0x8b, 0x9c, 0xc6, 0xc0,
	0x1a, 0x36, 0xf9, 0xc3, 0x30, 0x7b, 0x0e, 0x9d, 0xe0, 0x40, 0xdb, 0x50, 0x7e, 0xc5, 0xcc, 0xb1,
	0x35, 0xc9, 0x31, 0xe0, 0xfd, 0xb4, 0xe0, 0xe9, 0x02, 0x73, 0x25, 0x33, 0xb1, 0x5b, 0x92, 0x20,
	0x9c, 0xab, 0x78, 0x7a, 0x8b, 0x19, 0x31, 0x0f, 0xec, 0x42, 0xba, 0x56, 0xd4, 0x1d, 0xf7, 0x7d,
	0x33, 0x90, 0x5f, 0x68, 0xe3, 0x3a, 0xc7, 0x5e, 0x43, 0x53, 0x15, 0x4d, 0x5a, 0x5c, 0x7f, 0xec,
	0xfa, 0xa5, 0x05, 0xd3, 0xf9, 0xf2, 0xc3, 0x6a, 0x31, 0xe5, 0xcb, 0xeb, 0x8f, 0xc1, 0xd5, 0x6a,
	0x19, 0x06, 0xe1, 0x94, 0x97, 0x85, 0x6c, 0x02, 0x6d, 0x71, 0x43, 0xc9, 0x6d, 0x42, 0xf7, 0x5a,
	0x70, 0x77, 0xfc, 0xca, 0x37, 0xbe, 0xfa, 0x7f, 0x08, 0x09, 0xaa, 0x32, 0x23, 0x88, 0xd7, 0x8d,
	0xde, 0x2f, 0x0b, 0x7a, 0x93, 0xad, 0xa0, 0x5a, 0xab, 0x03, 0xff, 0x55, 0xf2, 0x2a, 0x03, 0x0d,
	0x64, 0x0c, 0xec, 0x4d, 0x2e, 0x53, 0x2d, 0xb0, 0xc3, 0xf5, 0x99, 0xf5, 0xe1, 0x8c, 0xa4, 0xbe,
	0xbd, 0xc3, 0xcf, 0x48, 0x16, 0xdd, 0x98, 0xca, 0x2f, 0x09, 0x2a, 0xc7, 0x1e, 0x34, 0x86, 0x36,
	0x37, 0xb0, 0x78, 0x99, 0x14, 0x33, 0x4a, 0x64, 0xa6, 0x9c, 0xa6, 0x4e, 0xd5, 0x98, 0xbd, 0x05,
	0xb8, 0xc9, 0x51, 0x10, 0x46, 0x2b, 0x41, 0x4e, 0x4b, 0xcf, 0xe2, 0xfa, 0xe5, 0xaa, 0x1c, 0x47,
	0x0a, 0xcd, 0xaa, 0xf0, 0x4e, 0x55, 0x1d, 0x90, 0x87, 0xf0, 0xe2, 0xaf, 0xa3, 0xb2, 0x0b, 0xe8,
	0x9a, 0x61, 0x57, 0x49, 0xa4, 0x67, 0x6a, 0x70, 0x30, 0xa1, 0x59, 0xc4, 0x5e, 0xc2, 0xff, 0x75,
	0x01, 0xe1, 0x1d, 0x55, 0xdb, 0xd1, 0x33, 0xc1, 0x10, 0xef, 0xc8, 0x7b, 0x07, 0xe7, 0xe1, 0xb6,
	0xb8, 0x33, 0xbf, 0x92, 0xf1, 0x75, 0x56, 0x52, 0x33, 0xb0, 0x73, 0x29, 0xd3, 0xca, 0x27, 0x7d,
	0x66, 0x4f, 0xa0, 0x49, 0x7a, 0x3d, 0x4a, 0x96, 0x12, 0x78, 0x9f, 0x01, 0x4e, 0xfa, 0x5c, 0x68,
	0x17, 0x4f, 0x9e, 0x89, 0xd4, 0x78, 0x5c, 0xe3, 0x22, 0xb7, 0x17, 0x4a, 0x7d, 0x93, 0x79, 0x54,
	0x51, 0xd4, 0xf8, 0xc8, 0xdd, 0x38, 0xe5, 0xfe, 0x0e, 0xed, 0x93, 0x61, 0x6d, 0xba, 0xdf, 0x97,
	0xac, 0xfd, 0x71, 0xf7, 0x64, 0x87, 0xb8, 0x4e, 0x14, 0x92, 0x23, 0x41, 0xa2, 0xa2, 0xd6, 0xe7,
	0x07, 0xee, 0x37, 0xfe, 0xc1, 0xfd, 0x75, 0x4b, 0xa7, 0xdf, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0xfc, 0xf5, 0x35, 0x35, 0xed, 0x03, 0x00, 0x00,
}