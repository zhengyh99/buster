// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

//同步玩家ID
type SyncPid struct {
	Pid                  int32    `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncPid) Reset()         { *m = SyncPid{} }
func (m *SyncPid) String() string { return proto.CompactTextString(m) }
func (*SyncPid) ProtoMessage()    {}
func (*SyncPid) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *SyncPid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPid.Unmarshal(m, b)
}
func (m *SyncPid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPid.Marshal(b, m, deterministic)
}
func (m *SyncPid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPid.Merge(m, src)
}
func (m *SyncPid) XXX_Size() int {
	return xxx_messageInfo_SyncPid.Size(m)
}
func (m *SyncPid) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPid.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPid proto.InternalMessageInfo

func (m *SyncPid) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

//位置信息
type Position struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=X,proto3" json:"X,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=Z,proto3" json:"Z,omitempty"`
	V                    float32  `protobuf:"fixed32,4,opt,name=V,proto3" json:"V,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Position) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Position) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *Position) GetV() float32 {
	if m != nil {
		return m.V
	}
	return 0
}

//广播消息
type BroadCast struct {
	Pid int32 `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Tp  int32 `protobuf:"varint,2,opt,name=Tp,proto3" json:"Tp,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*BroadCast_Content
	//	*BroadCast_P
	//	*BroadCast_ActionData
	Data                 isBroadCast_Data `protobuf_oneof:"Data"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BroadCast) Reset()         { *m = BroadCast{} }
func (m *BroadCast) String() string { return proto.CompactTextString(m) }
func (*BroadCast) ProtoMessage()    {}
func (*BroadCast) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *BroadCast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadCast.Unmarshal(m, b)
}
func (m *BroadCast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadCast.Marshal(b, m, deterministic)
}
func (m *BroadCast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadCast.Merge(m, src)
}
func (m *BroadCast) XXX_Size() int {
	return xxx_messageInfo_BroadCast.Size(m)
}
func (m *BroadCast) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadCast.DiscardUnknown(m)
}

var xxx_messageInfo_BroadCast proto.InternalMessageInfo

func (m *BroadCast) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *BroadCast) GetTp() int32 {
	if m != nil {
		return m.Tp
	}
	return 0
}

type isBroadCast_Data interface {
	isBroadCast_Data()
}

type BroadCast_Content struct {
	Content string `protobuf:"bytes,3,opt,name=Content,proto3,oneof"`
}

type BroadCast_P struct {
	P *Position `protobuf:"bytes,4,opt,name=P,proto3,oneof"`
}

type BroadCast_ActionData struct {
	ActionData int32 `protobuf:"varint,5,opt,name=ActionData,proto3,oneof"`
}

func (*BroadCast_Content) isBroadCast_Data() {}

func (*BroadCast_P) isBroadCast_Data() {}

func (*BroadCast_ActionData) isBroadCast_Data() {}

func (m *BroadCast) GetData() isBroadCast_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BroadCast) GetContent() string {
	if x, ok := m.GetData().(*BroadCast_Content); ok {
		return x.Content
	}
	return ""
}

func (m *BroadCast) GetP() *Position {
	if x, ok := m.GetData().(*BroadCast_P); ok {
		return x.P
	}
	return nil
}

func (m *BroadCast) GetActionData() int32 {
	if x, ok := m.GetData().(*BroadCast_ActionData); ok {
		return x.ActionData
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BroadCast) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BroadCast_Content)(nil),
		(*BroadCast_P)(nil),
		(*BroadCast_ActionData)(nil),
	}
}

func init() {
	proto.RegisterType((*SyncPid)(nil), "pb.SyncPid")
	proto.RegisterType((*Position)(nil), "pb.Position")
	proto.RegisterType((*BroadCast)(nil), "pb.BroadCast")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xb1, 0x4e, 0x86, 0x30,
	0x10, 0xc7, 0x69, 0xbf, 0xaf, 0x1f, 0x72, 0x12, 0x63, 0x3a, 0x35, 0xea, 0x40, 0x98, 0x9c, 0x18,
	0xf4, 0x09, 0x04, 0x07, 0xc6, 0xa6, 0x12, 0x02, 0x6c, 0x2d, 0x18, 0xd3, 0xc1, 0xb6, 0x81, 0x2e,
	0x3e, 0x86, 0xaf, 0xe1, 0x53, 0x9a, 0x96, 0x90, 0x98, 0x38, 0xdd, 0xfd, 0xfe, 0x97, 0xfc, 0xee,
	0x72, 0x90, 0x7d, 0x6e, 0x1f, 0x95, 0x5b, 0xad, 0xb7, 0x14, 0x3b, 0x55, 0xde, 0x43, 0xfa, 0xf6,
	0x65, 0x66, 0xae, 0x17, 0x7a, 0x0b, 0x27, 0xae, 0x17, 0x86, 0x0a, 0xf4, 0x48, 0x44, 0x68, 0xcb,
	0x1a, 0xae, 0xb8, 0xdd, 0xb4, 0xd7, 0xd6, 0xd0, 0x1c, 0xd0, 0x10, 0x67, 0x58, 0xa0, 0x21, 0xd0,
	0xc8, 0xf0, 0x4e, 0x63, 0xa0, 0x89, 0x9d, 0x76, 0x9a, 0x02, 0xf5, 0xec, 0xbc, 0x53, 0x5f, 0x7e,
	0x23, 0xc8, 0xea, 0xd5, 0xca, 0xa5, 0x91, 0x9b, 0xff, 0xbf, 0x83, 0xde, 0x00, 0xee, 0x5c, 0x54,
	0x11, 0x81, 0x3b, 0x47, 0xef, 0x20, 0x6d, 0xac, 0xf1, 0xef, 0xc6, 0x47, 0x63, 0xd6, 0x26, 0xe2,
	0x08, 0xe8, 0x03, 0x20, 0x1e, 0xcd, 0xd7, 0x4f, 0x79, 0xe5, 0x54, 0x75, 0x1c, 0xd7, 0x26, 0x02,
	0x71, 0x5a, 0x00, 0xbc, 0xcc, 0x01, 0x5f, 0xa5, 0x97, 0x8c, 0x04, 0x63, 0x9b, 0x88, 0x3f, 0x59,
	0x7d, 0x81, 0x73, 0xac, 0xe4, 0x07, 0x63, 0xa7, 0xd4, 0x25, 0xbe, 0xe1, 0xf9, 0x37, 0x00, 0x00,
	0xff, 0xff, 0xeb, 0x24, 0xba, 0x4a, 0x13, 0x01, 0x00, 0x00,
}
