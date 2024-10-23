package models

import (
	"time"

	"github.com/google/uuid"
)

type LedgerEntry struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	TransactionID   string    `json:"transaction_id"`
	Description     string    `json:"description"`
	Debit           float64   `json:"debit"`
	Credit          float64   `json:"credit"`
	Balance         float64   `json:"balance"`
	TransactionDate time.Time `json:"transaction_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
