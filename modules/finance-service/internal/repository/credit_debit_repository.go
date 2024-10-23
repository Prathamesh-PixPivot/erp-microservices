package repository

import (
	"finance-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreditDebitNoteRepository interface {
	Create(note *models.CreditDebitNote) error
	GetByID(id uuid.UUID) (*models.CreditDebitNote, error)
	Update(note *models.CreditDebitNote) error
	Delete(id uuid.UUID) error
	ListByInvoiceID(invoiceID uuid.UUID) ([]*models.CreditDebitNote, error)
	ListAll() ([]*models.CreditDebitNote, error)
}

type creditDebitNoteRepository struct {
	db *gorm.DB
}

func NewCreditDebitNoteRepository(db *gorm.DB) CreditDebitNoteRepository {
	return &creditDebitNoteRepository{db: db}
}

func (r *creditDebitNoteRepository) Create(note *models.CreditDebitNote) error {
	return r.db.Create(note).Error
}

func (r *creditDebitNoteRepository) GetByID(id uuid.UUID) (*models.CreditDebitNote, error) {
	var note models.CreditDebitNote
	err := r.db.Where("id = ?", id).First(&note).Error
	return &note, err
}

func (r *creditDebitNoteRepository) Update(note *models.CreditDebitNote) error {
	return r.db.Save(note).Error
}

func (r *creditDebitNoteRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.CreditDebitNote{}, id).Error
}

func (r *creditDebitNoteRepository) ListByInvoiceID(invoiceID uuid.UUID) ([]*models.CreditDebitNote, error) {
	var notes []*models.CreditDebitNote
	err := r.db.Where("invoice_id = ?", invoiceID).Find(&notes).Error
	return notes, err
}

func (r *creditDebitNoteRepository) ListAll() ([]*models.CreditDebitNote, error) {
	var notes []*models.CreditDebitNote
	err := r.db.Find(&notes).Error
	return notes, err
}
