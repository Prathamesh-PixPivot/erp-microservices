package grpc

import (
	"context"
	"pantry-service/internal/models"
	service "pantry-service/internal/services"

	pb "pantry-service/grpc/pantry_pb" // Import the generated gRPC code from .proto
)

// PantryServiceServer is the gRPC server struct
type PantryServiceServer struct {
	service service.PantryService
	pb.UnimplementedPantryServiceServer
}

// NewPantryServiceServer creates a new gRPC server instance
func NewPantryServiceServer(svc service.PantryService) *PantryServiceServer {
	return &PantryServiceServer{
		service: svc,
	}
}

// CRUD Functions for pantry items

func (s *PantryServiceServer) CreatePantryItem(ctx context.Context, req *pb.CreatePantryItemRequest) (*pb.CreatePantryItemResponse, error) {

	item := &models.PantryItem{
		ProductName: req.Item.ProductName,
		Price:       float64(req.Item.Price),
		Category:    req.Item.Category,
	}

	createdItem, err := s.service.CreatePantryItem(item)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePantryItemResponse{
		Item: &pb.PantryItem{
			ProductName: createdItem.ProductName,
			Price:       float32(createdItem.Price),
			Category:    createdItem.Category,
			CreatedAt:   createdItem.CreatedAt.String(),
			UpdatedAt:   createdItem.UpdatedAt.String(),
		},
		Success: true,
	}, nil

}

func (s *PantryServiceServer) GetPantryItem(ctx context.Context, req *pb.GetPantryItemRequest) (*pb.GetPantryItemResponse, error) {
	item, err := s.service.GetPantryItem(req.ProductId)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, nil
	}

	return &pb.GetPantryItemResponse{
		Item: &pb.PantryItem{
			ProductName: item.ProductName,
			Price:       float32(item.Price),
			Category:    item.Category,
		},
		Success: true,
	}, nil
}

func (s *PantryServiceServer) UpdatePantryItem(ctx context.Context, req *pb.UpdatePantryItemRequest) (*pb.UpdatePantryItemResponse, error) {
	item := &models.PantryItem{
		ProductName: req.Item.ProductName,
		Price:       float64(req.Item.Price),
		Category:    req.Item.Category,
	}

	updatedItem, err := s.service.UpdatePantryItem(item)
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePantryItemResponse{
		Item: &pb.PantryItem{
			ProductName: updatedItem.ProductName,
			Price:       float32(updatedItem.Price),
			Category:    updatedItem.Category,
		},
		Success: true,
	}, nil
}

func (s *PantryServiceServer) DeletePantryItem(ctx context.Context, req *pb.DeletePantryItemRequest) (*pb.DeletePantryItemResponse, error) {
	err := s.service.DeletePantryItem(req.ProductId)
	if err != nil {
		return nil, err
	}
	return &pb.DeletePantryItemResponse{
		Success: true,
	}, nil
}

func (s *PantryServiceServer) ListPantryItems(ctx context.Context, req *pb.ListPantryItemsRequest) (*pb.ListPantryItemsResponse, error) {
	items, err := s.service.ListPantryItems(int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	var pbItems []*pb.PantryItem
	for _, item := range items {
		pbItems = append(pbItems, &pb.PantryItem{
			ProductName: item.ProductName,
			Price:       float32(item.Price),
			Category:    item.Category,
		})
	}

	return &pb.ListPantryItemsResponse{
		Items:   pbItems,
		Success: true,
	}, nil
}

// CRUD Functions for pantry bucket items

func (s *PantryServiceServer) CreateBucketItem(ctx context.Context, req *pb.CreateBucketItemRequest) (*pb.CreateBucketItemResponse, error) {
	item := &models.BucketItem{
		ProductId: req.Item.ProductId,
		Qty:       float64(req.Item.Qty),
		Price:     float64(req.Item.Price),
		Total:     float64(req.Item.Total),
		PaidBy:    req.Item.PaidBy,
	}

	createdItem, err := s.service.CreateBucketItem(item)
	if err != nil {
		return nil, err
	}

	return &pb.CreateBucketItemResponse{
		Item: &pb.BucketItem{
			ProductId: createdItem.ProductId,
			Qty:       float32(createdItem.Qty),
			Price:     float32(createdItem.Price),
			Total:     float32(createdItem.Total),
			PaidBy:    createdItem.PaidBy,
		},
		Success: true,
	}, nil
}

func (s *PantryServiceServer) GetBucketItem(ctx context.Context, req *pb.GetBucketItemRequest) (*pb.GetBucketItemResponse, error) {
	item, err := s.service.GetBucketItem(req.ProductId)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, nil
	}

	return &pb.GetBucketItemResponse{
		Item: &pb.BucketItem{
			ProductId: item.ProductId,
			Qty:       float32(item.Qty),
			Price:     float32(item.Price),
			Total:     float32(item.Total),
			PaidBy:    item.PaidBy,
		},
		Success: true,
	}, nil
}

