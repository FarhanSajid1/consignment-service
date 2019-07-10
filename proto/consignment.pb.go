// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment.proto

package go_micro_srv_consignment

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

type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_774f0b29a0c4f8e9, []int{0}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_774f0b29a0c4f8e9, []int{1}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_774f0b29a0c4f8e9, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=consignments,proto3" json:"consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_774f0b29a0c4f8e9, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "go.micro.srv.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "go.micro.srv.consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "go.micro.srv.consignment.Response")
}

func init() { proto.RegisterFile("proto/consignment.proto", fileDescriptor_774f0b29a0c4f8e9) }

var fileDescriptor_774f0b29a0c4f8e9 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0xff, 0x67, 0x52, 0x14, 0xf7, 0x60, 0x17, 0x3d, 0x18, 0x22, 0x42, 0x4f, 0x11,
	0xea, 0x23, 0xe4, 0x50, 0x72, 0xdd, 0xde, 0x95, 0x9a, 0x0c, 0xe9, 0x80, 0xd9, 0x4d, 0x77, 0xb7,
	0xf5, 0xd5, 0xf4, 0xed, 0xa4, 0x9b, 0x46, 0x57, 0xa4, 0xd2, 0x5b, 0xbe, 0x6f, 0xe6, 0x9b, 0xfc,
	0x76, 0x18, 0x98, 0x35, 0x5a, 0x59, 0xf5, 0x58, 0x28, 0x69, 0xa8, 0x92, 0x35, 0x4a, 0x9b, 0x3a,
	0x87, 0xf1, 0x4a, 0xa5, 0x35, 0x15, 0x5a, 0xa5, 0x46, 0xef, 0x53, 0xaf, 0x9e, 0x7c, 0x06, 0x10,
	0x65, 0x3f, 0x9a, 0x5d, 0x40, 0x8f, 0x4a, 0x1e, 0xc4, 0xc1, 0x3c, 0x14, 0x3d, 0x2a, 0x59, 0x0c,
	0x51, 0x89, 0xa6, 0xd0, 0xd4, 0x58, 0x52, 0x92, 0xf7, 0x5c, 0xc1, 0xb7, 0xd8, 0x35, 0x8c, 0xde,
	0x91, 0xaa, 0x8d, 0xe5, 0xfd, 0x38, 0x98, 0x0f, 0xc5, 0x51, 0xb1, 0x0c, 0xa0, 0x50, 0xd2, 0xae,
	0x49, 0xa2, 0x36, 0x7c, 0x10, 0xf7, 0xe7, 0xd1, 0xe2, 0x3e, 0x3d, 0x05, 0x92, 0x66, 0x5d, 0xaf,
	0xf0, 0x62, 0xec, 0x16, 0xc2, 0x3d, 0x1a, 0x83, 0x6f, 0x2f, 0x54, 0xf2, 0xa1, 0xfb, 0xf9, 0xa4,
	0x35, 0xf2, 0x32, 0xa9, 0x21, 0xfc, 0x4e, 0xfd, 0x01, 0xbf, 0x83, 0xa8, 0xd8, 0x19, 0xab, 0x6a,
	0xd4, 0x87, 0x6c, 0x0b, 0x0e, 0x9d, 0x95, 0x97, 0x07, 0x6e, 0xa5, 0xa9, 0x22, 0xe9, 0xb8, 0x43,
	0x71, 0x54, 0x6c, 0x06, 0xe3, 0x9d, 0x69, 0x43, 0x83, 0xb6, 0x70, 0x90, 0x79, 0x99, 0x4c, 0x01,
	0x96, 0x68, 0x05, 0x6e, 0x77, 0x68, 0x6c, 0xf2, 0x11, 0xc0, 0x44, 0xa0, 0x69, 0x94, 0x34, 0xc8,
	0x38, 0x8c, 0x0b, 0x8d, 0x6b, 0x8b, 0x2d, 0xc1, 0x44, 0x74, 0x92, 0x2d, 0x21, 0xf2, 0x5e, 0xe9,
	0x30, 0xa2, 0xc5, 0xc3, 0xbf, 0x6b, 0xe8, 0xbe, 0x85, 0x9f, 0x64, 0x39, 0x4c, 0x3d, 0x69, 0x78,
	0xdf, 0x2d, 0xf4, 0xcc, 0x49, 0xbf, 0xa2, 0x8b, 0x2d, 0x5c, 0xae, 0x36, 0xd4, 0x34, 0x24, 0xab,
	0x15, 0xea, 0x3d, 0x15, 0xc8, 0x9e, 0xe1, 0x2a, 0x73, 0xc4, 0xfe, 0x2d, 0x9c, 0x37, 0xfc, 0x26,
	0x39, 0xdd, 0xd6, 0x2d, 0xe8, 0x75, 0xe4, 0xee, 0xf0, 0xe9, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xc0,
	0x58, 0x8f, 0x05, 0xa2, 0x02, 0x00, 0x00,
}