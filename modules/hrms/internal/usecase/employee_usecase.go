
package usecase

import (
	"context"
	"hrms/internal/dto"
	"hrms/internal/domain"
	"go.uber.org/zap"
)

// CreateEmployee handles employee creation
func (u *HrmsUsecase) CreateEmployee(ctx context.Context, req dto.CreateEmployeeRequest) (*dto.EmployeeDTO, error) {
	employee := domain.Employee{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          req.Phone,
		DateOfBirth:    req.DateOfBirth,
		EmploymentType: req.EmploymentType,
		HiredDate:      req.HiredDate,
		OrganizationID: req.OrganizationID,
		DepartmentID:   req.DepartmentID,
		DesignationID:  req.DesignationID,
		ReportsTo:      req.ReportsTo,
		Status:         "Active",
	}

	result, err := u.HrmsRepo.CreateEmployee(ctx, &employee)
	if err != nil {
		u.Logger.Error("Failed to create employee", zap.Error(err))
		return nil, err
	}

	return mapToEmployeeDTO(result), nil
}

// GetEmployee retrieves employee details
func (u *HrmsUsecase) GetEmployee(ctx context.Context, employeeID uint) (*dto.EmployeeDTO, error) {
	employee, err := u.HrmsRepo.GetEmployeeByID(ctx, employeeID)
	if err != nil {
		u.Logger.Error("Failed to fetch employee", zap.Error(err))
		return nil, err
	}
	return mapToEmployeeDTO(employee), nil
}

// UpdateEmployee updates employee details
func (u *HrmsUsecase) UpdateEmployee(ctx context.Context, employeeID uint, req dto.UpdateEmployeeRequest) error {
	updates := map[string]interface{}{}
	if req.FirstName != nil {
		updates["first_name"] = *req.FirstName
	}
	if req.LastName != nil {
		updates["last_name"] = *req.LastName
	}
	if req.Phone != nil {
		updates["phone"] = *req.Phone
	}
	if req.EmploymentType != nil {
		updates["employment_type"] = *req.EmploymentType
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if err := u.HrmsRepo.UpdateEmployeeProfile(ctx, employeeID, updates); err != nil {
		u.Logger.Error("Failed to update employee", zap.Error(err))
		return err
	}
	return nil
}

// DeleteEmployee soft deletes an employee
func (u *HrmsUsecase) DeleteEmployee(ctx context.Context, employeeID uint, reason string) error {
	if err := u.HrmsRepo.DeleteEmployee(ctx, employeeID, reason); err != nil {
		u.Logger.Error("Failed to delete employee", zap.Error(err))
		return err
	}
	return nil
}

// mapToEmployeeDTO converts domain model to DTO
func mapToEmployeeDTO(emp *domain.Employee) *dto.EmployeeDTO {
	return &dto.EmployeeDTO{
		ID:             emp.ID,
		FirstName:      emp.FirstName,
		LastName:       emp.LastName,
		Email:          emp.Email,
		Phone:          emp.Phone,
		DateOfBirth:    emp.DateOfBirth,
		EmploymentType: emp.EmploymentType,
		Status:         emp.Status,
		HiredDate:      emp.HiredDate,
		OrganizationID: emp.OrganizationID,
		DepartmentID:   emp.DepartmentID,
		DesignationID:  emp.DesignationID,
		ReportsTo:      emp.ReportsTo,
	}
}
