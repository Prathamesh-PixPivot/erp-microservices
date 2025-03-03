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

// CreateAttendance 📌 Adds a new attendance record (check-in)
func (r *HrmsRepository) CreateAttendance(ctx context.Context, attendance *domain.Attendance) (*domain.Attendance, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Ensure employee exists
	var employee domain.Employee
	if err := tx.Where("id = ?", attendance.EmployeeID).First(&employee).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, hrmsErrors.ErrInvalidAttendance
		}
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 2: Prevent multiple check-ins on the same day
	var existingAttendance domain.Attendance
	if err := tx.Where("employee_id = ? AND date = ?", attendance.EmployeeID, attendance.Date).First(&existingAttendance).Error; err == nil {
		tx.Rollback()
		return nil, hrmsErrors.ErrDuplicateCheckIn
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 3: Insert attendance record
	if err := tx.Create(attendance).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create attendance record", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return attendance, nil
}

// GetAttendanceByID 📌 Fetch attendance record by ID
func (r *HrmsRepository) GetAttendanceByID(ctx context.Context, attendanceID uint) (*domain.Attendance, error) {
	var attendance domain.Attendance
	err := r.DB.WithContext(ctx).Where("id = ?", attendanceID).First(&attendance).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrAttendanceNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetAttendanceByID", zap.Error(err))
		return nil, err
	}
	return &attendance, nil
}

// CheckOutAttendance 📌 Updates checkout time, work hours & overtime
func (r *HrmsRepository) CheckOutAttendance(ctx context.Context, employeeID uint, checkOutTime time.Time) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Find the latest attendance record for today
	var attendance domain.Attendance
	err := tx.Where("employee_id = ? AND date = ?", employeeID, checkOutTime.Format("2006-01-02")).
		First(&attendance).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return hrmsErrors.ErrInvalidCheckOut
	} else if err != nil {
		tx.Rollback()
		return err
	}

	// ✅ Step 2: Validate that check-in exists before checking out
	if attendance.CheckIn == nil {
		tx.Rollback()
		return hrmsErrors.ErrInvalidCheckOut
	}

	// ✅ Step 3: Calculate work hours
	workDuration := checkOutTime.Sub(*attendance.CheckIn).Hours()
	overtime := 0.0
	if workDuration > 8 {
		overtime = workDuration - 8
	}

	// ✅ Step 4: Update the attendance record
	updateData := map[string]interface{}{
		"check_out": checkOutTime,
		"work_hours": workDuration,
		"overtime": overtime,
	}

	if err := tx.Model(&attendance).Updates(updateData).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to update checkout in attendance", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// ListAttendance 📌 Fetch attendance records with filters (date range, employee, remote)
func (r *HrmsRepository) ListAttendance(ctx context.Context, employeeID uint, startDate, endDate time.Time, isRemote *bool, limit, offset int) ([]domain.Attendance, int64, error) {
	var attendances []domain.Attendance
	var totalCount int64

	query := r.DB.WithContext(ctx).Model(&domain.Attendance{}).
		Where("employee_id = ? AND date BETWEEN ? AND ?", employeeID, startDate, endDate)

	if isRemote != nil {
		query = query.Where("is_remote = ?", *isRemote)
	}

	query.Count(&totalCount)

	if err := query.Limit(limit).Offset(offset).Find(&attendances).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch attendance records", zap.Error(err))
		return nil, 0, err
	}

	return attendances, totalCount, nil
}

// DeleteAttendance 📌 Soft delete an attendance record
func (r *HrmsRepository) DeleteAttendance(ctx context.Context, attendanceID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Where("id = ?", attendanceID).Delete(&domain.Attendance{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete attendance record", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}
