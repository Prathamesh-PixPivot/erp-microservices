package models

import (
	"time"
)

type PantryItem struct {
	Id          uint    `gorm:"primaryKey"` //foreign key for items_bucket
	ProductName string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Category    string  `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time
}
