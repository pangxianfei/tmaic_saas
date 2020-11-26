//protoc gen go生成的代码。不要编辑。
//资料来源：pbs原型
package pbs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//这是一个编译时断言，以确保生成的文件
//与正在编译的proto包兼容。
//这一行的编译错误可能意味着
//proto包需要更新。
//请升级proto包
const _ = proto.ProtoPackageIsVersion3 

type UserRegistered struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AffiliationFromCode  string   `protobuf:"bytes,2,opt,name=affiliation_from_code,json=affiliationFromCode,proto3" json:"affiliation_from_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegistered) Reset()         { *m = UserRegistered{} }
func (m *UserRegistered) String() string { return proto.CompactTextString(m) }
func (*UserRegistered) ProtoMessage()    {}
func (*UserRegistered) Descriptor() ([]byte, []int) {
	return fileDescriptor_f262ac8d8cf0a40b, []int{0}
}

func (m *UserRegistered) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegistered.Unmarshal(m, b)
}
func (m *UserRegistered) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegistered.Marshal(b, m, deterministic)
}
func (m *UserRegistered) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegistered.Merge(m, src)
}
func (m *UserRegistered) XXX_Size() int {
	return xxx_messageInfo_UserRegistered.Size(m)
}
func (m *UserRegistered) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegistered.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegistered proto.InternalMessageInfo

func (m *UserRegistered) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserRegistered) GetAffiliationFromCode() string {
	if m != nil {
		return m.AffiliationFromCode
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRegistered)(nil), "UserRegistered")
}

func init() { proto.RegisterFile("pbs.proto", fileDescriptor_f262ac8d8cf0a40b) }

var fileDescriptor_f262ac8d8cf0a40b = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x48, 0x2a, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x8a, 0xe5, 0xe2, 0x0b, 0x2d, 0x4e, 0x2d, 0x0a, 0x4a, 0x4d,
	0xcf, 0x2c, 0x2e, 0x49, 0x2d, 0x4a, 0x4d, 0x11, 0x12, 0xe7, 0x62, 0x2f, 0x2d, 0x4e, 0x2d, 0x8a,
	0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d, 0x62, 0x03, 0x71, 0x3d, 0x53, 0x84, 0x8c,
	0xb8, 0x44, 0x13, 0xd3, 0xd2, 0x32, 0x73, 0x32, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0xe2, 0xd3, 0x8a,
	0xf2, 0x73, 0xe3, 0x93, 0xf3, 0x53, 0x52, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x84, 0x91,
	0x24, 0xdd, 0x8a, 0xf2, 0x73, 0x9d, 0xf3, 0x53, 0x52, 0x93, 0xd8, 0xc0, 0xb6, 0x18, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x6e, 0x83, 0x59, 0x54, 0x72, 0x00, 0x00, 0x00,
}
