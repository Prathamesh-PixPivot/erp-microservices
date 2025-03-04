package dto

// CreateEmployeePerkRequest represents the request to add a perk for an employee
type CreateEmployeePerkRequest struct {
	EmployeeID uint   `json:"employee_id" validate:"required"`
	Perk       string `json:"perk" validate:"required,max=50"`
}

// UpdateEmployeePerkRequest represents the request to update an employee's perk
type UpdateEmployeePerkRequest struct {
	Perk string `json:"perk" validate:"required,max=50"`
}

// EmployeePerkDTO represents the response structure for employee perks
type EmployeePerkDTO struct {
	ID         uint   `json:"id"`
	EmployeeID uint   `json:"employee_id"`
	Perk       string `json:"perk"`
}
