package services

import (
	"amaa/internal/models"
	"amaa/internal/repository"
)

type LicenseService interface {
	RegisterLicense(license *models.License) (*models.License, error)
	GetLicense(licenseID string) (*models.License, error)
	UpdateLicense(license *models.License) (*models.License, error)
	DeleteLicense(licenseID string) error
}

type licenseService struct {
	repo repository.LicenseRepository
}

func NewLicenseService(repo repository.LicenseRepository) LicenseService {
	return &licenseService{repo: repo}
}

func (s *licenseService) RegisterLicense(license *models.License) (*models.License, error) {
	return s.repo.CreateLicense(license)
}

func (s *licenseService) GetLicense(licenseID string) (*models.License, error) {
	return s.repo.GetLicense(licenseID)
}

func (s *licenseService) UpdateLicense(license *models.License) (*models.License, error) {
	return s.repo.UpdateLicense(license)
}

func (s *licenseService) DeleteLicense(licenseID string) error {
	return s.repo.DeleteLicense(licenseID)
}
