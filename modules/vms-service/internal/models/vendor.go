package models

import (
    "time"
    "github.com/google/uuid"
)

type Vendor struct {
    ID               uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
    Name             string     `json:"name"`
    Category         string     `json:"category"`
    Service          string     `json:"service"`
    Industry         string     `json:"industry"`
    GSTIN            string     `json:"gstin" gorm:"unique_index"`
    Certifications   string     `json:"certifications"`
    Licenses         string     `json:"licenses"`
    CreatedAt        time.Time  `json:"created_at"`
    UpdatedAt        time.Time  `json:"updated_at"`
    ExpiryDate       time.Time  `json:"expiry_date"`
    IsCompliant      bool       `json:"is_compliant"`
    PerformanceScore float64    `json:"performance_score"`
    ContractID       uuid.UUID  `json:"contract_id"`
    RiskAssessment   string     `json:"risk_assessment"`
}
