// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/skill_development.proto

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
	SkillDevelopmentService_CreateSkillDevelopment_FullMethodName = "/hrms.SkillDevelopmentService/CreateSkillDevelopment"
	SkillDevelopmentService_GetSkillDevelopment_FullMethodName    = "/hrms.SkillDevelopmentService/GetSkillDevelopment"
	SkillDevelopmentService_ListSkillDevelopments_FullMethodName  = "/hrms.SkillDevelopmentService/ListSkillDevelopments"
	SkillDevelopmentService_UpdateSkillDevelopment_FullMethodName = "/hrms.SkillDevelopmentService/UpdateSkillDevelopment"
	SkillDevelopmentService_DeleteSkillDevelopment_FullMethodName = "/hrms.SkillDevelopmentService/DeleteSkillDevelopment"
)

// SkillDevelopmentServiceClient is the client API for SkillDevelopmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// gRPC Service for Skill Development Management
type SkillDevelopmentServiceClient interface {
	CreateSkillDevelopment(ctx context.Context, in *CreateSkillDevelopmentRequest, opts ...grpc.CallOption) (*SkillDevelopmentResponse, error)
	GetSkillDevelopment(ctx context.Context, in *GetSkillDevelopmentRequest, opts ...grpc.CallOption) (*SkillDevelopmentResponse, error)
	ListSkillDevelopments(ctx context.Context, in *ListSkillDevelopmentsRequest, opts ...grpc.CallOption) (*ListSkillDevelopmentsResponse, error)
	UpdateSkillDevelopment(ctx context.Context, in *UpdateSkillDevelopmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteSkillDevelopment(ctx context.Context, in *DeleteSkillDevelopmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type skillDevelopmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSkillDevelopmentServiceClient(cc grpc.ClientConnInterface) SkillDevelopmentServiceClient {
	return &skillDevelopmentServiceClient{cc}
}

func (c *skillDevelopmentServiceClient) CreateSkillDevelopment(ctx context.Context, in *CreateSkillDevelopmentRequest, opts ...grpc.CallOption) (*SkillDevelopmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SkillDevelopmentResponse)
	err := c.cc.Invoke(ctx, SkillDevelopmentService_CreateSkillDevelopment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillDevelopmentServiceClient) GetSkillDevelopment(ctx context.Context, in *GetSkillDevelopmentRequest, opts ...grpc.CallOption) (*SkillDevelopmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SkillDevelopmentResponse)
	err := c.cc.Invoke(ctx, SkillDevelopmentService_GetSkillDevelopment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillDevelopmentServiceClient) ListSkillDevelopments(ctx context.Context, in *ListSkillDevelopmentsRequest, opts ...grpc.CallOption) (*ListSkillDevelopmentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSkillDevelopmentsResponse)
	err := c.cc.Invoke(ctx, SkillDevelopmentService_ListSkillDevelopments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillDevelopmentServiceClient) UpdateSkillDevelopment(ctx context.Context, in *UpdateSkillDevelopmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SkillDevelopmentService_UpdateSkillDevelopment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillDevelopmentServiceClient) DeleteSkillDevelopment(ctx context.Context, in *DeleteSkillDevelopmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SkillDevelopmentService_DeleteSkillDevelopment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SkillDevelopmentServiceServer is the server API for SkillDevelopmentService service.
// All implementations must embed UnimplementedSkillDevelopmentServiceServer
// for forward compatibility.
//
// gRPC Service for Skill Development Management
type SkillDevelopmentServiceServer interface {
	CreateSkillDevelopment(context.Context, *CreateSkillDevelopmentRequest) (*SkillDevelopmentResponse, error)
	GetSkillDevelopment(context.Context, *GetSkillDevelopmentRequest) (*SkillDevelopmentResponse, error)
	ListSkillDevelopments(context.Context, *ListSkillDevelopmentsRequest) (*ListSkillDevelopmentsResponse, error)
	UpdateSkillDevelopment(context.Context, *UpdateSkillDevelopmentRequest) (*emptypb.Empty, error)
	DeleteSkillDevelopment(context.Context, *DeleteSkillDevelopmentRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSkillDevelopmentServiceServer()
}

// UnimplementedSkillDevelopmentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSkillDevelopmentServiceServer struct{}

func (UnimplementedSkillDevelopmentServiceServer) CreateSkillDevelopment(context.Context, *CreateSkillDevelopmentRequest) (*SkillDevelopmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSkillDevelopment not implemented")
}
func (UnimplementedSkillDevelopmentServiceServer) GetSkillDevelopment(context.Context, *GetSkillDevelopmentRequest) (*SkillDevelopmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkillDevelopment not implemented")
}
func (UnimplementedSkillDevelopmentServiceServer) ListSkillDevelopments(context.Context, *ListSkillDevelopmentsRequest) (*ListSkillDevelopmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSkillDevelopments not implemented")
}
func (UnimplementedSkillDevelopmentServiceServer) UpdateSkillDevelopment(context.Context, *UpdateSkillDevelopmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSkillDevelopment not implemented")
}
func (UnimplementedSkillDevelopmentServiceServer) DeleteSkillDevelopment(context.Context, *DeleteSkillDevelopmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSkillDevelopment not implemented")
}
func (UnimplementedSkillDevelopmentServiceServer) mustEmbedUnimplementedSkillDevelopmentServiceServer() {
}
func (UnimplementedSkillDevelopmentServiceServer) testEmbeddedByValue() {}

// UnsafeSkillDevelopmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SkillDevelopmentServiceServer will
// result in compilation errors.
type UnsafeSkillDevelopmentServiceServer interface {
	mustEmbedUnimplementedSkillDevelopmentServiceServer()
}

func RegisterSkillDevelopmentServiceServer(s grpc.ServiceRegistrar, srv SkillDevelopmentServiceServer) {
	// If the following call pancis, it indicates UnimplementedSkillDevelopmentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SkillDevelopmentService_ServiceDesc, srv)
}

func _SkillDevelopmentService_CreateSkillDevelopment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSkillDevelopmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillDevelopmentServiceServer).CreateSkillDevelopment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillDevelopmentService_CreateSkillDevelopment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillDevelopmentServiceServer).CreateSkillDevelopment(ctx, req.(*CreateSkillDevelopmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillDevelopmentService_GetSkillDevelopment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSkillDevelopmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillDevelopmentServiceServer).GetSkillDevelopment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillDevelopmentService_GetSkillDevelopment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillDevelopmentServiceServer).GetSkillDevelopment(ctx, req.(*GetSkillDevelopmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillDevelopmentService_ListSkillDevelopments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSkillDevelopmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillDevelopmentServiceServer).ListSkillDevelopments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillDevelopmentService_ListSkillDevelopments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillDevelopmentServiceServer).ListSkillDevelopments(ctx, req.(*ListSkillDevelopmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillDevelopmentService_UpdateSkillDevelopment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSkillDevelopmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillDevelopmentServiceServer).UpdateSkillDevelopment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillDevelopmentService_UpdateSkillDevelopment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillDevelopmentServiceServer).UpdateSkillDevelopment(ctx, req.(*UpdateSkillDevelopmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillDevelopmentService_DeleteSkillDevelopment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSkillDevelopmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillDevelopmentServiceServer).DeleteSkillDevelopment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillDevelopmentService_DeleteSkillDevelopment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillDevelopmentServiceServer).DeleteSkillDevelopment(ctx, req.(*DeleteSkillDevelopmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SkillDevelopmentService_ServiceDesc is the grpc.ServiceDesc for SkillDevelopmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SkillDevelopmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hrms.SkillDevelopmentService",
	HandlerType: (*SkillDevelopmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSkillDevelopment",
			Handler:    _SkillDevelopmentService_CreateSkillDevelopment_Handler,
		},
		{
			MethodName: "GetSkillDevelopment",
			Handler:    _SkillDevelopmentService_GetSkillDevelopment_Handler,
		},
		{
			MethodName: "ListSkillDevelopments",
			Handler:    _SkillDevelopmentService_ListSkillDevelopments_Handler,
		},
		{
			MethodName: "UpdateSkillDevelopment",
			Handler:    _SkillDevelopmentService_UpdateSkillDevelopment_Handler,
		},
		{
			MethodName: "DeleteSkillDevelopment",
			Handler:    _SkillDevelopmentService_DeleteSkillDevelopment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/skill_development.proto",
}
