package services

import (
	"finance-service/internal/models"
	"finance-service/internal/repository"

	"github.com/google/uuid"
)

type LedgerService interface {
	CreateLedgerEntry(ledgerEntry *models.LedgerEntry) error
	GetLedgerEntryByID(id uuid.UUID) (*models.LedgerEntry, error)
	UpdateLedgerEntry(ledgerEntry *models.LedgerEntry) error
	DeleteLedgerEntry(id uuid.UUID) error
	ListAllLedgerEntries() ([]*models.LedgerEntry, error)
}

type ledgerService struct {
	ledgerRepo repository.LedgerRepository
}

func NewLedgerService(repo repository.LedgerRepository) LedgerService {
	return &ledgerService{ledgerRepo: repo}
}

func (s *ledgerService) CreateLedgerEntry(ledgerEntry *models.LedgerEntry) error {
	return s.ledgerRepo.Create(ledgerEntry)
}

func (s *ledgerService) GetLedgerEntryByID(id uuid.UUID) (*models.LedgerEntry, error) {
	return s.ledgerRepo.GetByID(id)
}

func (s *ledgerService) UpdateLedgerEntry(ledgerEntry *models.LedgerEntry) error {
	return s.ledgerRepo.Update(ledgerEntry)
}

func (s *ledgerService) DeleteLedgerEntry(id uuid.UUID) error {
	return s.ledgerRepo.Delete(id)
}

func (s *ledgerService) ListAllLedgerEntries() ([]*models.LedgerEntry, error) {
	return s.ledgerRepo.ListAll()
}
