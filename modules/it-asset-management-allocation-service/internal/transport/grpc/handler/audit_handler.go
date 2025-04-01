package handler

import (
	"context"

	"amaa/internal/models"
	"amaa/internal/services"
	pbaudit "amaa/internal/transport/grpc/proto"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuditHandler struct {
	pbaudit.UnimplementedAuditServiceServer
	auditService services.AuditService
	logger       *zap.Logger
}

func NewAuditHandler(service services.AuditService, logger *zap.Logger) *AuditHandler {
	return &AuditHandler{
		auditService: service,
		logger:       logger,
	}
}

func (h *AuditHandler) AuditAsset(ctx context.Context, req *pbaudit.AuditAssetRequest) (*pbaudit.AuditResponse, error) {
	h.logger.Info("AuditAsset request", zap.String("asset_id", req.AssetId))
	audit := &models.Audit{
		AssetID:   req.AssetId,
		AuditedBy: req.AuditedBy,
		Condition: req.Condition,
		Remarks:   req.Remarks,
	}
	created, err := h.auditService.AuditAsset(audit)
	if err != nil {
		h.logger.Error("AuditAsset error", zap.Error(err))
		return nil, err
	}
	return &pbaudit.AuditResponse{
		Message:   "Audit recorded: " + created.Condition,
		AuditDate: timestamppb.New(created.AuditDate),
	}, nil
}

func (h *AuditHandler) GetAuditHistory(ctx context.Context, req *pbaudit.GetAuditHistoryRequest) (*pbaudit.GetAuditHistoryResponse, error) {
	h.logger.Info("GetAuditHistory request", zap.String("asset_id", req.AssetId))
	history, err := h.auditService.GetAuditHistory(req.AssetId)
	if err != nil {
		h.logger.Error("GetAuditHistory error", zap.Error(err))
		return nil, err
	}
	var responses []*pbaudit.AuditResponse
	for _, a := range history {
		responses = append(responses, &pbaudit.AuditResponse{
			Message:   "Audit recorded: " + a.Condition,
			AuditDate: timestamppb.New(a.AuditDate),
		})
	}
	return &pbaudit.GetAuditHistoryResponse{Audits: responses}, nil
}
