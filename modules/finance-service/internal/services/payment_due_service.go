package services

import (
	"finance-service/internal/models"
	"finance-service/internal/repository"

	"github.com/google/uuid"
)

type PaymentDueService interface {
	CreatePaymentDue(paymentDue *models.PaymentDue) error
	GetPaymentDueByID(id uuid.UUID) (*models.PaymentDue, error)
	UpdatePaymentDue(paymentDue *models.PaymentDue) error
	DeletePaymentDue(id uuid.UUID) error
	ListAllPaymentDues() ([]*models.PaymentDue, error)
	SearchAndSortPaymentDues(query string, sortBy string, sortDirection string) ([]*models.PaymentDue, error)
	ListOverduePayments() ([]*models.PaymentDue, error)
}

type paymentDueService struct {
	paymentDueRepo repository.PaymentDueRepository
}

func NewPaymentDueService(repo repository.PaymentDueRepository) PaymentDueService {
	return &paymentDueService{paymentDueRepo: repo}
}

func (s *paymentDueService) CreatePaymentDue(paymentDue *models.PaymentDue) error {
	return s.paymentDueRepo.Create(paymentDue)
}

func (s *paymentDueService) GetPaymentDueByID(id uuid.UUID) (*models.PaymentDue, error) {
	return s.paymentDueRepo.GetByID(id)
}

func (s *paymentDueService) UpdatePaymentDue(paymentDue *models.PaymentDue) error {
	return s.paymentDueRepo.Update(paymentDue)
}

func (s *paymentDueService) DeletePaymentDue(id uuid.UUID) error {
	return s.paymentDueRepo.Delete(id)
}

func (s *paymentDueService) ListAllPaymentDues() ([]*models.PaymentDue, error) {
	return s.paymentDueRepo.ListAll()
}

func (s *paymentDueService) SearchAndSortPaymentDues(query string, sortBy string, sortDirection string) ([]*models.PaymentDue, error) {
	return s.paymentDueRepo.SearchAndSort(query, sortBy, sortDirection)
}

func (s *paymentDueService) ListOverduePayments() ([]*models.PaymentDue, error) {
	return s.paymentDueRepo.ListOverduePayments()
}
