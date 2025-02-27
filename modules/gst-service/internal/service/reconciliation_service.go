package service

import (
	"fmt"
	"gst-service/internal/domain"
	"gst-service/internal/infrastructure/external"
	"gst-service/internal/repository"
)

// ReconciliationService contains the logic for reconciling GSTR1 with GSTR2A
type ReconciliationService struct {
	gstr1Repo  *repository.GSTRRepository
	gstr2ARepo *repository.GSTR2ARepository
	gstClient  *external.GSTClient
}

// NewReconciliationService creates a new reconciliation service
func NewReconciliationService(gstr1Repo *repository.GSTRRepository, gstr2ARepo *repository.GSTR2ARepository, gstClient *external.GSTClient) *ReconciliationService {
	return &ReconciliationService{
		gstr1Repo:  gstr1Repo,
		gstr2ARepo: gstr2ARepo,
		gstClient:  gstClient,
	}
}

// ReconcileGSTR1WithGSTR2A performs reconciliation between GSTR1 and GSTR2A
func (s *ReconciliationService) ReconcileGSTR1WithGSTR2A(gstin, returnPeriod string) (*domain.GSTR1ReconcileResponse, error) {
	// Fetch GSTR1 invoices from database
	gstr1Invoices, err := s.gstr1Repo.GetInvoicesByGSTIN(gstin, returnPeriod)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch GSTR1 invoices: %w", err)
	}

	// Fetch GSTR2A invoices from GST Portal
	gstr2AData, err := s.gstClient.ReconcileGSTR1WithGSTR2A(gstin, returnPeriod)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch GSTR2A invoices from GST Portal: %w", err)
	}

	gstr2AInvoices, ok := gstr2AData["invoices"].([]map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid GSTR2A response format")
	}

	mismatches := []domain.Invoice{}

	// Convert GSTR2A API response to domain.Invoice format
	gstr2AInvoiceMap := make(map[string]domain.Invoice)
	for _, inv := range gstr2AInvoices {
		invoice := domain.Invoice{
			InvoiceNumber: inv["invoice_number"].(string),
			InvoiceDate:   inv["invoice_date"].(string),
			TaxableValue:  inv["taxable_value"].(string),
			TaxAmount:     inv["tax_amount"].(string),
		}
		gstr2AInvoiceMap[invoice.InvoiceNumber] = invoice
	}

	// Compare GSTR1 invoices with GSTR2A invoices
	for _, gstr1Inv := range gstr1Invoices {
		if gstr2AInv, exists := gstr2AInvoiceMap[gstr1Inv.InvoiceNumber]; exists {
			// Check for mismatches
			if gstr1Inv.TaxableValue != gstr2AInv.TaxableValue || gstr1Inv.TaxAmount != gstr2AInv.TaxAmount {
				mismatches = append(mismatches, gstr1Inv)
			}
		} else {
			// Invoice present in GSTR1 but missing in GSTR2A
			mismatches = append(mismatches, gstr1Inv)
		}
	}

	return &domain.GSTR1ReconcileResponse{
		Status:                "Completed",
		ReconciliationDetails: mismatches,
	}, nil
}
