package models

import (
	"time"
)

// OrderItem represents an item within an order.
type OrderItem struct {
	ID        uint   `gorm:"primaryKey"`
	OrderID   string `gorm:"not null"` // Links to the Order
	ProductID string `gorm:"not null"` // Links to the InventoryItem
	Quantity  int32  `gorm:"not null"` // Quantity ordered
	CreatedAt time.Time
	UpdatedAt time.Time
}
