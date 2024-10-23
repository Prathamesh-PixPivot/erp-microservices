package models

import (
    "time"
    "github.com/google/uuid"
)

type Payment struct {
    ID             uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
    PurchaseOrderID uuid.UUID `gorm:"type:uuid;index" json:"purchase_order_id"`
    PurchaseOrder  PurchaseOrder `gorm:"foreignKey:PurchaseOrderID"`
    Amount         float64    `json:"amount"`
    Status         string     `json:"status"`
    PaymentTerms   string     `json:"payment_terms"`
    PaidAt         time.Time  `json:"paid_at"`
}
