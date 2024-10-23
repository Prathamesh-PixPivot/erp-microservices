package repository

import (
    "vms-service/internal/models"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type PerformanceRepository struct {
    db *gorm.DB
}

func NewPerformanceRepository(db *gorm.DB) *PerformanceRepository {
    return &PerformanceRepository{db: db}
}

func (repo *PerformanceRepository) RecordPerformance(vp *models.VendorPerformance) error {
    vp.ID = uuid.New()
    return repo.db.Create(vp).Error
}

func (repo *PerformanceRepository) GetPerformanceByID(id uuid.UUID) (*models.VendorPerformance, error) {
    var vp models.VendorPerformance
    if err := repo.db.First(&vp, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &vp, nil
}

func (repo *PerformanceRepository) UpdatePerformance(vp *models.VendorPerformance) error {
    return repo.db.Save(vp).Error
}

func (repo *PerformanceRepository) DeletePerformance(id uuid.UUID) error {
    return repo.db.Delete(&models.VendorPerformance{}, "id = ?", id).Error
}
