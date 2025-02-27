package handler

import (
	"context"
	"fmt"
	pb "gst-service/internal/api/protobufs/gst-service" // Import generated protobuf code
	"gst-service/internal/domain"                       // Your domain models
	"gst-service/internal/service"                      // Adjust the import path according to your project
)

// GSTR1AHandler is the struct for the GSTR1A gRPC handler
type GSTR1AHandler struct {
	pb.UnimplementedGSTR1AServiceServer
	gstr1AService *service.GSTR1AService
}

// NewGSTR1AHandler creates a new GSTR1A handler
func NewGSTR1AHandler(gstr1AService *service.GSTR1AService) *GSTR1AHandler {
	return &GSTR1AHandler{
		gstr1AService: gstr1AService,
	}
}

// ConvertProtoToDomainGSTR1A converts a protobuf GSTR1ARequest to a domain GSTR1ARequest
func ConvertProtoToDomainGSTR1A(req *pb.GSTR1ARequest) *domain.GSTR1ARequest {
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

	return &domain.GSTR1ARequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		Invoices:     invoices,
	}
}

// SaveGSTR1A handles the SaveGSTR1A gRPC method
func (h *GSTR1AHandler) SaveGSTR1A(ctx context.Context, req *pb.GSTR1ARequest) (*pb.GSTR1AResponse, error) {
	// Convert the protobuf request to the domain request
	domainReq := ConvertProtoToDomainGSTR1A(req)

	// Call the service layer to save the GSTR1A data
	refID, err := h.gstr1AService.SaveGSTR1AData(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to save GSTR1A data: %w", err)
	}

	return &pb.GSTR1AResponse{
		RefId:  refID,
		Status: "Success",
	}, nil
}

// SubmitGSTR1A handles the SubmitGSTR1A gRPC method
func (h *GSTR1AHandler) SubmitGSTR1A(ctx context.Context, req *pb.GSTR1ASubmitRequest) (*pb.GSTR1ASubmitResponse, error) {
	domainReq := &domain.GSTR1ASubmitRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
	}

	arn, err := h.gstr1AService.SubmitGSTR1AData(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to submit GSTR1A data: %w", err)
	}

	return &pb.GSTR1ASubmitResponse{
		Arn:    arn,
		Status: "Success",
	}, nil
}

// FileGSTR1A handles the FileGSTR1A gRPC method
func (h *GSTR1AHandler) FileGSTR1A(ctx context.Context, req *pb.GSTR1AFileRequest) (*pb.GSTR1AFileResponse, error) {
	domainReq := &domain.GSTR1AFileRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		ARN:          req.GetArn(),
	}

	filingStatus, message, err := h.gstr1AService.FileGSTR1AData(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to file GSTR1A: %w", err)
	}

	return &pb.GSTR1AFileResponse{
		FilingStatus: filingStatus,
		Message:      message,
	}, nil
}

// GetGSTR1AStatus handles the GetGSTR1AStatus gRPC method
func (h *GSTR1AHandler) GetGSTR1AStatus(ctx context.Context, req *pb.GSTR1AStatusRequest) (*pb.GSTR1AStatusResponse, error) {
	domainReq := &domain.GSTR1AStatusRequest{
		GSTIN:        req.GetGstin(),
		ReturnPeriod: req.GetReturnPeriod(),
		ARN:          req.GetArn(),
	}

	status, message, err := h.gstr1AService.GetGSTR1AStatus(domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get GSTR1A status: %w", err)
	}

	return &pb.GSTR1AStatusResponse{
		Status:  status,
		Message: message,
	}, nil
}
