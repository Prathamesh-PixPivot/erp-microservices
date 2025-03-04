package usecase

import (
	"context"
	"hrms/internal/dto"
	"hrms/internal/domain"

	"go.uber.org/zap"
)

// CreateDesignation handles adding a new designation
func (u *HrmsUsecase) CreateDesignation(ctx context.Context, req dto.CreateDesignationRequest) (*dto.DesignationDTO, error) {
	designation := domain.Designation{
		Title:          req.Title,
		Level:          req.Level,
		HierarchyLevel: req.HierarchyLevel,
		DepartmentID:   req.DepartmentID,
	}

	result, err := u.HrmsRepo.CreateDesignation(ctx, &designation)
	if err != nil {
		u.Logger.Error("Failed to create designation", zap.Error(err))
		return nil, err
	}

	return mapToDesignationDTO(result), nil
}

// GetDesignationByID fetches a designation by ID
func (u *HrmsUsecase) GetDesignationByID(ctx context.Context, id uint) (*dto.DesignationDTO, error) {
	designation, err := u.HrmsRepo.GetDesignationByID(ctx, id)
	if err != nil {
		u.Logger.Error("Failed to fetch designation", zap.Error(err))
		return nil, err
	}

	return mapToDesignationDTO(designation), nil
}

// UpdateDesignation updates a designation's details
func (u *HrmsUsecase) UpdateDesignation(ctx context.Context, id uint, req dto.UpdateDesignationRequest) error {
	updates := make(map[string]interface{})

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Level != "" {
		updates["level"] = req.Level
	}
	if req.HierarchyLevel != 0 {
		updates["hierarchy_level"] = req.HierarchyLevel
	}

	err := u.HrmsRepo.UpdateDesignation(ctx, id, updates)
	if err != nil {
		u.Logger.Error("Failed to update designation", zap.Error(err))
		return err
	}
	return nil
}

// DeleteDesignation removes a designation record
func (u *HrmsUsecase) DeleteDesignation(ctx context.Context, id uint) error {
	if err := u.HrmsRepo.DeleteDesignation(ctx, id); err != nil {
		u.Logger.Error("Failed to delete designation", zap.Error(err))
		return err
	}
	return nil
}

// ListDesignations fetches multiple designations with pagination & search
func (u *HrmsUsecase) ListDesignations(ctx context.Context, departmentID uint, limit, offset int, search string) ([]dto.DesignationDTO, int64, error) {
	designations, totalCount, err := u.HrmsRepo.ListDesignations(ctx, departmentID, limit, offset, search)
	if err != nil {
		u.Logger.Error("Failed to list designations", zap.Error(err))
		return nil, 0, err
	}

	dtoDesignations := make([]dto.DesignationDTO, len(designations))
	for i, designation := range designations {
		dtoDesignations[i] = *mapToDesignationDTO(&designation)
	}

	return dtoDesignations, totalCount, nil
}

// mapToDesignationDTO converts a domain model to a DTO
func mapToDesignationDTO(d *domain.Designation) *dto.DesignationDTO {
	return &dto.DesignationDTO{
		ID:             d.ID,
		Title:          d.Title,
		Level:          d.Level,
		HierarchyLevel: d.HierarchyLevel,
		DepartmentID:   d.DepartmentID,
	}
}
