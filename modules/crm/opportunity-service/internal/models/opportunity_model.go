package models

import (
	"time"
)

type Opportunity struct {
    Id          uint      `gorm:"primaryKey" json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Stage       string    `json:"stage"`
    Amount      float64   `json:"amount"`
    CloseDate   time.Time `json:"close_date"`
    Probability float64   `json:"probability"`
    LeadID      uint      `json:"lead_id"`
    AccountID   uint      `json:"account_id"`
    OwnerID     uint      `json:"owner_id"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
