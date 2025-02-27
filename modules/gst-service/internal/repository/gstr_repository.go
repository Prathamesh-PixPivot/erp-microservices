package repository

import (
	"gst-service/internal/domain" // Adjust the import path based on your project

	"gorm.io/gorm"
)

// GSTR1ARepository interacts with the GSTR1A data in the database
type GSTR1ARepository struct {
	db *gorm.DB
}

// GSTRRepository is responsible for database interactions related to GSTR
type GSTRRepository struct {
	db *gorm.DB
}

// GSTR2ARepository interacts with the GSTR2A data in the database
type GSTR2ARepository struct {
	db *gorm.DB
}

// GSTR3BRepository interacts with the GSTR3B data in the database
type GSTR3BRepository struct {
	db *gorm.DB
}

// GSTR9Repository interacts with the GSTR9 data in the database
type GSTR9Repository struct {
	db *gorm.DB
}

// GSTR9CRepository interacts with the GSTR9C data in the database
type GSTR9CRepository struct {
	db *gorm.DB
}

// NewGSTRRepository creates a new GSTRRepository
func NewGSTRRepository(db *gorm.DB) *GSTRRepository {
	return &GSTRRepository{db: db}
}

// NewGSTR2ARepository creates a new GSTR2ARepository
func NewGSTR2ARepository(db *gorm.DB) *GSTR2ARepository {
	return &GSTR2ARepository{db: db}
}

// NewGSTR1ARepository creates a new GSTR1ARepository
func NewGSTR1ARepository(db *gorm.DB) *GSTR1ARepository {
	return &GSTR1ARepository{db: db}
}

// NewGSTR3BRepository creates a new GSTR3BRepository
func NewGSTR3BRepository(db *gorm.DB) *GSTR3BRepository {
	return &GSTR3BRepository{db: db}
}

// NewGSTR9Repository creates a new GSTR9Repository
func NewGSTR9Repository(db *gorm.DB) *GSTR9Repository {
	return &GSTR9Repository{db: db}
}

// NewGSTR9CRepository creates a new GSTR9CRepository
func NewGSTR9CRepository(db *gorm.DB) *GSTR9CRepository {
	return &GSTR9CRepository{db: db}
}

// SaveGSTR1 saves the GSTR1 data to the database
func (r *GSTRRepository) SaveGSTR1(request *domain.GSTR1Request) error {
	// Here you would insert the data into your database (using GORM)
	return r.db.Create(&request).Error
}

// SaveGSTR1A saves the GSTR1A data into the database
func (r *GSTR1ARepository) SaveGSTR1A(request *domain.GSTR1ARequest) error {
	return r.db.Create(&request).Error
}

// SaveInvoice saves individual invoices (if needed separately)
func (r *GSTR1ARepository) SaveInvoice(invoice *domain.Invoice) error {
	return r.db.Create(&invoice).Error
}

// SaveGSTR2A saves the GSTR2A data into the database
func (r *GSTR2ARepository) SaveGSTR2A(request *domain.GSTR2ARequest) error {
	return r.db.Create(&request).Error
}

// SaveInvoice saves individual invoices (if needed separately)
func (r *GSTR2ARepository) SaveInvoice(invoice *domain.Invoice) error {
	return r.db.Create(&invoice).Error
}

// SaveGSTR3B saves the GSTR3B data into the database
func (r *GSTR3BRepository) SaveGSTR3B(request *domain.GSTR3BRequest) error {
	return r.db.Create(&request).Error
}

// SaveGSTR9 saves the GSTR9 data into the database
func (r *GSTR9Repository) SaveGSTR9(request *domain.GSTR9Request) error {
	return r.db.Create(&request).Error
}

// SaveGSTR9C saves the GSTR9C data into the database
func (r *GSTR9CRepository) SaveGSTR9C(request *domain.GSTR9CRequest) error {
	return r.db.Create(&request).Error
}

// GetInvoicesByGSTIN fetches invoices for a given GSTIN and return period
func (r *GSTRRepository) GetInvoicesByGSTIN(gstin, returnPeriod string) ([]domain.Invoice, error) {
	var invoices []domain.Invoice
	err := r.db.Where("gstin = ? AND return_period = ?", gstin, returnPeriod).Find(&invoices).Error
	return invoices, err
}
