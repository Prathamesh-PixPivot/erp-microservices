package dto

import "time"

// CreateOrganizationDTO 📌 Request DTO for creating an organization
type CreateOrganizationDTO struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}

// UpdateOrganizationDTO 📌 Request DTO for updating an organization
type UpdateOrganizationDTO struct {
	Name    *string `json:"name,omitempty"`
	Address *string `json:"address,omitempty"`
	Phone   *string `json:"phone,omitempty"`
}

// OrganizationResponseDTO 📌 Response DTO for an organization
type OrganizationResponseDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PaginatedOrganizationsResponse 📌 Response DTO for paginated organizations
type PaginatedOrganizationsResponse struct {
	Total        int64                     `json:"total"`
	Limit        int                        `json:"limit"`
	Offset       int                        `json:"offset"`
	Organizations []OrganizationResponseDTO `json:"organizations"`
}
