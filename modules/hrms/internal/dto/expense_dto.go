package dto

import "time"

// ExpenseRequestDTO represents the request to create an expense
type ExpenseRequestDTO struct {
	EmployeeID  uint      `json:"employee_id" binding:"required"`
	ExpenseType string    `json:"expense_type" binding:"required"`
	Amount      float64   `json:"amount" binding:"required,min=0.01"`
	Date        time.Time `json:"date"`
	ReceiptURL  string    `json:"receipt_url,omitempty"`
}

// ExpenseResponseDTO represents the response structure for an expense
type ExpenseResponseDTO struct {
	ID          uint      `json:"id"`
	EmployeeID  uint      `json:"employee_id"`
	ExpenseType string    `json:"expense_type"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	Status      string    `json:"status"`
	ApproverID  uint      `json:"approver_id,omitempty"`
	ReceiptURL  string    `json:"receipt_url,omitempty"`
}

// ExpenseStatusUpdateDTO is used to update the status of an expense
type ExpenseStatusUpdateDTO struct {
	ExpenseID  uint   `json:"expense_id" binding:"required"`
	ApproverID uint   `json:"approver_id" binding:"required"`
	Status     string `json:"status" binding:"required,oneof=Pending Approved Rejected"`
}
