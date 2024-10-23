package repository

import (
	"finance-service/internal/models"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoiceRepository interface {
	Create(invoice *models.Invoice) error
	FindByID(id uuid.UUID) (*models.Invoice, error)
	FindAll(page, pageSize int) ([]*models.Invoice, error)
	Update(invoice *models.Invoice) error
	Delete(id uuid.UUID) error
	GetLastInvoiceNumber(organizationID uuid.UUID) (int, error)
	GetOrganizationByID(organizationID uuid.UUID) (*models.Organization, error)
}

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) InvoiceRepository {
	return &invoiceRepository{db: db}
}

func (r *invoiceRepository) Create(invoice *models.Invoice) error {
	return r.db.Create(invoice).Error
}

func (r *invoiceRepository) FindByID(id uuid.UUID) (*models.Invoice, error) {
	var invoice models.Invoice
	if err := r.db.Preload("Items").First(&invoice, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (r *invoiceRepository) FindAll(page, pageSize int) ([]*models.Invoice, error) {
	var invoices []*models.Invoice
	offset := (page - 1) * pageSize
	if err := r.db.Preload("Items").Limit(pageSize).Offset(offset).Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}

func (r *invoiceRepository) Update(invoice *models.Invoice) error {
	return r.db.Save(invoice).Error
}

func (r *invoiceRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Invoice{}, "id = ?", id).Error
}

// Fetch the latest invoice number from the database
func (r *invoiceRepository) GetLastInvoiceNumber(organizationID uuid.UUID) (int, error) {
	var lastInvoice models.Invoice
	err := r.db.Where("organization_id = ?", organizationID).Order("created_at desc").First(&lastInvoice).Error
	if err != nil {
		return 0, err
	}
	invoiceNumber, err := strconv.Atoi(lastInvoice.InvoiceNumber)
	if err != nil {
		return 0, err
	}
	return invoiceNumber, nil
}

// Fetch organization details by ID
func (r *invoiceRepository) GetOrganizationByID(organizationID uuid.UUID) (*models.Organization, error) {
	var organization models.Organization
	err := r.db.First(&organization, "id = ?", organizationID).Error
	if err != nil {
		return nil, err
	}
	return &organization, nil
}
