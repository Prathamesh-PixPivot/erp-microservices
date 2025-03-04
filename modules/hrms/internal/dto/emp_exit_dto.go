package dto

import "time"

// CreateEmployeeExitRequest represents the request to add an exit record
type CreateEmployeeExitRequest struct {
	EmployeeID      uint      `json:"employee_id" validate:"required"`
	ExitType        string    `json:"exit_type" validate:"required,oneof=Resignation Termination Retirement"`
	ExitDate        time.Time `json:"exit_date" validate:"required"`
	ClearanceStatus string    `json:"clearance_status" validate:"required,oneof=Pending Completed"`
}

// UpdateClearanceStatusRequest represents the request to update clearance status
type UpdateClearanceStatusRequest struct {
	ClearanceStatus string `json:"clearance_status" validate:"required,oneof=Pending Completed"`
}

// EmployeeExitDTO represents the response structure for employee exit records
type EmployeeExitDTO struct {
	ID              uint      `json:"id"`
	EmployeeID      uint      `json:"employee_id"`
	ExitType        string    `json:"exit_type"`
	ExitDate        time.Time `json:"exit_date"`
	ClearanceStatus string    `json:"clearance_status"`
}
