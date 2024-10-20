package repository

import (
	"fmt"
	"organization-service/config"
	"organization-service/internal/models"
)

// CreateOrganization stores a new organization in the database
func CreateOrganization(org *models.Organization) error {
	if config.DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	return config.DB.Create(org).Error
}

// GetOrganizationByID fetches an organization by ID from the database
func GetOrganizationByID(id uint) (*models.Organization, error) {
	var org models.Organization
	err := config.DB.First(&org, id).Error
	return &org, err
}
