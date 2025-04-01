package handler

import (
	"context"

	"amaa/internal/models"
	"amaa/internal/services"
	pballocation "amaa/internal/transport/grpc/proto"
	common "amaa/internal/transport/grpc/proto/common"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AllocationHandler struct {
	pballocation.UnimplementedAllocationServiceServer
	allocationService services.AllocationService
	logger            *zap.Logger
}

func NewAllocationHandler(service services.AllocationService, logger *zap.Logger) *AllocationHandler {
	return &AllocationHandler{
		allocationService: service,
		logger:            logger,
	}
}

func (h *AllocationHandler) AllocateAsset(ctx context.Context, req *pballocation.AllocateAssetRequest) (*pballocation.AllocationResponse, error) {
	h.logger.Info("AllocateAsset request", zap.String("asset_id", req.AssetId))
	allocation := &models.Allocation{
		AssetID:    req.AssetId,
		AssignedTo: req.AssignedTo,
	}
	created, err := h.allocationService.AllocateAsset(allocation)
	if err != nil {
		h.logger.Error("AllocateAsset error", zap.Error(err))
		return nil, err
	}
	return &pballocation.AllocationResponse{
		Id:             created.ID,
		AssetId:        created.AssetID,
		AssignedTo:     created.AssignedTo,
		AssignmentDate: timestamppb.New(created.AssignmentDate),
	}, nil
}

func (h *AllocationHandler) ReallocateAsset(ctx context.Context, req *pballocation.ReallocateAssetRequest) (*common.GenericResponse, error) {
	h.logger.Info("ReallocateAsset request", zap.String("asset_id", req.AssetId))
	err := h.allocationService.ReallocateAsset(req.AssetId, req.NewAssignedTo)
	if err != nil {
		h.logger.Error("ReallocateAsset error", zap.Error(err))
		return nil, err
	}
	return &common.GenericResponse{Message: "Asset reallocated successfully"}, nil
}

func (h *AllocationHandler) DeallocateAsset(ctx context.Context, req *pballocation.DeallocateAssetRequest) (*common.GenericResponse, error) {
	h.logger.Info("DeallocateAsset request", zap.String("allocation_id", req.AllocationId))
	err := h.allocationService.DeallocateAsset(req.AllocationId)
	if err != nil {
		h.logger.Error("DeallocateAsset error", zap.Error(err))
		return nil, err
	}
	return &common.GenericResponse{Message: "Asset deallocated successfully"}, nil
}
