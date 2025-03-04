package repository

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreateLeaveBalance 📌 Creates or initializes leave balance for an employee
func (r *HrmsRepository) CreateLeaveBalance(ctx context.Context, balance *domain.LeaveBalance) (*domain.LeaveBalance, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Check if balance already exists
	var existingBalance domain.LeaveBalance
	if err := tx.Where("employee_id = ? AND leave_type = ?", balance.EmployeeID, balance.LeaveType).First(&existingBalance).Error; err == nil {
		tx.Rollback()
		return nil, hrmsErrors.ErrInvalidLeaveType
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 2: Insert leave balance record
	balance.Remaining = balance.TotalLeaves
	if err := tx.Create(balance).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create leave balance", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return balance, nil
}

// GetLeaveBalance 📌 Fetch leave balance for an employee
func (r *HrmsRepository) GetLeaveBalance(ctx context.Context, employeeID uint, leaveType domain.LeaveType) (*domain.LeaveBalance, error) {
	var balance domain.LeaveBalance
	err := r.DB.WithContext(ctx).
		Where("employee_id = ? AND leave_type = ?", employeeID, leaveType).
		First(&balance).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrLeaveBalanceNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetLeaveBalance", zap.Error(err))
		return nil, err
	}
	return &balance, nil
}

// DeductLeaveBalance 📌 Deducts leave balance on approval
func (r *HrmsRepository) DeductLeaveBalance(ctx context.Context, employeeID uint, leaveType domain.LeaveType, leaveDays float64) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Fetch existing leave balance
	var balance domain.LeaveBalance
	err := tx.Where("employee_id = ? AND leave_type = ?", employeeID, leaveType).First(&balance).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return hrmsErrors.ErrLeaveBalanceNotFound
	} else if err != nil {
		tx.Rollback()
		return err
	}

	// ✅ Step 2: Ensure sufficient leave balance
	if balance.Remaining < leaveDays {
		tx.Rollback()
		return hrmsErrors.ErrInsufficientLeave
	}

	// ✅ Step 3: Deduct leave
	balance.UsedLeaves += leaveDays
	balance.Remaining -= leaveDays

	if err := tx.Save(&balance).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to deduct leave balance", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// RestoreLeaveBalance 📌 Restores leave balance on leave rejection/cancellation
func (r *HrmsRepository) RestoreLeaveBalance(ctx context.Context, employeeID uint, leaveType domain.LeaveType, leaveDays float64) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Fetch existing leave balance
	var balance domain.LeaveBalance
	err := tx.Where("employee_id = ? AND leave_type = ?", employeeID, leaveType).First(&balance).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return hrmsErrors.ErrLeaveBalanceNotFound
	} else if err != nil {
		tx.Rollback()
		return err
	}

	// ✅ Step 2: Restore leave
	balance.UsedLeaves -= leaveDays
	balance.Remaining += leaveDays

	if err := tx.Save(&balance).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to restore leave balance", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// ListLeaveBalances 📌 Fetch leave balances for all employees (with filters)
func (r *HrmsRepository) ListLeaveBalances(ctx context.Context, employeeID *uint, limit, offset int) ([]domain.LeaveBalance, int64, error) {
	var balances []domain.LeaveBalance
	var totalCount int64

	query := r.DB.WithContext(ctx).Model(&domain.LeaveBalance{})

	if employeeID != nil {
		query = query.Where("employee_id = ?", *employeeID)
	}

	query.Count(&totalCount)

	if err := query.Limit(limit).Offset(offset).Find(&balances).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch leave balances", zap.Error(err))
		return nil, 0, err
	}

	return balances, totalCount, nil
}

// DeleteLeaveBalance 📌 Deletes leave balance for an employee (use cautiously)
func (r *HrmsRepository) DeleteLeaveBalance(ctx context.Context, employeeID uint, leaveType domain.LeaveType) error {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Where("employee_id = ? AND leave_type = ?", employeeID, leaveType).Delete(&domain.LeaveBalance{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete leave balance", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}
