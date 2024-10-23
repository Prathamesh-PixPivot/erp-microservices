package services

import (
	"finance-service/internal/models"
	"finance-service/internal/repository"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type InvoiceService interface {
	CreateInvoice(invoice *models.Invoice) error
	GetInvoiceByID(id uuid.UUID) (*models.Invoice, error)
	ListInvoices(page, pageSize int) ([]*models.Invoice, error)
	UpdateInvoice(invoice *models.Invoice) error
	DeleteInvoice(id uuid.UUID) error
	CalculateTotalAmount(items []models.InvoiceItem) float64
	CalculateTaxes(totalAmount float64, invoiceType string) (cgst, sgst, igst float64)
	GenerateInvoiceNumber(organizationID uuid.UUID) (string, error)
}

type invoiceService struct {
	repo repository.InvoiceRepository
}

func NewInvoiceService(repo repository.InvoiceRepository) InvoiceService {
	return &invoiceService{repo: repo}
}

func (s *invoiceService) CreateInvoice(invoice *models.Invoice) error {
	return s.repo.Create(invoice)
}

func (s *invoiceService) GetInvoiceByID(id uuid.UUID) (*models.Invoice, error) {
	return s.repo.FindByID(id)
}

func (s *invoiceService) ListInvoices(page, pageSize int) ([]*models.Invoice, error) {
	return s.repo.FindAll(page, pageSize)
}

func (s *invoiceService) UpdateInvoice(invoice *models.Invoice) error {
	return s.repo.Update(invoice)
}

func (s *invoiceService) DeleteInvoice(id uuid.UUID) error {
	return s.repo.Delete(id)
}

// CalculateTotalAmount calculates the total amount from the provided items
func (s *invoiceService) CalculateTotalAmount(items []models.InvoiceItem) float64 {
	var total float64
	for _, item := range items {
		total += item.Total
	}
	return total
}

// CalculateTaxes calculates CGST, SGST, and IGST based on the total amount and the invoice type
func (s *invoiceService) CalculateTaxes(totalAmount float64, invoiceType string) (cgst, sgst, igst float64) {
	// Assuming standard GST rates. These can be adjusted based on requirements.
	if invoiceType == "purchase" {
		// IGST for purchase invoices
		igst = totalAmount * 0.18
	} else {
		// CGST + SGST for other types of invoices
		cgst = totalAmount * 0.09
		sgst = totalAmount * 0.09
	}
	return
}

// GenerateInvoiceNumber generates an invoice number based on a default pattern or custom pattern
func (s *invoiceService) GenerateInvoiceNumber(organizationID uuid.UUID) (string, error) {
	// Fetch organization details (for org initials, etc.)
	org, err := s.repo.GetOrganizationByID(organizationID)
	if err != nil {
		return "", err
	}

	// Extract organization initials from first 2 words of the organization name if available
	initials := ""
	if len(org.Name) > 0 {
		words := strings.Fields(org.Name)
		if len(words) > 1 {
			initials = strings.ToUpper(string(words[0][0]) + string(words[1][0]))
		} else {
			initials = strings.ToUpper(string(words[0][0]) + string(words[0][1]))
		}
	}

	// Get the current financial year (e.g., 2023-2024)
	year := time.Now().Year()
	nextYear := year + 1
	financialYear := fmt.Sprintf("%d-%d", year, nextYear)

	// Fetch the latest invoice number from the database
	lastInvoiceNo, err := s.repo.GetLastInvoiceNumber(organizationID)
	if err != nil {
		return "", err
	}
	newInvoiceNo := lastInvoiceNo + 1

	// Format the invoice number as per the default pattern
	invoiceNumber := fmt.Sprintf("%s/%d/%s", initials, newInvoiceNo, financialYear)

	return invoiceNumber, nil
}
