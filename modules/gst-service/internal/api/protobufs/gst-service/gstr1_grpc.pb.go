// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: internal/api/protobufs/gstr1.proto

package protobufs

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GSTR1Service_SaveGSTR1_FullMethodName      = "/GSTR1Service/SaveGSTR1"
	GSTR1Service_SubmitGSTR1_FullMethodName    = "/GSTR1Service/SubmitGSTR1"
	GSTR1Service_FileGSTR1_FullMethodName      = "/GSTR1Service/FileGSTR1"
	GSTR1Service_GetGSTR1Status_FullMethodName = "/GSTR1Service/GetGSTR1Status"
	GSTR1Service_ReconcileGSTR1_FullMethodName = "/GSTR1Service/ReconcileGSTR1"
)

// GSTR1ServiceClient is the client API for GSTR1Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GSTR1ServiceClient interface {
	SaveGSTR1(ctx context.Context, in *GSTR1Request, opts ...grpc.CallOption) (*GSTR1Response, error)
	SubmitGSTR1(ctx context.Context, in *GSTR1SubmitRequest, opts ...grpc.CallOption) (*GSTR1SubmitResponse, error)
	FileGSTR1(ctx context.Context, in *GSTR1FileRequest, opts ...grpc.CallOption) (*GSTR1FileResponse, error)
	GetGSTR1Status(ctx context.Context, in *GSTR1StatusRequest, opts ...grpc.CallOption) (*GSTR1StatusResponse, error)
	ReconcileGSTR1(ctx context.Context, in *GSTR1ReconcileRequest, opts ...grpc.CallOption) (*GSTR1ReconcileResponse, error)
}

type gSTR1ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGSTR1ServiceClient(cc grpc.ClientConnInterface) GSTR1ServiceClient {
	return &gSTR1ServiceClient{cc}
}

