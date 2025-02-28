package service

import (
	"pantry-service/internal/models"
	"pantry-service/internal/repository"
)

// PantryService defines the interface for the service layer.
type PantryService interface {
	// CRUD for Pantry Items
	CreatePantryItem(item *models.PantryItem) (*models.PantryItem, error)
	GetPantryItem(productID string) (*models.PantryItem, error)
	UpdatePantryItem(item *models.PantryItem) (*models.PantryItem, error)
	DeletePantryItem(productID string) error
	ListPantryItems(limit int, offset int) ([]models.PantryItem, error)

	// CRUD for Pantry bucket_items
	CreateBucketItem(item *models.BucketItem) (*models.BucketItem, error)
	GetBucketItem(productID string) (*models.BucketItem, error)
	UpdateBucketItem(item *models.BucketItem) (*models.BucketItem, error)
	DeleteBucketItem(productID string) error
	ListBucketItems(limit int, offset int) ([]models.BucketItem, error)

	// CRUD for Pantry expense logs
	CreateExpenseLog(item *models.ExpenseLog) (*models.ExpenseLog, error)
	GetExpenseLog(productID string) (*models.ExpenseLog, error)
	UpdateExpenseLog(item *models.ExpenseLog) (*models.ExpenseLog, error)
	DeleteExpenseLog(productID string) error
	ListExpenseLogs(limit int, offset int) ([]models.ExpenseLog, error)
}

// pantryService is the concrete implementation of PantryService.
type pantryService struct {
	repo repository.PantryRepository
}

// NewPantryService creates a new PantryService.
func NewPantryService(repo repository.PantryRepository) PantryService {
	return &pantryService{repo: repo}
}

// CRUD Methods for pantry_items
func (s *pantryService) CreatePantryItem(item *models.PantryItem) (*models.PantryItem, error) {
	return s.repo.CreatePantryItem(item)
}

func (s *pantryService) GetPantryItem(productID string) (*models.PantryItem, error) {
	return s.repo.GetPantryItem(productID)
}

func (s *pantryService) UpdatePantryItem(item *models.PantryItem) (*models.PantryItem, error) {
	return s.repo.UpdatePantryItem(item)
}

func (s *pantryService) DeletePantryItem(productID string) error {
	return s.repo.DeletePantryItem(productID)
}

func (s *pantryService) ListPantryItems(limit int, offset int) ([]models.PantryItem, error) {
	return s.repo.ListPantryItems(limit, offset)
}

// CRUD Methods for pantry bucket_items
func (s *pantryService) CreateBucketItem(item *models.BucketItem) (*models.BucketItem, error) {
	return s.repo.CreateBucketItem(item)
}

func (s *pantryService) GetBucketItem(productID string) (*models.BucketItem, error) {
	return s.repo.GetBucketItem(productID)
}

func (s *pantryService) UpdateBucketItem(item *models.BucketItem) (*models.BucketItem, error) {
	return s.repo.UpdateBucketItem(item)
}

func (s *pantryService) DeleteBucketItem(productID string) error {
	return s.repo.DeleteBucketItem(productID)
}

func (s *pantryService) ListBucketItems(limit int, offset int) ([]models.BucketItem, error) {
	return s.repo.ListBucketItems(limit, offset)
}

// CRUD Methods for pantry expense_logs
func (s *pantryService) CreateExpenseLog(item *models.ExpenseLog) (*models.ExpenseLog, error) {
	return s.repo.CreateExpenseLog(item)
}

func (s *pantryService) GetExpenseLog(productID string) (*models.ExpenseLog, error) {
	return s.repo.GetExpenseLog(productID)
}

func (s *pantryService) UpdateExpenseLog(item *models.ExpenseLog) (*models.ExpenseLog, error) {
	return s.repo.UpdateExpenseLog(item)
}

func (s *pantryService) DeleteExpenseLog(productID string) error {
	return s.repo.DeleteExpenseLog(productID)
}

func (s *pantryService) ListExpenseLogs(limit int, offset int) ([]models.ExpenseLog, error) {
	return s.repo.ListExpenseLogs(limit, offset)
}
