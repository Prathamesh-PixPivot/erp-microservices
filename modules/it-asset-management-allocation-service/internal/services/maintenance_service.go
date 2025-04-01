package services

import (
	"amaa/internal/models"
	"amaa/internal/repository"
)

type MaintenanceService interface {
	ScheduleMaintenance(record *models.Maintenance) (*models.Maintenance, error)
	GetMaintenanceRecords(assetID string) ([]*models.Maintenance, error)
}

type maintenanceService struct {
	repo repository.MaintenanceRepository
}

func NewMaintenanceService(repo repository.MaintenanceRepository) MaintenanceService {
	return &maintenanceService{repo: repo}
}

func (s *maintenanceService) ScheduleMaintenance(record *models.Maintenance) (*models.Maintenance, error) {
	return s.repo.CreateMaintenance(record)
}

func (s *maintenanceService) GetMaintenanceRecords(assetID string) ([]*models.Maintenance, error) {
	return s.repo.GetRecordsByAssetID(assetID)
}
