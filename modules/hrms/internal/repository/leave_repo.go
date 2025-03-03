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

// CreateLeave 📌 Creates a new leave request
func (r *HrmsRepository) CreateLeave(ctx context.Context, leave *domain.Leave) (*domain.Leave, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Ensure employee exists
	var employee domain.Employee
	if err := tx.Where("id = ?", leave.EmployeeID).First(&employee).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, hrmsErrors.ErrLeaveNotFound
		}
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 2: Validate leave dates
	if leave.EndDate.Before(leave.StartDate) {
		tx.Rollback()
		return nil, hrmsErrors.ErrInvalidLeaveDates
	}

	// ✅ Step 3: Check for conflicting approved leaves
	var existingLeave domain.Leave
	if err := tx.Where("employee_id = ? AND status = 'Approved' AND start_date <= ? AND end_date >= ?",
		leave.EmployeeID, leave.EndDate, leave.StartDate).First(&existingLeave).Error; err == nil {
		tx.Rollback()
		return nil, hrmsErrors.ErrLeaveConflict
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 4: Insert leave request
	leave.Status = "Pending" // Default status
	if err := tx.Create(leave).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create leave request", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return leave, nil
}

// GetLeaveByID 📌 Fetch leave request by ID
func (r *HrmsRepository) GetLeaveByID(ctx context.Context, leaveID uint) (*domain.Leave, error) {
	var leave domain.Leave
	err := r.DB.WithContext(ctx).Where("id = ?", leaveID).First(&leave).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrLeaveNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetLeaveByID", zap.Error(err))
		return nil, err
	}
	return &leave, nil
}

// UpdateLeaveStatus 📌 Approve/Reject a leave request
func (r *HrmsRepository) UpdateLeaveStatus(ctx context.Context, leaveID uint, approverID uint, status string, comments string) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Fetch leave request
	var leave domain.Leave
	err := tx.Where("id = ?", leaveID).First(&leave).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return hrmsErrors.ErrLeaveNotFound
	} else if err != nil {
		tx.Rollback()
		return err
	}

	// ✅ Step 2: Validate if the approver has the required clearance
	if leave.ApproverID != approverID {
		tx.Rollback()
		return hrmsErrors.ErrUnauthorizedAction
	}

	// ✅ Step 3: Update leave status
	leave.Status = status
	leave.Comments = comments

	if err := tx.Save(&leave).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to update leave status", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// ListLeaves 📌 Fetch leave records with filters (date range, status, employee)
func (r *HrmsRepository) ListLeaves(ctx context.Context, employeeID *uint, status *string, startDate, endDate time.Time, limit, offset int) ([]domain.Leave, int64, error) {
	var leaves []domain.Leave
	var totalCount int64

	query := r.DB.WithContext(ctx).Model(&domain.Leave{}).
		Where("start_date >= ? AND end_date <= ?", startDate, endDate)

	if employeeID != nil {
		query = query.Where("employee_id = ?", *employeeID)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	query.Count(&totalCount)

	if err := query.Limit(limit).Offset(offset).Find(&leaves).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch leave records", zap.Error(err))
		return nil, 0, err
	}

	return leaves, totalCount, nil
}

// DeleteLeave 📌 Soft delete a leave request
func (r *HrmsRepository) DeleteLeave(ctx context.Context, leaveID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Where("id = ?", leaveID).Delete(&domain.Leave{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete leave request", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}
