// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: internal/api/protobufs/gstr9.proto

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
	GSTR9Service_SaveGSTR9_FullMethodName      = "/GSTR9Service/SaveGSTR9"
	GSTR9Service_SubmitGSTR9_FullMethodName    = "/GSTR9Service/SubmitGSTR9"
	GSTR9Service_FileGSTR9_FullMethodName      = "/GSTR9Service/FileGSTR9"
	GSTR9Service_GetGSTR9Status_FullMethodName = "/GSTR9Service/GetGSTR9Status"
)

// GSTR9ServiceClient is the client API for GSTR9Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GSTR9ServiceClient interface {
	SaveGSTR9(ctx context.Context, in *GSTR9Request, opts ...grpc.CallOption) (*GSTR9Response, error)
	SubmitGSTR9(ctx context.Context, in *GSTR9SubmitRequest, opts ...grpc.CallOption) (*GSTR9SubmitResponse, error)
	FileGSTR9(ctx context.Context, in *GSTR9FileRequest, opts ...grpc.CallOption) (*GSTR9FileResponse, error)
	GetGSTR9Status(ctx context.Context, in *GSTR9StatusRequest, opts ...grpc.CallOption) (*GSTR9StatusResponse, error)
}

type gSTR9ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGSTR9ServiceClient(cc grpc.ClientConnInterface) GSTR9ServiceClient {
	return &gSTR9ServiceClient{cc}
}

func (c *gSTR9ServiceClient) SaveGSTR9(ctx context.Context, in *GSTR9Request, opts ...grpc.CallOption) (*GSTR9Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GSTR9Response)
	err := c.cc.Invoke(ctx, GSTR9Service_SaveGSTR9_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSTR9ServiceClient) SubmitGSTR9(ctx context.Context, in *GSTR9SubmitRequest, opts ...grpc.CallOption) (*GSTR9SubmitResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GSTR9SubmitResponse)
	err := c.cc.Invoke(ctx, GSTR9Service_SubmitGSTR9_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSTR9ServiceClient) FileGSTR9(ctx context.Context, in *GSTR9FileRequest, opts ...grpc.CallOption) (*GSTR9FileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GSTR9FileResponse)
	err := c.cc.Invoke(ctx, GSTR9Service_FileGSTR9_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gSTR9ServiceClient) GetGSTR9Status(ctx context.Context, in *GSTR9StatusRequest, opts ...grpc.CallOption) (*GSTR9StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GSTR9StatusResponse)
	err := c.cc.Invoke(ctx, GSTR9Service_GetGSTR9Status_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GSTR9ServiceServer is the server API for GSTR9Service service.
// All implementations must embed UnimplementedGSTR9ServiceServer
// for forward compatibility.
type GSTR9ServiceServer interface {
	SaveGSTR9(context.Context, *GSTR9Request) (*GSTR9Response, error)
	SubmitGSTR9(context.Context, *GSTR9SubmitRequest) (*GSTR9SubmitResponse, error)
	FileGSTR9(context.Context, *GSTR9FileRequest) (*GSTR9FileResponse, error)
	GetGSTR9Status(context.Context, *GSTR9StatusRequest) (*GSTR9StatusResponse, error)
	mustEmbedUnimplementedGSTR9ServiceServer()
}

// UnimplementedGSTR9ServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGSTR9ServiceServer struct{}

func (UnimplementedGSTR9ServiceServer) SaveGSTR9(context.Context, *GSTR9Request) (*GSTR9Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveGSTR9 not implemented")
}
func (UnimplementedGSTR9ServiceServer) SubmitGSTR9(context.Context, *GSTR9SubmitRequest) (*GSTR9SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitGSTR9 not implemented")
}
func (UnimplementedGSTR9ServiceServer) FileGSTR9(context.Context, *GSTR9FileRequest) (*GSTR9FileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileGSTR9 not implemented")
}
func (UnimplementedGSTR9ServiceServer) GetGSTR9Status(context.Context, *GSTR9StatusRequest) (*GSTR9StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGSTR9Status not implemented")
}
func (UnimplementedGSTR9ServiceServer) mustEmbedUnimplementedGSTR9ServiceServer() {}
func (UnimplementedGSTR9ServiceServer) testEmbeddedByValue()                      {}

// UnsafeGSTR9ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GSTR9ServiceServer will
// result in compilation errors.
type UnsafeGSTR9ServiceServer interface {
	mustEmbedUnimplementedGSTR9ServiceServer()
}

func RegisterGSTR9ServiceServer(s grpc.ServiceRegistrar, srv GSTR9ServiceServer) {
	// If the following call pancis, it indicates UnimplementedGSTR9ServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GSTR9Service_ServiceDesc, srv)
}

func _GSTR9Service_SaveGSTR9_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GSTR9Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSTR9ServiceServer).SaveGSTR9(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSTR9Service_SaveGSTR9_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSTR9ServiceServer).SaveGSTR9(ctx, req.(*GSTR9Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSTR9Service_SubmitGSTR9_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GSTR9SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSTR9ServiceServer).SubmitGSTR9(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSTR9Service_SubmitGSTR9_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSTR9ServiceServer).SubmitGSTR9(ctx, req.(*GSTR9SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSTR9Service_FileGSTR9_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GSTR9FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSTR9ServiceServer).FileGSTR9(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSTR9Service_FileGSTR9_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSTR9ServiceServer).FileGSTR9(ctx, req.(*GSTR9FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GSTR9Service_GetGSTR9Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GSTR9StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GSTR9ServiceServer).GetGSTR9Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GSTR9Service_GetGSTR9Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GSTR9ServiceServer).GetGSTR9Status(ctx, req.(*GSTR9StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GSTR9Service_ServiceDesc is the grpc.ServiceDesc for GSTR9Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GSTR9Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GSTR9Service",
	HandlerType: (*GSTR9ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveGSTR9",
			Handler:    _GSTR9Service_SaveGSTR9_Handler,
		},
		{
			MethodName: "SubmitGSTR9",
			Handler:    _GSTR9Service_SubmitGSTR9_Handler,
		},
		{
			MethodName: "FileGSTR9",
			Handler:    _GSTR9Service_FileGSTR9_Handler,
		},
		{
			MethodName: "GetGSTR9Status",
			Handler:    _GSTR9Service_GetGSTR9Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/api/protobufs/gstr9.proto",
}
