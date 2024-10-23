package repository

import (
	"finance-service/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentDueRepository interface {
	Create(paymentDue *models.PaymentDue) error
	GetByID(id uuid.UUID) (*models.PaymentDue, error)
	Update(paymentDue *models.PaymentDue) error
	Delete(id uuid.UUID) error
	ListAll() ([]*models.PaymentDue, error)
	SearchAndSort(query string, sortBy string, sortDirection string) ([]*models.PaymentDue, error)
	ListOverduePayments() ([]*models.PaymentDue, error)
}

type paymentDueRepository struct {
	db *gorm.DB
}

func NewPaymentDueRepository(db *gorm.DB) PaymentDueRepository {
	return &paymentDueRepository{db: db}
}

func (r *paymentDueRepository) Create(paymentDue *models.PaymentDue) error {
	return r.db.Create(paymentDue).Error
}

func (r *paymentDueRepository) GetByID(id uuid.UUID) (*models.PaymentDue, error) {
	var paymentDue models.PaymentDue
	err := r.db.Where("id = ?", id).First(&paymentDue).Error
	return &paymentDue, err
}

func (r *paymentDueRepository) Update(paymentDue *models.PaymentDue) error {
	return r.db.Save(paymentDue).Error
}

func (r *paymentDueRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.PaymentDue{}, id).Error
}

func (r *paymentDueRepository) ListAll() ([]*models.PaymentDue, error) {
	var paymentDues []*models.PaymentDue
	err := r.db.Find(&paymentDues).Error
	return paymentDues, err
}

func (r *paymentDueRepository) SearchAndSort(query string, sortBy string, sortDirection string) ([]*models.PaymentDue, error) {
	var paymentDues []*models.PaymentDue
	query = "%" + query + "%"
	sortDirection = normalizeSortDirection(sortDirection)

	err := r.db.Where("invoice_id LIKE ? OR status LIKE ?", query, query).
		Order(sortBy + " " + sortDirection).
		Find(&paymentDues).Error
	return paymentDues, err
}

func (r *paymentDueRepository) ListOverduePayments() ([]*models.PaymentDue, error) {
	var paymentDues []*models.PaymentDue
	err := r.db.Where("due_date < ? AND status = ?", time.Now(), "unpaid").Find(&paymentDues).Error
	return paymentDues, err
}

func normalizeSortDirection(direction string) string {
	if direction != "asc" && direction != "desc" {
		return "asc"
	}
	return direction
}
