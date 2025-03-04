package repository

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreateDepartment 📌 Adds a new department
func (r *HrmsRepository) CreateDepartment(ctx context.Context, department *domain.Department) (*domain.Department, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Check if the organization exists
	var org domain.Organization
	if err := tx.Where("id = ?", department.OrganizationID).First(&org).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, hrmsErrors.ErrInvalidOrganization
		}
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 2: Ensure department name is unique within the organization
	var existingDept domain.Department
	if err := tx.Where("name = ? AND organization_id = ?", department.Name, department.OrganizationID).First(&existingDept).Error; err == nil {
		tx.Rollback()
		return nil, hrmsErrors.ErrDepartmentAlreadyExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 3: Create Department
	if err := tx.Create(department).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create department", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return department, nil
}

// GetDepartmentByID 📌 Fetch department by ID
func (r *HrmsRepository) GetDepartmentByID(ctx context.Context, deptID uint) (*domain.Department, error) {
	var department domain.Department
	err := r.DB.WithContext(ctx).Where("id = ?", deptID).Preload("Employees").First(&department).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrDepartmentNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetDepartmentByID", zap.Error(err))
		return nil, err
	}
	return &department, nil
}

// UpdateDepartment 📌 Update department details
func (r *HrmsRepository) UpdateDepartment(ctx context.Context, deptID uint, updates map[string]interface{}) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Check if the department exists
	var department domain.Department
	if err := tx.Where("id = ?", deptID).First(&department).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return hrmsErrors.ErrDepartmentNotFound
		}
		tx.Rollback()
		return err
	}

	// ✅ Step 2: Ensure unique name within the same organization (if name update is requested)
	if newName, exists := updates["name"]; exists {
		var existingDept domain.Department
		if err := tx.Where("name = ? AND organization_id = ? AND id != ?", newName, department.OrganizationID, deptID).First(&existingDept).Error; err == nil {
			tx.Rollback()
			return hrmsErrors.ErrDepartmentAlreadyExists
		}
	}

	// ✅ Step 3: Update department
	if err := tx.Model(&department).Updates(updates).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to update department", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// DeleteDepartment 📌 Soft delete department
func (r *HrmsRepository) DeleteDepartment(ctx context.Context, deptID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Where("id = ?", deptID).Delete(&domain.Department{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete department", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// ListDepartments 📌 Fetch departments with pagination & search
func (r *HrmsRepository) ListDepartments(ctx context.Context, organizationID uint, limit, offset int, search string) ([]domain.Department, int64, error) {
	var departments []domain.Department
	var totalCount int64

	query := r.DB.WithContext(ctx).Model(&domain.Department{}).Where("organization_id = ?", organizationID)

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	query.Count(&totalCount)

	if err := query.Limit(limit).Offset(offset).Find(&departments).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch departments", zap.Error(err))
		return nil, 0, err
	}

	return departments, totalCount, nil
}
