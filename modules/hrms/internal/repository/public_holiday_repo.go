package repository

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreatePublicHoliday 📌 Adds a new public holiday
func (r *HrmsRepository) CreatePublicHoliday(ctx context.Context, holiday *domain.PublicHoliday) (*domain.PublicHoliday, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Check if a holiday already exists on the same date for the organization
	var existingHoliday domain.PublicHoliday
	if err := tx.Where("organization_id = ? AND date = ?", holiday.OrganizationID, holiday.Date).
		First(&existingHoliday).Error; err == nil {
		tx.Rollback()
		return nil, hrmsErrors.ErrPublicHolidayExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 2: Insert the holiday
	if err := tx.Create(holiday).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create public holiday", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return holiday, nil
}

// GetPublicHoliday 📌 Fetch a public holiday by ID
func (r *HrmsRepository) GetPublicHoliday(ctx context.Context, holidayID uint) (*domain.PublicHoliday, error) {
	var holiday domain.PublicHoliday
	err := r.DB.WithContext(ctx).Where("id = ?", holidayID).First(&holiday).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrPublicHolidayNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetPublicHoliday", zap.Error(err))
		return nil, err
	}

	return &holiday, nil
}

// ListPublicHolidays 📌 Fetch all holidays for an organization (with optional year filter)
func (r *HrmsRepository) ListPublicHolidays(ctx context.Context, organizationID uint, year *int) ([]domain.PublicHoliday, error) {
	var holidays []domain.PublicHoliday

	query := r.DB.WithContext(ctx).Where("organization_id = ?", organizationID)

	if year != nil {
		start := time.Date(*year, 1, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(*year, 12, 31, 23, 59, 59, 999, time.UTC)
		query = query.Where("date BETWEEN ? AND ?", start, end)
	}

	if err := query.Order("date ASC").Find(&holidays).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch public holidays", zap.Error(err))
		return nil, err
	}

	return holidays, nil
}

// UpdatePublicHoliday 📌 Update holiday details (e.g., rename or reschedule)
func (r *HrmsRepository) UpdatePublicHoliday(ctx context.Context, holidayID uint, updates map[string]interface{}) error {
	tx := r.DB.WithContext(ctx).Begin()

	// Ensure no date duplication within the same organization
	if date, ok := updates["date"].(time.Time); ok {
		var holiday domain.PublicHoliday
		if err := tx.Where("id != ? AND date = ?", holidayID, date).First(&holiday).Error; err == nil {
			tx.Rollback()
			return hrmsErrors.ErrPublicHolidayExists
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return err
		}
	}

	// ✅ Update holiday
	if err := tx.Model(&domain.PublicHoliday{}).Where("id = ?", holidayID).Updates(updates).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to update public holiday", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// DeletePublicHoliday 📌 Remove a holiday
func (r *HrmsRepository) DeletePublicHoliday(ctx context.Context, holidayID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Where("id = ?", holidayID).Delete(&domain.PublicHoliday{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete public holiday", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}
