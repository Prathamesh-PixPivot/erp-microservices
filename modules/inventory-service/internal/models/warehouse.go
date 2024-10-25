package models

import "time"

// Warehouse represents a storage location for inventory.
type Warehouse struct {
	ID          uint   `gorm:"primaryKey"`
	WarehouseID string `gorm:"unique;not null"` // Unique identifier for the warehouse
	Name        string `gorm:"not null"`
	Location    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// WarehouseStock represents the stock level of an inventory item in a specific warehouse.
type WarehouseStock struct {
	ID              uint  `gorm:"primaryKey"`
	InventoryItemID uint  `gorm:"not null"` // Foreign key to InventoryItem
	WarehouseID     uint  `gorm:"not null"` // Foreign key to Warehouse
	StockLevel      int32 `gorm:"not null"` // The available stock level in this warehouse
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
