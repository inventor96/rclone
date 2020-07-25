// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: datarepair.proto

package pb

import (
	fmt "fmt"
	math "math"
	time "time"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// InjuredSegment is the queue item used for the data repair queue.
type InjuredSegment struct {
	Path                 []byte    `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	LostPieces           []int32   `protobuf:"varint,2,rep,packed,name=lost_pieces,json=lostPieces,proto3" json:"lost_pieces,omitempty"`
	InsertedTime         time.Time `protobuf:"bytes,3,opt,name=inserted_time,json=insertedTime,proto3,stdtime" json:"inserted_time"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *InjuredSegment) Reset()         { *m = InjuredSegment{} }
func (m *InjuredSegment) String() string { return proto.CompactTextString(m) }
func (*InjuredSegment) ProtoMessage()    {}
func (*InjuredSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b08e6fe9398aa6, []int{0}
}
func (m *InjuredSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InjuredSegment.Unmarshal(m, b)
}
func (m *InjuredSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InjuredSegment.Marshal(b, m, deterministic)
}
func (m *InjuredSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InjuredSegment.Merge(m, src)
}
func (m *InjuredSegment) XXX_Size() int {
	return xxx_messageInfo_InjuredSegment.Size(m)
}
func (m *InjuredSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_InjuredSegment.DiscardUnknown(m)
}

var xxx_messageInfo_InjuredSegment proto.InternalMessageInfo

func (m *InjuredSegment) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *InjuredSegment) GetLostPieces() []int32 {
	if m != nil {
		return m.LostPieces
	}
	return nil
}

func (m *InjuredSegment) GetInsertedTime() time.Time {
	if m != nil {
		return m.InsertedTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*InjuredSegment)(nil), "repair.InjuredSegment")
}

func init() { proto.RegisterFile("datarepair.proto", fileDescriptor_b1b08e6fe9398aa6) }

var fileDescriptor_b1b08e6fe9398aa6 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0x3f, 0x4e, 0xc3, 0x30,
	0x14, 0x87, 0x31, 0x85, 0x0a, 0xb9, 0x05, 0x21, 0x8b, 0x21, 0xca, 0x92, 0x88, 0x29, 0x93, 0x2d,
	0xc1, 0x0d, 0xba, 0x75, 0x43, 0x81, 0x89, 0xa5, 0x72, 0x9a, 0x87, 0x71, 0x55, 0xfb, 0x59, 0xf6,
	0xeb, 0x3d, 0x72, 0x2c, 0x4e, 0x01, 0x57, 0x41, 0xb1, 0x95, 0xed, 0xfd, 0xbe, 0xf7, 0xef, 0xe3,
	0x8f, 0xa3, 0x26, 0x1d, 0x21, 0x68, 0x1b, 0x65, 0x88, 0x48, 0x28, 0xd6, 0x25, 0xd5, 0xdc, 0xa0,
	0xc1, 0xc2, 0xea, 0xc6, 0x20, 0x9a, 0x33, 0xa8, 0x9c, 0x86, 0xcb, 0x97, 0x22, 0xeb, 0x20, 0x91,
	0x76, 0xa1, 0x0c, 0x3c, 0x4f, 0x8c, 0x3f, 0xec, 0xfd, 0xe9, 0x12, 0x61, 0x7c, 0x07, 0xe3, 0xc0,
	0x93, 0x10, 0xfc, 0x26, 0x68, 0xfa, 0xae, 0x58, 0xcb, 0xba, 0x6d, 0x9f, 0x6b, 0xd1, 0xf0, 0xcd,
	0x19, 0x13, 0x1d, 0x82, 0x85, 0x23, 0xa4, 0xea, 0xba, 0x5d, 0x75, 0xb7, 0x3d, 0x9f, 0xd1, 0x5b,
	0x26, 0x62, 0xcf, 0xef, 0xad, 0x4f, 0x10, 0x09, 0xc6, 0xc3, 0xfc, 0xa3, 0x5a, 0xb5, 0xac, 0xdb,
	0xbc, 0xd4, 0xb2, 0x08, 0xc8, 0x45, 0x40, 0x7e, 0x2c, 0x02, 0xbb, 0xbb, 0x9f, 0xdf, 0xe6, 0x6a,
	0xfa, 0x6b, 0x58, 0xbf, 0x5d, 0x56, 0xe7, 0xe6, 0xee, 0xe9, 0x53, 0x24, 0xc2, 0x78, 0x92, 0x16,
	0xd5, 0x11, 0x9d, 0x43, 0xaf, 0xc2, 0x30, 0xac, 0xf3, 0x85, 0xd7, 0xff, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1f, 0x51, 0xa8, 0x4d, 0xf8, 0x00, 0x00, 0x00,
}