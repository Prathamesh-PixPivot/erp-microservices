// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/expense.proto

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
	ExpenseService_CreateExpense_FullMethodName       = "/hrms.ExpenseService/CreateExpense"
	ExpenseService_GetExpense_FullMethodName          = "/hrms.ExpenseService/GetExpense"
	ExpenseService_GetEmployeeExpenses_FullMethodName = "/hrms.ExpenseService/GetEmployeeExpenses"
	ExpenseService_UpdateExpenseStatus_FullMethodName = "/hrms.ExpenseService/UpdateExpenseStatus"
	ExpenseService_DeleteExpense_FullMethodName       = "/hrms.ExpenseService/DeleteExpense"
)

// ExpenseServiceClient is the client API for ExpenseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// gRPC Service for Expense Management
type ExpenseServiceClient interface {
	CreateExpense(ctx context.Context, in *CreateExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error)
	GetExpense(ctx context.Context, in *GetExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error)
	GetEmployeeExpenses(ctx context.Context, in *GetEmployeeExpensesRequest, opts ...grpc.CallOption) (*EmployeeExpensesResponse, error)
	UpdateExpenseStatus(ctx context.Context, in *UpdateExpenseStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteExpense(ctx context.Context, in *DeleteExpenseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type expenseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExpenseServiceClient(cc grpc.ClientConnInterface) ExpenseServiceClient {
	return &expenseServiceClient{cc}
}

func (c *expenseServiceClient) CreateExpense(ctx context.Context, in *CreateExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpenseResponse)
	err := c.cc.Invoke(ctx, ExpenseService_CreateExpense_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) GetExpense(ctx context.Context, in *GetExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpenseResponse)
	err := c.cc.Invoke(ctx, ExpenseService_GetExpense_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) GetEmployeeExpenses(ctx context.Context, in *GetEmployeeExpensesRequest, opts ...grpc.CallOption) (*EmployeeExpensesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmployeeExpensesResponse)
	err := c.cc.Invoke(ctx, ExpenseService_GetEmployeeExpenses_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) UpdateExpenseStatus(ctx context.Context, in *UpdateExpenseStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ExpenseService_UpdateExpenseStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) DeleteExpense(ctx context.Context, in *DeleteExpenseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ExpenseService_DeleteExpense_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpenseServiceServer is the server API for ExpenseService service.
// All implementations must embed UnimplementedExpenseServiceServer
// for forward compatibility.
//
// gRPC Service for Expense Management
type ExpenseServiceServer interface {
	CreateExpense(context.Context, *CreateExpenseRequest) (*ExpenseResponse, error)
	GetExpense(context.Context, *GetExpenseRequest) (*ExpenseResponse, error)
	GetEmployeeExpenses(context.Context, *GetEmployeeExpensesRequest) (*EmployeeExpensesResponse, error)
	UpdateExpenseStatus(context.Context, *UpdateExpenseStatusRequest) (*emptypb.Empty, error)
	DeleteExpense(context.Context, *DeleteExpenseRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedExpenseServiceServer()
}

// UnimplementedExpenseServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExpenseServiceServer struct{}

func (UnimplementedExpenseServiceServer) CreateExpense(context.Context, *CreateExpenseRequest) (*ExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExpense not implemented")
}
func (UnimplementedExpenseServiceServer) GetExpense(context.Context, *GetExpenseRequest) (*ExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpense not implemented")
}
func (UnimplementedExpenseServiceServer) GetEmployeeExpenses(context.Context, *GetEmployeeExpensesRequest) (*EmployeeExpensesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeExpenses not implemented")
}
func (UnimplementedExpenseServiceServer) UpdateExpenseStatus(context.Context, *UpdateExpenseStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExpenseStatus not implemented")
}
func (UnimplementedExpenseServiceServer) DeleteExpense(context.Context, *DeleteExpenseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExpense not implemented")
}
func (UnimplementedExpenseServiceServer) mustEmbedUnimplementedExpenseServiceServer() {}
func (UnimplementedExpenseServiceServer) testEmbeddedByValue()                        {}

// UnsafeExpenseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExpenseServiceServer will
// result in compilation errors.
type UnsafeExpenseServiceServer interface {
	mustEmbedUnimplementedExpenseServiceServer()
}

func RegisterExpenseServiceServer(s grpc.ServiceRegistrar, srv ExpenseServiceServer) {
	// If the following call pancis, it indicates UnimplementedExpenseServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExpenseService_ServiceDesc, srv)
}

func _ExpenseService_CreateExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).CreateExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExpenseService_CreateExpense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).CreateExpense(ctx, req.(*CreateExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_GetExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).GetExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExpenseService_GetExpense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).GetExpense(ctx, req.(*GetExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_GetEmployeeExpenses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmployeeExpensesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).GetEmployeeExpenses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExpenseService_GetEmployeeExpenses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).GetEmployeeExpenses(ctx, req.(*GetEmployeeExpensesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_UpdateExpenseStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExpenseStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).UpdateExpenseStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExpenseService_UpdateExpenseStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).UpdateExpenseStatus(ctx, req.(*UpdateExpenseStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_DeleteExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).DeleteExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExpenseService_DeleteExpense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).DeleteExpense(ctx, req.(*DeleteExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExpenseService_ServiceDesc is the grpc.ServiceDesc for ExpenseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExpenseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hrms.ExpenseService",
	HandlerType: (*ExpenseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExpense",
			Handler:    _ExpenseService_CreateExpense_Handler,
		},
		{
			MethodName: "GetExpense",
			Handler:    _ExpenseService_GetExpense_Handler,
		},
		{
			MethodName: "GetEmployeeExpenses",
			Handler:    _ExpenseService_GetEmployeeExpenses_Handler,
		},
		{
			MethodName: "UpdateExpenseStatus",
			Handler:    _ExpenseService_UpdateExpenseStatus_Handler,
		},
		{
			MethodName: "DeleteExpense",
			Handler:    _ExpenseService_DeleteExpense_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/expense.proto",
}
