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

// CreatePayroll 📌 Adds a new payroll record (Prevents duplicate payroll entries)
func (r *HrmsRepository) CreatePayroll(ctx context.Context, payroll *domain.Payroll) (*domain.Payroll, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Check if payroll already exists for the given employee & month
	var existingPayroll domain.Payroll
	firstDayOfMonth := time.Date(payroll.PaymentDate.Year(), payroll.PaymentDate.Month(), 1, 0, 0, 0, 0, time.UTC)

	if err := tx.Where("employee_id = ? AND payment_date >= ? AND payment_date < ?", 
		payroll.EmployeeID, firstDayOfMonth, firstDayOfMonth.AddDate(0, 1, 0)).
		First(&existingPayroll).Error; err == nil {
		tx.Rollback()
		return nil, hrmsErrors.ErrPayrollAlreadyProcessed
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 2: Validate payroll data (Ensure net salary calculation is correct)
	if payroll.NetSalary != (payroll.Salary + payroll.Allowances - payroll.Tax - payroll.Deductions) {
		tx.Rollback()
		return nil, hrmsErrors.ErrInvalidPayrollData
	}

	// ✅ Step 3: Insert the new payroll record
	if err := tx.Create(payroll).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create payroll", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return payroll, nil
}

// GetPayroll 📌 Fetch payroll record by ID
func (r *HrmsRepository) GetPayroll(ctx context.Context, payrollID uint) (*domain.Payroll, error) {
	var payroll domain.Payroll
	err := r.DB.WithContext(ctx).Where("id = ?", payrollID).First(&payroll).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrPayrollNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetPayroll", zap.Error(err))
		return nil, err
	}

	return &payroll, nil
}

// ListPayrolls 📌 Fetch all payroll records (with optional filters)
func (r *HrmsRepository) ListPayrolls(ctx context.Context, employeeID uint, month time.Time) ([]domain.Payroll, error) {
	var payrolls []domain.Payroll
	query := r.DB.WithContext(ctx).Model(&domain.Payroll{})

	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
	}

	if !month.IsZero() {
		firstDayOfMonth := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.UTC)
		query = query.Where("payment_date >= ? AND payment_date < ?", firstDayOfMonth, firstDayOfMonth.AddDate(0, 1, 0))
	}

	if err := query.Order("payment_date DESC").Find(&payrolls).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch payroll records", zap.Error(err))
		return nil, err
	}

	return payrolls, nil
}

// UpdatePayroll 📌 Update payroll details (e.g., status, payslip URL)
func (r *HrmsRepository) UpdatePayroll(ctx context.Context, payrollID uint, updates map[string]interface{}) error {
	tx := r.DB.WithContext(ctx).Begin()

	// Prevent modification if payroll is already processed
	if status, ok := updates["status"].(string); ok && status == "Processed" {
		var payroll domain.Payroll
		if err := tx.Where("id = ?", payrollID).First(&payroll).Error; err == nil && payroll.Status == "Processed" {
			tx.Rollback()
			return hrmsErrors.ErrPayrollAlreadyProcessed
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return err
		}
	}

	// ✅ Update payroll details
	if err := tx.Model(&domain.Payroll{}).Where("id = ?", payrollID).Updates(updates).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to update payroll", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// DeletePayroll 📌 Remove a payroll record (Only if not processed)
func (r *HrmsRepository) DeletePayroll(ctx context.Context, payrollID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	var payroll domain.Payroll
	if err := tx.Where("id = ?", payrollID).First(&payroll).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return hrmsErrors.ErrPayrollNotFound
		}
		r.Logger.Error("❌ Error finding payroll for deletion", zap.Error(err))
		return err
	}

	// Prevent deleting processed payroll
	if payroll.Status == "Processed" {
		tx.Rollback()
		return hrmsErrors.ErrPayrollAlreadyProcessed
	}

	if err := tx.Delete(&domain.Payroll{}, payrollID).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete payroll", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}
