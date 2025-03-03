package repository

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreateDesignation 📌 Adds a new designation
func (r *HrmsRepository) CreateDesignation(ctx context.Context, designation *domain.Designation) (*domain.Designation, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Check if the department exists
	var department domain.Department
	if err := tx.Where("id = ?", designation.DepartmentID).First(&department).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, hrmsErrors.ErrInvalidDepartment
		}
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 2: Ensure designation title & level are unique within the department
	var existingDesignation domain.Designation
	if err := tx.Where("title = ? AND level = ? AND department_id = ?", designation.Title, designation.Level, designation.DepartmentID).First(&existingDesignation).Error; err == nil {
		tx.Rollback()
		return nil, hrmsErrors.ErrDesignationAlreadyExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 3: Insert designation
	if err := tx.Create(designation).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create designation", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return designation, nil
}

// GetDesignationByID 📌 Fetch designation by ID
func (r *HrmsRepository) GetDesignationByID(ctx context.Context, designationID uint) (*domain.Designation, error) {
	var designation domain.Designation
	err := r.DB.WithContext(ctx).Where("id = ?", designationID).Preload("Employees").First(&designation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrDesignationNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetDesignationByID", zap.Error(err))
		return nil, err
	}
	return &designation, nil
}

// UpdateDesignation 📌 Update designation details
func (r *HrmsRepository) UpdateDesignation(ctx context.Context, designationID uint, updates map[string]interface{}) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Check if the designation exists
	var designation domain.Designation
	if err := tx.Where("id = ?", designationID).First(&designation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return hrmsErrors.ErrDesignationNotFound
		}
		tx.Rollback()
		return err
	}

	// ✅ Step 2: Ensure title & level uniqueness within the department (if title/level is being updated)
	if newTitle, titleExists := updates["title"]; titleExists {
		if newLevel, levelExists := updates["level"]; levelExists {
			var existingDesignation domain.Designation
			if err := tx.Where("title = ? AND level = ? AND department_id = ? AND id != ?", newTitle, newLevel, designation.DepartmentID, designationID).First(&existingDesignation).Error; err == nil {
				tx.Rollback()
				return hrmsErrors.ErrDesignationAlreadyExists
			}
		}
	}

	// ✅ Step 3: Update designation
	if err := tx.Model(&designation).Updates(updates).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to update designation", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// DeleteDesignation 📌 Soft delete designation
func (r *HrmsRepository) DeleteDesignation(ctx context.Context, designationID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Where("id = ?", designationID).Delete(&domain.Designation{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete designation", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// ListDesignations 📌 Fetch designations with pagination & search
func (r *HrmsRepository) ListDesignations(ctx context.Context, departmentID uint, limit, offset int, search string) ([]domain.Designation, int64, error) {
	var designations []domain.Designation
	var totalCount int64

	query := r.DB.WithContext(ctx).Model(&domain.Designation{}).Where("department_id = ?", departmentID)

	if search != "" {
		query = query.Where("title ILIKE ?", "%"+search+"%")
	}

	query.Count(&totalCount)

	if err := query.Limit(limit).Offset(offset).Find(&designations).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch designations", zap.Error(err))
		return nil, 0, err
	}

	return designations, totalCount, nil
}
