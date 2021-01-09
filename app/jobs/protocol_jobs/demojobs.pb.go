// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: demojobs.proto

package jobs

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DemoJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query         string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber    int32  `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage int32  `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
}

func (x *DemoJob) Reset() {
	*x = DemoJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demojobs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoJob) ProtoMessage() {}

func (x *DemoJob) ProtoReflect() protoreflect.Message {
	mi := &file_demojobs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoJob.ProtoReflect.Descriptor instead.
func (*DemoJob) Descriptor() ([]byte, []int) {
	return file_demojobs_proto_rawDescGZIP(), []int{0}
}

func (x *DemoJob) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *DemoJob) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *DemoJob) GetResultPerPage() int32 {
	if x != nil {
		return x.ResultPerPage
	}
	return 0
}

var File_demojobs_proto protoreflect.FileDescriptor

var file_demojobs_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x65, 0x6d, 0x6f, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0x68, 0x0a, 0x07, 0x64, 0x65, 0x6d, 0x6f, 0x4a, 0x6f,
	0x62, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demojobs_proto_rawDescOnce sync.Once
	file_demojobs_proto_rawDescData = file_demojobs_proto_rawDesc
)

func file_demojobs_proto_rawDescGZIP() []byte {
	file_demojobs_proto_rawDescOnce.Do(func() {
		file_demojobs_proto_rawDescData = protoimpl.X.CompressGZIP(file_demojobs_proto_rawDescData)
	})
	return file_demojobs_proto_rawDescData
}

var file_demojobs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_demojobs_proto_goTypes = []interface{}{
	(*DemoJob)(nil), // 0: jobs.demoJob
}
var file_demojobs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_demojobs_proto_init() }
func file_demojobs_proto_init() {
	if File_demojobs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demojobs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoJob); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_demojobs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_demojobs_proto_goTypes,
		DependencyIndexes: file_demojobs_proto_depIdxs,
		MessageInfos:      file_demojobs_proto_msgTypes,
	}.Build()
	File_demojobs_proto = out.File
	file_demojobs_proto_rawDesc = nil
	file_demojobs_proto_goTypes = nil
	file_demojobs_proto_depIdxs = nil
}
