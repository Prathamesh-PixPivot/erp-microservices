package models

import (
    "time"
    "github.com/google/uuid"
)

type VendorPerformance struct {
    ID         uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
    VendorID   uuid.UUID  `gorm:"type:uuid;index" json:"vendor_id"`
    Vendor     Vendor     `gorm:"foreignKey:VendorID"`
    Score      float64    `json:"score"`
    RiskLevel  string     `json:"risk_level"`
    EvaluatedAt time.Time `json:"evaluated_at"`
}
