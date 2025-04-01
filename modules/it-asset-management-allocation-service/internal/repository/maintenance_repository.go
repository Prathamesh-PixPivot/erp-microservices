package repository

import (
	"errors"
	"sync"
	"time"

	"amaa/internal/models"
)

type MaintenanceRepository interface {
	CreateMaintenance(record *models.Maintenance) (*models.Maintenance, error)
	GetRecordsByAssetID(assetID string) ([]*models.Maintenance, error)
}

type inMemoryMaintenanceRepo struct {
	mu      sync.RWMutex
	records map[string][]*models.Maintenance // keyed by AssetID
}

func NewMaintenanceRepository() MaintenanceRepository {
	return &inMemoryMaintenanceRepo{
		records: make(map[string][]*models.Maintenance),
	}
}

func (repo *inMemoryMaintenanceRepo) CreateMaintenance(record *models.Maintenance) (*models.Maintenance, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	record.MaintenanceDate = time.Now()
	repo.records[record.AssetID] = append(repo.records[record.AssetID], record)
	return record, nil
}

func (repo *inMemoryMaintenanceRepo) GetRecordsByAssetID(assetID string) ([]*models.Maintenance, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	records, exists := repo.records[assetID]
	if !exists {
		return nil, errors.New("no maintenance records found")
	}
	return records, nil
}
