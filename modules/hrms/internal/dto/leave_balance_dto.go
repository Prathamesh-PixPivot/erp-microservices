package dto

import "hrms/internal/domain"

// LeaveBalanceDTO represents the data transfer object for leave balance operations
type LeaveBalanceDTO struct {
	EmployeeID  uint             `json:"employee_id"`
	LeaveType   domain.LeaveType `json:"leave_type"`
	TotalLeaves float64          `json:"total_leaves"`
	UsedLeaves  float64          `json:"used_leaves"`
	Remaining   float64          `json:"remaining"`
}

// CreateLeaveBalanceRequest represents the request payload for creating leave balance
type CreateLeaveBalanceRequest struct {
	EmployeeID  uint             `json:"employee_id" validate:"required"`
	LeaveType   domain.LeaveType `json:"leave_type" validate:"required"`
	TotalLeaves float64          `json:"total_leaves" validate:"required,gte=0"`
}

// DeductLeaveBalanceRequest represents the request payload for deducting leave balance
type DeductLeaveBalanceRequest struct {
	EmployeeID uint             `json:"employee_id" validate:"required"`
	LeaveType  domain.LeaveType `json:"leave_type" validate:"required"`
	LeaveDays  float64          `json:"leave_days" validate:"required,gte=0"`
}

// RestoreLeaveBalanceRequest represents the request payload for restoring leave balance
type RestoreLeaveBalanceRequest struct {
	EmployeeID uint             `json:"employee_id" validate:"required"`
	LeaveType  domain.LeaveType `json:"leave_type" validate:"required"`
	LeaveDays  float64          `json:"leave_days" validate:"required,gte=0"`
}

// LeaveBalanceResponse represents the response structure for leave balance queries
type LeaveBalanceResponse struct {
	EmployeeID  uint             `json:"employee_id"`
	LeaveType   domain.LeaveType `json:"leave_type"`
	TotalLeaves float64          `json:"total_leaves"`
	UsedLeaves  float64          `json:"used_leaves"`
	Remaining   float64          `json:"remaining"`
}
