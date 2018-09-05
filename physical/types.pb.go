// Code generated by protoc-gen-go. DO NOT EDIT.
// source: physical/types.proto

package physical // import "github.com/hashicorp/vault/physical"

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

type SealWrapEntry struct {
	Ciphertext           []byte   `protobuf:"bytes,1,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	Wrapped              bool     `protobuf:"varint,4,opt,name=wrapped,proto3" json:"wrapped,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SealWrapEntry) Reset()         { *m = SealWrapEntry{} }
func (m *SealWrapEntry) String() string { return proto.CompactTextString(m) }
func (*SealWrapEntry) ProtoMessage()    {}
func (*SealWrapEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_80e2b3c13a71f106, []int{0}
}
func (m *SealWrapEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SealWrapEntry.Unmarshal(m, b)
}
func (m *SealWrapEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SealWrapEntry.Marshal(b, m, deterministic)
}
func (dst *SealWrapEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SealWrapEntry.Merge(dst, src)
}
func (m *SealWrapEntry) XXX_Size() int {
	return xxx_messageInfo_SealWrapEntry.Size(m)
}
func (m *SealWrapEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SealWrapEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SealWrapEntry proto.InternalMessageInfo

func (m *SealWrapEntry) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
	}
	return nil
}

func (m *SealWrapEntry) GetWrapped() bool {
	if m != nil {
		return m.Wrapped
	}
	return false
}

func init() {
	proto.RegisterType((*SealWrapEntry)(nil), "physical.SealWrapEntry")
}

func init() { proto.RegisterFile("physical/types.proto", fileDescriptor_types_80e2b3c13a71f106) }

var fileDescriptor_types_80e2b3c13a71f106 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0xc8, 0xa8, 0x2c,
	0xce, 0x4c, 0x4e, 0xcc, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x80, 0x89, 0x2a, 0x79, 0x72, 0xf1, 0x06, 0xa7, 0x26, 0xe6, 0x84, 0x17, 0x25, 0x16,
	0xb8, 0xe6, 0x95, 0x14, 0x55, 0x0a, 0xc9, 0x71, 0x71, 0x25, 0x67, 0x16, 0x64, 0xa4, 0x16, 0x95,
	0xa4, 0x56, 0x94, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x21, 0x89, 0x08, 0x49, 0x70, 0xb1,
	0x97, 0x17, 0x25, 0x16, 0x14, 0xa4, 0xa6, 0x48, 0xb0, 0x28, 0x30, 0x6a, 0x70, 0x04, 0xc1, 0xb8,
	0x4e, 0xaa, 0x51, 0xca, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x19,
	0x89, 0xc5, 0x19, 0x99, 0xc9, 0xf9, 0x45, 0x05, 0xfa, 0x65, 0x89, 0xa5, 0x39, 0x25, 0xfa, 0x30,
	0x1b, 0x93, 0xd8, 0xc0, 0x4e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x71, 0xf9, 0x7a, 0x09,
	0x9a, 0x00, 0x00, 0x00,
}
