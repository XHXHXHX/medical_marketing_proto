// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1api

import (
	context "context"
	common "github.com/XHXHXHX/medical_marketing_proto/gen/go/proto/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ApiReportServiceClient is the client API for ApiReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiReportServiceClient interface {
	// 报单登记.
	ReportCreate(ctx context.Context, in *ReportCreateRequest, opts ...grpc.CallOption) (*common.Empty, error)
	// 报单撤销.
	ReportRecover(ctx context.Context, in *ReportRecoverRequest, opts ...grpc.CallOption) (*common.Empty, error)
	// 报单列表.
	ReportList(ctx context.Context, in *ReportListRequest, opts ...grpc.CallOption) (*ReportListResponse, error)
	// 报单到访时间修改.
	ReportChangeActualArrivedTime(ctx context.Context, in *ReportChangeActualArrivedTimeRequest, opts ...grpc.CallOption) (*common.Empty, error)
	// 报单匹配修改.
	ReportChangeMatch(ctx context.Context, in *ReportChangeMatchRequest, opts ...grpc.CallOption) (*common.Empty, error)
}

type apiReportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiReportServiceClient(cc grpc.ClientConnInterface) ApiReportServiceClient {
	return &apiReportServiceClient{cc}
}

func (c *apiReportServiceClient) ReportCreate(ctx context.Context, in *ReportCreateRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/v1api.report.apiReportService/ReportCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiReportServiceClient) ReportRecover(ctx context.Context, in *ReportRecoverRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/v1api.report.apiReportService/ReportRecover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiReportServiceClient) ReportList(ctx context.Context, in *ReportListRequest, opts ...grpc.CallOption) (*ReportListResponse, error) {
	out := new(ReportListResponse)
	err := c.cc.Invoke(ctx, "/v1api.report.apiReportService/ReportList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiReportServiceClient) ReportChangeActualArrivedTime(ctx context.Context, in *ReportChangeActualArrivedTimeRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/v1api.report.apiReportService/ReportChangeActualArrivedTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiReportServiceClient) ReportChangeMatch(ctx context.Context, in *ReportChangeMatchRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/v1api.report.apiReportService/ReportChangeMatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiReportServiceServer is the server API for ApiReportService service.
// All implementations must embed UnimplementedApiReportServiceServer
// for forward compatibility
type ApiReportServiceServer interface {
	// 报单登记.
	ReportCreate(context.Context, *ReportCreateRequest) (*common.Empty, error)
	// 报单撤销.
	ReportRecover(context.Context, *ReportRecoverRequest) (*common.Empty, error)
	// 报单列表.
	ReportList(context.Context, *ReportListRequest) (*ReportListResponse, error)
	// 报单到访时间修改.
	ReportChangeActualArrivedTime(context.Context, *ReportChangeActualArrivedTimeRequest) (*common.Empty, error)
	// 报单匹配修改.
	ReportChangeMatch(context.Context, *ReportChangeMatchRequest) (*common.Empty, error)
	mustEmbedUnimplementedApiReportServiceServer()
}

// UnimplementedApiReportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiReportServiceServer struct {
}

func (UnimplementedApiReportServiceServer) ReportCreate(context.Context, *ReportCreateRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportCreate not implemented")
}
func (UnimplementedApiReportServiceServer) ReportRecover(context.Context, *ReportRecoverRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportRecover not implemented")
}
func (UnimplementedApiReportServiceServer) ReportList(context.Context, *ReportListRequest) (*ReportListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportList not implemented")
}
func (UnimplementedApiReportServiceServer) ReportChangeActualArrivedTime(context.Context, *ReportChangeActualArrivedTimeRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportChangeActualArrivedTime not implemented")
}
func (UnimplementedApiReportServiceServer) ReportChangeMatch(context.Context, *ReportChangeMatchRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportChangeMatch not implemented")
}
func (UnimplementedApiReportServiceServer) mustEmbedUnimplementedApiReportServiceServer() {}

// UnsafeApiReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiReportServiceServer will
// result in compilation errors.
type UnsafeApiReportServiceServer interface {
	mustEmbedUnimplementedApiReportServiceServer()
}

func RegisterApiReportServiceServer(s grpc.ServiceRegistrar, srv ApiReportServiceServer) {
	s.RegisterService(&ApiReportService_ServiceDesc, srv)
}

func _ApiReportService_ReportCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiReportServiceServer).ReportCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1api.report.apiReportService/ReportCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiReportServiceServer).ReportCreate(ctx, req.(*ReportCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiReportService_ReportRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiReportServiceServer).ReportRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1api.report.apiReportService/ReportRecover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiReportServiceServer).ReportRecover(ctx, req.(*ReportRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiReportService_ReportList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiReportServiceServer).ReportList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1api.report.apiReportService/ReportList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiReportServiceServer).ReportList(ctx, req.(*ReportListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiReportService_ReportChangeActualArrivedTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportChangeActualArrivedTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiReportServiceServer).ReportChangeActualArrivedTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1api.report.apiReportService/ReportChangeActualArrivedTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiReportServiceServer).ReportChangeActualArrivedTime(ctx, req.(*ReportChangeActualArrivedTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiReportService_ReportChangeMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportChangeMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiReportServiceServer).ReportChangeMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1api.report.apiReportService/ReportChangeMatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiReportServiceServer).ReportChangeMatch(ctx, req.(*ReportChangeMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiReportService_ServiceDesc is the grpc.ServiceDesc for ApiReportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiReportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1api.report.apiReportService",
	HandlerType: (*ApiReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportCreate",
			Handler:    _ApiReportService_ReportCreate_Handler,
		},
		{
			MethodName: "ReportRecover",
			Handler:    _ApiReportService_ReportRecover_Handler,
		},
		{
			MethodName: "ReportList",
			Handler:    _ApiReportService_ReportList_Handler,
		},
		{
			MethodName: "ReportChangeActualArrivedTime",
			Handler:    _ApiReportService_ReportChangeActualArrivedTime_Handler,
		},
		{
			MethodName: "ReportChangeMatch",
			Handler:    _ApiReportService_ReportChangeMatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1api/report.proto",
}
