package domain

import (
	"time"

	"gorm.io/gorm"
)

// Organization Model
type Organization struct {
	gorm.Model
	Name        string       `json:"name"`
	Address     string       `json:"address"`
	Phone       string       `json:"phone"`
	Email       string       `json:"email"`
	Employees   []Employee   `gorm:"foreignKey:OrganizationID"`
	Departments []Department `gorm:"foreignKey:OrganizationID"`
}

// Employee Model
type Employee struct {
	gorm.Model
	FirstName        string     `json:"first_name"`
	LastName         string     `json:"last_name"`
	Email            string     `json:"email" gorm:"unique"`
	Phone            string     `json:"phone"`
	DateOfBirth      time.Time  `json:"date_of_birth"`
	EmploymentType   string     `json:"employment_type" gorm:"type:varchar(20)"` // Full-time, Part-time, etc.
	Status           string     `json:"status" gorm:"type:varchar(20)"`          // Active, Resigned, etc.
	HiredDate        *time.Time `json:"hired_date,omitempty"`
	ProbationEndDate *time.Time `json:"probation_end_date,omitempty"`
	ContractEndDate  *time.Time `json:"contract_end_date,omitempty"`
	IsConfirmed      bool       `json:"is_confirmed"`

	OrganizationID uint `json:"organization_id" gorm:"index"`
	DepartmentID   uint `json:"department_id" gorm:"index"`
	DesignationID  uint `json:"designation_id" gorm:"index"`
	PayrollID      uint `json:"payroll_id" gorm:"index"`

	ReportsTo   *uint         `json:"reports_to,omitempty" gorm:"index"` // Self-referencing
	WorkHistory []WorkHistory `gorm:"foreignKey:EmployeeID"`

	Attendance    []Attendance       `gorm:"foreignKey:EmployeeID"`
	Leaves        []Leave            `gorm:"foreignKey:EmployeeID"`
	LeaveBalances []LeaveBalance     `gorm:"foreignKey:EmployeeID"`
	Documents     []EmployeeDocument `gorm:"foreignKey:EmployeeID"`
	Benefits      EmployeeBenefits   `gorm:"foreignKey:EmployeeID"`
	ShiftID       uint               `json:"shift_id" gorm:"index"`
}

// Shift Model (Updated)
type Shift struct {
	gorm.Model
	Name       string     `json:"name"`
	ShiftType  ShiftType  `json:"shift_type" gorm:"type:varchar(20)"`
	StartTime  string     `json:"start_time"`
	EndTime    string     `json:"end_time"`
	DaysOfWeek string     `json:"days_of_week"` // Stores days shift is active (e.g., "Mon-Fri")
	Employees  []Employee `gorm:"foreignKey:ShiftID"`
}

// Work History
type WorkHistory struct {
	gorm.Model
	EmployeeID    uint       `json:"employee_id" gorm:"index"`
	Company       string     `json:"company"`
	Designation   string     `json:"designation"`
	StartDate     time.Time  `json:"start_date"`
	EndDate       *time.Time `json:"end_date,omitempty"`
	ReasonForExit string     `json:"reason_for_exit"`
}

// Employee Document Model (with Expiry Tracking)
type EmployeeDocument struct {
	gorm.Model
	EmployeeID   uint       `json:"employee_id"`
	DocumentName string     `json:"document_name"`
	DocumentURL  string     `json:"document_url"`
	ExpiryDate   *time.Time `json:"expiry_date,omitempty"`
}

// Employee Exit & Clearance Process
type EmployeeExit struct {
	gorm.Model
	EmployeeID      uint      `json:"employee_id"`
	ExitType        string    `json:"exit_type" gorm:"type:varchar(20)"`
	ExitDate        time.Time `json:"exit_date"`
	ClearanceStatus string    `json:"clearance_status" gorm:"type:varchar(20)"`
}

