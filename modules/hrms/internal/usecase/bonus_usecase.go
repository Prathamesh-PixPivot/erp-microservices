package usecase

import (
	"context"
	"hrms/internal/dto"
	"hrms/internal/domain"

	"go.uber.org/zap"
)

// CreateBonus handles adding a new bonus
func (u *HrmsUsecase) CreateBonus(ctx context.Context, req dto.CreateBonusRequest) (*dto.BonusDTO, error) {
	bonus := domain.Bonus{
		EmployeeID:  req.EmployeeID,
		Amount:      req.Amount,
		BonusType:   req.BonusType,
		Description: req.Description,
		ApprovedBy:  req.ApprovedBy,
		IssueDate:   req.IssueDate,
		Status:      req.Status,
	}

	// Set Approval Date only if approved
	if req.Status == "Approved" && req.ApprovalDate != nil {
		bonus.ApprovalDate = *req.ApprovalDate
	}

	result, err := u.HrmsRepo.CreateBonus(ctx, &bonus)
	if err != nil {
		u.Logger.Error("Failed to create bonus", zap.Error(err))
		return nil, err
	}

	return mapToBonusDTO(result), nil
}

// GetBonusByID fetches a bonus record by ID
func (u *HrmsUsecase) GetBonusByID(ctx context.Context, id uint) (*dto.BonusDTO, error) {
	bonus, err := u.HrmsRepo.GetBonusByID(ctx, id)
	if err != nil {
		u.Logger.Error("Failed to fetch bonus", zap.Error(err))
		return nil, err
	}

	return mapToBonusDTO(bonus), nil
}

// ListBonuses fetches multiple bonuses with optional filters
func (u *HrmsUsecase) ListBonuses(ctx context.Context, employeeID uint, status string) ([]dto.BonusDTO, error) {
	bonuses, err := u.HrmsRepo.ListBonuses(ctx, employeeID, status)
	if err != nil {
		u.Logger.Error("Failed to list bonuses", zap.Error(err))
		return nil, err
	}

	dtoBonuses := make([]dto.BonusDTO, len(bonuses))
	for i, b := range bonuses {
		dtoBonuses[i] = *mapToBonusDTO(&b)
	}

	return dtoBonuses, nil
}

// UpdateBonus updates a bonus record (e.g., approval, status change)
func (u *HrmsUsecase) UpdateBonus(ctx context.Context, id uint, req dto.UpdateBonusRequest) error {
	updates := make(map[string]interface{})

	if req.Status != "" {
		updates["status"] = req.Status
	}
	if req.ApprovalDate != nil {
		updates["approval_date"] = *req.ApprovalDate
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}

	err := u.HrmsRepo.UpdateBonus(ctx, id, updates)
	if err != nil {
		u.Logger.Error("Failed to update bonus", zap.Error(err))
		return err
	}
	return nil
}

// DeleteBonus removes a bonus record
func (u *HrmsUsecase) DeleteBonus(ctx context.Context, id uint) error {
	if err := u.HrmsRepo.DeleteBonus(ctx, id); err != nil {
		u.Logger.Error("Failed to delete bonus", zap.Error(err))
		return err
	}
	return nil
}

// mapToBonusDTO converts a domain model to a DTO
func mapToBonusDTO(b *domain.Bonus) *dto.BonusDTO {
	return &dto.BonusDTO{
		ID:          b.ID,
		EmployeeID:  b.EmployeeID,
		Amount:      b.Amount,
		BonusType:   b.BonusType,
		Description: b.Description,
		ApprovedBy:  b.ApprovedBy,
		ApprovalDate: &b.ApprovalDate,
		IssueDate:   b.IssueDate,
		Status:      b.Status,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}
