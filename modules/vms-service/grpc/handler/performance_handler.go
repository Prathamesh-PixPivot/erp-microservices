package handler

import (
	"context"
	"log"
	"vms-service/grpc/vms_pb"
	"vms-service/internal/models"
	"vms-service/internal/services"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PerformanceHandler struct {
	vms_pb.UnimplementedPerformanceServiceServer
	service *services.PerformanceService
}

func NewPerformanceHandler(service *services.PerformanceService) *PerformanceHandler {
	return &PerformanceHandler{service: service}
}

func (h *PerformanceHandler) RecordPerformance(ctx context.Context, req *vms_pb.RecordPerformanceRequest) (*vms_pb.PerformanceResponse, error) {
	performance := &models.VendorPerformance{
		ID:          uuid.New(),
		VendorID:    uuid.MustParse(req.VendorId),
		Score:       float64(req.Score),
		RiskLevel:   req.RiskLevel,
		EvaluatedAt: req.EvaluatedAt.AsTime(),
	}

	if err := h.service.RecordPerformance(performance); err != nil {
		log.Printf("Failed to record performance: %v", err)
		return nil, err
	}

	return &vms_pb.PerformanceResponse{
		VendorId:    performance.VendorID.String(),
		Score:       float32(performance.Score),
		RiskLevel:   performance.RiskLevel,
		EvaluatedAt: timestamppb.New(performance.EvaluatedAt).String(),
	}, nil
}

func (h *PerformanceHandler) GetPerformanceByID(ctx context.Context, req *vms_pb.GetPerformanceByIDRequest) (*vms_pb.PerformanceResponse, error) {
	performance, err := h.service.GetPerformanceByID(uuid.MustParse(req.Id))
	if err != nil {
		log.Printf("Failed to get performance: %v", err)
		return nil, err
	}

	return &vms_pb.PerformanceResponse{
		VendorId:    performance.VendorID.String(),
		Score:       float32(performance.Score),
		RiskLevel:   performance.RiskLevel,
		EvaluatedAt: timestamppb.New(performance.EvaluatedAt).String(),
	}, nil
}

func (h *PerformanceHandler) UpdatePerformance(ctx context.Context, req *vms_pb.UpdatePerformanceRequest) (*vms_pb.PerformanceResponse, error) {
	performance := &models.VendorPerformance{
		ID:          uuid.MustParse(req.Performance.Id),
		VendorID:    uuid.MustParse(req.Performance.VendorId),
		Score:       float64(req.Performance.Score),
		RiskLevel:   req.Performance.RiskLevel,
		EvaluatedAt: req.Performance.EvaluatedAt.AsTime(),
	}

	if err := h.service.UpdatePerformance(performance); err != nil {
		log.Printf("Failed to update performance: %v", err)
		return nil, err
	}

	return &vms_pb.PerformanceResponse{
		VendorId:    performance.VendorID.String(),
		Score:       float32(performance.Score),
		RiskLevel:   performance.RiskLevel,
		EvaluatedAt: timestamppb.New(performance.EvaluatedAt).String(),
	}, nil
}

func (h *PerformanceHandler) DeletePerformance(ctx context.Context, req *vms_pb.DeletePerformanceRequest) (*vms_pb.DeletePerformanceResponse, error) {
	if err := h.service.DeletePerformance(uuid.MustParse(req.Id)); err != nil {
		log.Printf("Failed to delete performance: %v", err)
		return nil, err
	}

	return &vms_pb.DeletePerformanceResponse{Message: "Performance record deleted successfully"}, nil
}
