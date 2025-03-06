package grpc

import (
	"context"
	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto" 
	"hrms/internal/usecase"

	"google.golang.org/protobuf/types/known/emptypb"
)

// PerformanceKPIHandler handles gRPC requests for Performance KPI.
type PerformanceKPIHandler struct {
	proto.UnimplementedPerformanceKPIServiceServer
	Usecase *usecase.HrmsUsecase
}

// CreatePerformanceKPI handles the gRPC request to create a new KPI.
func (h *PerformanceKPIHandler) CreatePerformanceKPI(ctx context.Context, req *proto.CreatePerformanceKPIRequest) (*proto.PerformanceKPIResponse, error) {
	kpiDTO := dto.PerformanceKPIDTO{
		ReviewID: uint(req.ReviewId),
		KPIName:  req.KpiName,
		Score:    req.Score,
		Comments: req.Comments,
	}

	createdKPI, err := h.Usecase.CreatePerformanceKPI(ctx, kpiDTO)
	if err != nil {
		return nil, err
	}

	return &proto.PerformanceKPIResponse{
		Kpi: h.convertToProtoKPI(createdKPI),
	}, nil
}

// GetPerformanceKPI retrieves a KPI by ID.
func (h *PerformanceKPIHandler) GetPerformanceKPI(ctx context.Context, req *proto.GetPerformanceKPIRequest) (*proto.PerformanceKPIResponse, error) {
	kpi, err := h.Usecase.GetPerformanceKPIByID(ctx, uint(req.KpiId))
	if err != nil {
		return nil, err
	}

	return &proto.PerformanceKPIResponse{
		Kpi: h.convertToProtoKPI(kpi),
	}, nil
}

// UpdatePerformanceKPI updates a KPI.
func (h *PerformanceKPIHandler) UpdatePerformanceKPI(ctx context.Context, req *proto.UpdatePerformanceKPIRequest) (*emptypb.Empty, error) {
	updates := make(map[string]interface{})

	if req.KpiName != nil {
		updates["kpi_name"] = *req.KpiName
	}
	if req.Score != nil {
		updates["score"] = *req.Score
	}
	if req.Comments != nil {
		updates["comments"] = *req.Comments
	}

	err := h.Usecase.UpdatePerformanceKPI(ctx, uint(req.KpiId), updates)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeletePerformanceKPI deletes a KPI.
func (h *PerformanceKPIHandler) DeletePerformanceKPI(ctx context.Context, req *proto.DeletePerformanceKPIRequest) (*emptypb.Empty, error) {
	err := h.Usecase.DeletePerformanceKPI(ctx, uint(req.KpiId))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// ListPerformanceKPIs retrieves all KPIs for a review.
func (h *PerformanceKPIHandler) ListPerformanceKPIs(ctx context.Context, req *proto.ListPerformanceKPIsRequest) (*proto.ListPerformanceKPIsResponse, error) {
	kpis, total, err := h.Usecase.ListPerformanceKPIs(ctx, uint(req.ReviewId), int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	var kpiResponses []*proto.PerformanceKPI
	for _, kpi := range kpis {
		kpiResponses = append(kpiResponses, h.convertToProtoKPI(&kpi))
	}

	return &proto.ListPerformanceKPIsResponse{
		Total:  int32(total),
		Limit:  req.Limit,
		Offset: req.Offset,
		Kpis:   kpiResponses,
	}, nil
}

// convertToProtoKPI converts a PerformanceKPIDTO to its proto representation.
func (h *PerformanceKPIHandler) convertToProtoKPI(kpi *dto.PerformanceKPIDTO) *proto.PerformanceKPI {
	if kpi == nil {
		return nil
	}
	return &proto.PerformanceKPI{
		Id:       uint64(kpi.ID),
		ReviewId: uint64(kpi.ReviewID),
		KpiName:  kpi.KPIName,
		Score:    kpi.Score,
		Comments: kpi.Comments,
	}
}
