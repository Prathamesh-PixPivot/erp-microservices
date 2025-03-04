package repository

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreateShift 📌 Create a new shift
func (r *HrmsRepository) CreateShift(ctx context.Context, shift *domain.Shift) (*domain.Shift, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Check if shift with same name exists
	var existingShift domain.Shift
	if err := tx.Where("name = ?", shift.Name).First(&existingShift).Error; err == nil {
		tx.Rollback()
		r.Logger.Warn("⚠️ Shift already exists", zap.String("name", shift.Name))
		return nil, hrmsErrors.ErrShiftAlreadyExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err // Some other DB error
	}

	// ✅ Step 2: Insert Shift
	if err := tx.Create(shift).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create shift", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return shift, nil
}

// GetShiftByID 📌 Fetch shift by ID
func (r *HrmsRepository) GetShiftByID(ctx context.Context, shiftID uint) (*domain.Shift, error) {
	var shift domain.Shift
	err := r.DB.WithContext(ctx).Where("id = ?", shiftID).First(&shift).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrShiftNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetShiftByID", zap.Error(err))
		return nil, err
	}
	return &shift, nil
}

// UpdateShift 📌 Update shift details
func (r *HrmsRepository) UpdateShift(ctx context.Context, shiftID uint, updates map[string]interface{}) error {
	if err := r.DB.WithContext(ctx).Model(&domain.Shift{}).
		Where("id = ?", shiftID).
		Updates(updates).Error; err != nil {
		r.Logger.Error("❌ Failed to update shift", zap.Error(err))
		return err
	}
	return nil
}

// DeleteShift 📌 Soft delete shift
func (r *HrmsRepository) DeleteShift(ctx context.Context, shiftID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Where("id = ?", shiftID).Delete(&domain.Shift{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete shift", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// ListShifts 📌 Fetch shifts with pagination & search
func (r *HrmsRepository) ListShifts(ctx context.Context, limit, offset int, search string) ([]domain.Shift, int64, error) {
	var shifts []domain.Shift
	var totalCount int64

	query := r.DB.WithContext(ctx).Model(&domain.Shift{})

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	query.Count(&totalCount)

	if err := query.Limit(limit).Offset(offset).Find(&shifts).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch shifts", zap.Error(err))
		return nil, 0, err
	}

	return shifts, totalCount, nil
}
