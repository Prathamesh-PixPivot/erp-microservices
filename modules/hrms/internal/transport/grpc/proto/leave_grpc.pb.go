// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/leave.proto

package hrms

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LeaveService_CreateLeave_FullMethodName       = "/hrms.LeaveService/CreateLeave"
	LeaveService_GetLeave_FullMethodName          = "/hrms.LeaveService/GetLeave"
	LeaveService_UpdateLeaveStatus_FullMethodName = "/hrms.LeaveService/UpdateLeaveStatus"
	LeaveService_ListLeaves_FullMethodName        = "/hrms.LeaveService/ListLeaves"
	LeaveService_DeleteLeave_FullMethodName       = "/hrms.LeaveService/DeleteLeave"
)

// LeaveServiceClient is the client API for LeaveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// gRPC Service for Leave Management
type LeaveServiceClient interface {
	CreateLeave(ctx context.Context, in *CreateLeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error)
	GetLeave(ctx context.Context, in *GetLeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error)
	UpdateLeaveStatus(ctx context.Context, in *UpdateLeaveStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListLeaves(ctx context.Context, in *ListLeavesRequest, opts ...grpc.CallOption) (*ListLeavesResponse, error)
	DeleteLeave(ctx context.Context, in *DeleteLeaveRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type leaveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLeaveServiceClient(cc grpc.ClientConnInterface) LeaveServiceClient {
	return &leaveServiceClient{cc}
}

func (c *leaveServiceClient) CreateLeave(ctx context.Context, in *CreateLeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LeaveResponse)
	err := c.cc.Invoke(ctx, LeaveService_CreateLeave_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaveServiceClient) GetLeave(ctx context.Context, in *GetLeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LeaveResponse)
	err := c.cc.Invoke(ctx, LeaveService_GetLeave_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaveServiceClient) UpdateLeaveStatus(ctx context.Context, in *UpdateLeaveStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LeaveService_UpdateLeaveStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaveServiceClient) ListLeaves(ctx context.Context, in *ListLeavesRequest, opts ...grpc.CallOption) (*ListLeavesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLeavesResponse)
	err := c.cc.Invoke(ctx, LeaveService_ListLeaves_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaveServiceClient) DeleteLeave(ctx context.Context, in *DeleteLeaveRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LeaveService_DeleteLeave_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeaveServiceServer is the server API for LeaveService service.
// All implementations must embed UnimplementedLeaveServiceServer
// for forward compatibility.
//
// gRPC Service for Leave Management
type LeaveServiceServer interface {
	CreateLeave(context.Context, *CreateLeaveRequest) (*LeaveResponse, error)
	GetLeave(context.Context, *GetLeaveRequest) (*LeaveResponse, error)
	UpdateLeaveStatus(context.Context, *UpdateLeaveStatusRequest) (*emptypb.Empty, error)
	ListLeaves(context.Context, *ListLeavesRequest) (*ListLeavesResponse, error)
	DeleteLeave(context.Context, *DeleteLeaveRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedLeaveServiceServer()
}

// UnimplementedLeaveServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLeaveServiceServer struct{}

func (UnimplementedLeaveServiceServer) CreateLeave(context.Context, *CreateLeaveRequest) (*LeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLeave not implemented")
}
func (UnimplementedLeaveServiceServer) GetLeave(context.Context, *GetLeaveRequest) (*LeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeave not implemented")
}
func (UnimplementedLeaveServiceServer) UpdateLeaveStatus(context.Context, *UpdateLeaveStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLeaveStatus not implemented")
}
func (UnimplementedLeaveServiceServer) ListLeaves(context.Context, *ListLeavesRequest) (*ListLeavesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLeaves not implemented")
}
func (UnimplementedLeaveServiceServer) DeleteLeave(context.Context, *DeleteLeaveRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLeave not implemented")
}
func (UnimplementedLeaveServiceServer) mustEmbedUnimplementedLeaveServiceServer() {}
func (UnimplementedLeaveServiceServer) testEmbeddedByValue()                      {}

// UnsafeLeaveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeaveServiceServer will
// result in compilation errors.
type UnsafeLeaveServiceServer interface {
	mustEmbedUnimplementedLeaveServiceServer()
}

func RegisterLeaveServiceServer(s grpc.ServiceRegistrar, srv LeaveServiceServer) {
	// If the following call pancis, it indicates UnimplementedLeaveServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LeaveService_ServiceDesc, srv)
}

func _LeaveService_CreateLeave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaveServiceServer).CreateLeave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LeaveService_CreateLeave_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaveServiceServer).CreateLeave(ctx, req.(*CreateLeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeaveService_GetLeave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaveServiceServer).GetLeave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LeaveService_GetLeave_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaveServiceServer).GetLeave(ctx, req.(*GetLeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeaveService_UpdateLeaveStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLeaveStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaveServiceServer).UpdateLeaveStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LeaveService_UpdateLeaveStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaveServiceServer).UpdateLeaveStatus(ctx, req.(*UpdateLeaveStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeaveService_ListLeaves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLeavesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaveServiceServer).ListLeaves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LeaveService_ListLeaves_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaveServiceServer).ListLeaves(ctx, req.(*ListLeavesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeaveService_DeleteLeave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaveServiceServer).DeleteLeave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LeaveService_DeleteLeave_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaveServiceServer).DeleteLeave(ctx, req.(*DeleteLeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LeaveService_ServiceDesc is the grpc.ServiceDesc for LeaveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LeaveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hrms.LeaveService",
	HandlerType: (*LeaveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLeave",
			Handler:    _LeaveService_CreateLeave_Handler,
		},
		{
			MethodName: "GetLeave",
			Handler:    _LeaveService_GetLeave_Handler,
		},
		{
			MethodName: "UpdateLeaveStatus",
			Handler:    _LeaveService_UpdateLeaveStatus_Handler,
		},
		{
			MethodName: "ListLeaves",
			Handler:    _LeaveService_ListLeaves_Handler,
		},
		{
			MethodName: "DeleteLeave",
			Handler:    _LeaveService_DeleteLeave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/leave.proto",
}
