// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types.proto

package macaroon

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
	time "time"
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

type Caveat struct {
	// if any of these three are set, disallow that type of access
	DisallowReads   bool `protobuf:"varint,1,opt,name=disallow_reads,json=disallowReads,proto3" json:"disallow_reads,omitempty"`
	DisallowWrites  bool `protobuf:"varint,2,opt,name=disallow_writes,json=disallowWrites,proto3" json:"disallow_writes,omitempty"`
	DisallowLists   bool `protobuf:"varint,3,opt,name=disallow_lists,json=disallowLists,proto3" json:"disallow_lists,omitempty"`
	DisallowDeletes bool `protobuf:"varint,4,opt,name=disallow_deletes,json=disallowDeletes,proto3" json:"disallow_deletes,omitempty"`
	// if one of these two lists are not empty, restrict accesses
	// to the given values
	Buckets               [][]byte `protobuf:"bytes,10,rep,name=buckets,proto3" json:"buckets,omitempty"`
	EncryptedPathPrefixes [][]byte `protobuf:"bytes,11,rep,name=encrypted_path_prefixes,json=encryptedPathPrefixes,proto3" json:"encrypted_path_prefixes,omitempty"`
	// if set, the validity time window
	NotAfter             *time.Time `protobuf:"bytes,20,opt,name=not_after,json=notAfter,proto3,stdtime" json:"not_after,omitempty"`
	NotBefore            *time.Time `protobuf:"bytes,21,opt,name=not_before,json=notBefore,proto3,stdtime" json:"not_before,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Caveat) Reset()         { *m = Caveat{} }
func (m *Caveat) String() string { return proto.CompactTextString(m) }
func (*Caveat) ProtoMessage()    {}
func (*Caveat) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}
func (m *Caveat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Caveat.Unmarshal(m, b)
}
func (m *Caveat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Caveat.Marshal(b, m, deterministic)
}
func (m *Caveat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Caveat.Merge(m, src)
}
func (m *Caveat) XXX_Size() int {
	return xxx_messageInfo_Caveat.Size(m)
}
func (m *Caveat) XXX_DiscardUnknown() {
	xxx_messageInfo_Caveat.DiscardUnknown(m)
}

var xxx_messageInfo_Caveat proto.InternalMessageInfo

func (m *Caveat) GetDisallowReads() bool {
	if m != nil {
		return m.DisallowReads
	}
	return false
}

func (m *Caveat) GetDisallowWrites() bool {
	if m != nil {
		return m.DisallowWrites
	}
	return false
}

func (m *Caveat) GetDisallowLists() bool {
	if m != nil {
		return m.DisallowLists
	}
	return false
}

func (m *Caveat) GetDisallowDeletes() bool {
	if m != nil {
		return m.DisallowDeletes
	}
	return false
}

func (m *Caveat) GetBuckets() [][]byte {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *Caveat) GetEncryptedPathPrefixes() [][]byte {
	if m != nil {
		return m.EncryptedPathPrefixes
	}
	return nil
}

func (m *Caveat) GetNotAfter() *time.Time {
	if m != nil {
		return m.NotAfter
	}
	return nil
}

func (m *Caveat) GetNotBefore() *time.Time {
	if m != nil {
		return m.NotBefore
	}
	return nil
}

func init() {
	proto.RegisterType((*Caveat)(nil), "macaroon.Caveat")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x49, 0x5b, 0xfa, 0xf6, 0x9d, 0xfa, 0x8f, 0xc5, 0xe2, 0xd2, 0x4b, 0x8b, 0x20, 0xd6,
	0x4b, 0x0a, 0x0a, 0xde, 0x44, 0xac, 0x1e, 0x3d, 0x94, 0x20, 0x78, 0x0c, 0x9b, 0x66, 0x92, 0x06,
	0x93, 0x4c, 0xd8, 0x9d, 0x5a, 0xfb, 0x2d, 0xfc, 0x86, 0x7e, 0x11, 0x0f, 0xb2, 0x1b, 0x13, 0xe8,
	0xcd, 0xe3, 0xfc, 0xe6, 0xf7, 0xcc, 0xc0, 0x03, 0x43, 0xde, 0x55, 0x68, 0xfc, 0x4a, 0x13, 0x93,
	0x18, 0x14, 0x6a, 0xa5, 0x34, 0x51, 0x39, 0x86, 0x94, 0x52, 0xaa, 0xe9, 0x78, 0x92, 0x12, 0xa5,
	0x39, 0xce, 0xdd, 0x14, 0x6d, 0x92, 0x39, 0x67, 0x05, 0x1a, 0x56, 0x45, 0x55, 0x0b, 0xe7, 0xdf,
	0x1d, 0xe8, 0x3f, 0xaa, 0x77, 0x54, 0x2c, 0x2e, 0xe0, 0x28, 0xce, 0x8c, 0xca, 0x73, 0xda, 0x86,
	0x1a, 0x55, 0x6c, 0xa4, 0x37, 0xf5, 0x66, 0x83, 0xe0, 0xb0, 0xa1, 0x81, 0x85, 0xe2, 0x12, 0x8e,
	0x5b, 0x6d, 0xab, 0x33, 0x46, 0x23, 0x3b, 0xce, 0x6b, 0xd3, 0xaf, 0x8e, 0xee, 0xdd, 0xcb, 0x33,
	0xc3, 0x46, 0x76, 0xf7, 0xef, 0x3d, 0x5b, 0x28, 0xae, 0xe0, 0xa4, 0xd5, 0x62, 0xcc, 0xd1, 0x1e,
	0xec, 0x39, 0xb1, 0xfd, 0xf3, 0x54, 0x63, 0x21, 0xe1, 0x5f, 0xb4, 0x59, 0xbd, 0x21, 0x1b, 0x09,
	0xd3, 0xee, 0xec, 0x20, 0x68, 0x46, 0x71, 0x0b, 0x67, 0x58, 0xae, 0xf4, 0xae, 0x62, 0x8c, 0xc3,
	0x4a, 0xf1, 0x3a, 0xac, 0x34, 0x26, 0xd9, 0x07, 0x1a, 0x39, 0x74, 0xe6, 0xa8, 0x5d, 0x2f, 0x15,
	0xaf, 0x97, 0xbf, 0x4b, 0x71, 0x07, 0xff, 0x4b, 0xe2, 0x50, 0x25, 0x8c, 0x5a, 0x9e, 0x4e, 0xbd,
	0xd9, 0xf0, 0x7a, 0xec, 0xd7, 0x9d, 0xf9, 0x4d, 0x67, 0xfe, 0x4b, 0xd3, 0xd9, 0xa2, 0xf7, 0xf9,
	0x35, 0xf1, 0x82, 0x41, 0x49, 0xfc, 0x60, 0x13, 0xe2, 0x1e, 0xc0, 0xc6, 0x23, 0x4c, 0x48, 0xa3,
	0x1c, 0xfd, 0x31, 0x6f, 0x5f, 0x2e, 0x5c, 0x24, 0xea, 0x3b, 0xe9, 0xe6, 0x27, 0x00, 0x00, 0xff,
	0xff, 0xa2, 0x5a, 0x6d, 0xad, 0xcb, 0x01, 0x00, 0x00,
}
