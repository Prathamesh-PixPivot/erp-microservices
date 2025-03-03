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

// CreateLoanAdvance 📌 Employee requests a loan/advance salary
func (r *HrmsRepository) CreateLoanAdvance(ctx context.Context, loan *domain.LoanAdvance) (*domain.LoanAdvance, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Ensure the employee exists
	var employee domain.Employee
	if err := tx.Where("id = ?", loan.EmployeeID).First(&employee).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, hrmsErrors.ErrEmployeeNotFound
		}
		tx.Rollback()
		r.Logger.Error("❌ Failed to find employee for loan request", zap.Error(err))
		return nil, err
	}

	// ✅ Step 2: Insert LoanAdvance Request
	loan.Status = "Pending" // Default status
	if err := tx.Create(loan).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create loan request", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return loan, nil
}

// GetLoanAdvanceByID 📌 Fetch a loan/advance request by ID
func (r *HrmsRepository) GetLoanAdvanceByID(ctx context.Context, loanID uint) (*domain.LoanAdvance, error) {
	var loan domain.LoanAdvance
	err := r.DB.WithContext(ctx).Preload("Employee").Preload("Approver").Where("id = ?", loanID).First(&loan).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrLoanNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetLoanAdvanceByID", zap.Error(err))
		return nil, err
	}
	return &loan, nil
}

// ListLoanAdvances 📌 Fetch all loan/advance requests (with optional filters)
func (r *HrmsRepository) ListLoanAdvances(ctx context.Context, status string, employeeID *uint) ([]domain.LoanAdvance, error) {
	var loans []domain.LoanAdvance
	query := r.DB.WithContext(ctx).Model(&domain.LoanAdvance{}).Preload("Employee").Preload("Approver")

	if status != "" {
		query = query.Where("status = ?", status)
	}
	if employeeID != nil {
		query = query.Where("employee_id = ?", *employeeID)
	}

	if err := query.Find(&loans).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch loan/advance requests", zap.Error(err))
		return nil, err
	}
	return loans, nil
}

// ApproveLoanAdvance 📌 Approves a loan/advance request
func (r *HrmsRepository) ApproveLoanAdvance(ctx context.Context, loanID, approverID uint, approvalDate, repaymentStart time.Time) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Verify the loan request exists
	var loan domain.LoanAdvance
	if err := tx.Where("id = ?", loanID).First(&loan).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return hrmsErrors.ErrLoanNotFound
		}
		tx.Rollback()
		return err
	}

	// ✅ Step 2: Ensure the loan is still pending
	if loan.Status != "Pending" {
		tx.Rollback()
		return hrmsErrors.ErrLoanAlreadyProcessed
	}

	// ✅ Step 3: Update the loan status to "Approved"
	loan.Status = "Approved"
	loan.ApprovedBy = approverID
	loan.ApprovalDate = &approvalDate
	loan.RepaymentStart = &repaymentStart

	if err := tx.Save(&loan).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to approve loan", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// RejectLoanAdvance 📌 Rejects a loan/advance request
func (r *HrmsRepository) RejectLoanAdvance(ctx context.Context, loanID, approverID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Verify the loan request exists
	var loan domain.LoanAdvance
	if err := tx.Where("id = ?", loanID).First(&loan).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return hrmsErrors.ErrLoanNotFound
		}
		tx.Rollback()
		return err
	}

	// ✅ Step 2: Ensure the loan is still pending
	if loan.Status != "Pending" {
		tx.Rollback()
		return hrmsErrors.ErrLoanAlreadyProcessed
	}

	// ✅ Step 3: Update the loan status to "Rejected"
	loan.Status = "Rejected"
	loan.ApprovedBy = approverID

	if err := tx.Save(&loan).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to reject loan", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// DeleteLoanAdvance 📌 Delete a loan/advance request (Soft Delete)
func (r *HrmsRepository) DeleteLoanAdvance(ctx context.Context, loanID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Where("id = ?", loanID).Delete(&domain.LoanAdvance{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete loan request", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}
