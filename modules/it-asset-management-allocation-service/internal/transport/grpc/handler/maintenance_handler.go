package handler

import (
	"context"

	"amaa/internal/models"
	"amaa/internal/services"
	pbmaintenance "amaa/internal/transport/grpc/proto"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MaintenanceHandler struct {
	pbmaintenance.UnimplementedMaintenanceServiceServer
	maintenanceService services.MaintenanceService
	logger             *zap.Logger
}

func NewMaintenanceHandler(service services.MaintenanceService, logger *zap.Logger) *MaintenanceHandler {
	return &MaintenanceHandler{
		maintenanceService: service,
		logger:             logger,
	}
}

func (h *MaintenanceHandler) ScheduleMaintenance(ctx context.Context, req *pbmaintenance.ScheduleMaintenanceRequest) (*pbmaintenance.MaintenanceResponse, error) {
	h.logger.Info("ScheduleMaintenance request", zap.String("asset_id", req.AssetId))
	record := &models.Maintenance{
		AssetID:     req.AssetId,
		Description: req.Description,
		Cost:        req.Cost,
	}
	created, err := h.maintenanceService.ScheduleMaintenance(record)
	if err != nil {
		h.logger.Error("ScheduleMaintenance error", zap.Error(err))
		return nil, err
	}
	return &pbmaintenance.MaintenanceResponse{
		Id:              created.ID,
		AssetId:         created.AssetID,
		MaintenanceDate: timestamppb.New(created.MaintenanceDate),
		Description:     created.Description,
		Cost:            created.Cost,
	}, nil
}

func (h *MaintenanceHandler) GetMaintenanceRecords(ctx context.Context, req *pbmaintenance.GetMaintenanceRecordsRequest) (*pbmaintenance.GetMaintenanceRecordsResponse, error) {
	h.logger.Info("GetMaintenanceRecords request", zap.String("asset_id", req.AssetId))
	records, err := h.maintenanceService.GetMaintenanceRecords(req.AssetId)
	if err != nil {
		h.logger.Error("GetMaintenanceRecords error", zap.Error(err))
		return nil, err
	}
	var responses []*pbmaintenance.MaintenanceResponse
	for _, r := range records {
		responses = append(responses, &pbmaintenance.MaintenanceResponse{
			Id:              r.ID,
			AssetId:         r.AssetID,
			MaintenanceDate: timestamppb.New(r.MaintenanceDate),
			Description:     r.Description,
			Cost:            r.Cost,
		})
	}
	return &pbmaintenance.GetMaintenanceRecordsResponse{
		Records: responses,
	}, nil
}
