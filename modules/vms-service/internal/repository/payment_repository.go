package repository

import (
    "vms-service/internal/models"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type PaymentRepository struct {
    db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
    return &PaymentRepository{db: db}
}

func (repo *PaymentRepository) ProcessInvoice(payment *models.Payment) error {
    payment.ID = uuid.New()
    return repo.db.Create(payment).Error
}

func (repo *PaymentRepository) UpdatePaymentStatus(id uuid.UUID, status string) error {
    return repo.db.Model(&models.Payment{}).Where("id = ?", id).Update("status", status).Error
}

func (repo *PaymentRepository) GetPaymentByID(id uuid.UUID) (*models.Payment, error) {
    var payment models.Payment
    return &payment, repo.db.First(&payment, "id = ?", id).Error
}

func (repo *PaymentRepository) DeletePayment(id uuid.UUID) error {
    return repo.db.Delete(&models.Payment{}, "id = ?", id).Error
}
