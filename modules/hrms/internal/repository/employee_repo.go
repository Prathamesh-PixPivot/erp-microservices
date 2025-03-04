package repository

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"

	// "strings"
	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreateEmployee 📌 Create a new employee (Employee Signup)
func (r *HrmsRepository) CreateEmployee(ctx context.Context, employee *domain.Employee) (*domain.Employee, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Check if employee already exists
	var existingEmployee domain.Employee
	if err := tx.Where("email = ?", employee.Email).First(&existingEmployee).Error; err == nil {
		tx.Rollback()
		r.Logger.Warn("⚠️ Employee already exists", zap.String("email", employee.Email))
		return nil, hrmsErrors.ErrEmployeeAlreadyExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 2: Validate Foreign Keys
	if err := r.validateEmployeeForeignKeys(tx, employee); err != nil {
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 3: Insert Employee
	if err := tx.Create(employee).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create employee", zap.Error(err))
		return nil, err
	}

	tx.Commit() // ✅ Commit transaction
	return employee, nil
}

// validateEmployeeForeignKeys ensures Organization, Department, and Designation exist before creating employee
func (r *HrmsRepository) validateEmployeeForeignKeys(tx *gorm.DB, employee *domain.Employee) error {
	// Validate Organization
	if err := tx.First(&domain.Organization{}, employee.OrganizationID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return hrmsErrors.ErrOrganizationNotFound
	}

	// Validate Department
	if err := tx.First(&domain.Department{}, employee.DepartmentID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return hrmsErrors.ErrDepartmentNotFound
	}

	// Validate Designation
	if err := tx.First(&domain.Designation{}, employee.DesignationID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return hrmsErrors.ErrDesignationNotFound
	}

	// Validate ReportsTo (Self-referencing)
	if employee.ReportsTo != nil {
		if err := tx.First(&domain.Employee{}, *employee.ReportsTo).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return hrmsErrors.ErrEmployeeNotFound
		}
	}

	return nil
}

// GetEmployeeByID 📌
func (r *HrmsRepository) GetEmployeeByID(ctx context.Context, employeeID uint) (*domain.Employee, error) {
	var employee domain.Employee
	err := r.DB.WithContext(ctx).
		Preload("WorkHistory").
		Preload("Attendance").
		Preload("Leaves").
		Preload("LeaveBalances").
		Preload("Documents").
		Preload("Benefits").
		Where("id = ?", employeeID).
		First(&employee).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrEmployeeNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetEmployeeByID", zap.Error(err))
		return nil, err
	}
	return &employee, nil
}

// UpdateEmployeeProfile 📌 Update employee profile
func (r *HrmsRepository) UpdateEmployeeProfile(ctx context.Context, employeeID uint, updates map[string]interface{}) error {

	if err := r.DB.WithContext(ctx).Model(&domain.Employee{}).
		Where("id = ?", employeeID).
		Updates(updates).Error; err != nil {
		r.Logger.Error("❌ Failed to update employee profile", zap.Error(err))
		return err
	}
	return nil
}

// DeleteEmploee 📌 Delete employee (Soft Delete or Permanent)
func (r *HrmsRepository) DeleteEmployee(ctx context.Context, employeeID uint, reason string) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Soft Delete Employee
	if err := tx.Model(&domain.Employee{}).
		Where("id = ?", employeeID).
		Update("status", "Resigned").Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to update employee status before deletion", zap.Error(err))
		return err
	}

	if err := tx.Where("id = ?", employeeID).Delete(&domain.Employee{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to soft delete employee", zap.Error(err))
		return err
	}

	tx.Commit()
	r.Logger.Info("✅ Employee soft deleted", zap.Uint("employee_id", employeeID), zap.String("reason", reason))
	return nil
}

// ListEmployee 📌 Fetch employees with pagination & search
func (r *HrmsRepository) ListEmployees(ctx context.Context, limit, offset int, search string) ([]domain.Employee, int64, error) {
	var employees []domain.Employee
	var totalCount int64

	query := r.DB.WithContext(ctx).Model(&domain.Employee{}).
		Preload("Organization").
		Preload("Department").
		Preload("Designation").
		Order("created_at DESC") // Default sorting

	if search != "" {
		query = query.Where("email ILIKE ? OR first_name ILIKE ? OR last_name ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	query.Count(&totalCount)

	if err := query.Limit(limit).Offset(offset).Find(&employees).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch employees", zap.Error(err))
		return nil, 0, err
	}

	return employees, totalCount, nil
}
