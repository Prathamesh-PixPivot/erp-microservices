package handler

import (
	"context"
	"fmt"
	pb "gst-service/internal/api/protobufs/gst-service" // Import generated protobuf code
	"gst-service/internal/domain"                       // Your domain models
	"gst-service/internal/service"                      // Adjust the import path according to your project
)

// GSTR3BHandler is the struct for the GSTR3B gRPC handler
type GSTR3BHandler struct {
	pb.UnimplementedGSTR3BServiceServer
	gstr3BService *service.GSTR3BService
}

// NewGSTR3BHandler creates a new GSTR3B handler
func NewGSTR3BHandler(gstr3BService *service.GSTR3BService) *GSTR3BHandler {
	return &GSTR3BHandler{
		gstr3BService: gstr3BService,
	}
}

// ConvertProtoToDomainGSTR3B converts a protobuf GSTR3BRequest to a domain GSTR3BRequest
func ConvertProtoToDomainGSTR3B(req *pb.GSTR3BRequest) *domain.GSTR3BRequest {
	return &domain.GSTR3BRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		TaxableValue: req.GetTaxableValue(),
		TaxLiability: req.GetTaxLiability(),
		ITCClaimed:   req.GetItcClaimed(),
	}
}

// SaveGSTR3B handles the SaveGSTR3B gRPC method
func (h *GSTR3BHandler) SaveGSTR3B(ctx context.Context, req *pb.GSTR3BRequest) (*pb.GSTR3BResponse, error) {
	// Convert the protobuf request to the domain request
	domainReq := ConvertProtoToDomainGSTR3B(req)

	// Call the service layer to save the GSTR3B data
	refID, err := h.gstr3BService.SaveGSTR3BData(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to save GSTR3B data: %w", err)
	}

	return &pb.GSTR3BResponse{
		RefId:  refID,
		Status: "Success",
	}, nil
}

// SubmitGSTR3B handles the SubmitGSTR3B gRPC method
func (h *GSTR3BHandler) SubmitGSTR3B(ctx context.Context, req *pb.GSTR3BSubmitRequest) (*pb.GSTR3BSubmitResponse, error) {
	domainReq := &domain.GSTR3BSubmitRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
	}

	arn, err := h.gstr3BService.SubmitGSTR3BData(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to submit GSTR3B data: %w", err)
	}

	return &pb.GSTR3BSubmitResponse{
		Arn:    arn,
		Status: "Success",
	}, nil
}

// FileGSTR3B handles the FileGSTR3B gRPC method
func (h *GSTR3BHandler) FileGSTR3B(ctx context.Context, req *pb.GSTR3BFileRequest) (*pb.GSTR3BFileResponse, error) {
	domainReq := &domain.GSTR3BFileRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		ARN:          req.GetArn(),
	}

	filingStatus, message, err := h.gstr3BService.FileGSTR3BData(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to file GSTR3B: %w", err)
	}

	return &pb.GSTR3BFileResponse{
		FilingStatus: filingStatus,
		Message:      message,
	}, nil
}

// GetGSTR3BStatus handles the GetGSTR3BStatus gRPC method
func (h *GSTR3BHandler) GetGSTR3BStatus(ctx context.Context, req *pb.GSTR3BStatusRequest) (*pb.GSTR3BStatusResponse, error) {
	domainReq := &domain.GSTR3BStatusRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		ARN:          req.GetArn(),
	}

	status, message, err := h.gstr3BService.GetGSTR3BStatus(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get GSTR3B status: %w", err)
	}

	return &pb.GSTR3BStatusResponse{
		Status:  status,
		Message: message,
	}, nil
}
