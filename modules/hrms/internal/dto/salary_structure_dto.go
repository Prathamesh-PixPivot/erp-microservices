package dto

// SalaryStructureDTO represents a structured data transfer object for Salary Structures.
type SalaryStructureDTO struct {
	ID             uint    `json:"id,omitempty"`
	OrganizationID uint    `json:"organization_id"`
	DesignationID  uint    `json:"designation_id"`
	BaseSalary     float64 `json:"base_salary"`
	Allowances     float64 `json:"allowances"`
	TaxPercentage  float64 `json:"tax_percentage"`
	Deductions     float64 `json:"deductions"`
}
