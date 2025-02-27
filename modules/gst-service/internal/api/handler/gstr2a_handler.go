package handler

import (
	"context"
	"fmt"
	pb "gst-service/internal/api/protobufs/gst-service" // Import generated protobuf code
	"gst-service/internal/domain"                       // Your domain models
	"gst-service/internal/service"                      // Adjust the import path according to your project
)

// GSTR2AHandler is the struct for the GSTR2A gRPC handler
type GSTR2AHandler struct {
	pb.UnimplementedGSTR2AServiceServer
	gstr2AService *service.GSTR2AService
}

// NewGSTR2AHandler creates a new GSTR2A handler
func NewGSTR2AHandler(gstr2AService *service.GSTR2AService) *GSTR2AHandler {
	return &GSTR2AHandler{
		gstr2AService: gstr2AService,
	}
}

// ConvertProtoToDomainGSTR2A converts a protobuf GSTR2ARequest to a domain GSTR2ARequest
func ConvertProtoToDomainGSTR2A(req *pb.GSTR2ARequest) *domain.GSTR2ARequest {
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

	return &domain.GSTR2ARequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		Invoices:     invoices,
	}
}

// SaveGSTR2A handles the SaveGSTR2A gRPC method
func (h *GSTR2AHandler) SaveGSTR2A(ctx context.Context, req *pb.GSTR2ARequest) (*pb.GSTR2AResponse, error) {
	// Convert the protobuf request to the domain request
	domainReq := ConvertProtoToDomainGSTR2A(req)

	// Call the service layer to save the GSTR2A data
	refID, err := h.gstr2AService.SaveGSTR2AData(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to save GSTR2A data: %w", err)
	}

	return &pb.GSTR2AResponse{
		RefId:  refID,
		Status: "Success",
	}, nil
}

// SubmitGSTR2A handles the SubmitGSTR2A gRPC method
func (h *GSTR2AHandler) SubmitGSTR2A(ctx context.Context, req *pb.GSTR2ASubmitRequest) (*pb.GSTR2ASubmitResponse, error) {
	domainReq := &domain.GSTR2ASubmitRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
	}

	arn, err := h.gstr2AService.SubmitGSTR2AData(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to submit GSTR2A data: %w", err)
	}

	return &pb.GSTR2ASubmitResponse{
		Arn:    arn,
		Status: "Success",
	}, nil
}

// FileGSTR2A handles the FileGSTR2A gRPC method
func (h *GSTR2AHandler) FileGSTR2A(ctx context.Context, req *pb.GSTR2AFileRequest) (*pb.GSTR2AFileResponse, error) {
	domainReq := &domain.GSTR2AFileRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		ARN:          req.GetArn(),
	}

	filingStatus, message, err := h.gstr2AService.FileGSTR2AData(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to file GSTR2A: %w", err)
	}

	return &pb.GSTR2AFileResponse{
		FilingStatus: filingStatus,
		Message:      message,
	}, nil
}

// GetGSTR2AStatus handles the GetGSTR2AStatus gRPC method
func (h *GSTR2AHandler) GetGSTR2AStatus(ctx context.Context, req *pb.GSTR2AStatusRequest) (*pb.GSTR2AStatusResponse, error) {
	domainReq := &domain.GSTR2AStatusRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		ARN:          req.GetArn(),
	}

	status, message, err := h.gstr2AService.GetGSTR2AStatus(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get GSTR2A status: %w", err)
	}

	return &pb.GSTR2AStatusResponse{
		Status:  status,
		Message: message,
	}, nil
}
