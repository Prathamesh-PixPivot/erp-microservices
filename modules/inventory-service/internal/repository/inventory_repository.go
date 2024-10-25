package repository

import (
	"errors"
	"inventory-service/internal/models"
	"log"

	"gorm.io/gorm"
)

type InventoryRepository interface {
	// CRUD for Inventory Items
	CreateInventoryItem(item *models.InventoryItem) (*models.InventoryItem, error)
	GetInventoryItem(productID string) (*models.InventoryItem, error)
	UpdateInventoryItem(item *models.InventoryItem) (*models.InventoryItem, error)
	DeleteInventoryItem(productID string) error
	ListInventoryItems(limit int, offset int) ([]models.InventoryItem, error)

	// Stock Management Methods
	TrackInventory(productID string) (*models.InventoryItem, error)
	SetReorderPoint(productID string, reorderPoint int32) error
	ManageWarehouses(warehouses []models.Warehouse) error
	AddOrUpdateWarehouseStock(stock *models.WarehouseStock) error

	// Order Fulfillment Methods
	ProcessOrder(orderItems []models.OrderItem) error
	GeneratePickingList(orderID string) ([]models.PickingItem, error)
	UpdateInventoryStock(productID string, quantity int32, warehouseID uint) error

	// Vendor and Finance Integrations
	PlaceVendorOrder(vendorID string, orderItems []models.OrderItem) error
	NotifyFinanceForOrder(orderID string, totalAmount float64) error
}

type inventoryRepository struct {
	db *gorm.DB
}

// NewInventoryRepository creates a new InventoryRepository
func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{
		db: db,
	}
}

// CRUD for Inventory Items

func (r *inventoryRepository) CreateInventoryItem(item *models.InventoryItem) (*models.InventoryItem, error) {
	result := r.db.Create(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

func (r *inventoryRepository) GetInventoryItem(productID string) (*models.InventoryItem, error) {
	var item models.InventoryItem
	result := r.db.Preload("WarehouseStocks").Where("product_id = ?", productID).First(&item)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &item, nil
}

func (r *inventoryRepository) UpdateInventoryItem(item *models.InventoryItem) (*models.InventoryItem, error) {
	result := r.db.Save(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

func (r *inventoryRepository) DeleteInventoryItem(productID string) error {
	result := r.db.Where("product_id = ?", productID).Delete(&models.InventoryItem{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *inventoryRepository) ListInventoryItems(limit int, offset int) ([]models.InventoryItem, error) {
	var items []models.InventoryItem
	result := r.db.Limit(limit).Offset(offset).Preload("WarehouseStocks").Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// Stock Management Methods

func (r *inventoryRepository) TrackInventory(productID string) (*models.InventoryItem, error) {
	return r.GetInventoryItem(productID)
}

func (r *inventoryRepository) SetReorderPoint(productID string, reorderPoint int32) error {
	var item models.InventoryItem
	result := r.db.Where("product_id = ?", productID).First(&item)
	if result.Error != nil {
		return result.Error
	}
	item.ReorderPoint = reorderPoint
	return r.db.Save(&item).Error
}

func (r *inventoryRepository) ManageWarehouses(warehouses []models.Warehouse) error {
	for _, warehouse := range warehouses {
		if err := r.db.Save(&warehouse).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *inventoryRepository) AddOrUpdateWarehouseStock(stock *models.WarehouseStock) error {
	var existingStock models.WarehouseStock
	result := r.db.Where("inventory_item_id = ? AND warehouse_id = ?", stock.InventoryItemID, stock.WarehouseID).First(&existingStock)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Create new stock entry
		return r.db.Create(stock).Error
	} else if result.Error != nil {
		return result.Error
	}

	// Update existing stock level
	existingStock.StockLevel = stock.StockLevel
	return r.db.Save(&existingStock).Error
}

// Order Fulfillment Methods

func (r *inventoryRepository) ProcessOrder(orderItems []models.OrderItem) error {
	tx := r.db.Begin()

	for _, orderItem := range orderItems {
		err := r.UpdateInventoryStock(orderItem.ProductID, -orderItem.Quantity, 0) // Assume warehouse is handled separately
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *inventoryRepository) GeneratePickingList(orderID string) ([]models.PickingItem, error) {
	var pickingList []models.PickingItem
	// Here, you'd generate the picking list by retrieving the required items
	// and their stock levels from the corresponding warehouse(s).
	// This is a mock implementation.
	log.Printf("Generating picking list for order %s", orderID)
	// You should add real logic here to fetch and create a picking list based on the order.
	return pickingList, nil
}

func (r *inventoryRepository) UpdateInventoryStock(productID string, quantity int32, warehouseID uint) error {
	var stock models.WarehouseStock
	result := r.db.Where("inventory_item_id = ? AND warehouse_id = ?", productID, warehouseID).First(&stock)
	if result.Error != nil {
		return result.Error
	}
	stock.StockLevel += quantity
	return r.db.Save(&stock).Error
}

// Vendor and Finance Integrations

func (r *inventoryRepository) PlaceVendorOrder(vendorID string, orderItems []models.OrderItem) error {
	// This is a stub; you'd call the Vendor Management Service here.
	// For now, we'll just log it.
	for _, orderItem := range orderItems {
		log.Printf("Placing order with vendor %s for product %s, quantity %d", vendorID, orderItem.ProductID, orderItem.Quantity)
	}
	return nil
}

func (r *inventoryRepository) NotifyFinanceForOrder(orderID string, totalAmount float64) error {
	// This is a stub; you'd call the Finance Service here.
	log.Printf("Notifying finance for order %s, total amount: %f", orderID, totalAmount)
	return nil
}
