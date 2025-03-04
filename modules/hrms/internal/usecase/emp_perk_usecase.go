package usecase

import (
	"context"
	"hrms/internal/dto"
	"hrms/internal/domain"

	"go.uber.org/zap"
)

// CreateEmployeePerk adds a perk for an employee
func (u *HrmsUsecase) CreateEmployeePerk(ctx context.Context, req dto.CreateEmployeePerkRequest) (*dto.EmployeePerkDTO, error) {
	perk := domain.EmployeePerk{
		EmployeeID: req.EmployeeID,
		Perk:       req.Perk,
	}

	result, err := u.HrmsRepo.CreateEmployeePerk(ctx, &perk)
	if err != nil {
		u.Logger.Error("Failed to add employee perk", zap.Error(err))
		return nil, err
	}

	return mapToEmployeePerkDTO(result), nil
}

// GetEmployeePerks fetches all perks assigned to an employee
func (u *HrmsUsecase) GetEmployeePerks(ctx context.Context, employeeID uint) ([]dto.EmployeePerkDTO, error) {
	perks, err := u.HrmsRepo.GetEmployeePerks(ctx, employeeID)
	if err != nil {
		u.Logger.Error("Failed to fetch employee perks", zap.Error(err))
		return nil, err
	}

	dtoPerks := make([]dto.EmployeePerkDTO, len(perks))
	for i, perk := range perks {
		dtoPerks[i] = *mapToEmployeePerkDTO(&perk)
	}

	return dtoPerks, nil
}

// UpdateEmployeePerk updates an employee's perk details
func (u *HrmsUsecase) UpdateEmployeePerk(ctx context.Context, perkID uint, req dto.UpdateEmployeePerkRequest) error {
	err := u.HrmsRepo.UpdateEmployeePerk(ctx, perkID, req.Perk)
	if err != nil {
		u.Logger.Error("Failed to update employee perk", zap.Error(err))
		return err
	}
	return nil
}

// DeleteEmployeePerk removes a perk from an employee
func (u *HrmsUsecase) DeleteEmployeePerk(ctx context.Context, perkID uint) error {
	if err := u.HrmsRepo.DeleteEmployeePerk(ctx, perkID); err != nil {
		u.Logger.Error("Failed to delete employee perk", zap.Error(err))
		return err
	}
	return nil
}

// mapToEmployeePerkDTO converts a domain model to DTO
func mapToEmployeePerkDTO(p *domain.EmployeePerk) *dto.EmployeePerkDTO {
	return &dto.EmployeePerkDTO{
		ID:         p.ID,
		EmployeeID: p.EmployeeID,
		Perk:       p.Perk,
	}
}
