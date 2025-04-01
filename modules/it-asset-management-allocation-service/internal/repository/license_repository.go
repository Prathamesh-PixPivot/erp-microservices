package repository

import (
	"errors"
	"sync"
	"time"

	"amaa/internal/models"
)

type LicenseRepository interface {
	CreateLicense(license *models.License) (*models.License, error)
	GetLicense(licenseID string) (*models.License, error)
	UpdateLicense(license *models.License) (*models.License, error)
	DeleteLicense(licenseID string) error
}

type inMemoryLicenseRepo struct {
	mu       sync.RWMutex
	licenses map[string]*models.License
}

func NewLicenseRepository() LicenseRepository {
	return &inMemoryLicenseRepo{
		licenses: make(map[string]*models.License),
	}
}

func (repo *inMemoryLicenseRepo) CreateLicense(license *models.License) (*models.License, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// For demo, set expiry date to one year from now if not provided.
	if license.ExpiryDate.IsZero() {
		license.ExpiryDate = time.Now().AddDate(1, 0, 0)
	}
	repo.licenses[license.ID] = license
	return license, nil
}

func (repo *inMemoryLicenseRepo) GetLicense(licenseID string) (*models.License, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	license, exists := repo.licenses[licenseID]
	if !exists {
		return nil, errors.New("license not found")
	}
	return license, nil
}

func (repo *inMemoryLicenseRepo) UpdateLicense(license *models.License) (*models.License, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	_, exists := repo.licenses[license.ID]
	if !exists {
		return nil, errors.New("license not found")
	}
	repo.licenses[license.ID] = license
	return license, nil
}

func (repo *inMemoryLicenseRepo) DeleteLicense(licenseID string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	_, exists := repo.licenses[licenseID]
	if !exists {
		return errors.New("license not found")
	}
	delete(repo.licenses, licenseID)
	return nil
}
