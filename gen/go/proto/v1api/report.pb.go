// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.21.12
// source: proto/v1api/report.proto

package v1api

import (
	common "github.com/XHXHXHX/medical_marketing_proto/gen/go/proto/common"
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

type ReportListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*common.Report   `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Page *common.Pagination `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ReportListResponse) Reset() {
	*x = ReportListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1api_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportListResponse) ProtoMessage() {}

func (x *ReportListResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ReportListResponse.ProtoReflect.Descriptor instead.
func (*ReportListResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1api_report_proto_rawDescGZIP(), []int{0}
}

func (x *ReportListResponse) GetList() []*common.Report {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ReportListResponse) GetPage() *common.Pagination {
	if x != nil {
		return x.Page
	}
	return nil
}

type ReportListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerMobile  string       `protobuf:"bytes,1,opt,name=consumer_mobile,json=consumerMobile,proto3" json:"consumer_mobile,omitempty"`       // 客户电话
	CreateStartTime int64        `protobuf:"varint,2,opt,name=create_start_time,json=createStartTime,proto3" json:"create_start_time,omitempty"` // 报单开始时间
	CreatEndTime    int64        `protobuf:"varint,3,opt,name=creat_end_time,json=creatEndTime,proto3" json:"creat_end_time,omitempty"`          // 报单结束时间
	UserId          int64        `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                              // 员工id
	IsMatch         int64        `protobuf:"varint,5,opt,name=is_match,json=isMatch,proto3" json:"is_match,omitempty"`                           // 是否匹配 1是 2否 0啥也不是
	UserName        string       `protobuf:"bytes,6,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	RelationTask    bool         `protobuf:"varint,7,opt,name=relation_task,json=relationTask,proto3" json:"relation_task,omitempty"` // 是否额外返回关联任务字段
	Page            *common.Page `protobuf:"bytes,8,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ReportListRequest) Reset() {
	*x = ReportListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1api_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportListRequest) ProtoMessage() {}

func (x *ReportListRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ReportListRequest.ProtoReflect.Descriptor instead.
func (*ReportListRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1api_report_proto_rawDescGZIP(), []int{1}
}

func (x *ReportListRequest) GetConsumerMobile() string {
	if x != nil {
		return x.ConsumerMobile
	}
	return ""
}

func (x *ReportListRequest) GetCreateStartTime() int64 {
	if x != nil {
		return x.CreateStartTime
	}
	return 0
}

func (x *ReportListRequest) GetCreatEndTime() int64 {
	if x != nil {
		return x.CreatEndTime
	}
	return 0
}

func (x *ReportListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReportListRequest) GetIsMatch() int64 {
	if x != nil {
		return x.IsMatch
	}
	return 0
}

func (x *ReportListRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ReportListRequest) GetRelationTask() bool {
	if x != nil {
		return x.RelationTask
	}
	return false
}

func (x *ReportListRequest) GetPage() *common.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type ReportCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 客户电话
	ConsumerMobile string `protobuf:"bytes,1,opt,name=consumer_mobile,json=consumerMobile,proto3" json:"consumer_mobile,omitempty"`
	// 客户姓名
	ConsumerName string `protobuf:"bytes,2,opt,name=consumer_name,json=consumerName,proto3" json:"consumer_name,omitempty"`
	// 预期到访时间
	ExpectArriveTime int64 `protobuf:"varint,3,opt,name=expect_arrive_time,json=expectArriveTime,proto3" json:"expect_arrive_time,omitempty"`
}

func (x *ReportCreateRequest) Reset() {
	*x = ReportCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1api_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportCreateRequest) ProtoMessage() {}

func (x *ReportCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1api_report_proto_msgTypes[2]
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
	return file_proto_v1api_report_proto_rawDescGZIP(), []int{2}
}

func (x *ReportCreateRequest) GetConsumerMobile() string {
	if x != nil {
		return x.ConsumerMobile
	}
	return ""
}

func (x *ReportCreateRequest) GetConsumerName() string {
	if x != nil {
		return x.ConsumerName
	}
	return ""
}

func (x *ReportCreateRequest) GetExpectArriveTime() int64 {
	if x != nil {
		return x.ExpectArriveTime
	}
	return 0
}

type ReportRecoverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReportRecoverRequest) Reset() {
	*x = ReportRecoverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1api_report_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRecoverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRecoverRequest) ProtoMessage() {}

func (x *ReportRecoverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1api_report_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRecoverRequest.ProtoReflect.Descriptor instead.
func (*ReportRecoverRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1api_report_proto_rawDescGZIP(), []int{3}
}

func (x *ReportRecoverRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_v1api_report_proto protoreflect.FileDescriptor

var file_proto_v1api_report_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x76, 0x31, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x60, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x22, 0xa6, 0x02, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x69, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x20, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x91, 0x01, 0x0a,
	0x13, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x72, 0x72,
	0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x41, 0x72, 0x72, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x26, 0x0a, 0x14, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xc4, 0x02, 0x0a, 0x10, 0x61, 0x70, 0x69,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a,
	0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e,
	0x76, 0x31, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x61,
	0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12,
	0x22, 0x2e, 0x76, 0x31, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x6c, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1f, 0x2e, 0x76, 0x31, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x76, 0x31, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x42,
	0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x48,
	0x58, 0x48, 0x58, 0x48, 0x58, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_v1api_report_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_v1api_report_proto_goTypes = []interface{}{
	(*ReportListResponse)(nil),   // 0: v1api.report.ReportListResponse
	(*ReportListRequest)(nil),    // 1: v1api.report.ReportListRequest
	(*ReportCreateRequest)(nil),  // 2: v1api.report.ReportCreateRequest
	(*ReportRecoverRequest)(nil), // 3: v1api.report.ReportRecoverRequest
	(*common.Report)(nil),        // 4: common.Report
	(*common.Pagination)(nil),    // 5: common.Pagination
	(*common.Page)(nil),          // 6: common.Page
	(*common.Empty)(nil),         // 7: common.Empty
}
var file_proto_v1api_report_proto_depIdxs = []int32{
	4, // 0: v1api.report.ReportListResponse.list:type_name -> common.Report
	5, // 1: v1api.report.ReportListResponse.page:type_name -> common.Pagination
	6, // 2: v1api.report.ReportListRequest.page:type_name -> common.Page
	2, // 3: v1api.report.apiReportService.ReportCreate:input_type -> v1api.report.ReportCreateRequest
	3, // 4: v1api.report.apiReportService.ReportRecover:input_type -> v1api.report.ReportRecoverRequest
	1, // 5: v1api.report.apiReportService.ReportList:input_type -> v1api.report.ReportListRequest
	7, // 6: v1api.report.apiReportService.ReportCreate:output_type -> common.Empty
	7, // 7: v1api.report.apiReportService.ReportRecover:output_type -> common.Empty
	0, // 8: v1api.report.apiReportService.ReportList:output_type -> v1api.report.ReportListResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_v1api_report_proto_init() }
func file_proto_v1api_report_proto_init() {
	if File_proto_v1api_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1api_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportListResponse); i {
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
			switch v := v.(*ReportListRequest); i {
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
		file_proto_v1api_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_v1api_report_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRecoverRequest); i {
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
			NumMessages:   4,
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
