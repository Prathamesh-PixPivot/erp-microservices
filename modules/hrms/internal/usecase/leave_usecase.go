package usecase

import (
	"context"
	"hrms/internal/domain"
	"hrms/internal/dto"
	"hrms/internal/errors"
	"time"

	"go.uber.org/zap"
)

// CreateLeave handles the leave request creation
func (u *HrmsUsecase) CreateLeave(ctx context.Context, req dto.CreateLeaveRequest) (*dto.LeaveResponse, error) {
	// Validate input
	if req.EmployeeID == 0 || req.StartDate.IsZero() || req.EndDate.IsZero() || req.StartDate.After(req.EndDate) {
		return nil, errors.ErrInvalidLeaveDates
	}

	// Create leave entity
	leave := &domain.Leave{
		EmployeeID: req.EmployeeID,
		LeaveType:  req.LeaveType,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Status:     "Pending",
	}

	// Persist data
	createdLeave, err := u.HrmsRepo.CreateLeave(ctx, leave)
	if err != nil {
		u.Logger.Error("❌ Failed to create leave request", zap.Error(err))
		return nil, err
	}

	// Response mapping
	return &dto.LeaveResponse{
		ID:         createdLeave.ID,
		EmployeeID: createdLeave.EmployeeID,
		LeaveType:  createdLeave.LeaveType,
		StartDate:  createdLeave.StartDate,
		EndDate:    createdLeave.EndDate,
		Status:     createdLeave.Status,
		ApproverID: createdLeave.ApproverID,
		Comments:   createdLeave.Comments,
	}, nil
}

// GetLeaveByID fetches a leave request by ID
func (u *HrmsUsecase) GetLeaveByID(ctx context.Context, leaveID uint) (*dto.LeaveResponse, error) {
	// Fetch leave request from repository
	leave, err := u.HrmsRepo.GetLeaveByID(ctx, leaveID)
	if err != nil {
		u.Logger.Error("❌ Failed to fetch leave request", zap.Error(err))
		return nil, err
	}

	// Response mapping
	return &dto.LeaveResponse{
		ID:         leave.ID,
		EmployeeID: leave.EmployeeID,
		LeaveType:  leave.LeaveType,
		StartDate:  leave.StartDate,
		EndDate:    leave.EndDate,
		Status:     leave.Status,
		ApproverID: leave.ApproverID,
		Comments:   leave.Comments,
	}, nil
}

// UpdateLeaveStatus updates the status of a leave request
func (u *HrmsUsecase) UpdateLeaveStatus(ctx context.Context, leaveID uint, req dto.UpdateLeaveStatusRequest) error {
	// Validate input
	if req.Status != "Approved" && req.Status != "Rejected" {
		return errors.ErrInvalidLeaveStatus
	}

	// Update status in repository
	err := u.HrmsRepo.UpdateLeaveStatus(ctx, leaveID, req.ApproverID, req.Status, req.Comments)
	if err != nil {
		u.Logger.Error("❌ Failed to update leave status", zap.Error(err))
		return err
	}

	return nil
}

// ListLeaves fetches leave records with optional filters
func (u *HrmsUsecase) ListLeaves(ctx context.Context, employeeID *uint, status *string, startDate, endDate time.Time, limit, offset int) ([]dto.LeaveResponse, error) {
	// Fetch leave records from repository
	leaves, _, err := u.HrmsRepo.ListLeaves(ctx, employeeID, status, startDate, endDate, limit, offset)
	if err != nil {
		u.Logger.Error("❌ Failed to list leave records", zap.Error(err))
		return nil, err
	}

	// Convert domain models to DTOs
	var response []dto.LeaveResponse
	for _, leave := range leaves {
		response = append(response, dto.LeaveResponse{
			ID:         leave.ID,
			EmployeeID: leave.EmployeeID,
			LeaveType:  leave.LeaveType,
			StartDate:  leave.StartDate,
			EndDate:    leave.EndDate,
			Status:     leave.Status,
			ApproverID: leave.ApproverID,
			Comments:   leave.Comments,
		})
	}

	return response, nil
}

// DeleteLeave handles soft deletion of a leave request
func (u *HrmsUsecase) DeleteLeave(ctx context.Context, leaveID uint) error {
	// Validate input
	if leaveID == 0 {
		return errors.ErrInvalidLeaveID
	}

	// Delete leave record from repository
	err := u.HrmsRepo.DeleteLeave(ctx, leaveID)
	if err != nil {
		u.Logger.Error("❌ Failed to delete leave request", zap.Error(err))
		return err
	}

	return nil
}
