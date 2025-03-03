package repository

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreateEmployeeExit 📌 Process an employee exit request
func (r *HrmsRepository) CreateEmployeeExit(ctx context.Context, exit *domain.EmployeeExit) (*domain.EmployeeExit, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Ensure the employee exists
	var employee domain.Employee
	if err := tx.Where("id = ?", exit.EmployeeID).First(&employee).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, hrmsErrors.ErrEmployeeNotFound
		}
		tx.Rollback()
		r.Logger.Error("❌ Failed to find employee for exit process", zap.Error(err))
		return nil, err
	}

	// ✅ Step 2: Insert Employee Exit Record
	if err := tx.Create(exit).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create employee exit record", zap.Error(err))
		return nil, err
	}

	// ✅ Step 3: Update Employee Status to "Inactive"
	if err := tx.Model(&domain.Employee{}).
		Where("id = ?", exit.EmployeeID).
		Update("status", "Inactive").Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to update employee status", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return exit, nil
}

// GetEmployeeExitByID 📌 Fetch an employee exit record by ID
func (r *HrmsRepository) GetEmployeeExitByID(ctx context.Context, exitID uint) (*domain.EmployeeExit, error) {
	var exitRecord domain.EmployeeExit
	err := r.DB.WithContext(ctx).Where("id = ?", exitID).First(&exitRecord).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrEmployeeExitNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetEmployeeExitByID", zap.Error(err))
		return nil, err
	}
	return &exitRecord, nil
}

// GetExitRecordsByEmployee 📌 Fetch all exit records for a specific employee
func (r *HrmsRepository) GetExitRecordsByEmployee(ctx context.Context, employeeID uint) ([]domain.EmployeeExit, error) {
	var exits []domain.EmployeeExit
	err := r.DB.WithContext(ctx).Where("employee_id = ?", employeeID).Find(&exits).Error
	if err != nil {
		r.Logger.Error("❌ Failed to fetch employee exit records", zap.Uint("employee_id", employeeID), zap.Error(err))
		return nil, err
	}
	return exits, nil
}

// UpdateClearanceStatus 📌 Update clearance status for an exit record
func (r *HrmsRepository) UpdateClearanceStatus(ctx context.Context, exitID uint, status string) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Verify the exit record exists
	var exit domain.EmployeeExit
	if err := tx.Where("id = ?", exitID).First(&exit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return hrmsErrors.ErrEmployeeExitNotFound
		}
		tx.Rollback()
		return err
	}

	// ✅ Step 2: Update Clearance Status
	if err := tx.Model(&domain.EmployeeExit{}).
		Where("id = ?", exitID).
		Update("clearance_status", status).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to update clearance status", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// DeleteEmployeeExit 📌 Delete an employee exit record (Soft Delete)
func (r *HrmsRepository) DeleteEmployeeExit(ctx context.Context, exitID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Where("id = ?", exitID).Delete(&domain.EmployeeExit{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete employee exit record", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// GetPendingClearances 📌 Fetch all employees with pending clearance
func (r *HrmsRepository) GetPendingClearances(ctx context.Context) ([]domain.EmployeeExit, error) {
	var pendingClearances []domain.EmployeeExit
	err := r.DB.WithContext(ctx).Where("clearance_status = ?", "Pending").Find(&pendingClearances).Error
	if err != nil {
		r.Logger.Error("❌ Failed to fetch pending clearances", zap.Error(err))
		return nil, err
	}
	return pendingClearances, nil
}
