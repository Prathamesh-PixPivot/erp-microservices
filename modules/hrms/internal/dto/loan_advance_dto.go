package dto

import "time"

// LoanAdvanceRequestDTO 📌 DTO for requesting a loan/advance
type LoanAdvanceRequestDTO struct {
	EmployeeID      uint    `json:"employee_id" validate:"required"`
	Amount          float64 `json:"amount" validate:"required,gt=0"`
	Purpose         string  `json:"purpose" validate:"required"`
	RepaymentMonths int     `json:"repayment_months" validate:"required,gt=0"`
}

// LoanAdvanceResponseDTO 📌 DTO for returning loan request details
type LoanAdvanceResponseDTO struct {
	ID              uint       `json:"id"`
	EmployeeID      uint       `json:"employee_id"`
	Amount          float64    `json:"amount"`
	Purpose         string     `json:"purpose"`
	Status          string     `json:"status"`
	ApprovedBy      uint       `json:"approved_by,omitempty"`
	ApprovalDate    *time.Time `json:"approval_date,omitempty"`
	RepaymentStart  *time.Time `json:"repayment_start,omitempty"`
	RepaymentMonths int        `json:"repayment_months"`
	CreatedAt       time.Time  `json:"created_at"`
}

// ApproveLoanAdvanceDTO 📌 DTO for approving a loan request
type ApproveLoanAdvanceDTO struct {
	LoanID         uint      `json:"loan_id" validate:"required"`
	ApproverID     uint      `json:"approver_id" validate:"required"`
	ApprovalDate   time.Time `json:"approval_date" validate:"required"`
	RepaymentStart time.Time `json:"repayment_start" validate:"required"`
}

// RejectLoanAdvanceDTO 📌 DTO for rejecting a loan request
type RejectLoanAdvanceDTO struct {
	LoanID     uint `json:"loan_id" validate:"required"`
	ApproverID uint `json:"approver_id" validate:"required"`
}
