package usecase

import (
	"context"
	"hrms/internal/dto"
	"hrms/internal/domain"

	"go.uber.org/zap"
)

// CreateEmployeeExit processes an employee exit request
func (u *HrmsUsecase) CreateEmployeeExit(ctx context.Context, req dto.CreateEmployeeExitRequest) (*dto.EmployeeExitDTO, error) {
	exit := domain.EmployeeExit{
		EmployeeID:      req.EmployeeID,
		ExitType:        req.ExitType,
		ExitDate:        req.ExitDate,
		ClearanceStatus: req.ClearanceStatus,
	}

	result, err := u.HrmsRepo.CreateEmployeeExit(ctx, &exit)
	if err != nil {
		u.Logger.Error("Failed to process employee exit", zap.Error(err))
		return nil, err
	}

	return mapToEmployeeExitDTO(result), nil
}

// GetEmployeeExitByID fetches a specific exit record by ID
func (u *HrmsUsecase) GetEmployeeExitByID(ctx context.Context, exitID uint) (*dto.EmployeeExitDTO, error) {
	exitRecord, err := u.HrmsRepo.GetEmployeeExitByID(ctx, exitID)
	if err != nil {
		u.Logger.Error("Failed to fetch employee exit record", zap.Error(err))
		return nil, err
	}

	return mapToEmployeeExitDTO(exitRecord), nil
}

// GetExitRecordsByEmployee fetches all exit records for a specific employee
func (u *HrmsUsecase) GetExitRecordsByEmployee(ctx context.Context, employeeID uint) ([]dto.EmployeeExitDTO, error) {
	exitRecords, err := u.HrmsRepo.GetExitRecordsByEmployee(ctx, employeeID)
	if err != nil {
		u.Logger.Error("Failed to fetch employee exit records", zap.Error(err))
		return nil, err
	}

	dtoExits := make([]dto.EmployeeExitDTO, len(exitRecords))
	for i, exit := range exitRecords {
		dtoExits[i] = *mapToEmployeeExitDTO(&exit)
	}

	return dtoExits, nil
}

// UpdateClearanceStatus updates the clearance status of an employee exit record
func (u *HrmsUsecase) UpdateClearanceStatus(ctx context.Context, exitID uint, req dto.UpdateClearanceStatusRequest) error {
	err := u.HrmsRepo.UpdateClearanceStatus(ctx, exitID, req.ClearanceStatus)
	if err != nil {
		u.Logger.Error("Failed to update clearance status", zap.Error(err))
		return err
	}
	return nil
}

// DeleteEmployeeExit removes a specific employee exit record
func (u *HrmsUsecase) DeleteEmployeeExit(ctx context.Context, exitID uint) error {
	if err := u.HrmsRepo.DeleteEmployeeExit(ctx, exitID); err != nil {
		u.Logger.Error("Failed to delete employee exit record", zap.Error(err))
		return err
	}
	return nil
}

// GetPendingClearances fetches all pending employee exit clearances
func (u *HrmsUsecase) GetPendingClearances(ctx context.Context) ([]dto.EmployeeExitDTO, error) {
	pendingClearances, err := u.HrmsRepo.GetPendingClearances(ctx)
	if err != nil {
		u.Logger.Error("Failed to fetch pending clearances", zap.Error(err))
		return nil, err
	}

	dtoExits := make([]dto.EmployeeExitDTO, len(pendingClearances))
	for i, exit := range pendingClearances {
		dtoExits[i] = *mapToEmployeeExitDTO(&exit)
	}

	return dtoExits, nil
}

// mapToEmployeeExitDTO converts domain model to DTO
func mapToEmployeeExitDTO(e *domain.EmployeeExit) *dto.EmployeeExitDTO {
	return &dto.EmployeeExitDTO{
		ID:              e.ID,
		EmployeeID:      e.EmployeeID,
		ExitType:        e.ExitType,
		ExitDate:        e.ExitDate,
		ClearanceStatus: e.ClearanceStatus,
	}
}