func (c *gSTR1ServiceClient) SaveGSTR1(ctx context.Context, in *GSTR1Request, opts ...grpc.CallOption) (*GSTR1Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GSTR1Response)
	err := c.cc.Invoke(ctx, GSTR1Service_SaveGSTR1_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSTR1ServiceClient) SubmitGSTR1(ctx context.Context, in *GSTR1SubmitRequest, opts ...grpc.CallOption) (*GSTR1SubmitResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GSTR1SubmitResponse)
	err := c.cc.Invoke(ctx, GSTR1Service_SubmitGSTR1_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSTR1ServiceClient) FileGSTR1(ctx context.Context, in *GSTR1FileRequest, opts ...grpc.CallOption) (*GSTR1FileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GSTR1FileResponse)
	err := c.cc.Invoke(ctx, GSTR1Service_FileGSTR1_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSTR1ServiceClient) GetGSTR1Status(ctx context.Context, in *GSTR1StatusRequest, opts ...grpc.CallOption) (*GSTR1StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GSTR1StatusResponse)
	err := c.cc.Invoke(ctx, GSTR1Service_GetGSTR1Status_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSTR1ServiceClient) ReconcileGSTR1(ctx context.Context, in *GSTR1ReconcileRequest, opts ...grpc.CallOption) (*GSTR1ReconcileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GSTR1ReconcileResponse)
	err := c.cc.Invoke(ctx, GSTR1Service_ReconcileGSTR1_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GSTR1ServiceServer is the server API for GSTR1Service service.
// All implementations must embed UnimplementedGSTR1ServiceServer
// for forward compatibility.
type GSTR1ServiceServer interface {
	SaveGSTR1(context.Context, *GSTR1Request) (*GSTR1Response, error)
	SubmitGSTR1(context.Context, *GSTR1SubmitRequest) (*GSTR1SubmitResponse, error)
	FileGSTR1(context.Context, *GSTR1FileRequest) (*GSTR1FileResponse, error)
	GetGSTR1Status(context.Context, *GSTR1StatusRequest) (*GSTR1StatusResponse, error)
	ReconcileGSTR1(context.Context, *GSTR1ReconcileRequest) (*GSTR1ReconcileResponse, error)
	mustEmbedUnimplementedGSTR1ServiceServer()
}

// UnimplementedGSTR1ServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGSTR1ServiceServer struct{}

func (UnimplementedGSTR1ServiceServer) SaveGSTR1(context.Context, *GSTR1Request) (*GSTR1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveGSTR1 not implemented")
}
func (UnimplementedGSTR1ServiceServer) SubmitGSTR1(context.Context, *GSTR1SubmitRequest) (*GSTR1SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitGSTR1 not implemented")
}
func (UnimplementedGSTR1ServiceServer) FileGSTR1(context.Context, *GSTR1FileRequest) (*GSTR1FileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileGSTR1 not implemented")
}
func (UnimplementedGSTR1ServiceServer) GetGSTR1Status(context.Context, *GSTR1StatusRequest) (*GSTR1StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGSTR1Status not implemented")
}
func (UnimplementedGSTR1ServiceServer) ReconcileGSTR1(context.Context, *GSTR1ReconcileRequest) (*GSTR1ReconcileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReconcileGSTR1 not implemented")
}
func (UnimplementedGSTR1ServiceServer) mustEmbedUnimplementedGSTR1ServiceServer() {}
func (UnimplementedGSTR1ServiceServer) testEmbeddedByValue()                      {}

// UnsafeGSTR1ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GSTR1ServiceServer will
// result in compilation errors.
type UnsafeGSTR1ServiceServer interface {
	mustEmbedUnimplementedGSTR1ServiceServer()
}

func RegisterGSTR1ServiceServer(s grpc.ServiceRegistrar, srv GSTR1ServiceServer) {
	// If the following call pancis, it indicates UnimplementedGSTR1ServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GSTR1Service_ServiceDesc, srv)
}

func _GSTR1Service_SaveGSTR1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GSTR1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSTR1ServiceServer).SaveGSTR1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSTR1Service_SaveGSTR1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSTR1ServiceServer).SaveGSTR1(ctx, req.(*GSTR1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSTR1Service_SubmitGSTR1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GSTR1SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSTR1ServiceServer).SubmitGSTR1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSTR1Service_SubmitGSTR1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSTR1ServiceServer).SubmitGSTR1(ctx, req.(*GSTR1SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSTR1Service_FileGSTR1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GSTR1FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSTR1ServiceServer).FileGSTR1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSTR1Service_FileGSTR1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSTR1ServiceServer).FileGSTR1(ctx, req.(*GSTR1FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSTR1Service_GetGSTR1Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GSTR1StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSTR1ServiceServer).GetGSTR1Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSTR1Service_GetGSTR1Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSTR1ServiceServer).GetGSTR1Status(ctx, req.(*GSTR1StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSTR1Service_ReconcileGSTR1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GSTR1ReconcileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSTR1ServiceServer).ReconcileGSTR1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSTR1Service_ReconcileGSTR1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSTR1ServiceServer).ReconcileGSTR1(ctx, req.(*GSTR1ReconcileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GSTR1Service_ServiceDesc is the grpc.ServiceDesc for GSTR1Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GSTR1Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GSTR1Service",
	HandlerType: (*GSTR1ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveGSTR1",
			Handler:    _GSTR1Service_SaveGSTR1_Handler,
		},
		{
			MethodName: "SubmitGSTR1",
			Handler:    _GSTR1Service_SubmitGSTR1_Handler,
		},
		{
			MethodName: "FileGSTR1",
			Handler:    _GSTR1Service_FileGSTR1_Handler,
		},
		{
			MethodName: "GetGSTR1Status",
			Handler:    _GSTR1Service_GetGSTR1Status_Handler,
		},
		{
			MethodName: "ReconcileGSTR1",
			Handler:    _GSTR1Service_ReconcileGSTR1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/api/protobufs/gstr1.proto",
}
