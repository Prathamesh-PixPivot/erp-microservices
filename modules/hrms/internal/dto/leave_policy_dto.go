package dto

import "hrms/internal/domain"

// LeavePolicyDTO represents the data transfer object for leave policy operations
type LeavePolicyDTO struct {
	OrganizationID uint             `json:"organization_id"`
	LeaveType      domain.LeaveType `json:"leave_type"`
	MaxDays        int              `json:"max_days"`
	CarryForward   bool             `json:"carry_forward"`
}

// CreateLeavePolicyRequest represents the request payload for creating a leave policy
type CreateLeavePolicyRequest struct {
	OrganizationID uint             `json:"organization_id" validate:"required"`
	LeaveType      domain.LeaveType `json:"leave_type" validate:"required"`
	MaxDays        int              `json:"max_days" validate:"required,gte=0"`
	CarryForward   bool             `json:"carry_forward"`
}

// UpdateLeavePolicyRequest represents the request payload for updating a leave policy
type UpdateLeavePolicyRequest struct {
	MaxDays      *int  `json:"max_days,omitempty"`
	CarryForward *bool `json:"carry_forward,omitempty"`
}

// LeavePolicyResponse represents the response structure for leave policy queries
type LeavePolicyResponse struct {
	ID             uint             `json:"id"`
	OrganizationID uint             `json:"organization_id"`
	LeaveType      domain.LeaveType `json:"leave_type"`
	MaxDays        int              `json:"max_days"`
	CarryForward   bool             `json:"carry_forward"`
}
