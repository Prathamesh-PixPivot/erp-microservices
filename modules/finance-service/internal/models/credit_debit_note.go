package models

import (
	"time"

	"github.com/google/uuid"
)

type CreditDebitNote struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	InvoiceID uuid.UUID `json:"invoice_id"`
	Type      string    `json:"type"`
	Reason    string    `json:"reason"`
	Amount    float64   `json:"amount"`
	NoteDate  time.Time `json:"note_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
