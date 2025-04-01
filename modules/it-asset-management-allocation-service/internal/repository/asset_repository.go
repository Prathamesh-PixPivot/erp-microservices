package repository

import (
	"amaa/internal/models"

	"gorm.io/gorm"
)

type AssetRepository interface {
	CreateAsset(asset *models.Asset) (*models.Asset, error)
	GetAsset(id string) (*models.Asset, error)
	UpdateAsset(asset *models.Asset) error
	DeleteAsset(id string) error
}

type gormAssetRepo struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) AssetRepository {
	return &gormAssetRepo{db: db}
}

func (r *gormAssetRepo) CreateAsset(asset *models.Asset) (*models.Asset, error) {
	// Create a new asset record
	if err := r.db.Create(asset).Error; err != nil {
		return nil, err
	}
	return asset, nil
}

func (r *gormAssetRepo) GetAsset(id string) (*models.Asset, error) {
	var asset models.Asset
	if err := r.db.First(&asset, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &asset, nil
}

func (r *gormAssetRepo) UpdateAsset(asset *models.Asset) error {
	return r.db.Save(asset).Error
}

func (r *gormAssetRepo) DeleteAsset(id string) error {
	return r.db.Delete(&models.Asset{}, "id = ?", id).Error
}
