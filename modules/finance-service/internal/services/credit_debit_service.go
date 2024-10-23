package services

import (
	"finance-service/internal/models"
	"finance-service/internal/repository"

	"github.com/google/uuid"
)

type CreditDebitNoteService interface {
	CreateCreditDebitNote(note *models.CreditDebitNote) error
	GetCreditDebitNoteByID(id uuid.UUID) (*models.CreditDebitNote, error)
	UpdateCreditDebitNote(note *models.CreditDebitNote) error
	DeleteCreditDebitNote(id uuid.UUID) error
	ListCreditDebitNotesByInvoiceID(invoiceID uuid.UUID) ([]*models.CreditDebitNote, error)
	ListAllCreditDebitNotes() ([]*models.CreditDebitNote, error)
}

type creditDebitNoteService struct {
	noteRepo repository.CreditDebitNoteRepository
}

func NewCreditDebitNoteService(repo repository.CreditDebitNoteRepository) CreditDebitNoteService {
	return &creditDebitNoteService{noteRepo: repo}
}

func (s *creditDebitNoteService) CreateCreditDebitNote(note *models.CreditDebitNote) error {
	return s.noteRepo.Create(note)
}

func (s *creditDebitNoteService) GetCreditDebitNoteByID(id uuid.UUID) (*models.CreditDebitNote, error) {
	return s.noteRepo.GetByID(id)
}

func (s *creditDebitNoteService) UpdateCreditDebitNote(note *models.CreditDebitNote) error {
	return s.noteRepo.Update(note)
}

func (s *creditDebitNoteService) DeleteCreditDebitNote(id uuid.UUID) error {
	return s.noteRepo.Delete(id)
}

func (s *creditDebitNoteService) ListCreditDebitNotesByInvoiceID(invoiceID uuid.UUID) ([]*models.CreditDebitNote, error) {
	return s.noteRepo.ListByInvoiceID(invoiceID)
}

func (s *creditDebitNoteService) ListAllCreditDebitNotes() ([]*models.CreditDebitNote, error) {
	return s.noteRepo.ListAll()
}
