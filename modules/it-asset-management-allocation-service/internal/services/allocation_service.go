package services

import (
	"amaa/internal/models"
	"amaa/internal/repository"
	"time"

	"github.com/google/uuid"
)

type AllocationService interface {
	AllocateAsset(allocation *models.Allocation) (*models.Allocation, error)
	GetAllocationByAssetID(assetID string) (*models.Allocation, error)
	ReallocateAsset(allocationID, newAssignedTo string) error
	DeallocateAsset(allocationID string) error
}

type allocationService struct {
	repo repository.AllocationRepository
}

func NewAllocationService(repo repository.AllocationRepository) AllocationService {
	return &allocationService{repo: repo}
}

func (s *allocationService) AllocateAsset(allocation *models.Allocation) (*models.Allocation, error) {
	allocation.ID = uuid.New().String()
	allocation.AssignmentDate = time.Now()
	return s.repo.Create(allocation)
}

func (s *allocationService) GetAllocationByAssetID(assetID string) (*models.Allocation, error) {
	return s.repo.GetByAssetID(assetID)
}

func (s *allocationService) ReallocateAsset(allocationID, newAssignedTo string) error {
	// For demo purposes, assume allocationID is the assetID for lookup.
	alloc, err := s.repo.GetByAssetID(allocationID)
	if err != nil {
		return err
	}
	alloc.AssignedTo = newAssignedTo
	alloc.AssignmentDate = time.Now()
	return s.repo.Update(alloc)
}

func (s *allocationService) DeallocateAsset(allocationID string) error {
	alloc, err := s.repo.GetByAssetID(allocationID)
	if err != nil {
		return err
	}
	now := time.Now()
	alloc.ReleaseDate = &now
	return s.repo.Update(alloc)
}
