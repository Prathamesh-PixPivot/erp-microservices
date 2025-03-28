package models

import (
	"time"
)

type ExpenseLog struct {
	Id          uint      `gorm:"primaryKey"`
	AmtReceived float64   `gorm:"not null"`
	AmtSpend    float64   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time
}
