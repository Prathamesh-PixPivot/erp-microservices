package models

import (
	"time"

	"github.com/google/uuid"
)

type PaymentDue struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	InvoiceID uuid.UUID  `json:"invoice_id"`
	DueDate   time.Time  `json:"due_date"`
	AmountDue float64    `json:"amount_due"`
	Status    string     `json:"status"` // e.g., 'paid', 'unpaid'
	PaidDate  *time.Time `json:"paid_date,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
