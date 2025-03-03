package dto

import (
	"time"
)

// LeaveDTO represents the data transfer object for leave requests
type LeaveDTO struct {
	EmployeeID     uint      `json:"employee_id"`
	LeaveType      string    `json:"leave_type"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	Status         string    `json:"status"`
	ApproverID     uint      `json:"approver_id"`
	MultiLevelStep int       `json:"multi_level_step"`
	Comments       string    `json:"comments,omitempty"`
}

// CreateLeaveRequest represents the request payload for creating a leave request
type CreateLeaveRequest struct {
	EmployeeID uint      `json:"employee_id" validate:"required"`
	LeaveType  string    `json:"leave_type" validate:"required"`
	StartDate  time.Time `json:"start_date" validate:"required"`
	EndDate    time.Time `json:"end_date" validate:"required"`
}

// UpdateLeaveStatusRequest represents the request payload for updating leave status
type UpdateLeaveStatusRequest struct {
	ApproverID uint   `json:"approver_id" validate:"required"`
	Status     string `json:"status" validate:"required,oneof=Pending Approved Rejected"`
	Comments   string `json:"comments,omitempty"`
}

// LeaveResponse represents the response structure for leave requests
type LeaveResponse struct {
	ID             uint      `json:"id"`
	EmployeeID     uint      `json:"employee_id"`
	LeaveType      string    `json:"leave_type"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	Status         string    `json:"status"`
	ApproverID     uint      `json:"approver_id"`
	MultiLevelStep int       `json:"multi_level_step"`
	Comments       string    `json:"comments,omitempty"`
}
