package usecase

import (
	"context"
	"hrms/internal/dto"
	"hrms/internal/domain"

	"go.uber.org/zap"
)

// CreateDepartment handles adding a new department
func (u *HrmsUsecase) CreateDepartment(ctx context.Context, req dto.CreateDepartmentRequest) (*dto.DepartmentDTO, error) {
	department := domain.Department{
		Name:           req.Name,
		OrganizationID: req.OrganizationID,
	}

	result, err := u.HrmsRepo.CreateDepartment(ctx, &department)
	if err != nil {
		u.Logger.Error("Failed to create department", zap.Error(err))
		return nil, err
	}

	return mapToDepartmentDTO(result), nil
}

// GetDepartmentByID fetches a department by ID
func (u *HrmsUsecase) GetDepartmentByID(ctx context.Context, id uint) (*dto.DepartmentDTO, error) {
	department, err := u.HrmsRepo.GetDepartmentByID(ctx, id)
	if err != nil {
		u.Logger.Error("Failed to fetch department", zap.Error(err))
		return nil, err
	}

	return mapToDepartmentDTO(department), nil
}

// UpdateDepartment updates a department's details
func (u *HrmsUsecase) UpdateDepartment(ctx context.Context, id uint, req dto.UpdateDepartmentRequest) error {
	updates := make(map[string]interface{})

	if req.Name != "" {
		updates["name"] = req.Name
	}

	err := u.HrmsRepo.UpdateDepartment(ctx, id, updates)
	if err != nil {
		u.Logger.Error("Failed to update department", zap.Error(err))
		return err
	}
	return nil
}

// DeleteDepartment removes a department record
func (u *HrmsUsecase) DeleteDepartment(ctx context.Context, id uint) error {
	if err := u.HrmsRepo.DeleteDepartment(ctx, id); err != nil {
		u.Logger.Error("Failed to delete department", zap.Error(err))
		return err
	}
	return nil
}

// ListDepartments fetches multiple departments with pagination & search
func (u *HrmsUsecase) ListDepartments(ctx context.Context, organizationID uint, limit, offset int, search string) ([]dto.DepartmentDTO, int64, error) {
	departments, totalCount, err := u.HrmsRepo.ListDepartments(ctx, organizationID, limit, offset, search)
	if err != nil {
		u.Logger.Error("Failed to list departments", zap.Error(err))
		return nil, 0, err
	}

	dtoDepartments := make([]dto.DepartmentDTO, len(departments))
	for i, dept := range departments {
		dtoDepartments[i] = *mapToDepartmentDTO(&dept)
	}

	return dtoDepartments, totalCount, nil
}

// mapToDepartmentDTO converts a domain model to a DTO
func mapToDepartmentDTO(d *domain.Department) *dto.DepartmentDTO {
	return &dto.DepartmentDTO{
		ID:             d.ID,
		Name:           d.Name,
		OrganizationID: d.OrganizationID,
	}
}
