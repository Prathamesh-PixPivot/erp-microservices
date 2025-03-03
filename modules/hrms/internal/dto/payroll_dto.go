package dto

import "time"

// CreatePayrollDTO 📌 Request DTO for creating a payroll record
type CreatePayrollDTO struct {
	EmployeeID uint      `json:"employee_id" validate:"required"`
	Salary     float64   `json:"salary" validate:"required"`
	Tax        float64   `json:"tax"`
	Allowances float64   `json:"allowances"`
	Deductions float64   `json:"deductions"`
	NetSalary  float64   `json:"net_salary" validate:"required"`
	PaymentDate time.Time `json:"payment_date" validate:"required"`
	Status      string    `json:"status" validate:"required,oneof=Pending Processed"`
	PayslipURL  string    `json:"payslip_url,omitempty"`
	BankName    string    `json:"bank_name" validate:"required"`
	BankAccountNumber string `json:"bank_account_number,omitempty"`
	BranchCode  string    `json:"branch_code" validate:"required"`
}

// UpdatePayrollDTO 📌 Request DTO for updating payroll details
type UpdatePayrollDTO struct {
	Status     *string `json:"status,omitempty" validate:"omitempty,oneof=Pending Processed"`
	PayslipURL *string `json:"payslip_url,omitempty"`
}

// PayrollResponseDTO 📌 Response DTO for payroll data
type PayrollResponseDTO struct {
	ID               uint      `json:"id"`
	EmployeeID       uint      `json:"employee_id"`
	Salary           float64   `json:"salary"`
	Tax              float64   `json:"tax"`
	Allowances       float64   `json:"allowances"`
	Deductions       float64   `json:"deductions"`
	NetSalary        float64   `json:"net_salary"`
	PaymentDate      time.Time `json:"payment_date"`
	Status           string    `json:"status"`
	PayslipURL       string    `json:"payslip_url,omitempty"`
	BankName         string    `json:"bank_name"`
	BankAccountNumber string    `json:"bank_account_number,omitempty"`
	BranchCode       string    `json:"branch_code"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// PaginatedPayrollResponse 📌 Response DTO for paginated payroll results
type PaginatedPayrollResponse struct {
	Total    int64               `json:"total"`
	Limit    int                 `json:"limit"`
	Offset   int                 `json:"offset"`
	Payrolls []PayrollResponseDTO `json:"payrolls"`
}
