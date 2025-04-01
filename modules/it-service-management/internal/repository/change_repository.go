// internal/repository/change_repository.go
package repository

import (
	"errors"
	"time"

	"itsm/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChangeRepository interface {
	CreateChange(change *models.Change) (*models.Change, error)
	GetChange(id string) (*models.Change, error)
	UpdateChange(change *models.Change) error
	DeleteChange(id string) error
	ListChanges() ([]*models.Change, error)
}

type gormChangeRepo struct {
	db *gorm.DB
}

func NewChangeRepository(db *gorm.DB) ChangeRepository {
	db.AutoMigrate(&models.Change{})
	return &gormChangeRepo{db: db}
}

func (r *gormChangeRepo) CreateChange(change *models.Change) (*models.Change, error) {
	change.ID = uuid.New().String()
	now := time.Now()
	change.CreatedAt = now
	change.UpdatedAt = now
	change.Status = "pending"
	if err := r.db.Create(change).Error; err != nil {
		return nil, err
	}
	return change, nil
}

func (r *gormChangeRepo) GetChange(id string) (*models.Change, error) {
	var change models.Change
	if err := r.db.First(&change, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("change not found")
		}
		return nil, err
	}
	return &change, nil
}

func (r *gormChangeRepo) UpdateChange(change *models.Change) error {
	change.UpdatedAt = time.Now()
	return r.db.Save(change).Error
}

func (r *gormChangeRepo) DeleteChange(id string) error {
	return r.db.Delete(&models.Change{}, "id = ?", id).Error
}

func (r *gormChangeRepo) ListChanges() ([]*models.Change, error) {
	var changes []*models.Change
	if err := r.db.Find(&changes).Error; err != nil {
		return nil, err
	}
	return changes, nil
}
