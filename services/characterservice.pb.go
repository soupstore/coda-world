// Code generated by protoc-gen-go. DO NOT EDIT.
// source: characterservice.proto

/*
Package services is a generated protocol buffer package.

It is generated from these files:
	characterservice.proto

It has these top-level messages:
	CommandMessage
	EventMessage
	CharacterDescription
	RoomDescriptionEvent
	CharacterWakesUpEvent
	CharacterSleepsEvent
*/
package services

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type CommandType int32

const (
	CommandType_CmdLook CommandType = 0
)

var CommandType_name = map[int32]string{
	0: "CmdLook",
}
var CommandType_value = map[string]int32{
	"CmdLook": 0,
}

func (x CommandType) String() string {
	return proto.EnumName(CommandType_name, int32(x))
}
func (CommandType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EventType int32

const (
	EventType_EvtRoomDescription  EventType = 0
	EventType_EvtCharacterWakesUp EventType = 1
	EventType_EvtCharacterSleeps  EventType = 2
)

var EventType_name = map[int32]string{
	0: "EvtRoomDescription",
	1: "EvtCharacterWakesUp",
	2: "EvtCharacterSleeps",
}
var EventType_value = map[string]int32{
	"EvtRoomDescription":  0,
	"EvtCharacterWakesUp": 1,
	"EvtCharacterSleeps":  2,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}
func (EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CommandMessage struct {
	Type    CommandType `protobuf:"varint,1,opt,name=type,enum=services.CommandType" json:"type,omitempty"`
	Payload []byte      `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *CommandMessage) Reset()                    { *m = CommandMessage{} }
func (m *CommandMessage) String() string            { return proto.CompactTextString(m) }
func (*CommandMessage) ProtoMessage()               {}
func (*CommandMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CommandMessage) GetType() CommandType {
	if m != nil {
		return m.Type
	}
	return CommandType_CmdLook
}

func (m *CommandMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type EventMessage struct {
	Type    EventType `protobuf:"varint,1,opt,name=type,enum=services.EventType" json:"type,omitempty"`
	Payload []byte    `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *EventMessage) Reset()                    { *m = EventMessage{} }
func (m *EventMessage) String() string            { return proto.CompactTextString(m) }
func (*EventMessage) ProtoMessage()               {}
func (*EventMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EventMessage) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_EvtRoomDescription
}

func (m *EventMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Events
type CharacterDescription struct {
	Id    int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Awake bool   `protobuf:"varint,3,opt,name=awake" json:"awake,omitempty"`
}

func (m *CharacterDescription) Reset()                    { *m = CharacterDescription{} }
func (m *CharacterDescription) String() string            { return proto.CompactTextString(m) }
func (*CharacterDescription) ProtoMessage()               {}
func (*CharacterDescription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CharacterDescription) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CharacterDescription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CharacterDescription) GetAwake() bool {
	if m != nil {
		return m.Awake
	}
	return false
}

type RoomDescriptionEvent struct {
	Name        string                  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string                  `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Characters  []*CharacterDescription `protobuf:"bytes,3,rep,name=characters" json:"characters,omitempty"`
}

func (m *RoomDescriptionEvent) Reset()                    { *m = RoomDescriptionEvent{} }
func (m *RoomDescriptionEvent) String() string            { return proto.CompactTextString(m) }
func (*RoomDescriptionEvent) ProtoMessage()               {}
func (*RoomDescriptionEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RoomDescriptionEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoomDescriptionEvent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RoomDescriptionEvent) GetCharacters() []*CharacterDescription {
	if m != nil {
		return m.Characters
	}
	return nil
}

type CharacterWakesUpEvent struct {
	Character *CharacterDescription `protobuf:"bytes,1,opt,name=character" json:"character,omitempty"`
}

func (m *CharacterWakesUpEvent) Reset()                    { *m = CharacterWakesUpEvent{} }
func (m *CharacterWakesUpEvent) String() string            { return proto.CompactTextString(m) }
func (*CharacterWakesUpEvent) ProtoMessage()               {}
func (*CharacterWakesUpEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CharacterWakesUpEvent) GetCharacter() *CharacterDescription {
	if m != nil {
		return m.Character
	}
	return nil
}

type CharacterSleepsEvent struct {
	Character *CharacterDescription `protobuf:"bytes,1,opt,name=character" json:"character,omitempty"`
}

func (m *CharacterSleepsEvent) Reset()                    { *m = CharacterSleepsEvent{} }
func (m *CharacterSleepsEvent) String() string            { return proto.CompactTextString(m) }
func (*CharacterSleepsEvent) ProtoMessage()               {}
func (*CharacterSleepsEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CharacterSleepsEvent) GetCharacter() *CharacterDescription {
	if m != nil {
		return m.Character
	}
	return nil
}

func init() {
	proto.RegisterType((*CommandMessage)(nil), "services.CommandMessage")
	proto.RegisterType((*EventMessage)(nil), "services.EventMessage")
	proto.RegisterType((*CharacterDescription)(nil), "services.CharacterDescription")
	proto.RegisterType((*RoomDescriptionEvent)(nil), "services.RoomDescriptionEvent")
	proto.RegisterType((*CharacterWakesUpEvent)(nil), "services.CharacterWakesUpEvent")
	proto.RegisterType((*CharacterSleepsEvent)(nil), "services.CharacterSleepsEvent")
	proto.RegisterEnum("services.CommandType", CommandType_name, CommandType_value)
	proto.RegisterEnum("services.EventType", EventType_name, EventType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Character service

type CharacterClient interface {
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (Character_SubscribeClient, error)
}

type characterClient struct {
	cc *grpc.ClientConn
}

func NewCharacterClient(cc *grpc.ClientConn) CharacterClient {
	return &characterClient{cc}
}

func (c *characterClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (Character_SubscribeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Character_serviceDesc.Streams[0], c.cc, "/services.Character/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &characterSubscribeClient{stream}
	return x, nil
}

type Character_SubscribeClient interface {
	Send(*CommandMessage) error
	Recv() (*EventMessage, error)
	grpc.ClientStream
}

type characterSubscribeClient struct {
	grpc.ClientStream
}

func (x *characterSubscribeClient) Send(m *CommandMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *characterSubscribeClient) Recv() (*EventMessage, error) {
	m := new(EventMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Character service

type CharacterServer interface {
	Subscribe(Character_SubscribeServer) error
}

func RegisterCharacterServer(s *grpc.Server, srv CharacterServer) {
	s.RegisterService(&_Character_serviceDesc, srv)
}

func _Character_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CharacterServer).Subscribe(&characterSubscribeServer{stream})
}

type Character_SubscribeServer interface {
	Send(*EventMessage) error
	Recv() (*CommandMessage, error)
	grpc.ServerStream
}

type characterSubscribeServer struct {
	grpc.ServerStream
}

func (x *characterSubscribeServer) Send(m *EventMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *characterSubscribeServer) Recv() (*CommandMessage, error) {
	m := new(CommandMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Character_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.Character",
	HandlerType: (*CharacterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Character_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "characterservice.proto",
}

func init() { proto.RegisterFile("characterservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0xbb, 0x50, 0x6d, 0x19, 0x9a, 0x86, 0x6c, 0xff, 0x48, 0x7a, 0x30, 0x84, 0x8b, 0xd8,
	0x43, 0x63, 0xea, 0xd5, 0x78, 0xc1, 0xde, 0x34, 0xa9, 0xb4, 0x8d, 0xe7, 0x2d, 0x4c, 0x94, 0xb4,
	0xb0, 0x04, 0x10, 0xd3, 0xef, 0xe0, 0x87, 0x36, 0x6e, 0xf9, 0x27, 0x9a, 0xe8, 0xc1, 0x1b, 0x33,
	0x79, 0xf3, 0xcb, 0x9b, 0xe1, 0x2d, 0x8c, 0xdd, 0x17, 0x16, 0x33, 0x37, 0xc5, 0x38, 0xc1, 0x38,
	0xf3, 0x5d, 0x9c, 0x45, 0x31, 0x4f, 0x39, 0xed, 0xe6, 0x65, 0x62, 0x6e, 0xa0, 0x6f, 0xf3, 0x20,
	0x60, 0xa1, 0xf7, 0x80, 0x49, 0xc2, 0x9e, 0x91, 0x5e, 0x42, 0x3b, 0x3d, 0x44, 0xa8, 0x13, 0x83,
	0x58, 0xfd, 0xf9, 0x68, 0x56, 0x48, 0x67, 0xb9, 0x6e, 0x7d, 0x88, 0xd0, 0x11, 0x12, 0xaa, 0x43,
	0x27, 0x62, 0x87, 0x3d, 0x67, 0x9e, 0x2e, 0x19, 0xc4, 0xea, 0x39, 0x45, 0x69, 0x3e, 0x42, 0x6f,
	0x91, 0x61, 0x98, 0x16, 0xd0, 0x8b, 0x2f, 0xd0, 0x41, 0x05, 0x15, 0xaa, 0x3f, 0x21, 0x97, 0x30,
	0xb4, 0x8b, 0x6d, 0xee, 0x30, 0x71, 0x63, 0x3f, 0x4a, 0x7d, 0x1e, 0xd2, 0x3e, 0x48, 0xbe, 0x27,
	0xc0, 0xb2, 0x23, 0xf9, 0x1e, 0xa5, 0xd0, 0x0e, 0x59, 0x80, 0x62, 0x5c, 0x71, 0xc4, 0x37, 0x1d,
	0xc2, 0x09, 0x7b, 0x63, 0x3b, 0xd4, 0x65, 0x83, 0x58, 0x5d, 0xe7, 0x58, 0x98, 0xef, 0x04, 0x86,
	0x0e, 0xe7, 0x41, 0x8d, 0x26, 0xec, 0x94, 0x08, 0x52, 0x43, 0x18, 0xa0, 0x7a, 0x95, 0x2e, 0xa7,
	0xd7, 0x5b, 0xf4, 0x16, 0xa0, 0x3a, 0xb7, 0x2e, 0x1b, 0xb2, 0xa5, 0xce, 0xcf, 0x6b, 0xe7, 0xfb,
	0xc1, 0xbc, 0x53, 0x9b, 0x30, 0x37, 0x30, 0x2a, 0x35, 0x4f, 0x6c, 0x87, 0xc9, 0x26, 0x3a, 0xda,
	0xb9, 0x01, 0xa5, 0x94, 0x09, 0x4f, 0xbf, 0x73, 0xab, 0x01, 0x73, 0x5d, 0xbb, 0xdb, 0x6a, 0x8f,
	0x18, 0x25, 0xff, 0x40, 0x9d, 0x4e, 0x40, 0xad, 0xe5, 0x81, 0xaa, 0xd0, 0xb1, 0x03, 0xef, 0x9e,
	0xf3, 0x9d, 0xd6, 0x9a, 0xae, 0x41, 0x29, 0x7f, 0x2b, 0x1d, 0x03, 0x5d, 0x64, 0x69, 0xe3, 0xcc,
	0x5a, 0x8b, 0x9e, 0xc1, 0x60, 0x91, 0xa5, 0xcd, 0x85, 0x35, 0x92, 0x0f, 0x34, 0x2c, 0x6b, 0xd2,
	0x7c, 0x09, 0x4a, 0xd9, 0xa4, 0x36, 0x28, 0xab, 0xd7, 0xed, 0x27, 0x6f, 0x8b, 0x54, 0xff, 0x96,
	0xd1, 0x3c, 0x76, 0x93, 0x71, 0x23, 0x68, 0x79, 0xdf, 0x6c, 0x59, 0xe4, 0x8a, 0x6c, 0x4f, 0xc5,
	0x63, 0xb8, 0xfe, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xfb, 0xdc, 0x25, 0x26, 0x03, 0x00, 0x00,
}
