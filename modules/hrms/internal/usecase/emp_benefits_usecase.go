package usecase

import (
	"context"
	"hrms/internal/dto"
	"hrms/internal/domain"

	"go.uber.org/zap"
)

// CreateEmployeeBenefits adds new employee benefits
func (u *HrmsUsecase) CreateEmployeeBenefits(ctx context.Context, req dto.CreateEmployeeBenefitsRequest) (*dto.EmployeeBenefitsDTO, error) {
	benefits := domain.EmployeeBenefits{
		EmployeeID:     req.EmployeeID,
		HealthPlan:     req.HealthPlan,
		RetirementPlan: req.RetirementPlan,
	}

	result, err := u.HrmsRepo.CreateEmployeeBenefits(ctx, &benefits)
	if err != nil {
		u.Logger.Error("Failed to create employee benefits", zap.Error(err))
		return nil, err
	}

	return mapToEmployeeBenefitsDTO(result), nil
}

// GetEmployeeBenefits fetches benefits by Employee ID
func (u *HrmsUsecase) GetEmployeeBenefits(ctx context.Context, employeeID uint) (*dto.EmployeeBenefitsDTO, error) {
	benefits, err := u.HrmsRepo.GetEmployeeBenefits(ctx, employeeID)
	if err != nil {
		u.Logger.Error("Failed to fetch employee benefits", zap.Error(err))
		return nil, err
	}

	return mapToEmployeeBenefitsDTO(benefits), nil
}

// UpdateEmployeeBenefits updates an employee’s benefits
func (u *HrmsUsecase) UpdateEmployeeBenefits(ctx context.Context, employeeID uint, req dto.UpdateEmployeeBenefitsRequest) error {
	updates := make(map[string]interface{})

	if req.HealthPlan != "" {
		updates["health_plan"] = req.HealthPlan
	}
	if req.RetirementPlan != "" {
		updates["retirement_plan"] = req.RetirementPlan
	}

	err := u.HrmsRepo.UpdateEmployeeBenefits(ctx, employeeID, updates)
	if err != nil {
		u.Logger.Error("Failed to update employee benefits", zap.Error(err))
		return err
	}
	return nil
}

// DeleteEmployeeBenefits removes employee benefits
func (u *HrmsUsecase) DeleteEmployeeBenefits(ctx context.Context, employeeID uint) error {
	if err := u.HrmsRepo.DeleteEmployeeBenefits(ctx, employeeID); err != nil {
		u.Logger.Error("Failed to delete employee benefits", zap.Error(err))
		return err
	}
	return nil
}

// ListEmployeeBenefits retrieves all employee benefits with filtering
func (u *HrmsUsecase) ListEmployeeBenefits(ctx context.Context, healthPlan, retirementPlan string) ([]dto.EmployeeBenefitsDTO, error) {
	benefits, err := u.HrmsRepo.ListEmployeeBenefits(ctx, healthPlan, retirementPlan)
	if err != nil {
		u.Logger.Error("Failed to list employee benefits", zap.Error(err))
		return nil, err
	}

	dtoBenefits := make([]dto.EmployeeBenefitsDTO, len(benefits))
	for i, benefit := range benefits {
		dtoBenefits[i] = *mapToEmployeeBenefitsDTO(&benefit)
	}

	return dtoBenefits, nil
}

// mapToEmployeeBenefitsDTO converts domain model to DTO
func mapToEmployeeBenefitsDTO(b *domain.EmployeeBenefits) *dto.EmployeeBenefitsDTO {
	return &dto.EmployeeBenefitsDTO{
		ID:             b.ID,
		EmployeeID:     b.EmployeeID,
		HealthPlan:     b.HealthPlan,
		RetirementPlan: b.RetirementPlan,
	}
}
