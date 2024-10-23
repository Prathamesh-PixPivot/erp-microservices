package services

import (
    "vms-service/internal/models"
    "vms-service/internal/repository"
    "github.com/google/uuid"
)

type VendorService struct {
    repo *repository.VendorRepository
}

func NewVendorService(repo *repository.VendorRepository) *VendorService {
    return &VendorService{repo: repo}
}

func (service *VendorService) CreateVendor(vendor *models.Vendor) error {
    return service.repo.CreateVendor(vendor)
}

func (service *VendorService) GetVendorByID(id uuid.UUID) (*models.Vendor, error) {
    return service.repo.GetVendorByID(id)
}

func (service *VendorService) UpdateVendor(vendor *models.Vendor) error {
    return service.repo.UpdateVendor(vendor)
}

func (service *VendorService) DeleteVendor(id uuid.UUID) error {
    return service.repo.DeleteVendor(id)
}

func (service *VendorService) SearchVendors(query string) ([]models.Vendor, error) {
    return service.repo.SearchVendors(query)
}
