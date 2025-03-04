package usecase

import (
	"context"

	"hrms/internal/domain"
	"hrms/internal/dto"
)

// CreateWorkHistory 📌 Adds a new work history record.
func (uc *HrmsUsecase) CreateWorkHistory(ctx context.Context, workHistoryDTO dto.WorkHistoryDTO) (*dto.WorkHistoryDTO, error) {
	workHistory := &domain.WorkHistory{
		EmployeeID:    workHistoryDTO.EmployeeID,
		Company:       workHistoryDTO.Company,
		Designation:   workHistoryDTO.Designation,
		StartDate:     workHistoryDTO.StartDate,
		EndDate:       workHistoryDTO.EndDate,
		ReasonForExit: workHistoryDTO.ReasonForExit,
	}

	createdWorkHistory, err := uc.HrmsRepo.CreateWorkHistory(ctx, workHistory)
	if err != nil {
		return nil, err
	}

	return uc.convertWorkHistoryDomainToDTO(createdWorkHistory), nil
}

// GetWorkHistoryByID 📌 Fetch a specific work history record by its ID.
func (uc *HrmsUsecase) GetWorkHistoryByID(ctx context.Context, workHistoryID uint) (*dto.WorkHistoryDTO, error) {
	workHistory, err := uc.HrmsRepo.GetWorkHistoryByID(ctx, workHistoryID)
	if err != nil {
		return nil, err
	}

	return uc.convertWorkHistoryDomainToDTO(workHistory), nil
}

// GetWorkHistoryByEmployee 📌 Retrieves work history records for a specific employee.
func (uc *HrmsUsecase) GetWorkHistoryByEmployee(ctx context.Context, employeeID uint) ([]dto.WorkHistoryDTO, error) {
	workHistories, err := uc.HrmsRepo.GetWorkHistoryByEmployee(ctx, employeeID)
	if err != nil {
		return nil, err
	}

	var workHistoryDTOs []dto.WorkHistoryDTO
	for _, workHistory := range workHistories {
		workHistoryDTOs = append(workHistoryDTOs, *uc.convertWorkHistoryDomainToDTO(&workHistory))
	}

	return workHistoryDTOs, nil
}

// UpdateWorkHistory 📌 Updates an existing work history record.
func (uc *HrmsUsecase) UpdateWorkHistory(ctx context.Context, workHistoryID uint, updates map[string]interface{}) error {
	return uc.HrmsRepo.UpdateWorkHistory(ctx, workHistoryID, updates)
}

// DeleteWorkHistory 📌 Soft deletes a work history record.
func (uc *HrmsUsecase) DeleteWorkHistory(ctx context.Context, workHistoryID uint) error {
	return uc.HrmsRepo.DeleteWorkHistory(ctx, workHistoryID)
}

// convertWorkHistoryDomainToDTO converts a WorkHistory domain model to DTO.
func (uc *HrmsUsecase) convertWorkHistoryDomainToDTO(workHistory *domain.WorkHistory) *dto.WorkHistoryDTO {
	return &dto.WorkHistoryDTO{
		ID:            workHistory.ID,
		EmployeeID:    workHistory.EmployeeID,
		Company:       workHistory.Company,
		Designation:   workHistory.Designation,
		StartDate:     workHistory.StartDate,
		EndDate:       workHistory.EndDate,
		ReasonForExit: workHistory.ReasonForExit,
	}
}