func (s *PantryServiceServer) UpdateBucketItem(ctx context.Context, req *pb.UpdateBucketItemRequest) (*pb.UpdateBucketItemResponse, error) {
	item := &models.BucketItem{
		ProductId: req.Item.ProductId,
		Qty:       float64(req.Item.Qty),
		Price:     float64(req.Item.Price),
		Total:     float64(req.Item.Total),
		PaidBy:    req.Item.PaidBy,
	}

	updatedItem, err := s.service.UpdateBucketItem(item)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateBucketItemResponse{
		Item: &pb.BucketItem{
			ProductId: updatedItem.ProductId,
			Qty:       float32(updatedItem.Qty),
			Price:     float32(updatedItem.Price),
			Total:     float32(updatedItem.Total),
			PaidBy:    updatedItem.PaidBy,
		},
		Success: true,
	}, nil
}

func (s *PantryServiceServer) DeleteBucketItem(ctx context.Context, req *pb.DeleteBucketItemRequest) (*pb.DeleteBucketItemResponse, error) {
	err := s.service.DeleteBucketItem(req.ProductId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBucketItemResponse{
		Success: true,
	}, nil
}

func (s *PantryServiceServer) ListBucketItems(ctx context.Context, req *pb.ListBucketItemsRequest) (*pb.ListBucketItemsResponse, error) {
	items, err := s.service.ListBucketItems(int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	var pbItems []*pb.BucketItem
	for _, item := range items {
		pbItems = append(pbItems, &pb.BucketItem{
			ProductId: item.ProductId,
			Qty:       float32(item.Qty),
			Price:     float32(item.Price),
			Total:     float32(item.Total),
			PaidBy:    item.PaidBy,
		})
	}

	return &pb.ListBucketItemsResponse{
		Items:   pbItems,
		Success: true,
	}, nil
}

// CRUD Functions for pantry expense logs

func (s *PantryServiceServer) CreateExpenseLog(ctx context.Context, req *pb.CreateExpenseLogRequest) (*pb.CreateExpenseLogResponse, error) {
	item := &models.ExpenseLog{
		AmtReceived: float64(req.Log.AmtReceived),
		AmtSpend:    float64(req.Log.AmtSpend),
	}

	createdLog, err := s.service.CreateExpenseLog(item)
	if err != nil {
		return nil, err
	}

	return &pb.CreateExpenseLogResponse{
		Log: &pb.ExpenseLog{
			AmtReceived: float32(createdLog.AmtReceived),
			AmtSpend:    float32(createdLog.AmtSpend),
		},
		Success: true,
	}, nil
}

func (s *PantryServiceServer) GetExpenseLog(ctx context.Context, req *pb.GetExpenseLogRequest) (*pb.GetExpenseLogResponse, error) {
	log, err := s.service.GetExpenseLog(req.LogID)
	if err != nil {
		return nil, err
	}
	if log == nil {
		return nil, nil
	}

	return &pb.GetExpenseLogResponse{
		Log: &pb.ExpenseLog{
			AmtReceived: float32(log.AmtReceived),
			AmtSpend:    float32(log.AmtSpend),
		},
		Success: true,
	}, nil
}

func (s *PantryServiceServer) UpdateExpenseLog(ctx context.Context, req *pb.UpdateExpenseLogRequest) (*pb.UpdateExpenseLogResponse, error) {
	item := &models.ExpenseLog{
		AmtReceived: float64(req.Log.AmtReceived),
		AmtSpend:    float64(req.Log.AmtSpend),
	}

	updatedLog, err := s.service.UpdateExpenseLog(item)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateExpenseLogResponse{
		Log: &pb.ExpenseLog{
			AmtReceived: float32(updatedLog.AmtReceived),
			AmtSpend:    float32(updatedLog.AmtSpend),
		},
		Success: true,
	}, nil
}

func (s *PantryServiceServer) DeleteExpenseLog(ctx context.Context, req *pb.DeleteExpenseLogRequest) (*pb.DeleteExpenseLogResponse, error) {
	err := s.service.DeleteExpenseLog(req.LogID)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteExpenseLogResponse{
		Success: true,
	}, nil
}

func (s *PantryServiceServer) ListExpenseLogs(ctx context.Context, req *pb.ListExpenseLogsRequest) (*pb.ListExpenseLogsResponse, error) {
	logs, err := s.service.ListExpenseLogs(int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	var pbLogs []*pb.ExpenseLog
	for _, log := range logs {
		pbLogs = append(pbLogs, &pb.ExpenseLog{
			AmtReceived: float32(log.AmtReceived),
			AmtSpend:    float32(log.AmtSpend),
		})
	}

	return &pb.ListExpenseLogsResponse{
		Logs:    pbLogs,
		Success: true,
	}, nil
}
