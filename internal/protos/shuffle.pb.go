// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/shuffle.proto

package protos

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

type TrackReq struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Worker               string   `protobuf:"bytes,2,opt,name=worker,proto3" json:"worker,omitempty"`
	Partitions           []string `protobuf:"bytes,3,rep,name=partitions,proto3" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackReq) Reset()         { *m = TrackReq{} }
func (m *TrackReq) String() string { return proto.CompactTextString(m) }
func (*TrackReq) ProtoMessage()    {}
func (*TrackReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6c49aed730ba08e, []int{0}
}

func (m *TrackReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackReq.Unmarshal(m, b)
}
func (m *TrackReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackReq.Marshal(b, m, deterministic)
}
func (m *TrackReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackReq.Merge(m, src)
}
func (m *TrackReq) XXX_Size() int {
	return xxx_messageInfo_TrackReq.Size(m)
}
func (m *TrackReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackReq.DiscardUnknown(m)
}

var xxx_messageInfo_TrackReq proto.InternalMessageInfo

func (m *TrackReq) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *TrackReq) GetWorker() string {
	if m != nil {
		return m.Worker
	}
	return ""
}

func (m *TrackReq) GetPartitions() []string {
	if m != nil {
		return m.Partitions
	}
	return nil
}

func init() {
	proto.RegisterType((*TrackReq)(nil), "protos.TrackReq")
}

func init() { proto.RegisterFile("protos/shuffle.proto", fileDescriptor_a6c49aed730ba08e) }

var fileDescriptor_a6c49aed730ba08e = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x2f, 0xce, 0x28, 0x4d, 0x4b, 0xcb, 0x49, 0xd5, 0x03, 0x73, 0x85, 0xd8, 0x20,
	0xa2, 0x4a, 0x61, 0x5c, 0x1c, 0x21, 0x45, 0x89, 0xc9, 0xd9, 0x41, 0xa9, 0x85, 0x42, 0x42, 0x5c,
	0x2c, 0xa5, 0xa5, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x18,
	0x17, 0x5b, 0x79, 0x7e, 0x51, 0x76, 0x6a, 0x91, 0x04, 0x13, 0x58, 0x14, 0xca, 0x13, 0x92, 0xe3,
	0xe2, 0x2a, 0x48, 0x2c, 0x2a, 0xc9, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x96, 0x60, 0x56, 0x60, 0xd6,
	0xe0, 0x0c, 0x42, 0x12, 0x71, 0xe2, 0x8c, 0x62, 0xd7, 0x87, 0x58, 0x91, 0x04, 0xb1, 0xca, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x39, 0x0d, 0x23, 0xd7, 0x89, 0x00, 0x00, 0x00,
}
