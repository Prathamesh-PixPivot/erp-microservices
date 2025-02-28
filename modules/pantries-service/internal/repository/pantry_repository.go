package repository

import (
	"errors"
	"pantry-service/internal/models"

	"gorm.io/gorm"
)

type PantryRepository interface {
	// CRUD  for pantry Items
	CreatePantryItem(item *models.PantryItem) (*models.PantryItem, error)
	GetPantryItem(ProductId string) (*models.PantryItem, error)
	ListPantryItems(limit int, offset int) ([]models.PantryItem, error)
	UpdatePantryItem(item *models.PantryItem) (*models.PantryItem, error)
	DeletePantryItem(ProductId string) error

	// CRUD for pantry bucket_items
	CreateBucketItem(item *models.BucketItem) (*models.BucketItem, error)
	GetBucketItem(ProductId string) (*models.BucketItem, error)
	ListBucketItems(limit int, offset int) ([]models.BucketItem, error)
	UpdateBucketItem(item *models.BucketItem) (*models.BucketItem, error)
	DeleteBucketItem(ProductId string) error

	// CRUD for pantry expense_logs
	CreateExpenseLog(item *models.ExpenseLog) (*models.ExpenseLog, error)
	GetExpenseLog(ProductId string) (*models.ExpenseLog, error)
	ListExpenseLogs(limit int, offset int) ([]models.ExpenseLog, error)
	UpdateExpenseLog(item *models.ExpenseLog) (*models.ExpenseLog, error)
	DeleteExpenseLog(ProductId string) error
}

type pantryRepository struct {
	db *gorm.DB
}

// NewPantryRepository creates a new PantryRepository
func NewPantryRepository(db *gorm.DB) PantryRepository {
	return &pantryRepository{
		db: db,
	}
}

// CRUD Methods for pantry items
func (r *pantryRepository) CreatePantryItem(item *models.PantryItem) (*models.PantryItem, error) {
	result := r.db.Create(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

func (r *pantryRepository) GetPantryItem(ProductId string) (*models.PantryItem, error) {
	var item models.PantryItem
	result := r.db.Where("id = ?", ProductId).First(&item)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &item, nil
}

func (r *pantryRepository) ListPantryItems(limit int, offset int) ([]models.PantryItem, error) {
	var items []models.PantryItem
	result := r.db.Limit(limit).Offset(offset).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}
func (r *pantryRepository) UpdatePantryItem(item *models.PantryItem) (*models.PantryItem, error) {
	result := r.db.Save(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

func (r *pantryRepository) DeletePantryItem(ProductId string) error {
	result := r.db.Where("id = ?", ProductId).Delete(&models.PantryItem{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// CRUD Methods for pantry bucket_items
func (r *pantryRepository) CreateBucketItem(item *models.BucketItem) (*models.BucketItem, error) {
	// Start a new transaction
	tx := r.db.Begin()

	// Create the bucket item
	result := tx.Create(item)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	// Fetch the related pantry item based on some identifier (e.g., ProductId)
	var pantryItem models.PantryItem
	err := tx.Where("id = ?", item.ProductId).First(&pantryItem).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Compare the prices and update the pantry item if the price has changed
	if pantryItem.Price != item.Price {
		pantryItem.Price = item.Price
		err = tx.Save(&pantryItem).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Commit the transaction
	tx.Commit()

	return item, nil
}

func (r *pantryRepository) GetBucketItem(itemID string) (*models.BucketItem, error) {
	var item models.BucketItem
	result := r.db.Where("id = ?", itemID).First(&item)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &item, nil
}

func (r *pantryRepository) ListBucketItems(limit int, offset int) ([]models.BucketItem, error) {
	var items []models.BucketItem
	result := r.db.Limit(limit).Offset(offset).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}
func (r *pantryRepository) UpdateBucketItem(item *models.BucketItem) (*models.BucketItem, error) {
	// Start a new transaction
	tx := r.db.Begin()

	// Update the bucket item
	result := tx.Save(item)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	// Fetch the related pantry item based on some identifier (e.g., ProductId)
	var pantryItem models.PantryItem
	err := tx.Where("id = ?", item.ProductId).First(&pantryItem).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Compare the prices and update the pantry item if the price has changed
	if pantryItem.Price != item.Price {
		pantryItem.Price = item.Price
		err = tx.Save(&pantryItem).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Commit the transaction
	tx.Commit()

	return item, nil
}

func (r *pantryRepository) DeleteBucketItem(itemID string) error {
	result := r.db.Where("id = ?", itemID).Delete(&models.BucketItem{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// CRUD Methods for pantry expense_logs
func (r *pantryRepository) CreateExpenseLog(log *models.ExpenseLog) (*models.ExpenseLog, error) {
	result := r.db.Create(log)
	if result.Error != nil {
		return nil, result.Error
	}
	return log, nil
}

func (r *pantryRepository) GetExpenseLog(logID string) (*models.ExpenseLog, error) {
	var log models.ExpenseLog
	result := r.db.Where("id = ?", logID).First(&log)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &log, nil
}

func (r *pantryRepository) ListExpenseLogs(limit int, offset int) ([]models.ExpenseLog, error) {
	var logs []models.ExpenseLog
	result := r.db.Limit(limit).Offset(offset).Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}
	return logs, nil
}
func (r *pantryRepository) UpdateExpenseLog(log *models.ExpenseLog) (*models.ExpenseLog, error) {
	result := r.db.Save(log)
	if result.Error != nil {
		return nil, result.Error
	}
	return log, nil
}

func (r *pantryRepository) DeleteExpenseLog(logID string) error {
	result := r.db.Where("id = ?", logID).Delete(&models.ExpenseLog{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
