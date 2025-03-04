package dto

import "time"

// EmployeeDTO represents the employee response data
type EmployeeDTO struct {
	ID              uint      `json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	EmploymentType  string    `json:"employment_type"`
	Status          string    `json:"status"`
	HiredDate       *time.Time `json:"hired_date,omitempty"`
	OrganizationID  uint      `json:"organization_id"`
	DepartmentID    uint      `json:"department_id"`
	DesignationID   uint      `json:"designation_id"`
	ReportsTo       *uint     `json:"reports_to,omitempty"`
}

// CreateEmployeeRequest represents the request to create an employee
type CreateEmployeeRequest struct {
	FirstName      string     `json:"first_name" validate:"required"`
	LastName       string     `json:"last_name" validate:"required"`
	Email          string     `json:"email" validate:"required,email"`
	Phone          string     `json:"phone"`
	DateOfBirth    time.Time  `json:"date_of_birth"`
	EmploymentType string     `json:"employment_type" validate:"required"`
	HiredDate      *time.Time `json:"hired_date,omitempty"`
	OrganizationID uint       `json:"organization_id" validate:"required"`
	DepartmentID   uint       `json:"department_id" validate:"required"`
	DesignationID  uint       `json:"designation_id" validate:"required"`
	ReportsTo      *uint      `json:"reports_to,omitempty"`
}

// UpdateEmployeeRequest represents the request to update an employee profile
type UpdateEmployeeRequest struct {
	FirstName      *string `json:"first_name,omitempty"`
	LastName       *string `json:"last_name,omitempty"`
	Phone          *string `json:"phone,omitempty"`
	EmploymentType *string `json:"employment_type,omitempty"`
	Status         *string `json:"status,omitempty"`
}
