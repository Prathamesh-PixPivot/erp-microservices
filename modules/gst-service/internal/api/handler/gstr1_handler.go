package handler

import (
	"context"
	"fmt"
	pb "gst-service/internal/api/protobufs/gst-service" // Import generated protobuf code
	"gst-service/internal/domain"                       // Your domain models
	"gst-service/internal/infrastructure/logger"
	"gst-service/internal/service" // Adjust the import path according to your project

	"go.uber.org/zap"
)

// GSTR1Handler is the struct for the GSTR1 gRPC handler
type GSTR1Handler struct {
	pb.UnimplementedGSTR1ServiceServer // Embedding UnimplementedGSTServiceServer
	gstr1Service                       *service.GSTR1Service
}

// NewGSTR1Handler creates a new GSTR1 handler
func NewGSTR1Handler(gstr1Service *service.GSTR1Service) *GSTR1Handler {
	return &GSTR1Handler{
		gstr1Service: gstr1Service,
	}
}

// ConvertProtoToDomain converts a protobuf GSTR1Request to a domain GSTR1Request
func ConvertProtoToDomain(req *pb.GSTR1Request) *domain.GSTR1Request {
	var invoices []domain.Invoice
	for _, inv := range req.GetInvoices() {
		invoices = append(invoices, domain.Invoice{
			InvoiceNumber: inv.GetInvoiceNumber(),
			InvoiceDate:   inv.GetInvoiceDate(),
			GSTINSupplier: inv.GetGstinSupplier(),
			GSTINReceiver: inv.GetGstinReceiver(),
			TaxableValue:  inv.GetTaxableValue(),
			TaxAmount:     inv.GetTaxAmount(),
			HSNCode:       inv.GetHsnCode(),
			PlaceOfSupply: inv.GetPlaceOfSupply(),
		})
	}

	return &domain.GSTR1Request{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		Invoices:     invoices,
	}
}

// SaveGSTR1 handles the SaveGSTR1 gRPC method
func (h *GSTR1Handler) SaveGSTR1(ctx context.Context, req *pb.GSTR1Request) (*pb.GSTR1Response, error) {
	// Convert the protobuf request to the domain request
	domainReq := ConvertProtoToDomain(req)

	// Call the service layer to save the GSTR1 data
	refID, err := h.gstr1Service.SaveGSTR1Data(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to save GSTR1 data: %w", err)
	}

	logger.AppLogger.Info("GSTR1 data saved successfully", zap.String("refID", refID))
	return &pb.GSTR1Response{
		RefId:  refID,
		Status: "Success",
	}, nil
}

// SubmitGSTR1 handles the SubmitGSTR1 gRPC method
func (h *GSTR1Handler) SubmitGSTR1(ctx context.Context, req *pb.GSTR1SubmitRequest) (*pb.GSTR1SubmitResponse, error) {
	// Convert the protobuf request to the domain request
	domainReq := &domain.GSTR1SubmitRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
	}

	// Call the service layer to submit the GSTR1 data
	arn, err := h.gstr1Service.SubmitGSTR1Data(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to submit GSTR1 data: %w", err)
	}

	return &pb.GSTR1SubmitResponse{
		Arn:    arn,
		Status: "Success",
	}, nil
}

// FileGSTR1 handles the FileGSTR1 gRPC method
func (h *GSTR1Handler) FileGSTR1(ctx context.Context, req *pb.GSTR1FileRequest) (*pb.GSTR1FileResponse, error) {
	// Convert the protobuf request to the domain request
	domainReq := &domain.GSTR1FileRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		ARN:          req.GetArn(),
	}

	// Call the service layer to file the GSTR1 data
	filingStatus, message, err := h.gstr1Service.FileGSTR1Data(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to file GSTR1: %w", err)
	}

	return &pb.GSTR1FileResponse{
		FilingStatus: filingStatus,
		Message:      message,
	}, nil
}

// GetGSTR1Status handles the GetGSTR1Status gRPC method
func (h *GSTR1Handler) GetGSTR1Status(ctx context.Context, req *pb.GSTR1StatusRequest) (*pb.GSTR1StatusResponse, error) {
	// Convert the protobuf request to the domain request
	domainReq := &domain.GSTR1StatusRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		ARN:          req.GetArn(),
	}

	// Call the service layer to get GSTR1 status
	status, message, err := h.gstr1Service.GetGSTR1Status(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get GSTR1 status: %w", err)
	}

	return &pb.GSTR1StatusResponse{
		Status:  status,
		Message: message,
	}, nil
}
