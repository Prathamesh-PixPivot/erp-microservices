package usecase

import (
	"context"
	"hrms/internal/domain"
	"hrms/internal/dto"
	"hrms/internal/errors"

	"go.uber.org/zap"
)

// CreateLeaveBalance handles leave balance creation
func (u *HrmsUsecase) CreateLeaveBalance(ctx context.Context, req dto.CreateLeaveBalanceRequest) (*dto.LeaveBalanceResponse, error) {
	// Validate input
	if req.EmployeeID == 0 || req.TotalLeaves < 0 {
		return nil, errors.ErrInvalidLeaveType
	}

	// Create leave balance entity
	balance := &domain.LeaveBalance{
		EmployeeID:  req.EmployeeID,
		LeaveType:   req.LeaveType,
		TotalLeaves: req.TotalLeaves,
		UsedLeaves:  0,
		Remaining:   req.TotalLeaves,
	}

	// Persist data
	createdBalance, err := u.HrmsRepo.CreateLeaveBalance(ctx, balance)
	if err != nil {
		u.Logger.Error("❌ Failed to create leave balance", zap.Error(err))
		return nil, err
	}

	// Response mapping
	return &dto.LeaveBalanceResponse{
		EmployeeID:  createdBalance.EmployeeID,
		LeaveType:   createdBalance.LeaveType,
		TotalLeaves: createdBalance.TotalLeaves,
		UsedLeaves:  createdBalance.UsedLeaves,
		Remaining:   createdBalance.Remaining,
	}, nil
}

// DeductLeaveBalance handles deduction of leave balance on approval
func (u *HrmsUsecase) DeductLeaveBalance(ctx context.Context, req dto.DeductLeaveBalanceRequest) error {
	// Validate input
	if req.EmployeeID == 0 || req.LeaveDays <= 0 {
		return errors.ErrInvalidLeaveType
	}

	// Deduct balance in repository
	err := u.HrmsRepo.DeductLeaveBalance(ctx, req.EmployeeID, req.LeaveType, req.LeaveDays)
	if err != nil {
		u.Logger.Error("❌ Failed to deduct leave balance", zap.Error(err))
		return err
	}

	return nil
}

// RestoreLeaveBalance handles restoration of leave balance on cancellation
func (u *HrmsUsecase) RestoreLeaveBalance(ctx context.Context, req dto.RestoreLeaveBalanceRequest) error {
	// Validate input
	if req.EmployeeID == 0 || req.LeaveDays <= 0 {
		return errors.ErrInvalidLeaveType
	}

	// Restore balance in repository
	err := u.HrmsRepo.RestoreLeaveBalance(ctx, req.EmployeeID, req.LeaveType, req.LeaveDays)
	if err != nil {
		u.Logger.Error("❌ Failed to restore leave balance", zap.Error(err))
		return err
	}

	return nil
}

// GetLeaveBalance fetches leave balance for an employee
func (u *HrmsUsecase) GetLeaveBalance(ctx context.Context, employeeID uint, leaveType domain.LeaveType) (*dto.LeaveBalanceResponse, error) {
	// Fetch balance from repository
	balance, err := u.HrmsRepo.GetLeaveBalance(ctx, employeeID, leaveType)
	if err != nil {
		u.Logger.Error("❌ Failed to fetch leave balance", zap.Error(err))
		return nil, err
	}

	// Response mapping
	return &dto.LeaveBalanceResponse{
		EmployeeID:  balance.EmployeeID,
		LeaveType:   balance.LeaveType,
		TotalLeaves: balance.TotalLeaves,
		UsedLeaves:  balance.UsedLeaves,
		Remaining:   balance.Remaining,
	}, nil
}

// ListLeaveBalances fetches leave balances with filters
func (u *HrmsUsecase) ListLeaveBalances(ctx context.Context, employeeID *uint, limit, offset int) ([]dto.LeaveBalanceResponse, int64, error) {
	// Fetch balances from repository
	balances, total, err := u.HrmsRepo.ListLeaveBalances(ctx, employeeID, limit, offset)
	if err != nil {
		u.Logger.Error("❌ Failed to list leave balances", zap.Error(err))
		return nil, 0, err
	}

	// Convert domain models to DTOs
	var response []dto.LeaveBalanceResponse
	for _, balance := range balances {
		response = append(response, dto.LeaveBalanceResponse{
			EmployeeID:  balance.EmployeeID,
			LeaveType:   balance.LeaveType,
			TotalLeaves: balance.TotalLeaves,
			UsedLeaves:  balance.UsedLeaves,
			Remaining:   balance.Remaining,
		})
	}

	return response, total, nil
}

// DeleteLeaveBalance deletes an employee's leave balance (use cautiously)
func (u *HrmsUsecase) DeleteLeaveBalance(ctx context.Context, employeeID uint, leaveType domain.LeaveType) error {
	// Validate input
	if employeeID == 0 {
		return errors.ErrInvalidLeaveType
	}

	// Delete leave balance from repository
	err := u.HrmsRepo.DeleteLeaveBalance(ctx, employeeID, leaveType)
	if err != nil {
		u.Logger.Error("❌ Failed to delete leave balance", zap.Error(err))
		return err
	}

	return nil
}