// LoanAdvance Model - Tracks loans & advance salaries for employees
type LoanAdvance struct {
	gorm.Model
	EmployeeID      uint       `json:"employee_id" gorm:"index"`
	Employee        Employee   `gorm:"foreignKey:EmployeeID"`
	Amount          float64    `json:"amount"`
	Purpose         string     `json:"purpose"`
	Status          string     `json:"status" gorm:"type:varchar(20)"` // Pending, Approved, Rejected
	ApprovedBy      uint       `json:"approved_by"`
	Approver        Employee   `gorm:"foreignKey:ApprovedBy"` // Explicit foreign key reference to Employee
	RepaymentMonths int        `json:"repayment_months"`
	ApprovalDate    *time.Time `json:"approval_date,omitempty"`
	RepaymentStart  *time.Time `json:"repayment_start,omitempty"`
}

// Department Model
type Department struct {
	gorm.Model
	Name           string     `json:"name"`
	OrganizationID uint       `json:"organization_id"`
	Employees      []Employee `gorm:"foreignKey:DepartmentID"`
}

// Designation Model (With Hierarchy Level)
type Designation struct {
	gorm.Model
	Title          string     `json:"title"`
	Level          string     `json:"level"`
	HierarchyLevel int        `json:"hierarchy_level"`
	DepartmentID   uint       `json:"department_id"`
	Employees      []Employee `gorm:"foreignKey:DesignationID"`
}

// Attendance Model
type Attendance struct {
	gorm.Model
	EmployeeID  uint       `json:"employee_id" gorm:"index"`
	Date        time.Time  `json:"date"`
	CheckIn     *time.Time `json:"check_in,omitempty"`
	CheckOut    *time.Time `json:"check_out,omitempty"`
	WorkHours   float64    `json:"work_hours"`
	Overtime    float64    `json:"overtime"`
	BreakTime   float64    `json:"break_time"`
	Location    string     `json:"location"`
	IsRemote    bool       `json:"is_remote"`
	PunchMethod string     `json:"punch_method" gorm:"type:varchar(20)"` // Manual, Biometric, Geolocation
}

