package dto

// CreateEmployeeBenefitsRequest represents the request to create employee benefits
type CreateEmployeeBenefitsRequest struct {
	EmployeeID     uint   `json:"employee_id" validate:"required"`
	HealthPlan     string `json:"health_plan" validate:"required"`
	RetirementPlan string `json:"retirement_plan" validate:"required"`
}

// UpdateEmployeeBenefitsRequest represents the request to update employee benefits
type UpdateEmployeeBenefitsRequest struct {
	HealthPlan     string `json:"health_plan,omitempty"`
	RetirementPlan string `json:"retirement_plan,omitempty"`
}

// EmployeeBenefitsDTO represents the response structure for employee benefits
type EmployeeBenefitsDTO struct {
	ID             uint   `json:"id"`
	EmployeeID     uint   `json:"employee_id"`
	HealthPlan     string `json:"health_plan"`
	RetirementPlan string `json:"retirement_plan"`
}
