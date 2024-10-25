package grpc

import (
	"context"
	"inventory-service/internal/models"
	service "inventory-service/internal/services"
	"strconv"

	pb "inventory-service/grpc/inventory_pb" // Import the generated gRPC code from .proto
)

// InventoryServiceServer is the gRPC server struct
type InventoryServiceServer struct {
	service service.InventoryService
	pb.UnimplementedInventoryServiceServer
}

// NewInventoryServiceServer creates a new gRPC server instance
func NewInventoryServiceServer(svc service.InventoryService) *InventoryServiceServer {
	return &InventoryServiceServer{
		service: svc,
	}
}

// CRUD Functions

func (s *InventoryServiceServer) CreateInventoryItem(ctx context.Context, req *pb.CreateInventoryItemRequest) (*pb.CreateInventoryItemResponse, error) {
	item := &models.InventoryItem{
		ProductID:          req.Item.ProductId,
		ProductName:        req.Item.ProductName,
		ProductDescription: req.Item.ProductDescription,
		SKU:                req.Item.Sku,
		SupplierID:         req.Item.SupplierId,
		Category:           req.Item.Category,
		Price:              float64(req.Item.Price),
		AvailableQuantity:  req.Item.AvailableQuantity,
		ReorderPoint:       req.Item.ReorderPoint,
	}

	createdItem, err := s.service.CreateInventoryItem(item)
	if err != nil {
		return nil, err
	}

	return &pb.CreateInventoryItemResponse{
		Item: &pb.InventoryItem{
			ProductId:          createdItem.ProductID,
			ProductName:        createdItem.ProductName,
			ProductDescription: createdItem.ProductDescription,
			Sku:                createdItem.SKU,
			SupplierId:         createdItem.SupplierID,
			Category:           createdItem.Category,
			Price:              float32(createdItem.Price),
			AvailableQuantity:  createdItem.AvailableQuantity,
			ReorderPoint:       createdItem.ReorderPoint,
		},
	}, nil
}

func (s *InventoryServiceServer) GetInventoryItem(ctx context.Context, req *pb.GetInventoryItemRequest) (*pb.GetInventoryItemResponse, error) {
	item, err := s.service.GetInventoryItem(req.ProductId)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, nil
	}

	return &pb.GetInventoryItemResponse{
		Item: &pb.InventoryItem{
			ProductId:          item.ProductID,
			ProductName:        item.ProductName,
			ProductDescription: item.ProductDescription,
			Sku:                item.SKU,
			SupplierId:         item.SupplierID,
			Category:           item.Category,
			Price:              float32(item.Price),
			AvailableQuantity:  item.AvailableQuantity,
			ReorderPoint:       item.ReorderPoint,
		},
	}, nil
}

func (s *InventoryServiceServer) UpdateInventoryItem(ctx context.Context, req *pb.UpdateInventoryItemRequest) (*pb.UpdateInventoryItemResponse, error) {
	item := &models.InventoryItem{
		ProductID:          req.Item.ProductId,
		ProductName:        req.Item.ProductName,
		ProductDescription: req.Item.ProductDescription,
		SKU:                req.Item.Sku,
		SupplierID:         req.Item.SupplierId,
		Category:           req.Item.Category,
		Price:              float64(req.Item.Price),
		AvailableQuantity:  req.Item.AvailableQuantity,
		ReorderPoint:       req.Item.ReorderPoint,
	}

	updatedItem, err := s.service.UpdateInventoryItem(item)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateInventoryItemResponse{
		Item: &pb.InventoryItem{
			ProductId:          updatedItem.ProductID,
			ProductName:        updatedItem.ProductName,
			ProductDescription: updatedItem.ProductDescription,
			Sku:                updatedItem.SKU,
			SupplierId:         updatedItem.SupplierID,
			Category:           updatedItem.Category,
			Price:              float32(updatedItem.Price),
			AvailableQuantity:  updatedItem.AvailableQuantity,
			ReorderPoint:       updatedItem.ReorderPoint,
		},
	}, nil
}

