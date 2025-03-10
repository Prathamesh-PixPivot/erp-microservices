// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: inventory.proto

package inventory_pb

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
	InventoryService_CreateInventoryItem_FullMethodName      = "/inventory.InventoryService/CreateInventoryItem"
	InventoryService_GetInventoryItem_FullMethodName         = "/inventory.InventoryService/GetInventoryItem"
	InventoryService_UpdateInventoryItem_FullMethodName      = "/inventory.InventoryService/UpdateInventoryItem"
	InventoryService_DeleteInventoryItem_FullMethodName      = "/inventory.InventoryService/DeleteInventoryItem"
	InventoryService_ListInventoryItems_FullMethodName       = "/inventory.InventoryService/ListInventoryItems"
	InventoryService_TrackInventory_FullMethodName           = "/inventory.InventoryService/TrackInventory"
	InventoryService_SetReorderPoint_FullMethodName          = "/inventory.InventoryService/SetReorderPoint"
	InventoryService_ManageWarehouses_FullMethodName         = "/inventory.InventoryService/ManageWarehouses"
	InventoryService_AddOrUpdateInventoryItem_FullMethodName = "/inventory.InventoryService/AddOrUpdateInventoryItem"
	InventoryService_ProcessOrder_FullMethodName             = "/inventory.InventoryService/ProcessOrder"
	InventoryService_GeneratePickingList_FullMethodName      = "/inventory.InventoryService/GeneratePickingList"
	InventoryService_UpdateInventory_FullMethodName          = "/inventory.InventoryService/UpdateInventory"
	InventoryService_PlaceVendorOrder_FullMethodName         = "/inventory.InventoryService/PlaceVendorOrder"
	InventoryService_NotifyFinanceForOrder_FullMethodName    = "/inventory.InventoryService/NotifyFinanceForOrder"
)

// InventoryServiceClient is the client API for InventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Inventory Management Service
type InventoryServiceClient interface {
	// CRUD for Inventory Items
	CreateInventoryItem(ctx context.Context, in *CreateInventoryItemRequest, opts ...grpc.CallOption) (*CreateInventoryItemResponse, error)
	GetInventoryItem(ctx context.Context, in *GetInventoryItemRequest, opts ...grpc.CallOption) (*GetInventoryItemResponse, error)
	UpdateInventoryItem(ctx context.Context, in *UpdateInventoryItemRequest, opts ...grpc.CallOption) (*UpdateInventoryItemResponse, error)
	DeleteInventoryItem(ctx context.Context, in *DeleteInventoryItemRequest, opts ...grpc.CallOption) (*DeleteInventoryItemResponse, error)
	ListInventoryItems(ctx context.Context, in *ListInventoryItemsRequest, opts ...grpc.CallOption) (*ListInventoryItemsResponse, error)
	// Stock Management Methods
	TrackInventory(ctx context.Context, in *TrackInventoryRequest, opts ...grpc.CallOption) (*TrackInventoryResponse, error)
	SetReorderPoint(ctx context.Context, in *SetReorderPointRequest, opts ...grpc.CallOption) (*SetReorderPointResponse, error)
	ManageWarehouses(ctx context.Context, in *ManageWarehousesRequest, opts ...grpc.CallOption) (*ManageWarehousesResponse, error)
	AddOrUpdateInventoryItem(ctx context.Context, in *AddOrUpdateInventoryItemRequest, opts ...grpc.CallOption) (*AddOrUpdateInventoryItemResponse, error)
	// Order Fulfillment Methods
	ProcessOrder(ctx context.Context, in *ProcessOrderRequest, opts ...grpc.CallOption) (*ProcessOrderResponse, error)
	GeneratePickingList(ctx context.Context, in *GeneratePickingListRequest, opts ...grpc.CallOption) (*GeneratePickingListResponse, error)
	UpdateInventory(ctx context.Context, in *UpdateInventoryRequest, opts ...grpc.CallOption) (*UpdateInventoryResponse, error)
	// Integration with Vendor Management
	PlaceVendorOrder(ctx context.Context, in *PlaceVendorOrderRequest, opts ...grpc.CallOption) (*PlaceVendorOrderResponse, error)
	// Integration with Finance Service
	NotifyFinanceForOrder(ctx context.Context, in *NotifyFinanceRequest, opts ...grpc.CallOption) (*NotifyFinanceResponse, error)
}

type inventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryServiceClient(cc grpc.ClientConnInterface) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) CreateInventoryItem(ctx context.Context, in *CreateInventoryItemRequest, opts ...grpc.CallOption) (*CreateInventoryItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateInventoryItemResponse)
	err := c.cc.Invoke(ctx, InventoryService_CreateInventoryItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GetInventoryItem(ctx context.Context, in *GetInventoryItemRequest, opts ...grpc.CallOption) (*GetInventoryItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInventoryItemResponse)
	err := c.cc.Invoke(ctx, InventoryService_GetInventoryItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) UpdateInventoryItem(ctx context.Context, in *UpdateInventoryItemRequest, opts ...grpc.CallOption) (*UpdateInventoryItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateInventoryItemResponse)
	err := c.cc.Invoke(ctx, InventoryService_UpdateInventoryItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) DeleteInventoryItem(ctx context.Context, in *DeleteInventoryItemRequest, opts ...grpc.CallOption) (*DeleteInventoryItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteInventoryItemResponse)
	err := c.cc.Invoke(ctx, InventoryService_DeleteInventoryItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) ListInventoryItems(ctx context.Context, in *ListInventoryItemsRequest, opts ...grpc.CallOption) (*ListInventoryItemsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListInventoryItemsResponse)
	err := c.cc.Invoke(ctx, InventoryService_ListInventoryItems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) TrackInventory(ctx context.Context, in *TrackInventoryRequest, opts ...grpc.CallOption) (*TrackInventoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TrackInventoryResponse)
	err := c.cc.Invoke(ctx, InventoryService_TrackInventory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) SetReorderPoint(ctx context.Context, in *SetReorderPointRequest, opts ...grpc.CallOption) (*SetReorderPointResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetReorderPointResponse)
	err := c.cc.Invoke(ctx, InventoryService_SetReorderPoint_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) ManageWarehouses(ctx context.Context, in *ManageWarehousesRequest, opts ...grpc.CallOption) (*ManageWarehousesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ManageWarehousesResponse)
	err := c.cc.Invoke(ctx, InventoryService_ManageWarehouses_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) AddOrUpdateInventoryItem(ctx context.Context, in *AddOrUpdateInventoryItemRequest, opts ...grpc.CallOption) (*AddOrUpdateInventoryItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddOrUpdateInventoryItemResponse)
	err := c.cc.Invoke(ctx, InventoryService_AddOrUpdateInventoryItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) ProcessOrder(ctx context.Context, in *ProcessOrderRequest, opts ...grpc.CallOption) (*ProcessOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProcessOrderResponse)
	err := c.cc.Invoke(ctx, InventoryService_ProcessOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GeneratePickingList(ctx context.Context, in *GeneratePickingListRequest, opts ...grpc.CallOption) (*GeneratePickingListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GeneratePickingListResponse)
	err := c.cc.Invoke(ctx, InventoryService_GeneratePickingList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) UpdateInventory(ctx context.Context, in *UpdateInventoryRequest, opts ...grpc.CallOption) (*UpdateInventoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateInventoryResponse)
	err := c.cc.Invoke(ctx, InventoryService_UpdateInventory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) PlaceVendorOrder(ctx context.Context, in *PlaceVendorOrderRequest, opts ...grpc.CallOption) (*PlaceVendorOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlaceVendorOrderResponse)
	err := c.cc.Invoke(ctx, InventoryService_PlaceVendorOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) NotifyFinanceForOrder(ctx context.Context, in *NotifyFinanceRequest, opts ...grpc.CallOption) (*NotifyFinanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NotifyFinanceResponse)
	err := c.cc.Invoke(ctx, InventoryService_NotifyFinanceForOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServiceServer is the server API for InventoryService service.
// All implementations must embed UnimplementedInventoryServiceServer
// for forward compatibility.
//
// Inventory Management Service
type InventoryServiceServer interface {
	// CRUD for Inventory Items
	CreateInventoryItem(context.Context, *CreateInventoryItemRequest) (*CreateInventoryItemResponse, error)
	GetInventoryItem(context.Context, *GetInventoryItemRequest) (*GetInventoryItemResponse, error)
	UpdateInventoryItem(context.Context, *UpdateInventoryItemRequest) (*UpdateInventoryItemResponse, error)
	DeleteInventoryItem(context.Context, *DeleteInventoryItemRequest) (*DeleteInventoryItemResponse, error)
	ListInventoryItems(context.Context, *ListInventoryItemsRequest) (*ListInventoryItemsResponse, error)
	// Stock Management Methods
	TrackInventory(context.Context, *TrackInventoryRequest) (*TrackInventoryResponse, error)
	SetReorderPoint(context.Context, *SetReorderPointRequest) (*SetReorderPointResponse, error)
	ManageWarehouses(context.Context, *ManageWarehousesRequest) (*ManageWarehousesResponse, error)
	AddOrUpdateInventoryItem(context.Context, *AddOrUpdateInventoryItemRequest) (*AddOrUpdateInventoryItemResponse, error)
	// Order Fulfillment Methods
	ProcessOrder(context.Context, *ProcessOrderRequest) (*ProcessOrderResponse, error)
	GeneratePickingList(context.Context, *GeneratePickingListRequest) (*GeneratePickingListResponse, error)
	UpdateInventory(context.Context, *UpdateInventoryRequest) (*UpdateInventoryResponse, error)
	// Integration with Vendor Management
	PlaceVendorOrder(context.Context, *PlaceVendorOrderRequest) (*PlaceVendorOrderResponse, error)
	// Integration with Finance Service
	NotifyFinanceForOrder(context.Context, *NotifyFinanceRequest) (*NotifyFinanceResponse, error)
	mustEmbedUnimplementedInventoryServiceServer()
}

// UnimplementedInventoryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInventoryServiceServer struct{}

func (UnimplementedInventoryServiceServer) CreateInventoryItem(context.Context, *CreateInventoryItemRequest) (*CreateInventoryItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInventoryItem not implemented")
}
func (UnimplementedInventoryServiceServer) GetInventoryItem(context.Context, *GetInventoryItemRequest) (*GetInventoryItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInventoryItem not implemented")
}
func (UnimplementedInventoryServiceServer) UpdateInventoryItem(context.Context, *UpdateInventoryItemRequest) (*UpdateInventoryItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInventoryItem not implemented")
}
func (UnimplementedInventoryServiceServer) DeleteInventoryItem(context.Context, *DeleteInventoryItemRequest) (*DeleteInventoryItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInventoryItem not implemented")
}
func (UnimplementedInventoryServiceServer) ListInventoryItems(context.Context, *ListInventoryItemsRequest) (*ListInventoryItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInventoryItems not implemented")
}
func (UnimplementedInventoryServiceServer) TrackInventory(context.Context, *TrackInventoryRequest) (*TrackInventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackInventory not implemented")
}
func (UnimplementedInventoryServiceServer) SetReorderPoint(context.Context, *SetReorderPointRequest) (*SetReorderPointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetReorderPoint not implemented")
}
func (UnimplementedInventoryServiceServer) ManageWarehouses(context.Context, *ManageWarehousesRequest) (*ManageWarehousesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManageWarehouses not implemented")
}
func (UnimplementedInventoryServiceServer) AddOrUpdateInventoryItem(context.Context, *AddOrUpdateInventoryItemRequest) (*AddOrUpdateInventoryItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrUpdateInventoryItem not implemented")
}
func (UnimplementedInventoryServiceServer) ProcessOrder(context.Context, *ProcessOrderRequest) (*ProcessOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessOrder not implemented")
}
func (UnimplementedInventoryServiceServer) GeneratePickingList(context.Context, *GeneratePickingListRequest) (*GeneratePickingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePickingList not implemented")
}
func (UnimplementedInventoryServiceServer) UpdateInventory(context.Context, *UpdateInventoryRequest) (*UpdateInventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInventory not implemented")
}
func (UnimplementedInventoryServiceServer) PlaceVendorOrder(context.Context, *PlaceVendorOrderRequest) (*PlaceVendorOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceVendorOrder not implemented")
}
func (UnimplementedInventoryServiceServer) NotifyFinanceForOrder(context.Context, *NotifyFinanceRequest) (*NotifyFinanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyFinanceForOrder not implemented")
}
func (UnimplementedInventoryServiceServer) mustEmbedUnimplementedInventoryServiceServer() {}
func (UnimplementedInventoryServiceServer) testEmbeddedByValue()                          {}

// UnsafeInventoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServiceServer will
// result in compilation errors.
type UnsafeInventoryServiceServer interface {
	mustEmbedUnimplementedInventoryServiceServer()
}

func RegisterInventoryServiceServer(s grpc.ServiceRegistrar, srv InventoryServiceServer) {
	// If the following call pancis, it indicates UnimplementedInventoryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InventoryService_ServiceDesc, srv)
}

func _InventoryService_CreateInventoryItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInventoryItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).CreateInventoryItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_CreateInventoryItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).CreateInventoryItem(ctx, req.(*CreateInventoryItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GetInventoryItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInventoryItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetInventoryItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_GetInventoryItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetInventoryItem(ctx, req.(*GetInventoryItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_UpdateInventoryItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInventoryItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).UpdateInventoryItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_UpdateInventoryItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).UpdateInventoryItem(ctx, req.(*UpdateInventoryItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_DeleteInventoryItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInventoryItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).DeleteInventoryItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_DeleteInventoryItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).DeleteInventoryItem(ctx, req.(*DeleteInventoryItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_ListInventoryItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInventoryItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).ListInventoryItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_ListInventoryItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).ListInventoryItems(ctx, req.(*ListInventoryItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_TrackInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).TrackInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_TrackInventory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).TrackInventory(ctx, req.(*TrackInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_SetReorderPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReorderPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).SetReorderPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_SetReorderPoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).SetReorderPoint(ctx, req.(*SetReorderPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_ManageWarehouses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManageWarehousesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).ManageWarehouses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_ManageWarehouses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).ManageWarehouses(ctx, req.(*ManageWarehousesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_AddOrUpdateInventoryItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateInventoryItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).AddOrUpdateInventoryItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_AddOrUpdateInventoryItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).AddOrUpdateInventoryItem(ctx, req.(*AddOrUpdateInventoryItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_ProcessOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).ProcessOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_ProcessOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).ProcessOrder(ctx, req.(*ProcessOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GeneratePickingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneratePickingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GeneratePickingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_GeneratePickingList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GeneratePickingList(ctx, req.(*GeneratePickingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_UpdateInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).UpdateInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_UpdateInventory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).UpdateInventory(ctx, req.(*UpdateInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_PlaceVendorOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceVendorOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).PlaceVendorOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_PlaceVendorOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).PlaceVendorOrder(ctx, req.(*PlaceVendorOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_NotifyFinanceForOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyFinanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).NotifyFinanceForOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_NotifyFinanceForOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).NotifyFinanceForOrder(ctx, req.(*NotifyFinanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InventoryService_ServiceDesc is the grpc.ServiceDesc for InventoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInventoryItem",
			Handler:    _InventoryService_CreateInventoryItem_Handler,
		},
		{
			MethodName: "GetInventoryItem",
			Handler:    _InventoryService_GetInventoryItem_Handler,
		},
		{
			MethodName: "UpdateInventoryItem",
			Handler:    _InventoryService_UpdateInventoryItem_Handler,
		},
		{
			MethodName: "DeleteInventoryItem",
			Handler:    _InventoryService_DeleteInventoryItem_Handler,
		},
		{
			MethodName: "ListInventoryItems",
			Handler:    _InventoryService_ListInventoryItems_Handler,
		},
		{
			MethodName: "TrackInventory",
			Handler:    _InventoryService_TrackInventory_Handler,
		},
		{
			MethodName: "SetReorderPoint",
			Handler:    _InventoryService_SetReorderPoint_Handler,
		},
		{
			MethodName: "ManageWarehouses",
			Handler:    _InventoryService_ManageWarehouses_Handler,
		},
		{
			MethodName: "AddOrUpdateInventoryItem",
			Handler:    _InventoryService_AddOrUpdateInventoryItem_Handler,
		},
		{
			MethodName: "ProcessOrder",
			Handler:    _InventoryService_ProcessOrder_Handler,
		},
		{
			MethodName: "GeneratePickingList",
			Handler:    _InventoryService_GeneratePickingList_Handler,
		},
		{
			MethodName: "UpdateInventory",
			Handler:    _InventoryService_UpdateInventory_Handler,
		},
		{
			MethodName: "PlaceVendorOrder",
			Handler:    _InventoryService_PlaceVendorOrder_Handler,
		},
		{
			MethodName: "NotifyFinanceForOrder",
			Handler:    _InventoryService_NotifyFinanceForOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory.proto",
}
