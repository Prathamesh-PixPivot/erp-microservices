package handler

import (
	"context"
	"fmt"
	pb "gst-service/internal/api/protobufs/gst-service"
	"gst-service/internal/domain"
	"gst-service/internal/service"
)

// GSTR9CHandler is the struct for the GSTR9C gRPC handler
type GSTR9CHandler struct {
	pb.UnimplementedGSTR9CServiceServer
	gstr9CService *service.GSTR9CService
}

// NewGSTR9CHandler creates a new GSTR9C handler
func NewGSTR9CHandler(gstr9CService *service.GSTR9CService) *GSTR9CHandler {
	return &GSTR9CHandler{
		gstr9CService: gstr9CService,
	}
}

// SaveGSTR9C handles the SaveGSTR9C gRPC method
func (h *GSTR9CHandler) SaveGSTR9C(ctx context.Context, req *pb.GSTR9CRequest) (*pb.GSTR9CResponse, error) {
	domainReq := &domain.GSTR9CRequest{
		GSTIN:              req.GetGstin(),
		ReturnPeriod:       req.GetReturnPeriod(),
		AuditDetails:       req.GetAuditDetails(),
		ReconciliationStmt: req.GetReconciliationStatement(),
	}

	refID, err := h.gstr9CService.SaveGSTR9CData(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to save GSTR9C data: %w", err)
	}

	return &pb.GSTR9CResponse{
		RefId:  refID,
		Status: "Success",
	}, nil
}

// GetGSTR9CStatus handles the GetGSTR9CStatus gRPC method
func (h *GSTR9CHandler) GetGSTR9CStatus(ctx context.Context, req *pb.GSTR9CStatusRequest) (*pb.GSTR9CStatusResponse, error) {
	domainReq := &domain.GSTR9CStatusRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		ARN:          req.GetArn(),
	}

	status, message, err := h.gstr9CService.GetGSTR9CStatus(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get GSTR9C status: %w", err)
	}

	return &pb.GSTR9CStatusResponse{
		Status:  status,
		Message: message,
	}, nil
}
