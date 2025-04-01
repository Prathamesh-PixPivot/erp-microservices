package repository

import (
	"errors"
	"sync"
	"time"

	"amaa/internal/models"
)

type AllocationRepository interface {
	Create(allocation *models.Allocation) (*models.Allocation, error)
	GetByAssetID(assetID string) (*models.Allocation, error)
	Update(allocation *models.Allocation) error
	Delete(id string) error
}

type inMemoryAllocationRepo struct {
	mu          sync.RWMutex
	allocations map[string]*models.Allocation
}

func NewAllocationRepository() AllocationRepository {
	return &inMemoryAllocationRepo{
		allocations: make(map[string]*models.Allocation),
	}
}

func (repo *inMemoryAllocationRepo) Create(allocation *models.Allocation) (*models.Allocation, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// Set assignment date when creating the allocation
	allocation.AssignmentDate = time.Now()
	repo.allocations[allocation.ID] = allocation
	return allocation, nil
}

func (repo *inMemoryAllocationRepo) GetByAssetID(assetID string) (*models.Allocation, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	for _, alloc := range repo.allocations {
		// Return the active allocation (without a ReleaseDate)
		if alloc.AssetID == assetID && alloc.ReleaseDate == nil {
			return alloc, nil
		}
	}
	return nil, errors.New("allocation not found")
}

func (repo *inMemoryAllocationRepo) Update(allocation *models.Allocation) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	_, exists := repo.allocations[allocation.ID]
	if !exists {
		return errors.New("allocation not found")
	}
	repo.allocations[allocation.ID] = allocation
	return nil
}

func (repo *inMemoryAllocationRepo) Delete(id string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	_, exists := repo.allocations[id]
	if !exists {
		return errors.New("allocation not found")
	}
	delete(repo.allocations, id)
	return nil
}
