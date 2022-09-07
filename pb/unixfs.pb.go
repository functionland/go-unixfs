// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: unixfs.proto

package unixfs_pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Data_DataType int32

const (
	Data_Raw       Data_DataType = 0
	Data_Directory Data_DataType = 1
	Data_File      Data_DataType = 2
	Data_Metadata  Data_DataType = 3
	Data_Symlink   Data_DataType = 4
	Data_HAMTShard Data_DataType = 5
	Data_EncFile   Data_DataType = 6
)

var Data_DataType_name = map[int32]string{
	0: "Raw",
	1: "Directory",
	2: "File",
	3: "Metadata",
	4: "Symlink",
	5: "HAMTShard",
	6: "EncFile",
}

var Data_DataType_value = map[string]int32{
	"Raw":       0,
	"Directory": 1,
	"File":      2,
	"Metadata":  3,
	"Symlink":   4,
	"HAMTShard": 5,
	"EncFile":   6,
}

func (x Data_DataType) Enum() *Data_DataType {
	p := new(Data_DataType)
	*p = x
	return p
}

func (x Data_DataType) String() string {
	return proto.EnumName(Data_DataType_name, int32(x))
}

func (x *Data_DataType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Data_DataType_value, data, "Data_DataType")
	if err != nil {
		return err
	}
	*x = Data_DataType(value)
	return nil
}

func (Data_DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e2fd76cc44dfc7c3, []int{0, 0}
}

type Data struct {
	Type                 *Data_DataType `protobuf:"varint,1,req,name=Type,enum=unixfs.pb.Data_DataType" json:"Type,omitempty"`
	Data                 []byte         `protobuf:"bytes,2,opt,name=Data" json:"Data,omitempty"`
	Filesize             *uint64        `protobuf:"varint,3,opt,name=filesize" json:"filesize,omitempty"`
	Blocksizes           []uint64       `protobuf:"varint,4,rep,name=blocksizes" json:"blocksizes,omitempty"`
	HashType             *uint64        `protobuf:"varint,5,opt,name=hashType" json:"hashType,omitempty"`
	Fanout               *uint64        `protobuf:"varint,6,opt,name=fanout" json:"fanout,omitempty"`
	Jwe                  []byte         `protobuf:"bytes,7,opt,name=jwe" json:"jwe,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2fd76cc44dfc7c3, []int{0}
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

func (m *Data) GetType() Data_DataType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Data_Raw
}

func (m *Data) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Data) GetFilesize() uint64 {
	if m != nil && m.Filesize != nil {
		return *m.Filesize
	}
	return 0
}

func (m *Data) GetBlocksizes() []uint64 {
	if m != nil {
		return m.Blocksizes
	}
	return nil
}

func (m *Data) GetHashType() uint64 {
	if m != nil && m.HashType != nil {
		return *m.HashType
	}
	return 0
}

func (m *Data) GetFanout() uint64 {
	if m != nil && m.Fanout != nil {
		return *m.Fanout
	}
	return 0
}

func (m *Data) GetJwe() []byte {
	if m != nil {
		return m.Jwe
	}
	return nil
}

type Metadata struct {
	MimeType             *string  `protobuf:"bytes,1,opt,name=MimeType" json:"MimeType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2fd76cc44dfc7c3, []int{1}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetMimeType() string {
	if m != nil && m.MimeType != nil {
		return *m.MimeType
	}
	return ""
}

func init() {
	proto.RegisterEnum("unixfs.pb.Data_DataType", Data_DataType_name, Data_DataType_value)
	proto.RegisterType((*Data)(nil), "unixfs.pb.Data")
	proto.RegisterType((*Metadata)(nil), "unixfs.pb.Metadata")
}

func init() { proto.RegisterFile("unixfs.proto", fileDescriptor_e2fd76cc44dfc7c3) }

var fileDescriptor_e2fd76cc44dfc7c3 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x6e, 0xb3, 0x30,
	0x14, 0x86, 0x3f, 0xc0, 0xe1, 0xe7, 0x7c, 0xb4, 0xb2, 0xce, 0x50, 0x59, 0x1d, 0x2a, 0xc4, 0x50,
	0x31, 0x54, 0x0c, 0xbd, 0x83, 0x4a, 0x69, 0xd5, 0x85, 0xc5, 0xc9, 0x0d, 0x38, 0xc4, 0x08, 0x37,
	0x04, 0x22, 0x70, 0x94, 0xd2, 0x9b, 0xea, 0x2d, 0x56, 0x36, 0x81, 0x76, 0xb1, 0xfc, 0xf8, 0x3d,
	0x8f, 0xf5, 0xea, 0x40, 0x7c, 0x6e, 0xd5, 0x67, 0x35, 0xe4, 0xa7, 0xbe, 0xd3, 0x1d, 0x46, 0x33,
	0xed, 0xd2, 0x6f, 0x17, 0xc8, 0x5a, 0x68, 0x81, 0x4f, 0x40, 0xb6, 0xe3, 0x49, 0x32, 0x27, 0x71,
	0xb3, 0xdb, 0x67, 0x96, 0x2f, 0x23, 0xb9, 0x89, 0xed, 0x61, 0x72, 0x6e, 0xa7, 0x10, 0x27, 0x8b,
	0xb9, 0x89, 0x93, 0xc5, 0x7c, 0xfa, 0xe1, 0x1e, 0xc2, 0x4a, 0x35, 0x72, 0x50, 0x5f, 0x92, 0x79,
	0x89, 0x93, 0x11, 0xbe, 0x30, 0x3e, 0x00, 0xec, 0x9a, 0xae, 0x3c, 0x18, 0x18, 0x18, 0x49, 0xbc,
	0x8c, 0xf0, 0x3f, 0x2f, 0xc6, 0xad, 0xc5, 0x50, 0xdb, 0x06, 0xab, 0xc9, 0x9d, 0x19, 0xef, 0xc0,
	0xaf, 0x44, 0xdb, 0x9d, 0x35, 0xf3, 0x6d, 0x72, 0x25, 0xa4, 0xe0, 0x7d, 0x5c, 0x24, 0x0b, 0x6c,
	0x05, 0x73, 0x4d, 0x4b, 0x08, 0xe7, 0x9e, 0x18, 0x80, 0xc7, 0xc5, 0x85, 0xfe, 0xc3, 0x1b, 0x88,
	0xd6, 0xaa, 0x97, 0xa5, 0xee, 0xfa, 0x91, 0x3a, 0x18, 0x02, 0x79, 0x53, 0x8d, 0xa4, 0x2e, 0xc6,
	0x10, 0x16, 0x52, 0x8b, 0xbd, 0xd0, 0x82, 0x7a, 0xf8, 0x1f, 0x82, 0xcd, 0x78, 0x6c, 0x54, 0x7b,
	0xa0, 0xc4, 0x38, 0xef, 0x2f, 0xc5, 0x76, 0x53, 0x8b, 0x7e, 0x4f, 0x57, 0x26, 0x7b, 0x6d, 0x4b,
	0xab, 0xf9, 0xe9, 0xe3, 0xaf, 0x66, 0x6a, 0x17, 0xea, 0x28, 0xaf, 0x8b, 0x73, 0xb2, 0x88, 0x2f,
	0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x20, 0x25, 0xaf, 0x26, 0x73, 0x01, 0x00, 0x00,
}
