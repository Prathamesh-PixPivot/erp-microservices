package services

import (
	"vms-service/internal/models"
	"vms-service/internal/repository"

	"github.com/google/uuid"
)

type PaymentService struct {
	repo *repository.PaymentRepository
}

func NewPaymentService(repo *repository.PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

func (service *PaymentService) ProcessInvoice(payment *models.Payment) error {
	return service.repo.ProcessInvoice(payment)
}

func (service *PaymentService) UpdatePaymentStatus(id uuid.UUID, status string) error {
	return service.repo.UpdatePaymentStatus(id, status)
}

func (service *PaymentService) GetPaymentByID(id uuid.UUID) (*models.Payment, error) {
	return service.repo.GetPaymentByID(id)
}

func (service *PaymentService) DeletePayment(id uuid.UUID) error {
	return service.repo.DeletePayment(id)
}