// Leave Model
type Leave struct {
	gorm.Model
	EmployeeID     uint      `json:"employee_id" gorm:"index"`
	LeaveType      string    `json:"leave_type" gorm:"type:varchar(20)"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	Status         string    `json:"status" gorm:"type:varchar(20)"` // Pending, Approved, Rejected
	ApproverID     uint      `json:"approver_id"`
	MultiLevelStep int       `json:"multi_level_step"` // 1=Manager, 2=HR, 3=Director
	Comments       string    `json:"comments,omitempty"`
}

// LeaveBalance Model
type LeaveBalance struct {
	gorm.Model
	EmployeeID  uint      `json:"employee_id"`
	LeaveType   LeaveType `json:"leave_type"`
	TotalLeaves float64   `json:"total_leaves"`
	UsedLeaves  float64   `json:"used_leaves"`
	Remaining   float64   `json:"remaining"`
}

// Public Holidays Model
type PublicHoliday struct {
	gorm.Model
	OrganizationID uint      `json:"organization_id"`
	Name           string    `json:"name"`
	Date           time.Time `json:"date"`
}

// Leave Policy Model
type LeavePolicy struct {
	gorm.Model
	OrganizationID uint      `json:"organization_id"`
	LeaveType      LeaveType `json:"leave_type"`
	MaxDays        int       `json:"max_days"`
	CarryForward   bool      `json:"carry_forward"`
}

// Payroll Model
type Payroll struct {
	gorm.Model
	EmployeeID        uint      `json:"employee_id" gorm:"index"`
	Salary            float64   `json:"salary"`
	Tax               float64   `json:"tax"`
	Allowances        float64   `json:"allowances"`
	Deductions        float64   `json:"deductions"`
	NetSalary         float64   `json:"net_salary"`
	PaymentDate       time.Time `json:"payment_date"`
	Status            string    `json:"status" gorm:"type:varchar(20)"` // Pending, Processed, etc.
	PayslipURL        string    `json:"payslip_url,omitempty"`
	BankName          string    `json:"bank_name"`
	BankAccountNumber string    `json:"bank_account_number,omitempty"`
	BranchCode        string    `json:"branch_code"`
}

// Salary Structure Model
type SalaryStructure struct {
	gorm.Model
	OrganizationID uint    `json:"organization_id"`
	DesignationID  uint    `json:"designation_id"`
	BaseSalary     float64 `json:"base_salary"`
	Allowances     float64 `json:"allowances"`
	TaxPercentage  float64 `json:"tax_percentage"`
	Deductions     float64 `json:"deductions"`
}

// Bonus Model
type Bonus struct {
	gorm.Model
	EmployeeID   uint      `json:"employee_id" gorm:"index"`
	Amount       float64   `json:"amount"`
	BonusType    string    `json:"bonus_type" gorm:"type:varchar(20)"` // Performance, Festival, etc.
	Description  string    `json:"description"`
	ApprovedBy   uint      `json:"approved_by"`
	ApprovalDate time.Time `json:"approval_date"`
	IssueDate    time.Time `json:"issue_date"`
	Status       string    `json:"status" gorm:"type:varchar(20)"` // Pending, Approved, Rejected
}

// Employee Benefits Model
type EmployeeBenefits struct {
	gorm.Model
	EmployeeID     uint           `json:"employee_id" gorm:"index"`
	HealthPlan     string         `json:"health_plan" gorm:"type:varchar(20)"`     // Basic, Premium, etc.
	RetirementPlan string         `json:"retirement_plan" gorm:"type:varchar(20)"` // Pension, Stock Options, etc.
	Perks          []EmployeePerk `gorm:"foreignKey:EmployeeID"`
}

// Employee Perks
type EmployeePerk struct {
	gorm.Model
	EmployeeID uint   `json:"employee_id" gorm:"index"`
	Perk       string `json:"perk" gorm:"type:varchar(50)"` // Meal Card, Travel Allowance, etc.
}

// Expense Reimbursement Model
type Expense struct {
	gorm.Model
	EmployeeID  uint      `json:"employee_id" gorm:"index"`
	ExpenseType string    `json:"expense_type" gorm:"type:varchar(20)"` // Travel, Office Supplies, etc.
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	Status      string    `json:"status" gorm:"type:varchar(20)"`
	ApproverID  uint      `json:"approver_id"`
	ReceiptURL  string    `json:"receipt_url,omitempty"`
}

// Performance Review Model (Unified)
type PerformanceReview struct {
	gorm.Model
	EmployeeID    uint      `json:"employee_id" gorm:"index"` // Employee being reviewed
	ReviewerID    uint      `json:"reviewer_id" gorm:"index"` // Manager or HR
	ReviewDate    time.Time `json:"review_date"`
	ReviewPeriod  string    `json:"review_period" gorm:"type:varchar(20)"` // Monthly, Annual, etc.
	OverallRating int       `json:"overall_rating"`                        // 1-5 scale
	Feedback      string    `json:"feedback"`
	Promotion     bool      `json:"promotion"`

	// Relations
	KPIs              []PerformanceKPI   `gorm:"foreignKey:ReviewID"`
	SkillDevelopments []SkillDevelopment `gorm:"foreignKey:ReviewID"`
}

// Performance Review Model
type Performance struct {
	gorm.Model
	EmployeeID uint             `json:"employee_id"`
	ReviewDate time.Time        `json:"review_date"`
	Feedback   string           `json:"feedback"`
	Rating     int              `json:"rating"` // 1 to 5 scale
	Promotion  bool             `json:"promotion"`
	KPIs       []PerformanceKPI `gorm:"foreignKey:ReviewID"`
}

// Performance KPI Model
type PerformanceKPI struct {
	gorm.Model
	ReviewID uint    `json:"review_id" gorm:"index"`
	KPIName  string  `json:"kpi_name"`
	Score    float64 `json:"score"` // 0-100 or weighted score
	Comments string  `json:"comments,omitempty"`
}

// Skill Development Model
type SkillDevelopment struct {
	gorm.Model
	ReviewID uint   `json:"review_id" gorm:"index"`
	Skill    string `json:"skill"`
	Progress string `json:"progress"`
}
