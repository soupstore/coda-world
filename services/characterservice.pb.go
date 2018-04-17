// Code generated by protoc-gen-go. DO NOT EDIT.
// source: characterservice.proto

/*
Package services is a generated protocol buffer package.

It is generated from these files:
	characterservice.proto

It has these top-level messages:
	CommandMessage
	EventMessage
	SayCommand
	TakeCommand
	CharacterDescription
	ItemDescription
	ExitDescription
	RoomDescriptionEvent
	CharacterWakesUpEvent
	CharacterSleepsEvent
	CharacterSpeaksEvent
	CharacterArrivesEvent
	CharacterLeavesEvent
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

type Direction int32

const (
	Direction_North     Direction = 0
	Direction_NorthEast Direction = 1
	Direction_East      Direction = 2
	Direction_SouthEast Direction = 3
	Direction_South     Direction = 4
	Direction_SouthWest Direction = 5
	Direction_West      Direction = 6
	Direction_NorthWest Direction = 7
)

var Direction_name = map[int32]string{
	0: "North",
	1: "NorthEast",
	2: "East",
	3: "SouthEast",
	4: "South",
	5: "SouthWest",
	6: "West",
	7: "NorthWest",
}
var Direction_value = map[string]int32{
	"North":     0,
	"NorthEast": 1,
	"East":      2,
	"SouthEast": 3,
	"South":     4,
	"SouthWest": 5,
	"West":      6,
	"NorthWest": 7,
}

func (x Direction) String() string {
	return proto.EnumName(Direction_name, int32(x))
}
func (Direction) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CommandType int32

const (
	CommandType_CmdLook      CommandType = 0
	CommandType_CmdSay       CommandType = 1
	CommandType_CmdNorth     CommandType = 2
	CommandType_CmdNorthEast CommandType = 3
	CommandType_CmdEast      CommandType = 4
	CommandType_CmdSouthEast CommandType = 5
	CommandType_CmdSouth     CommandType = 6
	CommandType_CmdSouthWest CommandType = 7
	CommandType_CmdWest      CommandType = 8
	CommandType_CmdNorthWest CommandType = 9
	CommandType_CmdTake      CommandType = 10
)

var CommandType_name = map[int32]string{
	0:  "CmdLook",
	1:  "CmdSay",
	2:  "CmdNorth",
	3:  "CmdNorthEast",
	4:  "CmdEast",
	5:  "CmdSouthEast",
	6:  "CmdSouth",
	7:  "CmdSouthWest",
	8:  "CmdWest",
	9:  "CmdNorthWest",
	10: "CmdTake",
}
var CommandType_value = map[string]int32{
	"CmdLook":      0,
	"CmdSay":       1,
	"CmdNorth":     2,
	"CmdNorthEast": 3,
	"CmdEast":      4,
	"CmdSouthEast": 5,
	"CmdSouth":     6,
	"CmdSouthWest": 7,
	"CmdWest":      8,
	"CmdNorthWest": 9,
	"CmdTake":      10,
}

func (x CommandType) String() string {
	return proto.EnumName(CommandType_name, int32(x))
}
func (CommandType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type EventType int32

const (
	EventType_EvtRoomDescription       EventType = 0
	EventType_EvtCharacterWakesUp      EventType = 1
	EventType_EvtCharacterSleeps       EventType = 2
	EventType_EvtCharacterSpeaks       EventType = 3
	EventType_EvtCharacterArrives      EventType = 4
	EventType_EvtCharacterLeaves       EventType = 5
	EventType_EvtNoExitInThatDirection EventType = 6
	EventType_EvtItemNotHere           EventType = 7
)

var EventType_name = map[int32]string{
	0: "EvtRoomDescription",
	1: "EvtCharacterWakesUp",
	2: "EvtCharacterSleeps",
	3: "EvtCharacterSpeaks",
	4: "EvtCharacterArrives",
	5: "EvtCharacterLeaves",
	6: "EvtNoExitInThatDirection",
	7: "EvtItemNotHere",
}
var EventType_value = map[string]int32{
	"EvtRoomDescription":       0,
	"EvtCharacterWakesUp":      1,
	"EvtCharacterSleeps":       2,
	"EvtCharacterSpeaks":       3,
	"EvtCharacterArrives":      4,
	"EvtCharacterLeaves":       5,
	"EvtNoExitInThatDirection": 6,
	"EvtItemNotHere":           7,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}
func (EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

// Commands
type SayCommand struct {
	Content string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
}

func (m *SayCommand) Reset()                    { *m = SayCommand{} }
func (m *SayCommand) String() string            { return proto.CompactTextString(m) }
func (*SayCommand) ProtoMessage()               {}
func (*SayCommand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SayCommand) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type TakeCommand struct {
	Alias string `protobuf:"bytes,1,opt,name=alias" json:"alias,omitempty"`
}

func (m *TakeCommand) Reset()                    { *m = TakeCommand{} }
func (m *TakeCommand) String() string            { return proto.CompactTextString(m) }
func (*TakeCommand) ProtoMessage()               {}
func (*TakeCommand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TakeCommand) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
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
func (*CharacterDescription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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

type ItemDescription struct {
	Id   int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ItemDescription) Reset()                    { *m = ItemDescription{} }
func (m *ItemDescription) String() string            { return proto.CompactTextString(m) }
func (*ItemDescription) ProtoMessage()               {}
func (*ItemDescription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ItemDescription) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ItemDescription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ExitDescription struct {
	Direction Direction `protobuf:"varint,1,opt,name=direction,enum=services.Direction" json:"direction,omitempty"`
	Name      string    `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ExitDescription) Reset()                    { *m = ExitDescription{} }
func (m *ExitDescription) String() string            { return proto.CompactTextString(m) }
func (*ExitDescription) ProtoMessage()               {}
func (*ExitDescription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ExitDescription) GetDirection() Direction {
	if m != nil {
		return m.Direction
	}
	return Direction_North
}

func (m *ExitDescription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RoomDescriptionEvent struct {
	Name        string                  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string                  `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Characters  []*CharacterDescription `protobuf:"bytes,3,rep,name=characters" json:"characters,omitempty"`
	Items       []*ItemDescription      `protobuf:"bytes,4,rep,name=items" json:"items,omitempty"`
	Exits       []*ExitDescription      `protobuf:"bytes,5,rep,name=exits" json:"exits,omitempty"`
}

func (m *RoomDescriptionEvent) Reset()                    { *m = RoomDescriptionEvent{} }
func (m *RoomDescriptionEvent) String() string            { return proto.CompactTextString(m) }
func (*RoomDescriptionEvent) ProtoMessage()               {}
func (*RoomDescriptionEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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

func (m *RoomDescriptionEvent) GetItems() []*ItemDescription {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *RoomDescriptionEvent) GetExits() []*ExitDescription {
	if m != nil {
		return m.Exits
	}
	return nil
}

type CharacterWakesUpEvent struct {
	Character *CharacterDescription `protobuf:"bytes,1,opt,name=character" json:"character,omitempty"`
}

func (m *CharacterWakesUpEvent) Reset()                    { *m = CharacterWakesUpEvent{} }
func (m *CharacterWakesUpEvent) String() string            { return proto.CompactTextString(m) }
func (*CharacterWakesUpEvent) ProtoMessage()               {}
func (*CharacterWakesUpEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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
func (*CharacterSleepsEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CharacterSleepsEvent) GetCharacter() *CharacterDescription {
	if m != nil {
		return m.Character
	}
	return nil
}

type CharacterSpeaksEvent struct {
	Character *CharacterDescription `protobuf:"bytes,1,opt,name=character" json:"character,omitempty"`
	Content   string                `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
}

func (m *CharacterSpeaksEvent) Reset()                    { *m = CharacterSpeaksEvent{} }
func (m *CharacterSpeaksEvent) String() string            { return proto.CompactTextString(m) }
func (*CharacterSpeaksEvent) ProtoMessage()               {}
func (*CharacterSpeaksEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CharacterSpeaksEvent) GetCharacter() *CharacterDescription {
	if m != nil {
		return m.Character
	}
	return nil
}

func (m *CharacterSpeaksEvent) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CharacterArrivesEvent struct {
	Character *CharacterDescription `protobuf:"bytes,1,opt,name=character" json:"character,omitempty"`
	Direction Direction             `protobuf:"varint,2,opt,name=direction,enum=services.Direction" json:"direction,omitempty"`
}

func (m *CharacterArrivesEvent) Reset()                    { *m = CharacterArrivesEvent{} }
func (m *CharacterArrivesEvent) String() string            { return proto.CompactTextString(m) }
func (*CharacterArrivesEvent) ProtoMessage()               {}
func (*CharacterArrivesEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *CharacterArrivesEvent) GetCharacter() *CharacterDescription {
	if m != nil {
		return m.Character
	}
	return nil
}

func (m *CharacterArrivesEvent) GetDirection() Direction {
	if m != nil {
		return m.Direction
	}
	return Direction_North
}

type CharacterLeavesEvent struct {
	Character *CharacterDescription `protobuf:"bytes,1,opt,name=character" json:"character,omitempty"`
	Direction Direction             `protobuf:"varint,2,opt,name=direction,enum=services.Direction" json:"direction,omitempty"`
}

func (m *CharacterLeavesEvent) Reset()                    { *m = CharacterLeavesEvent{} }
func (m *CharacterLeavesEvent) String() string            { return proto.CompactTextString(m) }
func (*CharacterLeavesEvent) ProtoMessage()               {}
func (*CharacterLeavesEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *CharacterLeavesEvent) GetCharacter() *CharacterDescription {
	if m != nil {
		return m.Character
	}
	return nil
}

func (m *CharacterLeavesEvent) GetDirection() Direction {
	if m != nil {
		return m.Direction
	}
	return Direction_North
}

func init() {
	proto.RegisterType((*CommandMessage)(nil), "services.CommandMessage")
	proto.RegisterType((*EventMessage)(nil), "services.EventMessage")
	proto.RegisterType((*SayCommand)(nil), "services.SayCommand")
	proto.RegisterType((*TakeCommand)(nil), "services.TakeCommand")
	proto.RegisterType((*CharacterDescription)(nil), "services.CharacterDescription")
	proto.RegisterType((*ItemDescription)(nil), "services.ItemDescription")
	proto.RegisterType((*ExitDescription)(nil), "services.ExitDescription")
	proto.RegisterType((*RoomDescriptionEvent)(nil), "services.RoomDescriptionEvent")
	proto.RegisterType((*CharacterWakesUpEvent)(nil), "services.CharacterWakesUpEvent")
	proto.RegisterType((*CharacterSleepsEvent)(nil), "services.CharacterSleepsEvent")
	proto.RegisterType((*CharacterSpeaksEvent)(nil), "services.CharacterSpeaksEvent")
	proto.RegisterType((*CharacterArrivesEvent)(nil), "services.CharacterArrivesEvent")
	proto.RegisterType((*CharacterLeavesEvent)(nil), "services.CharacterLeavesEvent")
	proto.RegisterEnum("services.Direction", Direction_name, Direction_value)
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
	// 687 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xae, 0xd3, 0xa6, 0x6d, 0x4e, 0x4b, 0x67, 0x79, 0x3f, 0x04, 0x84, 0x50, 0x15, 0x24, 0x28,
	0xbb, 0x18, 0x30, 0xc4, 0x1d, 0x42, 0x42, 0x5d, 0x24, 0x26, 0x8d, 0x69, 0xa4, 0x9d, 0xc6, 0xad,
	0xd7, 0x58, 0x2c, 0xea, 0x12, 0x87, 0xc4, 0x0b, 0xeb, 0x13, 0xc0, 0xeb, 0xf0, 0x24, 0xbc, 0x09,
	0xcf, 0x80, 0xe2, 0x38, 0x89, 0x97, 0x4d, 0x62, 0x48, 0x93, 0xb8, 0xf3, 0xf1, 0xf9, 0xbe, 0x2f,
	0x9f, 0x4f, 0x8e, 0x8f, 0x61, 0x6b, 0x71, 0x46, 0x13, 0xba, 0x10, 0x2c, 0x49, 0x59, 0x92, 0x05,
	0x0b, 0xb6, 0x13, 0x27, 0x5c, 0x70, 0xd2, 0x57, 0x61, 0xea, 0x1c, 0xc3, 0x68, 0xca, 0xc3, 0x90,
	0x46, 0xfe, 0x47, 0x96, 0xa6, 0xf4, 0x0b, 0x23, 0xcf, 0xa1, 0x23, 0x56, 0x31, 0xb3, 0xd1, 0x18,
	0x4d, 0x46, 0xbb, 0x9b, 0x3b, 0x25, 0x74, 0x47, 0xe1, 0xe6, 0xab, 0x98, 0x79, 0x12, 0x42, 0x6c,
	0xe8, 0xc5, 0x74, 0x75, 0xce, 0xa9, 0x6f, 0x1b, 0x63, 0x34, 0x19, 0x7a, 0x65, 0xe8, 0x7c, 0x82,
	0xa1, 0x9b, 0xb1, 0x48, 0x94, 0xa2, 0xcf, 0xae, 0x88, 0xae, 0xd7, 0xa2, 0x12, 0x75, 0x2b, 0xc9,
	0xa7, 0x00, 0x33, 0xba, 0x52, 0x26, 0x72, 0xdc, 0x82, 0x47, 0x82, 0x45, 0x42, 0x6a, 0x5a, 0x5e,
	0x19, 0x3a, 0x4f, 0x60, 0x30, 0xa7, 0x4b, 0x56, 0x02, 0x37, 0xc0, 0xa4, 0xe7, 0x01, 0x4d, 0x15,
	0xac, 0x08, 0x9c, 0x23, 0xd8, 0x98, 0x96, 0xa5, 0xd9, 0x63, 0xe9, 0x22, 0x09, 0x62, 0x11, 0xf0,
	0x88, 0x8c, 0xc0, 0x08, 0x7c, 0x09, 0x6d, 0x7b, 0x46, 0xe0, 0x13, 0x02, 0x9d, 0x88, 0x86, 0x4c,
	0x7a, 0xb1, 0x3c, 0xb9, 0x96, 0x8a, 0xdf, 0xe8, 0x92, 0xd9, 0xed, 0x31, 0x9a, 0xf4, 0xbd, 0x22,
	0x70, 0xde, 0xc0, 0xda, 0xbe, 0x60, 0xe1, 0x3f, 0x8a, 0x39, 0x9f, 0x61, 0xcd, 0xbd, 0x0c, 0x84,
	0x4e, 0x7b, 0x05, 0x96, 0x1f, 0x24, 0x6c, 0x91, 0x07, 0xd7, 0x0b, 0xb6, 0x57, 0xa6, 0xbc, 0x1a,
	0x75, 0xa3, 0xf2, 0x6f, 0x04, 0x1b, 0x1e, 0xe7, 0xba, 0x23, 0x59, 0xec, 0x0a, 0x8c, 0xb4, 0x33,
	0x8d, 0x61, 0xe0, 0xd7, 0x38, 0xa5, 0xa3, 0x6f, 0x91, 0x77, 0x00, 0x75, 0x33, 0xd9, 0xed, 0x71,
	0x7b, 0x32, 0xd8, 0x7d, 0xac, 0x35, 0xc7, 0x0d, 0xd5, 0xf4, 0x34, 0x06, 0x79, 0x01, 0x66, 0x20,
	0x58, 0x98, 0xda, 0x1d, 0x49, 0x7d, 0x50, 0x53, 0x1b, 0x65, 0xf3, 0x0a, 0x5c, 0x4e, 0x60, 0x97,
	0x81, 0x48, 0x6d, 0xb3, 0x49, 0x68, 0x14, 0xcc, 0x2b, 0x70, 0xce, 0x31, 0x6c, 0x56, 0x2e, 0x4e,
	0xe8, 0x92, 0xa5, 0xc7, 0x71, 0x71, 0xe0, 0xb7, 0x60, 0x55, 0x46, 0xe4, 0xa9, 0xff, 0xee, 0xbc,
	0x26, 0x38, 0x73, 0xad, 0x55, 0x66, 0xe7, 0x8c, 0xc5, 0xe9, 0x5d, 0xa8, 0x46, 0xba, 0x6a, 0xcc,
	0xe8, 0xf2, 0x2e, 0x54, 0xf5, 0x5b, 0x61, 0x5c, 0xbd, 0x15, 0x3f, 0x90, 0x56, 0x9d, 0xf7, 0x49,
	0x12, 0x64, 0xec, 0x4e, 0xbe, 0x78, 0xa5, 0x59, 0x8d, 0xdb, 0x34, 0xab, 0xf3, 0x1d, 0x69, 0x67,
	0x3f, 0x60, 0xf4, 0xbf, 0x39, 0xd9, 0xfe, 0x0a, 0x56, 0xb5, 0x4f, 0x2c, 0x30, 0x0f, 0x79, 0x22,
	0xce, 0x70, 0x8b, 0xdc, 0x03, 0x4b, 0x2e, 0x5d, 0x9a, 0x0a, 0x8c, 0x48, 0x1f, 0x3a, 0x72, 0x65,
	0xe4, 0x89, 0x19, 0xbf, 0x50, 0x89, 0x76, 0x4e, 0x91, 0x21, 0xee, 0x54, 0x99, 0x13, 0x96, 0x0a,
	0x6c, 0xe6, 0x14, 0xb9, 0xea, 0x56, 0x5a, 0x32, 0xec, 0x6d, 0xff, 0x44, 0x30, 0xd0, 0x06, 0x29,
	0x19, 0x40, 0x6f, 0x1a, 0xfa, 0x07, 0x9c, 0x2f, 0x71, 0x8b, 0x00, 0x74, 0xa7, 0xa1, 0x3f, 0xa3,
	0x2b, 0x8c, 0xc8, 0x10, 0xfa, 0xd3, 0xd0, 0x2f, 0x1c, 0x19, 0x04, 0xc3, 0xb0, 0x8c, 0xd4, 0xb7,
	0x0b, 0xa2, 0x0c, 0x3a, 0x2a, 0x5d, 0x5b, 0x33, 0x15, 0xbd, 0x70, 0xd7, 0xd5, 0xf3, 0x85, 0x0f,
	0x45, 0x97, 0x41, 0x5f, 0x57, 0x97, 0x3b, 0x96, 0x4a, 0xe7, 0x73, 0x14, 0xc3, 0xf6, 0x2f, 0x04,
	0x56, 0x35, 0xa7, 0xc9, 0x16, 0x10, 0x37, 0x13, 0x8d, 0xc9, 0x82, 0x5b, 0xe4, 0x3e, 0xac, 0xbb,
	0x99, 0x68, 0xde, 0x40, 0x8c, 0x14, 0xa1, 0x71, 0x87, 0xb0, 0x71, 0x6d, 0x5f, 0xde, 0x02, 0xdc,
	0x6e, 0x0a, 0xa9, 0x66, 0xc5, 0x9d, 0x26, 0xa1, 0x68, 0x1d, 0x6c, 0x92, 0x47, 0x60, 0xbb, 0x99,
	0x38, 0xe4, 0xf9, 0x5c, 0xd8, 0x8f, 0xe6, 0x67, 0x54, 0x54, 0x7f, 0x15, 0x77, 0x09, 0x81, 0x91,
	0x9b, 0x89, 0x7c, 0xc8, 0x1c, 0x72, 0xf1, 0x81, 0x25, 0x0c, 0xf7, 0x76, 0x8f, 0xc0, 0xaa, 0x64,
	0xc8, 0x14, 0xac, 0xd9, 0xc5, 0x69, 0x7e, 0x94, 0x53, 0x46, 0xec, 0x6b, 0xef, 0x9d, 0x7a, 0xc2,
	0x1e, 0x6e, 0x35, 0x1e, 0x2d, 0xb5, 0xef, 0xb4, 0x26, 0xe8, 0x25, 0x3a, 0xed, 0xca, 0x87, 0xf5,
	0xf5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xc4, 0xdd, 0x23, 0x72, 0x07, 0x00, 0x00,
}
