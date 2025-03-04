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

// CreateExpense 📌 Submits an expense reimbursement request
func (r *HrmsRepository) CreateExpense(ctx context.Context, expense *domain.Expense) (*domain.Expense, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Validate Required Fields
	if expense.EmployeeID == 0 || expense.ExpenseType == "" || expense.Amount <= 0 {
		tx.Rollback()
		return nil, hrmsErrors.ErrInvalidExpense
	}

	// ✅ Default Status: "Pending"
	expense.Status = "Pending"
	expense.Date = time.Now()

	// ✅ Insert Expense Record
	if err := tx.Create(expense).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create expense record", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return expense, nil
}

// GetExpense 📌 Fetches a specific expense record by ID
func (r *HrmsRepository) GetExpense(ctx context.Context, expenseID uint) (*domain.Expense, error) {
	var expense domain.Expense
	err := r.DB.WithContext(ctx).Where("id = ?", expenseID).First(&expense).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrExpenseNotFound
	} else if err != nil {
		r.Logger.Error("❌ Failed to fetch expense record", zap.Error(err))
		return nil, err
	}

	return &expense, nil
}

// GetEmployeeExpenses 📌 Fetches all expenses for an employee
func (r *HrmsRepository) GetEmployeeExpenses(ctx context.Context, employeeID uint) ([]domain.Expense, error) {
	var expenses []domain.Expense
	err := r.DB.WithContext(ctx).Where("employee_id = ?", employeeID).Find(&expenses).Error

	if err != nil {
		r.Logger.Error("❌ Failed to fetch expenses", zap.Error(err))
		return nil, err
	}

	return expenses, nil
}

// UpdateExpenseStatus 📌 Approves or Rejects an expense reimbursement request
func (r *HrmsRepository) UpdateExpenseStatus(ctx context.Context, expenseID uint, approverID uint, newStatus string) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Validate Expense Exists
	var expense domain.Expense
	if err := tx.Where("id = ?", expenseID).First(&expense).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return hrmsErrors.ErrExpenseNotFound
		}
		r.Logger.Error("❌ Error finding expense for approval", zap.Error(err))
		return err
	}

	// ✅ Validate Status Transition
	validStatuses := map[string]bool{"Pending": true, "Approved": true, "Rejected": true}
	if !validStatuses[newStatus] {
		tx.Rollback()
		return hrmsErrors.ErrInvalidExpense
	}

	// ✅ Update Status & Approver
	expense.Status = newStatus
	expense.ApproverID = approverID

	if err := tx.Save(&expense).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to update expense status", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// DeleteExpense 📌 Deletes an expense record (only if pending)
func (r *HrmsRepository) DeleteExpense(ctx context.Context, expenseID uint, employeeID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Validate Expense Exists
	var expense domain.Expense
	if err := tx.Where("id = ?", expenseID).First(&expense).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return hrmsErrors.ErrExpenseNotFound
		}
		r.Logger.Error("❌ Error finding expense for deletion", zap.Error(err))
		return err
	}

	// ✅ Ensure Employee Can Only Delete Pending Expenses
	if expense.EmployeeID != employeeID || expense.Status != "Pending" {
		tx.Rollback()
		return hrmsErrors.ErrUnauthorized
	}

	// ✅ Delete Expense
	if err := tx.Delete(&expense).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete expense", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}
