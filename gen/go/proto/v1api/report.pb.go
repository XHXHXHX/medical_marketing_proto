// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: proto/v1api/report.proto

package report

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ReportCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportCreateRequest) Reset() {
	*x = ReportCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1api_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportCreateRequest) ProtoMessage() {}

func (x *ReportCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1api_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportCreateRequest.ProtoReflect.Descriptor instead.
func (*ReportCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1api_report_proto_rawDescGZIP(), []int{0}
}

type ReportCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportCreateResponse) Reset() {
	*x = ReportCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1api_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportCreateResponse) ProtoMessage() {}

func (x *ReportCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1api_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportCreateResponse.ProtoReflect.Descriptor instead.
func (*ReportCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1api_report_proto_rawDescGZIP(), []int{1}
}

var File_proto_v1api_report_proto protoreflect.FileDescriptor

var file_proto_v1api_report_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x16, 0x0a,
	0x14, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x88, 0x01, 0x0a, 0x10, 0x61, 0x70, 0x69, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x0c, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a,
	0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58,
	0x48, 0x58, 0x48, 0x58, 0x48, 0x58, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1api_report_proto_rawDescOnce sync.Once
	file_proto_v1api_report_proto_rawDescData = file_proto_v1api_report_proto_rawDesc
)

func file_proto_v1api_report_proto_rawDescGZIP() []byte {
	file_proto_v1api_report_proto_rawDescOnce.Do(func() {
		file_proto_v1api_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1api_report_proto_rawDescData)
	})
	return file_proto_v1api_report_proto_rawDescData
}

var file_proto_v1api_report_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_v1api_report_proto_goTypes = []interface{}{
	(*ReportCreateRequest)(nil),  // 0: proto.report.ReportCreateRequest
	(*ReportCreateResponse)(nil), // 1: proto.report.ReportCreateResponse
}
var file_proto_v1api_report_proto_depIdxs = []int32{
	0, // 0: proto.report.apiReportService.ReportCreate:input_type -> proto.report.ReportCreateRequest
	1, // 1: proto.report.apiReportService.ReportCreate:output_type -> proto.report.ReportCreateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_v1api_report_proto_init() }
func file_proto_v1api_report_proto_init() {
	if File_proto_v1api_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1api_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportCreateRequest); i {
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
		file_proto_v1api_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportCreateResponse); i {
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
			RawDescriptor: file_proto_v1api_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1api_report_proto_goTypes,
		DependencyIndexes: file_proto_v1api_report_proto_depIdxs,
		MessageInfos:      file_proto_v1api_report_proto_msgTypes,
	}.Build()
	File_proto_v1api_report_proto = out.File
	file_proto_v1api_report_proto_rawDesc = nil
	file_proto_v1api_report_proto_goTypes = nil
	file_proto_v1api_report_proto_depIdxs = nil
}
