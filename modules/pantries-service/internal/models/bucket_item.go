package models

import (
	"time"
)

type BucketItem struct {
	ID        uint      `gorm:"primaryKey"`
	ProductId string    `gorm:"foreignKey:PantryItem"`
	Qty       float64   `gorm:"not null"`
	Price     float64   `gorm:"not null"`
	Total     float64   `gorm:"not null"`
	PaidBy    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}
