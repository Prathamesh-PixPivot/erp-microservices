package usecase

import (
	"context"

	"hrms/internal/domain"
	"hrms/internal/dto"
)

// CreateShift 📌 Adds a new shift.
func (uc *HrmsUsecase) CreateShift(ctx context.Context, shiftDTO dto.ShiftDTO) (*dto.ShiftDTO, error) {
	shift := &domain.Shift{
		Name:       shiftDTO.Name,
		ShiftType:  domain.ShiftType(shiftDTO.ShiftType), // Convert string to ShiftType enum
		StartTime:  shiftDTO.StartTime,
		EndTime:    shiftDTO.EndTime,
		DaysOfWeek: shiftDTO.DaysOfWeek,
	}

	createdShift, err := uc.HrmsRepo.CreateShift(ctx, shift)
	if err != nil {
		return nil, err
	}

	return uc.convertShiftDomainToDTO(createdShift), nil
}

// GetShiftByID 📌 Fetch a specific shift by its ID.
func (uc *HrmsUsecase) GetShiftByID(ctx context.Context, shiftID uint) (*dto.ShiftDTO, error) {
	shift, err := uc.HrmsRepo.GetShiftByID(ctx, shiftID)
	if err != nil {
		return nil, err
	}

	return uc.convertShiftDomainToDTO(shift), nil
}

// ListShifts 📌 Retrieves all shifts with optional search & pagination.
func (uc *HrmsUsecase) ListShifts(ctx context.Context, limit, offset int, search string) ([]dto.ShiftDTO, int64, error) {
	shifts, total, err := uc.HrmsRepo.ListShifts(ctx, limit, offset, search)
	if err != nil {
		return nil, 0, err
	}

	var shiftDTOs []dto.ShiftDTO
	for _, shift := range shifts {
		shiftDTOs = append(shiftDTOs, *uc.convertShiftDomainToDTO(&shift))
	}

	return shiftDTOs, total, nil
}

// UpdateShift 📌 Updates an existing shift.
func (uc *HrmsUsecase) UpdateShift(ctx context.Context, shiftID uint, updates map[string]interface{}) error {
	return uc.HrmsRepo.UpdateShift(ctx, shiftID, updates)
}

// DeleteShift 📌 Soft deletes a shift.
func (uc *HrmsUsecase) DeleteShift(ctx context.Context, shiftID uint) error {
	return uc.HrmsRepo.DeleteShift(ctx, shiftID)
}

// convertShiftDomainToDTO converts a Shift domain model to DTO.
func (uc *HrmsUsecase) convertShiftDomainToDTO(shift *domain.Shift) *dto.ShiftDTO {
	return &dto.ShiftDTO{
		ID:         shift.ID,
		Name:       shift.Name,
		ShiftType:  string(shift.ShiftType),
		StartTime:  shift.StartTime,
		EndTime:    shift.EndTime,
		DaysOfWeek: shift.DaysOfWeek,
	}
}
