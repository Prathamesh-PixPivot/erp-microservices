package repository

import (
	"sync"
	"time"

	"amaa/internal/models"
)

type DisposalRepository interface {
	DecommissionAsset(disposal *models.Disposal) (*models.Disposal, error)
}

type inMemoryDisposalRepo struct {
	mu        sync.RWMutex
	disposals map[string]*models.Disposal // keyed by AssetID
}

func NewDisposalRepository() DisposalRepository {
	return &inMemoryDisposalRepo{
		disposals: make(map[string]*models.Disposal),
	}
}

func (repo *inMemoryDisposalRepo) DecommissionAsset(disposal *models.Disposal) (*models.Disposal, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	disposal.DecommissionDate = time.Now()
	repo.disposals[disposal.AssetID] = disposal
	return disposal, nil
}