func (s *InventoryServiceServer) DeleteInventoryItem(ctx context.Context, req *pb.DeleteInventoryItemRequest) (*pb.DeleteInventoryItemResponse, error) {
	err := s.service.DeleteInventoryItem(req.ProductId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteInventoryItemResponse{
		Success: true,
	}, nil
}

func (s *InventoryServiceServer) ListInventoryItems(ctx context.Context, req *pb.ListInventoryItemsRequest) (*pb.ListInventoryItemsResponse, error) {
	items, err := s.service.ListInventoryItems(int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	var pbItems []*pb.InventoryItem
	for _, item := range items {
		pbItems = append(pbItems, &pb.InventoryItem{
			ProductId:          item.ProductID,
			ProductName:        item.ProductName,
			ProductDescription: item.ProductDescription,
			Sku:                item.SKU,
			SupplierId:         item.SupplierID,
			Category:           item.Category,
			Price:              float32(item.Price),
			AvailableQuantity:  item.AvailableQuantity,
			ReorderPoint:       item.ReorderPoint,
		})
	}

	return &pb.ListInventoryItemsResponse{
		Items: pbItems,
	}, nil
}

// Stock Management

func (s *InventoryServiceServer) TrackInventory(ctx context.Context, req *pb.TrackInventoryRequest) (*pb.TrackInventoryResponse, error) {
	item, err := s.service.TrackInventory(req.ProductId)
	if err != nil {
		return nil, err
	}

	return &pb.TrackInventoryResponse{
		Item: &pb.InventoryItem{
			ProductId:          item.ProductID,
			ProductName:        item.ProductName,
			ProductDescription: item.ProductDescription,
			Sku:                item.SKU,
			SupplierId:         item.SupplierID,
			Category:           item.Category,
			Price:              float32(item.Price),
			AvailableQuantity:  item.AvailableQuantity,
			ReorderPoint:       item.ReorderPoint,
		},
	}, nil
}

func (s *InventoryServiceServer) SetReorderPoint(ctx context.Context, req *pb.SetReorderPointRequest) (*pb.SetReorderPointResponse, error) {
	err := s.service.SetReorderPoint(req.ProductId, req.ReorderPoint)
	if err != nil {
		return nil, err
	}
	return &pb.SetReorderPointResponse{
		Success: true,
	}, nil
}

func (s *InventoryServiceServer) ManageWarehouses(ctx context.Context, req *pb.ManageWarehousesRequest) (*pb.ManageWarehousesResponse, error) {
	var warehouses []models.Warehouse
	for _, warehouse := range req.Warehouses {
		warehouses = append(warehouses, models.Warehouse{
			WarehouseID: warehouse.WarehouseId,
			Name:        warehouse.WarehouseName,
			Location:    warehouse.Location,
		})
	}
	err := s.service.ManageWarehouses(warehouses)
	if err != nil {
		return nil, err
	}

	return &pb.ManageWarehousesResponse{
		Success: true,
	}, nil
}

func (s *InventoryServiceServer) AddOrUpdateInventoryItem(ctx context.Context, req *pb.AddOrUpdateInventoryItemRequest) (*pb.AddOrUpdateInventoryItemResponse, error) {
	// Create InventoryItem struct from the request data
	item := &models.InventoryItem{
		ProductID:          req.Item.ProductId,
		ProductName:        req.Item.ProductName,
		ProductDescription: req.Item.ProductDescription,
		SKU:                req.Item.Sku,
		SupplierID:         req.Item.SupplierId,
		Category:           req.Item.Category,
		Price:              float64(req.Item.Price),
		AvailableQuantity:  req.Item.AvailableQuantity,
		ReorderPoint:       req.Item.ReorderPoint,
	}

	// Call the service to add or update the inventory item
	_, err := s.service.AddOrUpdateInventoryItem(item)
	if err != nil {
		return nil, err
	}

	// Loop through warehouse stocks and add or update them
	for _, warehouseStock := range req.Item.WarehouseStocks {
		// Assuming the WarehouseStock model does not have InventoryItemId, just WarehouseID and StockLevel
		err := s.service.AddOrUpdateWarehouseStock(&models.WarehouseStock{
			InventoryItemID: uint(item.ID), // Reference the item ID after saving/updating
			WarehouseID: func() uint {
				id, _ := strconv.ParseUint(warehouseStock.WarehouseId, 10, 32)
				return uint(id)
			}(),
			StockLevel: warehouseStock.StockLevel,
		})
		if err != nil {
			return nil, err
		}
	}

	return &pb.AddOrUpdateInventoryItemResponse{
		Success: true,
	}, nil
}

// Order Fulfillment

func (s *InventoryServiceServer) ProcessOrder(ctx context.Context, req *pb.ProcessOrderRequest) (*pb.ProcessOrderResponse, error) {
	var orderItems []models.OrderItem
	for _, item := range req.Items {
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductId,
			Quantity:  item.Quantity,
		})
	}

	err := s.service.ProcessOrder(orderItems)
	if err != nil {
		return nil, err
	}

	return &pb.ProcessOrderResponse{
		Success: true,
	}, nil
}

func (s *InventoryServiceServer) GeneratePickingList(ctx context.Context, req *pb.GeneratePickingListRequest) (*pb.GeneratePickingListResponse, error) {
	pickingList, err := s.service.GeneratePickingList(req.OrderId)
	if err != nil {
		return nil, err
	}

	var pbPickingItems []*pb.PickingItem
	for _, item := range pickingList {
		pbPickingItems = append(pbPickingItems, &pb.PickingItem{
			ProductId:   item.ProductID,
			Quantity:    item.Quantity,
			WarehouseId: strconv.Itoa(int(item.WarehouseID)),
		})
	}

	return &pb.GeneratePickingListResponse{
		PickingList: pbPickingItems,
	}, nil
}

func (s *InventoryServiceServer) UpdateInventory(ctx context.Context, req *pb.UpdateInventoryRequest) (*pb.UpdateInventoryResponse, error) {
	warehouseID, err := strconv.ParseUint(req.WarehouseId, 10, 32)
	if err != nil {
		return nil, err
	}
	err = s.service.UpdateInventoryStock(req.ProductId, req.Quantity, uint(warehouseID))
	if err != nil {
		return nil, err
	}

	return &pb.UpdateInventoryResponse{
		Success: true,
	}, nil
}

// Vendor Management Integration

func (s *InventoryServiceServer) PlaceVendorOrder(ctx context.Context, req *pb.PlaceVendorOrderRequest) (*pb.PlaceVendorOrderResponse, error) {
	var orderItems []models.OrderItem
	for _, item := range req.Items {
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductId,
			Quantity:  item.Quantity,
		})
	}

	err := s.service.PlaceVendorOrder(req.VendorId, orderItems)
	if err != nil {
		return nil, err
	}

	return &pb.PlaceVendorOrderResponse{
		Success: true,
	}, nil
}

// Finance Service Integration

func (s *InventoryServiceServer) NotifyFinanceForOrder(ctx context.Context, req *pb.NotifyFinanceRequest) (*pb.NotifyFinanceResponse, error) {
	err := s.service.NotifyFinanceForOrder(req.OrderId, float64(req.TotalAmount))
	if err != nil {
		return nil, err
	}

	return &pb.NotifyFinanceResponse{
		Success: true,
	}, nil
}
