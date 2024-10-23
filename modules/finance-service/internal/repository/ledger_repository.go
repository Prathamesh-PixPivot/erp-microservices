package repository

import (
	"finance-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LedgerRepository interface {
	Create(ledgerEntry *models.LedgerEntry) error
	GetByID(id uuid.UUID) (*models.LedgerEntry, error)
	Update(ledgerEntry *models.LedgerEntry) error
	Delete(id uuid.UUID) error
	ListAll() ([]*models.LedgerEntry, error)
}

type ledgerRepository struct {
	db *gorm.DB
}

func NewLedgerRepository(db *gorm.DB) LedgerRepository {
	return &ledgerRepository{db: db}
}

func (r *ledgerRepository) Create(ledgerEntry *models.LedgerEntry) error {
	return r.db.Create(ledgerEntry).Error
}

func (r *ledgerRepository) GetByID(id uuid.UUID) (*models.LedgerEntry, error) {
	var entry models.LedgerEntry
	err := r.db.Where("id = ?", id).First(&entry).Error
	return &entry, err
}

func (r *ledgerRepository) Update(ledgerEntry *models.LedgerEntry) error {
	return r.db.Save(ledgerEntry).Error
}

func (r *ledgerRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.LedgerEntry{}, id).Error
}

func (r *ledgerRepository) ListAll() ([]*models.LedgerEntry, error) {
	var ledgerEntries []*models.LedgerEntry
	err := r.db.Find(&ledgerEntries).Error
	return ledgerEntries, err
}
