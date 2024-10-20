package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Password       string `json:"-"` // Not included in the JSON response
	Phone          string `json:"phone"`
	Role           string `json:"role"`            // e.g., admin, user
	OrganizationID uint   `json:"organization_id"` // ForeignKey to Organization
	KeycloakID     string `json:"keycloak_id"`
}
