package usecase

import (
	"context"
	"hrms/internal/dto"
	"hrms/internal/domain"
	"time"

	"go.uber.org/zap"
)

// CreateAttendance handles employee check-in
func (u *HrmsUsecase) CreateAttendance(ctx context.Context, req dto.CreateAttendanceRequest) (*dto.AttendanceDTO, error) {
	attendance := domain.Attendance{
		EmployeeID:  req.EmployeeID,
		Date:        req.Date,
		CheckIn:     &req.CheckIn,
		Location:    req.Location,
		IsRemote:    req.IsRemote,
		PunchMethod: req.PunchMethod,
	}

	result, err := u.HrmsRepo.CreateAttendance(ctx, &attendance)
	if err != nil {
		u.Logger.Error("Failed to create attendance", zap.Error(err))
		return nil, err
	}

	return mapToAttendanceDTO(result), nil
}

// CheckOutAttendance handles employee check-out
func (u *HrmsUsecase) CheckOutAttendance(ctx context.Context, req dto.CheckOutAttendanceRequest) error {
	err := u.HrmsRepo.CheckOutAttendance(ctx, req.EmployeeID, req.CheckOut)
	if err != nil {
		u.Logger.Error("Failed to process check-out", zap.Error(err))
		return err
	}
	return nil
}

// GetAttendanceByID fetches an attendance record by ID
func (u *HrmsUsecase) GetAttendanceByID(ctx context.Context, id uint) (*dto.AttendanceDTO, error) {
	attendance, err := u.HrmsRepo.GetAttendanceByID(ctx, id)
	if err != nil {
		u.Logger.Error("Failed to fetch attendance", zap.Error(err))
		return nil, err
	}

	return mapToAttendanceDTO(attendance), nil
}

// ListAttendances fetches multiple attendance records with optional filters
func (u *HrmsUsecase) ListAttendances(ctx context.Context, employeeID uint, startDate, endDate time.Time, isRemote *bool, limit, offset int) ([]dto.AttendanceDTO, int64, error) {
	attendances, total, err := u.HrmsRepo.ListAttendance(ctx, employeeID, startDate, endDate, isRemote, limit, offset)
	if err != nil {
		u.Logger.Error("Failed to list attendances", zap.Error(err))
		return nil, 0, err
	}

	dtoAttendances := make([]dto.AttendanceDTO, len(attendances))
	for i, a := range attendances {
		dtoAttendances[i] = *mapToAttendanceDTO(&a)
	}

	return dtoAttendances, total, nil
}

// DeleteAttendance removes an attendance record
func (u *HrmsUsecase) DeleteAttendance(ctx context.Context, id uint) error {
	if err := u.HrmsRepo.DeleteAttendance(ctx, id); err != nil {
		u.Logger.Error("Failed to delete attendance", zap.Error(err))
		return err
	}
	return nil
}

// mapToAttendanceDTO converts a domain model to a DTO
func mapToAttendanceDTO(a *domain.Attendance) *dto.AttendanceDTO {
	return &dto.AttendanceDTO{
		ID:          a.ID,
		EmployeeID:  a.EmployeeID,
		Date:        a.Date,
		CheckIn:     a.CheckIn,
		CheckOut:    a.CheckOut,
		WorkHours:   a.WorkHours,
		Overtime:    a.Overtime,
		BreakTime:   a.BreakTime,
		Location:    a.Location,
		IsRemote:    a.IsRemote,
		PunchMethod: a.PunchMethod,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}
