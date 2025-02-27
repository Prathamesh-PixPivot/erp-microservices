package handler

import (
	"context"
	"fmt"
	pb "gst-service/internal/api/protobufs/gst-service"
	"gst-service/internal/service"
)

// ReconciliationHandler handles gRPC requests for reconciliation
type ReconciliationHandler struct {
	pb.UnimplementedGSTR1ServiceServer
	reconciliationService *service.ReconciliationService
}

// NewReconciliationHandler creates a new handler for reconciliation
func NewReconciliationHandler(reconciliationService *service.ReconciliationService) *ReconciliationHandler {
	return &ReconciliationHandler{
		reconciliationService: reconciliationService,
	}
}

// ReconcileGSTR1WithGSTR2A handles the gRPC request for reconciliation
func (h *ReconciliationHandler) ReconcileGSTR1WithGSTR2A(ctx context.Context, req *pb.GSTR1ReconcileRequest) (*pb.GSTR1ReconcileResponse, error) {
	response, err := h.reconciliationService.ReconcileGSTR1WithGSTR2A(req.GetGstin(), req.GetReturnPeriod())
	if err != nil {
		return nil, fmt.Errorf("failed to reconcile GSTR1 with GSTR2A: %w", err)
	}

	var grpcInvoices []*pb.GSTR1Invoice
	for _, inv := range response.ReconciliationDetails {
		grpcInvoices = append(grpcInvoices, &pb.GSTR1Invoice{
			InvoiceNumber: inv.InvoiceNumber,
			InvoiceDate:   inv.InvoiceDate,
			TaxableValue:  inv.TaxableValue,
			TaxAmount:     inv.TaxAmount,
		})
	}

	return &pb.GSTR1ReconcileResponse{
		Status:                response.Status,
		ReconciliationDetails: grpcInvoices,
	}, nil
}
