package service

import (
	"inventory-service/internal/models"
	"inventory-service/internal/repository"
)

// InventoryService defines the interface for the service layer.
type InventoryService interface {
	// CRUD for Inventory Items
	CreateInventoryItem(item *models.InventoryItem) (*models.InventoryItem, error)
	GetInventoryItem(productID string) (*models.InventoryItem, error)
	UpdateInventoryItem(item *models.InventoryItem) (*models.InventoryItem, error)
	DeleteInventoryItem(productID string) error
	ListInventoryItems(limit int, offset int) ([]models.InventoryItem, error)

	// Stock Management
	TrackInventory(productID string) (*models.InventoryItem, error)
	SetReorderPoint(productID string, reorderPoint int32) error
	ManageWarehouses(warehouses []models.Warehouse) error
	AddOrUpdateWarehouseStock(stock *models.WarehouseStock) error
	AddOrUpdateInventoryItem(item *models.InventoryItem) (*models.InventoryItem, error)

	// Order Fulfillment
	ProcessOrder(orderItems []models.OrderItem) error
	GeneratePickingList(orderID string) ([]models.PickingItem, error)
	UpdateInventoryStock(productID string, quantity int32, warehouseID uint) error

	// Vendor and Finance Integrations
	PlaceVendorOrder(vendorID string, orderItems []models.OrderItem) error
	NotifyFinanceForOrder(orderID string, totalAmount float64) error
}

// inventoryService is the concrete implementation of InventoryService.
type inventoryService struct {
	repo repository.InventoryRepository
}

// NewInventoryService creates a new InventoryService.
func NewInventoryService(repo repository.InventoryRepository) InventoryService {
	return &inventoryService{repo: repo}
}

// CRUD for Inventory Items

func (s *inventoryService) CreateInventoryItem(item *models.InventoryItem) (*models.InventoryItem, error) {
	return s.repo.CreateInventoryItem(item)
}

func (s *inventoryService) GetInventoryItem(productID string) (*models.InventoryItem, error) {
	return s.repo.GetInventoryItem(productID)
}

func (s *inventoryService) UpdateInventoryItem(item *models.InventoryItem) (*models.InventoryItem, error) {
	return s.repo.UpdateInventoryItem(item)
}

func (s *inventoryService) DeleteInventoryItem(productID string) error {
	return s.repo.DeleteInventoryItem(productID)
}

func (s *inventoryService) ListInventoryItems(limit int, offset int) ([]models.InventoryItem, error) {
	return s.repo.ListInventoryItems(limit, offset)
}

// AddOrUpdateInventoryItem handles both creation and updating of an inventory item
func (s *inventoryService) AddOrUpdateInventoryItem(item *models.InventoryItem) (*models.InventoryItem, error) {
	existingItem, err := s.GetInventoryItem(item.ProductID)
	if err != nil {
		return nil, err
	}

	if existingItem == nil {
		// Item doesn't exist, create a new one
		return s.CreateInventoryItem(item)
	}

	// Item exists, update it
	item.ID = existingItem.ID // Ensure we use the same ID for the update
	return s.UpdateInventoryItem(item)
}

// Stock Management

func (s *inventoryService) TrackInventory(productID string) (*models.InventoryItem, error) {
	return s.repo.TrackInventory(productID)
}

func (s *inventoryService) SetReorderPoint(productID string, reorderPoint int32) error {
	return s.repo.SetReorderPoint(productID, reorderPoint)
}

func (s *inventoryService) ManageWarehouses(warehouses []models.Warehouse) error {
	return s.repo.ManageWarehouses(warehouses)
}

func (s *inventoryService) AddOrUpdateWarehouseStock(stock *models.WarehouseStock) error {
	return s.repo.AddOrUpdateWarehouseStock(stock)
}

// Order Fulfillment

func (s *inventoryService) ProcessOrder(orderItems []models.OrderItem) error {
	return s.repo.ProcessOrder(orderItems)
}

func (s *inventoryService) GeneratePickingList(orderID string) ([]models.PickingItem, error) {
	return s.repo.GeneratePickingList(orderID)
}

func (s *inventoryService) UpdateInventoryStock(productID string, quantity int32, warehouseID uint) error {
	return s.repo.UpdateInventoryStock(productID, quantity, warehouseID)
}

// Vendor and Finance Integrations

func (s *inventoryService) PlaceVendorOrder(vendorID string, orderItems []models.OrderItem) error {
	return s.repo.PlaceVendorOrder(vendorID, orderItems)
}

func (s *inventoryService) NotifyFinanceForOrder(orderID string, totalAmount float64) error {
	return s.repo.NotifyFinanceForOrder(orderID, totalAmount)
}
