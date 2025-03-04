package dto

// CreateDesignationRequest represents the request for creating a designation
type CreateDesignationRequest struct {
	Title          string `json:"title" validate:"required"`
	Level          string `json:"level" validate:"required"`
	HierarchyLevel int    `json:"hierarchy_level" validate:"required"`
	DepartmentID   uint   `json:"department_id" validate:"required"`
}

// UpdateDesignationRequest represents the request for updating a designation
type UpdateDesignationRequest struct {
	Title          string `json:"title,omitempty"`
	Level          string `json:"level,omitempty"`
	HierarchyLevel int    `json:"hierarchy_level,omitempty"`
}

// DesignationDTO represents the response structure for a designation
type DesignationDTO struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Level          string `json:"level"`
	HierarchyLevel int    `json:"hierarchy_level"`
	DepartmentID   uint   `json:"department_id"`
}
