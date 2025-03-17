package models

import (
	"time"
)

// InventoryItem represents an item in the inventory.
type InventoryItem struct {
	ID                 uint   `gorm:"primaryKey"`
	ProductID          string `gorm:"unique;not null"` // Unique product ID
	ProductName        string `gorm:"not null"`
	ProductDescription string `gorm:"type:text"`
	SKU                string `gorm:"not null;unique"`
	SupplierID         string `gorm:"not null"` // Foreign key to the Vendor Management Service
	Category           string
	Price              float64          `gorm:"not null"`
	AvailableQuantity  int32            `gorm:"not null"`
	ReorderPoint       int32            `gorm:"not null"`
	WarehouseStocks    []WarehouseStock `gorm:"foreignKey:InventoryItemID"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
