// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

package test

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
const _ = proto.ProtoPackageIsVersion3	// please upgrade the proto package

type Data struct {
	Val			string		`protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral	struct{}	`json:"-" bson:"-"`
	XXX_unrecognized	[]byte		`json:"-" bson:"-"`
	XXX_sizecache		int32		`json:"-" bson:"-"`
}

func (m *Data) Reset()		{ *m = Data{} }
func (m *Data) String() string	{ return proto.CompactTextString(m) }
func (*Data) ProtoMessage()	{}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type DataNested struct {
	Val2			string		`protobuf:"bytes,1,opt,name=val2,proto3" json:"val2,omitempty"`
	XXX_NoUnkeyedLiteral	struct{}	`json:"-" bson:"-"`
	XXX_unrecognized	[]byte		`json:"-" bson:"-"`
	XXX_sizecache		int32		`json:"-" bson:"-"`
}

func (m *DataNested) Reset()		{ *m = DataNested{} }
func (m *DataNested) String() string	{ return proto.CompactTextString(m) }
func (*DataNested) ProtoMessage()	{}
func (*DataNested) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0, 0}
}

func (m *DataNested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataNested.Unmarshal(m, b)
}
func (m *DataNested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataNested.Marshal(b, m, deterministic)
}
func (m *DataNested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataNested.Merge(m, src)
}
func (m *DataNested) XXX_Size() int {
	return xxx_messageInfo_DataNested.Size(m)
}
func (m *DataNested) XXX_DiscardUnknown() {
	xxx_messageInfo_DataNested.DiscardUnknown(m)
}

var xxx_messageInfo_DataNested proto.InternalMessageInfo

func (m *DataNested) GetVal2() string {
	if m != nil {
		return m.Val2
	}
	return ""
}

type DataNested_OneMore_Nested struct {
	Val3			string		`protobuf:"bytes,1,opt,name=val3,proto3" json:"val3,omitempty"`
	XXX_NoUnkeyedLiteral	struct{}	`json:"-" bson:"-"`
	XXX_unrecognized	[]byte		`json:"-" bson:"-"`
	XXX_sizecache		int32		`json:"-" bson:"-"`
}

func (m *DataNested_OneMore_Nested) Reset()		{ *m = DataNested_OneMore_Nested{} }
func (m *DataNested_OneMore_Nested) String() string	{ return proto.CompactTextString(m) }
func (*DataNested_OneMore_Nested) ProtoMessage()	{}
func (*DataNested_OneMore_Nested) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0, 0, 0}
}

func (m *DataNested_OneMore_Nested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataNested_OneMore_Nested.Unmarshal(m, b)
}
func (m *DataNested_OneMore_Nested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataNested_OneMore_Nested.Marshal(b, m, deterministic)
}
func (m *DataNested_OneMore_Nested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataNested_OneMore_Nested.Merge(m, src)
}
func (m *DataNested_OneMore_Nested) XXX_Size() int {
	return xxx_messageInfo_DataNested_OneMore_Nested.Size(m)
}
func (m *DataNested_OneMore_Nested) XXX_DiscardUnknown() {
	xxx_messageInfo_DataNested_OneMore_Nested.DiscardUnknown(m)
}

var xxx_messageInfo_DataNested_OneMore_Nested proto.InternalMessageInfo

func (m *DataNested_OneMore_Nested) GetVal3() string {
	if m != nil {
		return m.Val3
	}
	return ""
}

func init() {
	proto.RegisterType((*Data)(nil), "test.Data")
	proto.RegisterType((*DataNested)(nil), "test.Data.nested")
	proto.RegisterType((*DataNested_OneMore_Nested)(nil), "test.Data.nested.OneMore_Nested")
}

func init()	{ proto.RegisterFile("data.proto", fileDescriptor_871986018790d2fd) }

var fileDescriptor_871986018790d2fd = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0x2c, 0x49,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x49, 0x2d, 0x2e, 0x51, 0x8a, 0xe1, 0x62,
	0x71, 0x49, 0x2c, 0x49, 0x14, 0x12, 0xe0, 0x62, 0x2e, 0x4b, 0xcc, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x02, 0x31, 0xa5, 0x9c, 0xb8, 0xd8, 0xf2, 0x52, 0x8b, 0x4b, 0x52, 0x53, 0x84, 0x84,
	0xb8, 0x58, 0xca, 0x12, 0x73, 0x8c, 0xa0, 0x92, 0x60, 0xb6, 0x94, 0x0a, 0x17, 0x9f, 0x7f, 0x5e,
	0xaa, 0x6f, 0x7e, 0x51, 0x6a, 0xbc, 0x1f, 0xb2, 0x2a, 0x63, 0x24, 0x55, 0xc6, 0x49, 0x6c, 0x60,
	0xab, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x55, 0xe7, 0x48, 0x4c, 0x78, 0x00, 0x00, 0x00,
}
