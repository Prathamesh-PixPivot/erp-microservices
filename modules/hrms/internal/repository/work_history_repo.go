package repository

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreateWorkHistory 📌 Add a new work history record
func (r *HrmsRepository) CreateWorkHistory(ctx context.Context, workHistory *domain.WorkHistory) (*domain.WorkHistory, error) {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Create(workHistory).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create work history", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return workHistory, nil
}

// GetWorkHistoryByID 📌 Fetch a specific work history record by ID
func (r *HrmsRepository) GetWorkHistoryByID(ctx context.Context, workHistoryID uint) (*domain.WorkHistory, error) {
	var workHistory domain.WorkHistory
	err := r.DB.WithContext(ctx).Where("id = ?", workHistoryID).First(&workHistory).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrWorkHistoryNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetWorkHistoryByID", zap.Error(err))
		return nil, err
	}
	return &workHistory, nil
}

// GetWorkHistoryByEmployee 📌 Fetch work history records for an employee
func (r *HrmsRepository) GetWorkHistoryByEmployee(ctx context.Context, employeeID uint) ([]domain.WorkHistory, error) {
	var workHistories []domain.WorkHistory
	err := r.DB.WithContext(ctx).Where("employee_id = ?", employeeID).Find(&workHistories).Error
	if err != nil {
		r.Logger.Error("❌ Failed to fetch work history for employee", zap.Uint("employee_id", employeeID), zap.Error(err))
		return nil, err
	}
	return workHistories, nil
}

// UpdateWorkHistory 📌 Update a specific work history record
func (r *HrmsRepository) UpdateWorkHistory(ctx context.Context, workHistoryID uint, updates map[string]interface{}) error {
	if err := r.DB.WithContext(ctx).Model(&domain.WorkHistory{}).
		Where("id = ?", workHistoryID).
		Updates(updates).Error; err != nil {
		r.Logger.Error("❌ Failed to update work history", zap.Error(err))
		return err
	}
	return nil
}

// DeleteWorkHistory 📌 Soft delete a work history record
func (r *HrmsRepository) DeleteWorkHistory(ctx context.Context, workHistoryID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Where("id = ?", workHistoryID).Delete(&domain.WorkHistory{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete work history", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}
