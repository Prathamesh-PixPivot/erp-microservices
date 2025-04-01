package services

import (
	"amaa/internal/models"
	"amaa/internal/repository"

	"github.com/google/uuid"
)

type AssetService interface {
	CreateAsset(asset *models.Asset) (*models.Asset, error)
	GetAsset(id string) (*models.Asset, error)
	UpdateAsset(asset *models.Asset) error
	DeleteAsset(id string) error
}

type assetService struct {
	repo repository.AssetRepository
}

func NewAssetService(repo repository.AssetRepository) AssetService {
	return &assetService{repo: repo}
}

func (s *assetService) CreateAsset(asset *models.Asset) (*models.Asset, error) {
	// Generate a new UUID if not provided.
	if asset.ID == "" {
		asset.ID = uuid.New().String()
	}
	// Set default status if not provided.
	if asset.Status == "" {
		asset.Status = "active"
	}
	return s.repo.CreateAsset(asset)
}

func (s *assetService) GetAsset(id string) (*models.Asset, error) {
	return s.repo.GetAsset(id)
}

func (s *assetService) UpdateAsset(asset *models.Asset) error {
	return s.repo.UpdateAsset(asset)
}

func (s *assetService) DeleteAsset(id string) error {
	return s.repo.DeleteAsset(id)
}
