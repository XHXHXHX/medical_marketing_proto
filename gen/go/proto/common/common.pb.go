// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.21.12
// source: proto/common/common.proto

package common

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_common_common_proto_rawDescGZIP(), []int{0}
}

// 请求中使用的 Page.
type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前页号.
	CurrentPage int64 `protobuf:"varint,1,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	// 每页显示的产品数.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_proto_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *Page) GetCurrentPage() int64 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *Page) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// Response 中使用的 Pagination.
type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总页数.
	PageCount int32 `protobuf:"varint,1,opt,name=page_count,json=pageCount,proto3" json:"page_count,omitempty"`
	// 每页显示的产品数.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// 总数.
	Total int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	// 当前页号.
	CurrentPage int32 `protobuf:"varint,4,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_proto_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *Pagination) GetPageCount() int32 {
	if x != nil {
		return x.PageCount
	}
	return 0
}

func (x *Pagination) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Pagination) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Pagination) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                                    //
	UserId               int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                                              // 员工id
	UserName             string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`                                         // 员工名称
	ExceptArriveTime     int64  `protobuf:"varint,4,opt,name=except_arrive_time,json=exceptArriveTime,proto3" json:"except_arrive_time,omitempty"`              // 预期到达时间
	ConsumerMobile       string `protobuf:"bytes,5,opt,name=consumer_mobile,json=consumerMobile,proto3" json:"consumer_mobile,omitempty"`                       // 客户电话
	ConsumerName         string `protobuf:"bytes,6,opt,name=consumer_name,json=consumerName,proto3" json:"consumer_name,omitempty"`                             // 客户名称
	IsMatch              bool   `protobuf:"varint,7,opt,name=is_match,json=isMatch,proto3" json:"is_match,omitempty"`                                           // 是否匹配
	ActualArriveTime     int64  `protobuf:"varint,8,opt,name=actual_arrive_time,json=actualArriveTime,proto3" json:"actual_arrive_time,omitempty"`              // 客户到达时间
	ConsumerAmount       int64  `protobuf:"varint,9,opt,name=consumer_amount,json=consumerAmount,proto3" json:"consumer_amount,omitempty"`                      // 消费金额
	CreateTime           int64  `protobuf:"varint,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`                                 // 创建时间
	RelationTask         bool   `protobuf:"varint,11,opt,name=relation_task,json=relationTask,proto3" json:"relation_task,omitempty"`                           // 是否关联任务
	RelationTaskUserId   int64  `protobuf:"varint,12,opt,name=relation_task_user_id,json=relationTaskUserId,proto3" json:"relation_task_user_id,omitempty"`     // 关联用户id
	RelationTaskUsername int64  `protobuf:"varint,13,opt,name=relation_task_username,json=relationTaskUsername,proto3" json:"relation_task_username,omitempty"` // 关联用户名称
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_proto_common_common_proto_rawDescGZIP(), []int{3}
}

func (x *Report) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Report) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Report) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Report) GetExceptArriveTime() int64 {
	if x != nil {
		return x.ExceptArriveTime
	}
	return 0
}

func (x *Report) GetConsumerMobile() string {
	if x != nil {
		return x.ConsumerMobile
	}
	return ""
}

func (x *Report) GetConsumerName() string {
	if x != nil {
		return x.ConsumerName
	}
	return ""
}

func (x *Report) GetIsMatch() bool {
	if x != nil {
		return x.IsMatch
	}
	return false
}

func (x *Report) GetActualArriveTime() int64 {
	if x != nil {
		return x.ActualArriveTime
	}
	return 0
}

func (x *Report) GetConsumerAmount() int64 {
	if x != nil {
		return x.ConsumerAmount
	}
	return 0
}

func (x *Report) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Report) GetRelationTask() bool {
	if x != nil {
		return x.RelationTask
	}
	return false
}

func (x *Report) GetRelationTaskUserId() int64 {
	if x != nil {
		return x.RelationTaskUserId
	}
	return 0
}

func (x *Report) GetRelationTaskUsername() int64 {
	if x != nil {
		return x.RelationTaskUsername
	}
	return 0
}

// 通过的 error 信息
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 错误 code.
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// 错误提示信息, 可用于提示文案.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// 错误的详细信息, 可用于定位错误详情.
	Details string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_proto_common_common_proto_rawDescGZIP(), []int{4}
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

type CustomerServerHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务id
	TaskId int64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// 时间.
	Time int64 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	// 备注.
	Desc string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	// 客服id
	UserId int64 `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 客服名称
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CustomerServerHistory) Reset() {
	*x = CustomerServerHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerServerHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerServerHistory) ProtoMessage() {}

func (x *CustomerServerHistory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerServerHistory.ProtoReflect.Descriptor instead.
func (*CustomerServerHistory) Descriptor() ([]byte, []int) {
	return file_proto_common_common_proto_rawDescGZIP(), []int{5}
}

func (x *CustomerServerHistory) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *CustomerServerHistory) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *CustomerServerHistory) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CustomerServerHistory) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CustomerServerHistory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CustomerServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId int64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// 报单id
	ReportId int64 `protobuf:"varint,2,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
	// 对应客服id.
	UserId int64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 客服名称.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// 分配任务时间.
	DistributeTime int64 `protobuf:"varint,5,opt,name=distribute_time,json=distributeTime,proto3" json:"distribute_time,omitempty"`
	// 顾客姓名.
	CustomerName string `protobuf:"bytes,6,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	// 消费金额.
	CustomerAmount int64 `protobuf:"varint,7,opt,name=customer_amount,json=customerAmount,proto3" json:"customer_amount,omitempty"`
	// 是否完成. 1是 2否
	IsFinished int64 `protobuf:"varint,8,opt,name=is_finished,json=isFinished,proto3" json:"is_finished,omitempty"`
	// 任务历史.
	History []*CustomerServerHistory `protobuf:"bytes,9,rep,name=history,proto3" json:"history,omitempty"`
	// 最新备注.
	LastDesc string `protobuf:"bytes,10,opt,name=last_desc,json=lastDesc,proto3" json:"last_desc,omitempty"`
	// 完成时间
	FinishTime int64 `protobuf:"varint,11,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
	// 顾客电话
	CustomerMobile string `protobuf:"bytes,12,opt,name=customer_mobile,json=customerMobile,proto3" json:"customer_mobile,omitempty"`
}

func (x *CustomerServerInfo) Reset() {
	*x = CustomerServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerServerInfo) ProtoMessage() {}

func (x *CustomerServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerServerInfo.ProtoReflect.Descriptor instead.
func (*CustomerServerInfo) Descriptor() ([]byte, []int) {
	return file_proto_common_common_proto_rawDescGZIP(), []int{6}
}

func (x *CustomerServerInfo) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *CustomerServerInfo) GetReportId() int64 {
	if x != nil {
		return x.ReportId
	}
	return 0
}

func (x *CustomerServerInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CustomerServerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomerServerInfo) GetDistributeTime() int64 {
	if x != nil {
		return x.DistributeTime
	}
	return 0
}

func (x *CustomerServerInfo) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

func (x *CustomerServerInfo) GetCustomerAmount() int64 {
	if x != nil {
		return x.CustomerAmount
	}
	return 0
}

func (x *CustomerServerInfo) GetIsFinished() int64 {
	if x != nil {
		return x.IsFinished
	}
	return 0
}

func (x *CustomerServerInfo) GetHistory() []*CustomerServerHistory {
	if x != nil {
		return x.History
	}
	return nil
}

func (x *CustomerServerInfo) GetLastDesc() string {
	if x != nil {
		return x.LastDesc
	}
	return ""
}

func (x *CustomerServerInfo) GetFinishTime() int64 {
	if x != nil {
		return x.FinishTime
	}
	return 0
}

func (x *CustomerServerInfo) GetCustomerMobile() string {
	if x != nil {
		return x.CustomerMobile
	}
	return ""
}

var File_proto_common_common_proto protoreflect.FileDescriptor

var file_proto_common_common_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x46, 0x0a, 0x04,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x22, 0xeb, 0x03, 0x0a, 0x06, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x78, 0x63,
	0x65, 0x70, 0x74, 0x5f, 0x61, 0x72, 0x72, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x41, 0x72, 0x72,
	0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x2c, 0x0a, 0x12, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x61, 0x72, 0x72, 0x69, 0x76,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x61, 0x63,
	0x74, 0x75, 0x61, 0x6c, 0x41, 0x72, 0x72, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x31, 0x0a,
	0x15, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x34, 0x0a, 0x16, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x14, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x15, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xaf, 0x03, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x12, 0x37, 0x0a, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x58, 0x48, 0x58, 0x48, 0x58, 0x48, 0x58, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x5f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_common_common_proto_rawDescOnce sync.Once
	file_proto_common_common_proto_rawDescData = file_proto_common_common_proto_rawDesc
)

func file_proto_common_common_proto_rawDescGZIP() []byte {
	file_proto_common_common_proto_rawDescOnce.Do(func() {
		file_proto_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_common_common_proto_rawDescData)
	})
	return file_proto_common_common_proto_rawDescData
}

var file_proto_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_common_common_proto_goTypes = []interface{}{
	(*Empty)(nil),                 // 0: common.Empty
	(*Page)(nil),                  // 1: common.Page
	(*Pagination)(nil),            // 2: common.Pagination
	(*Report)(nil),                // 3: common.Report
	(*Error)(nil),                 // 4: common.Error
	(*CustomerServerHistory)(nil), // 5: common.CustomerServerHistory
	(*CustomerServerInfo)(nil),    // 6: common.CustomerServerInfo
}
var file_proto_common_common_proto_depIdxs = []int32{
	5, // 0: common.CustomerServerInfo.history:type_name -> common.CustomerServerHistory
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_common_common_proto_init() }
func file_proto_common_common_proto_init() {
	if File_proto_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_common_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_proto_common_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_proto_common_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
		file_proto_common_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_proto_common_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerServerHistory); i {
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
		file_proto_common_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerServerInfo); i {
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
			RawDescriptor: file_proto_common_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_common_common_proto_goTypes,
		DependencyIndexes: file_proto_common_common_proto_depIdxs,
		MessageInfos:      file_proto_common_common_proto_msgTypes,
	}.Build()
	File_proto_common_common_proto = out.File
	file_proto_common_common_proto_rawDesc = nil
	file_proto_common_common_proto_goTypes = nil
	file_proto_common_common_proto_depIdxs = nil
}
