package models

import "time"

// PickingItem represents an item to be picked from a warehouse for order fulfillment.
type PickingItem struct {
	ID            uint   `gorm:"primaryKey"`
	OrderID       string `gorm:"not null"` // The order ID this picking item belongs to
	ProductID     string `gorm:"not null"` // The product ID that needs to be picked
	ProductName   string `gorm:"not null"` // The name of the product to be picked
	Quantity      int32  `gorm:"not null"` // The quantity of the product to be picked
	WarehouseID   uint   `gorm:"not null"` // The warehouse where the item is located
	WarehouseName string `gorm:"not null"` // The name of the warehouse
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
