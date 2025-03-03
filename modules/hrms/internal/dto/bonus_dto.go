package dto

import "time"

// CreateBonusRequest represents the request for adding a bonus
type CreateBonusRequest struct {
	EmployeeID  uint      `json:"employee_id" validate:"required"`
	Amount      float64   `json:"amount" validate:"required,gt=0"`
	BonusType   string    `json:"bonus_type" validate:"required,oneof=Performance Festival Other"`
	Description string    `json:"description"`
	ApprovedBy  uint      `json:"approved_by,omitempty"`
	ApprovalDate *time.Time `json:"approval_date,omitempty"`
	IssueDate   time.Time `json:"issue_date" validate:"required"`
	Status      string    `json:"status" validate:"required,oneof=Pending Approved Rejected"`
}

// UpdateBonusRequest represents the request for updating a bonus
type UpdateBonusRequest struct {
	Status       string     `json:"status,omitempty" validate:"omitempty,oneof=Pending Approved Rejected"`
	ApprovalDate *time.Time `json:"approval_date,omitempty"`
	Description  string     `json:"description,omitempty"`
}

// BonusDTO represents the response structure for a bonus record
type BonusDTO struct {
	ID          uint       `json:"id"`
	EmployeeID  uint       `json:"employee_id"`
	Amount      float64    `json:"amount"`
	BonusType   string     `json:"bonus_type"`
	Description string     `json:"description"`
	ApprovedBy  uint       `json:"approved_by"`
	ApprovalDate *time.Time `json:"approval_date,omitempty"`
	IssueDate   time.Time  `json:"issue_date"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
