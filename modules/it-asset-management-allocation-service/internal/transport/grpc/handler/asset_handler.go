package handler

import (
	"context"

	"amaa/internal/models"
	"amaa/internal/services"
	pbasset "amaa/internal/transport/grpc/proto" // update the path as needed

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AssetHandler struct {
	pbasset.UnimplementedAssetServiceServer
	assetService services.AssetService
	logger       *zap.Logger
}

func NewAssetHandler(service services.AssetService, logger *zap.Logger) *AssetHandler {
	return &AssetHandler{
		assetService: service,
		logger:       logger,
	}
}

func (h *AssetHandler) CreateAsset(ctx context.Context, req *pbasset.CreateAssetRequest) (*pbasset.AssetResponse, error) {
	h.logger.Info("CreateAsset request", zap.String("name", req.Name))
	asset := &models.Asset{
		Name:             req.Name,
		Description:      req.Description,
		Category:         req.Category,
		PurchaseDate:     req.PurchaseDate.AsTime(),
		PurchasePrice:    req.PurchasePrice,
		CurrentValue:     req.PurchasePrice, // Initial current value equals purchase price
		Location:         req.Location,
		Status:           "active",
		DepreciationRate: req.DepreciationRate,
		Guidelines:       req.Guidelines,
	}
	created, err := h.assetService.CreateAsset(asset)
	if err != nil {
		h.logger.Error("CreateAsset error", zap.Error(err))
		return nil, err
	}
	return &pbasset.AssetResponse{
		Id:               created.ID,
		Name:             created.Name,
		Description:      created.Description,
		Category:         created.Category,
		PurchaseDate:     timestamppb.New(created.PurchaseDate),
		PurchasePrice:    created.PurchasePrice,
		CurrentValue:     created.CurrentValue,
		Location:         created.Location,
		Status:           created.Status,
		DepreciationRate: created.DepreciationRate,
		Guidelines:       created.Guidelines,
	}, nil
}

func (h *AssetHandler) GetAsset(ctx context.Context, req *pbasset.GetAssetRequest) (*pbasset.AssetResponse, error) {
	h.logger.Info("GetAsset request", zap.String("id", req.Id))
	asset, err := h.assetService.GetAsset(req.Id)
	if err != nil {
		h.logger.Error("GetAsset error", zap.Error(err))
		return nil, err
	}
	return &pbasset.AssetResponse{
		Id:               asset.ID,
		Name:             asset.Name,
		Description:      asset.Description,
		Category:         asset.Category,
		PurchaseDate:     timestamppb.New(asset.PurchaseDate),
		PurchasePrice:    asset.PurchasePrice,
		CurrentValue:     asset.CurrentValue,
		Location:         asset.Location,
		Status:           asset.Status,
		DepreciationRate: asset.DepreciationRate,
		Guidelines:       asset.Guidelines,
	}, nil
}

func (h *AssetHandler) UpdateAsset(ctx context.Context, req *pbasset.UpdateAssetRequest) (*pbasset.AssetResponse, error) {
	h.logger.Info("UpdateAsset request", zap.String("id", req.Id))
	asset := &models.Asset{
		ID:               req.Id,
		Name:             req.Name,
		Description:      req.Description,
		Category:         req.Category,
		PurchaseDate:     req.PurchaseDate.AsTime(),
		PurchasePrice:    req.PurchasePrice,
		CurrentValue:     req.CurrentValue,
		Location:         req.Location,
		Status:           req.Status,
		DepreciationRate: req.DepreciationRate,
		Guidelines:       req.Guidelines,
	}
	if err := h.assetService.UpdateAsset(asset); err != nil {
		h.logger.Error("UpdateAsset error", zap.Error(err))
		return nil, err
	}
	return &pbasset.AssetResponse{
		Id:               asset.ID,
		Name:             asset.Name,
		Description:      asset.Description,
		Category:         asset.Category,
		PurchaseDate:     timestamppb.New(asset.PurchaseDate),
		PurchasePrice:    asset.PurchasePrice,
		CurrentValue:     asset.CurrentValue,
		Location:         asset.Location,
		Status:           asset.Status,
		DepreciationRate: asset.DepreciationRate,
		Guidelines:       asset.Guidelines,
	}, nil
}

func (h *AssetHandler) DeleteAsset(ctx context.Context, req *pbasset.DeleteAssetRequest) (*pbasset.DeleteAssetResponse, error) {
	h.logger.Info("DeleteAsset request", zap.String("id", req.Id))
	if err := h.assetService.DeleteAsset(req.Id); err != nil {
		h.logger.Error("DeleteAsset error", zap.Error(err))
		return nil, err
	}
	return &pbasset.DeleteAssetResponse{
		Message: "Asset deleted successfully",
	}, nil
}
