package usecase

import (
	"context"
	"time"

	"go.uber.org/zap"
	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreateExpense handles the submission of an expense reimbursement request
func (u *HrmsUsecase) CreateExpense(ctx context.Context, expense *domain.Expense) (*domain.Expense, error) {
	u.Logger.Info("Creating expense request", zap.Uint("employee_id", expense.EmployeeID))

	// Validate Required Fields
	if expense.EmployeeID == 0 || expense.ExpenseType == "" || expense.Amount <= 0 {
		return nil, hrmsErrors.ErrInvalidExpense
	}

	// Set Default Values
	expense.Status = "Pending"
	expense.Date = time.Now()

	// Persist to Database
	savedExpense, err := u.HrmsRepo.CreateExpense(ctx, expense)
	if err != nil {
		u.Logger.Error("Failed to create expense", zap.Error(err))
		return nil, err
	}

	u.Logger.Info("Expense request created successfully", zap.Uint("expense_id", savedExpense.ID))
	return savedExpense, nil
}

// GetExpense retrieves a specific expense record
func (u *HrmsUsecase) GetExpense(ctx context.Context, expenseID uint) (*domain.Expense, error) {
	expense, err := u.HrmsRepo.GetExpense(ctx, expenseID)
	if err != nil {
		u.Logger.Error("Failed to retrieve expense", zap.Uint("expense_id", expenseID), zap.Error(err))
		return nil, err
	}
	return expense, nil
}

// GetEmployeeExpenses retrieves all expenses for a given employee
func (u *HrmsUsecase) GetEmployeeExpenses(ctx context.Context, employeeID uint) ([]domain.Expense, error) {
	expenses, err := u.HrmsRepo.GetEmployeeExpenses(ctx, employeeID)
	if err != nil {
		u.Logger.Error("Failed to retrieve expenses", zap.Uint("employee_id", employeeID), zap.Error(err))
		return nil, err
	}
	return expenses, nil
}

// UpdateExpenseStatus handles approval or rejection of an expense
func (u *HrmsUsecase) UpdateExpenseStatus(ctx context.Context, expenseID uint, approverID uint, newStatus string) error {
	if err := u.HrmsRepo.UpdateExpenseStatus(ctx, expenseID, approverID, newStatus); err != nil {
		u.Logger.Error("Failed to update expense status", zap.Uint("expense_id", expenseID), zap.Error(err))
		return err
	}
	return nil
}

// DeleteExpense allows an employee to delete a pending expense
func (u *HrmsUsecase) DeleteExpense(ctx context.Context, expenseID uint, employeeID uint) error {
	if err := u.HrmsRepo.DeleteExpense(ctx, expenseID, employeeID); err != nil {
		u.Logger.Error("Failed to delete expense", zap.Uint("expense_id", expenseID), zap.Error(err))
		return err
	}
	return nil
}
