// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.0--rc2
// source: asset_audit.proto

package proto

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
	AuditService_AuditAsset_FullMethodName      = "/audit.AuditService/AuditAsset"
	AuditService_GetAuditHistory_FullMethodName = "/audit.AuditService/GetAuditHistory"
)

// AuditServiceClient is the client API for AuditService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuditServiceClient interface {
	AuditAsset(ctx context.Context, in *AuditAssetRequest, opts ...grpc.CallOption) (*AuditResponse, error)
	GetAuditHistory(ctx context.Context, in *GetAuditHistoryRequest, opts ...grpc.CallOption) (*GetAuditHistoryResponse, error)
}

type auditServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuditServiceClient(cc grpc.ClientConnInterface) AuditServiceClient {
	return &auditServiceClient{cc}
}

func (c *auditServiceClient) AuditAsset(ctx context.Context, in *AuditAssetRequest, opts ...grpc.CallOption) (*AuditResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuditResponse)
	err := c.cc.Invoke(ctx, AuditService_AuditAsset_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditServiceClient) GetAuditHistory(ctx context.Context, in *GetAuditHistoryRequest, opts ...grpc.CallOption) (*GetAuditHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAuditHistoryResponse)
	err := c.cc.Invoke(ctx, AuditService_GetAuditHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditServiceServer is the server API for AuditService service.
// All implementations must embed UnimplementedAuditServiceServer
// for forward compatibility.
type AuditServiceServer interface {
	AuditAsset(context.Context, *AuditAssetRequest) (*AuditResponse, error)
	GetAuditHistory(context.Context, *GetAuditHistoryRequest) (*GetAuditHistoryResponse, error)
	mustEmbedUnimplementedAuditServiceServer()
}

// UnimplementedAuditServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuditServiceServer struct{}

func (UnimplementedAuditServiceServer) AuditAsset(context.Context, *AuditAssetRequest) (*AuditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuditAsset not implemented")
}
func (UnimplementedAuditServiceServer) GetAuditHistory(context.Context, *GetAuditHistoryRequest) (*GetAuditHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuditHistory not implemented")
}
func (UnimplementedAuditServiceServer) mustEmbedUnimplementedAuditServiceServer() {}
func (UnimplementedAuditServiceServer) testEmbeddedByValue()                      {}

// UnsafeAuditServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuditServiceServer will
// result in compilation errors.
type UnsafeAuditServiceServer interface {
	mustEmbedUnimplementedAuditServiceServer()
}

func RegisterAuditServiceServer(s grpc.ServiceRegistrar, srv AuditServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuditServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuditService_ServiceDesc, srv)
}

func _AuditService_AuditAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServiceServer).AuditAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuditService_AuditAsset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).AuditAsset(ctx, req.(*AuditAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuditService_GetAuditHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuditHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServiceServer).GetAuditHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuditService_GetAuditHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).GetAuditHistory(ctx, req.(*GetAuditHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuditService_ServiceDesc is the grpc.ServiceDesc for AuditService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuditService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "audit.AuditService",
	HandlerType: (*AuditServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuditAsset",
			Handler:    _AuditService_AuditAsset_Handler,
		},
		{
			MethodName: "GetAuditHistory",
			Handler:    _AuditService_GetAuditHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "asset_audit.proto",
}
