package dto

// CreateDepartmentRequest represents the request for creating a department
type CreateDepartmentRequest struct {
	Name           string `json:"name" validate:"required"`
	OrganizationID uint   `json:"organization_id" validate:"required"`
}

// UpdateDepartmentRequest represents the request for updating a department
type UpdateDepartmentRequest struct {
	Name string `json:"name,omitempty"`
}

// DepartmentDTO represents the response structure for a department
type DepartmentDTO struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	OrganizationID uint   `json:"organization_id"`
}
