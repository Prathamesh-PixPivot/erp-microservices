package services

import (
    "vms-service/internal/models"
    "vms-service/internal/repository"
    "github.com/google/uuid"
)

type PerformanceService struct {
    repo *repository.PerformanceRepository
}

func NewPerformanceService(repo *repository.PerformanceRepository) *PerformanceService {
    return &PerformanceService{repo: repo}
}

func (service *PerformanceService) RecordPerformance(performance *models.VendorPerformance) error {
    return service.repo.RecordPerformance(performance)
}

func (service *PerformanceService) GetPerformanceByID(id uuid.UUID) (*models.VendorPerformance, error) {
    return service.repo.GetPerformanceByID(id)
}

func (service *PerformanceService) UpdatePerformance(performance *models.VendorPerformance) error {
    return service.repo.UpdatePerformance(performance)
}

func (service *PerformanceService) DeletePerformance(id uuid.UUID) error {
    return service.repo.DeletePerformance(id)
}
