package pbs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.ProtoPackageIsVersion3

type Test struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_f262ac8d8cf0a40b, []int{0}
}

func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (m *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(m, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

func (m *Test) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Test) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*Test)(nil), "Test")
}

func init() { proto.RegisterFile("pbs.proto", fileDescriptor_f262ac8d8cf0a40b) }

var fileDescriptor_f262ac8d8cf0a40b = []byte{
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x48, 0x2a, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2, 0xe1, 0x62, 0x09, 0x49, 0x2d, 0x2e, 0x11, 0x12, 0xe7,
	0x62, 0x2f, 0x2d, 0x4e, 0x2d, 0x8a, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d, 0x62,
	0x03, 0x71, 0x3d, 0x53, 0x84, 0xa4, 0xb9, 0x38, 0xc1, 0x12, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x1c, 0x20, 0x01, 0xbf, 0xc4, 0xdc, 0xd4, 0x24, 0x36, 0xb0, 0x21,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x25, 0xfc, 0x02, 0x51, 0x00, 0x00, 0x00,
}
