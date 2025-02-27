package handler

import (
	"context"
	"fmt"
	pb "gst-service/internal/api/protobufs/gst-service" // Import generated protobuf code
	"gst-service/internal/domain"                       // Your domain models
	"gst-service/internal/service"                      // Adjust the import path according to your project
)

// GSTR9Handler is the struct for the GSTR9 gRPC handler
type GSTR9Handler struct {
	pb.UnimplementedGSTR9ServiceServer
	gstr9Service *service.GSTR9Service
}

// NewGSTR9Handler creates a new GSTR9 handler
func NewGSTR9Handler(gstr9Service *service.GSTR9Service) *GSTR9Handler {
	return &GSTR9Handler{
		gstr9Service: gstr9Service,
	}
}

// ConvertProtoToDomainGSTR9 converts a protobuf GSTR9Request to a domain GSTR9Request
func ConvertProtoToDomainGSTR9(req *pb.GSTR9Request) *domain.GSTR9Request {
	return &domain.GSTR9Request{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		Turnover:     req.GetTotalTurnover(),
		TaxPayable:   req.GetTaxPayable(),
	}
}

// SaveGSTR9 handles the SaveGSTR9 gRPC method
func (h *GSTR9Handler) SaveGSTR9(ctx context.Context, req *pb.GSTR9Request) (*pb.GSTR9Response, error) {
	// Convert the protobuf request to the domain request
	domainReq := ConvertProtoToDomainGSTR9(req)

	// Call the service layer to save the GSTR9 data
	refID, err := h.gstr9Service.SaveGSTR9Data(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to save GSTR9 data: %w", err)
	}

	return &pb.GSTR9Response{
		RefId:  refID,
		Status: "Success",
	}, nil
}

// SubmitGSTR9 handles the SubmitGSTR9 gRPC method
func (h *GSTR9Handler) SubmitGSTR9(ctx context.Context, req *pb.GSTR9SubmitRequest) (*pb.GSTR9SubmitResponse, error) {
	domainReq := &domain.GSTR9SubmitRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
	}

	arn, err := h.gstr9Service.SubmitGSTR9Data(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to submit GSTR9 data: %w", err)
	}

	return &pb.GSTR9SubmitResponse{
		Arn:    arn,
		Status: "Success",
	}, nil
}

// FileGSTR9 handles the FileGSTR9 gRPC method
func (h *GSTR9Handler) FileGSTR9(ctx context.Context, req *pb.GSTR9FileRequest) (*pb.GSTR9FileResponse, error) {
	domainReq := &domain.GSTR9FileRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		ARN:          req.GetArn(),
	}

	filingStatus, message, err := h.gstr9Service.FileGSTR9Data(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to file GSTR9: %w", err)
	}

	return &pb.GSTR9FileResponse{
		FilingStatus: filingStatus,
		Message:      message,
	}, nil
}

// GetGSTR9Status handles the GetGSTR9Status gRPC method
func (h *GSTR9Handler) GetGSTR9Status(ctx context.Context, req *pb.GSTR9StatusRequest) (*pb.GSTR9StatusResponse, error) {
	domainReq := &domain.GSTR9StatusRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		ARN:          req.GetArn(),
	}

	status, message, err := h.gstr9Service.GetGSTR9Status(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get GSTR9 status: %w", err)
	}

	return &pb.GSTR9StatusResponse{
		Status:  status,
		Message: message,
	}, nil
}
