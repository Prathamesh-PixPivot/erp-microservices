package dto

import "time"

// CreateEmployeeDocumentRequest represents the request to add a document
type CreateEmployeeDocumentRequest struct {
	EmployeeID   uint       `json:"employee_id" validate:"required"`
	DocumentName string     `json:"document_name" validate:"required"`
	DocumentURL  string     `json:"document_url" validate:"required"`
	ExpiryDate   *time.Time `json:"expiry_date,omitempty"`
}

// UpdateEmployeeDocumentRequest represents the request to update a document
type UpdateEmployeeDocumentRequest struct {
	DocumentName string     `json:"document_name,omitempty"`
	DocumentURL  string     `json:"document_url,omitempty"`
	ExpiryDate   *time.Time `json:"expiry_date,omitempty"`
}

// EmployeeDocumentDTO represents the response structure for employee documents
type EmployeeDocumentDTO struct {
	ID           uint       `json:"id"`
	EmployeeID   uint       `json:"employee_id"`
	DocumentName string     `json:"document_name"`
	DocumentURL  string     `json:"document_url"`
	ExpiryDate   *time.Time `json:"expiry_date,omitempty"`
}
