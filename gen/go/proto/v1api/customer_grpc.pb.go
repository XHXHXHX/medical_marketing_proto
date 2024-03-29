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

// ApiCustomerServerServiceClient is the client API for ApiCustomerServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiCustomerServerServiceClient interface {
	// 给客服分配报单.
	CustomerServerDistribute(ctx context.Context, in *CustomerDistributeRequest, opts ...grpc.CallOption) (*common.Empty, error)
	// 查看客服分配列表.
	CustomerServerList(ctx context.Context, in *CustomerReportListRequest, opts ...grpc.CallOption) (*CustomerReportListResponse, error)
	// 客服回访结果.
	CustomerServerResult(ctx context.Context, in *CustomerServerResultRequest, opts ...grpc.CallOption) (*common.Empty, error)
}

type apiCustomerServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiCustomerServerServiceClient(cc grpc.ClientConnInterface) ApiCustomerServerServiceClient {
	return &apiCustomerServerServiceClient{cc}
}

func (c *apiCustomerServerServiceClient) CustomerServerDistribute(ctx context.Context, in *CustomerDistributeRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/v1api.customer.apiCustomerServerService/CustomerServerDistribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiCustomerServerServiceClient) CustomerServerList(ctx context.Context, in *CustomerReportListRequest, opts ...grpc.CallOption) (*CustomerReportListResponse, error) {
	out := new(CustomerReportListResponse)
	err := c.cc.Invoke(ctx, "/v1api.customer.apiCustomerServerService/CustomerServerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiCustomerServerServiceClient) CustomerServerResult(ctx context.Context, in *CustomerServerResultRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/v1api.customer.apiCustomerServerService/CustomerServerResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiCustomerServerServiceServer is the server API for ApiCustomerServerService service.
// All implementations must embed UnimplementedApiCustomerServerServiceServer
// for forward compatibility
type ApiCustomerServerServiceServer interface {
	// 给客服分配报单.
	CustomerServerDistribute(context.Context, *CustomerDistributeRequest) (*common.Empty, error)
	// 查看客服分配列表.
	CustomerServerList(context.Context, *CustomerReportListRequest) (*CustomerReportListResponse, error)
	// 客服回访结果.
	CustomerServerResult(context.Context, *CustomerServerResultRequest) (*common.Empty, error)
	mustEmbedUnimplementedApiCustomerServerServiceServer()
}

// UnimplementedApiCustomerServerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiCustomerServerServiceServer struct {
}

func (UnimplementedApiCustomerServerServiceServer) CustomerServerDistribute(context.Context, *CustomerDistributeRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomerServerDistribute not implemented")
}
func (UnimplementedApiCustomerServerServiceServer) CustomerServerList(context.Context, *CustomerReportListRequest) (*CustomerReportListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomerServerList not implemented")
}
func (UnimplementedApiCustomerServerServiceServer) CustomerServerResult(context.Context, *CustomerServerResultRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomerServerResult not implemented")
}
func (UnimplementedApiCustomerServerServiceServer) mustEmbedUnimplementedApiCustomerServerServiceServer() {
}

// UnsafeApiCustomerServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiCustomerServerServiceServer will
// result in compilation errors.
type UnsafeApiCustomerServerServiceServer interface {
	mustEmbedUnimplementedApiCustomerServerServiceServer()
}

func RegisterApiCustomerServerServiceServer(s grpc.ServiceRegistrar, srv ApiCustomerServerServiceServer) {
	s.RegisterService(&ApiCustomerServerService_ServiceDesc, srv)
}

func _ApiCustomerServerService_CustomerServerDistribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerDistributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiCustomerServerServiceServer).CustomerServerDistribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1api.customer.apiCustomerServerService/CustomerServerDistribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiCustomerServerServiceServer).CustomerServerDistribute(ctx, req.(*CustomerDistributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiCustomerServerService_CustomerServerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerReportListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiCustomerServerServiceServer).CustomerServerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1api.customer.apiCustomerServerService/CustomerServerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiCustomerServerServiceServer).CustomerServerList(ctx, req.(*CustomerReportListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiCustomerServerService_CustomerServerResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerServerResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiCustomerServerServiceServer).CustomerServerResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1api.customer.apiCustomerServerService/CustomerServerResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiCustomerServerServiceServer).CustomerServerResult(ctx, req.(*CustomerServerResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiCustomerServerService_ServiceDesc is the grpc.ServiceDesc for ApiCustomerServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiCustomerServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1api.customer.apiCustomerServerService",
	HandlerType: (*ApiCustomerServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CustomerServerDistribute",
			Handler:    _ApiCustomerServerService_CustomerServerDistribute_Handler,
		},
		{
			MethodName: "CustomerServerList",
			Handler:    _ApiCustomerServerService_CustomerServerList_Handler,
		},
		{
			MethodName: "CustomerServerResult",
			Handler:    _ApiCustomerServerService_CustomerServerResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1api/customer.proto",
}
