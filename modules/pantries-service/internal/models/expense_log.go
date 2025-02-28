package models

import (
	"time"
)

type ExpenseLog struct {
	ID          uint      `gorm:"primaryKey"`
	AmtReceived float64   `gorm:"not null"`
	AmtSpend    float64   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoCreateTime"`
}
