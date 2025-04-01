package handler

import (
	"context"

	"amaa/internal/models"
	"amaa/internal/services"
	pbdisposal "amaa/internal/transport/grpc/proto"

	"go.uber.org/zap"
)

type DisposalHandler struct {
	pbdisposal.UnimplementedDisposalServiceServer
	disposalService services.DisposalService
	logger          *zap.Logger
}

func NewDisposalHandler(service services.DisposalService, logger *zap.Logger) *DisposalHandler {
	return &DisposalHandler{
		disposalService: service,
		logger:          logger,
	}
}

func (h *DisposalHandler) DecommissionAsset(ctx context.Context, req *pbdisposal.DecommissionAssetRequest) (*pbdisposal.DecommissionResponse, error) {
	h.logger.Info("DecommissionAsset request", zap.String("asset_id", req.AssetId))
	disposal := &models.Disposal{
		AssetID: req.AssetId,
		Reason:  req.Reason,
	}
	result, err := h.disposalService.DecommissionAsset(disposal)
	if err != nil {
		h.logger.Error("DecommissionAsset error", zap.Error(err))
		return nil, err
	}
	return &pbdisposal.DecommissionResponse{
		Message: "Asset " + result.AssetID + " decommissioned. Reason: " + result.Reason,
	}, nil
}
