package repository

import (
    "vms-service/internal/models"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type VendorRepository struct {
    db *gorm.DB
}

func NewVendorRepository(db *gorm.DB) *VendorRepository {
    return &VendorRepository{db: db}
}

func (repo *VendorRepository) CreateVendor(vendor *models.Vendor) error {
    vendor.ID = uuid.New()
    return repo.db.Create(vendor).Error
}

func (repo *VendorRepository) GetVendorByID(id uuid.UUID) (*models.Vendor, error) {
    var vendor models.Vendor
    if err := repo.db.First(&vendor, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &vendor, nil
}

func (repo *VendorRepository) UpdateVendor(vendor *models.Vendor) error {
    return repo.db.Save(vendor).Error
}

func (repo *VendorRepository) DeleteVendor(id uuid.UUID) error {
    return repo.db.Delete(&models.Vendor{}, "id = ?", id).Error
}

func (repo *VendorRepository) SearchVendors(query string) ([]models.Vendor, error) {
    var vendors []models.Vendor
    q := repo.db.Where("name ILIKE ? OR category ILIKE ?", "%"+query+"%", "%"+query+"%")
    if err := q.Find(&vendors).Error; err != nil {
        return nil, err
    }
    return vendors, nil
}
