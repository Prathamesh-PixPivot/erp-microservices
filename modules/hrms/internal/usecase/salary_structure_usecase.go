package usecase

import (
	"context"

	"hrms/internal/domain"
	"hrms/internal/dto"
)

// CreateSalaryStructure 📌 Adds a new salary structure.
func (uc *HrmsUsecase) CreateSalaryStructure(ctx context.Context, salaryDTO dto.SalaryStructureDTO) (*dto.SalaryStructureDTO, error) {
	salaryStructure := &domain.SalaryStructure{
		OrganizationID: salaryDTO.OrganizationID,
		DesignationID:  salaryDTO.DesignationID,
		BaseSalary:     salaryDTO.BaseSalary,
		Allowances:     salaryDTO.Allowances,
		TaxPercentage:  salaryDTO.TaxPercentage,
		Deductions:     salaryDTO.Deductions,
	}

	createdStructure, err := uc.HrmsRepo.CreateSalaryStructure(ctx, salaryStructure)
	if err != nil {
		return nil, err
	}

	return uc.convertSalaryDomainToDTO(createdStructure), nil
}

// GetSalaryStructure 📌 Fetch a specific salary structure by its ID.
func (uc *HrmsUsecase) GetSalaryStructure(ctx context.Context, salaryID uint) (*dto.SalaryStructureDTO, error) {
	salaryStructure, err := uc.HrmsRepo.GetSalaryStructure(ctx, salaryID)
	if err != nil {
		return nil, err
	}

	return uc.convertSalaryDomainToDTO(salaryStructure), nil
}

// ListSalaryStructures 📌 Retrieves all salary structures (with optional organization or designation filters).
func (uc *HrmsUsecase) ListSalaryStructures(ctx context.Context, organizationID, designationID uint) ([]dto.SalaryStructureDTO, error) {
	salaryStructures, err := uc.HrmsRepo.ListSalaryStructures(ctx, organizationID, designationID)
	if err != nil {
		return nil, err
	}

	var salaryDTOs []dto.SalaryStructureDTO
	for _, salary := range salaryStructures {
		salaryDTOs = append(salaryDTOs, *uc.convertSalaryDomainToDTO(&salary))
	}

	return salaryDTOs, nil
}

// UpdateSalaryStructure 📌 Updates an existing salary structure.
func (uc *HrmsUsecase) UpdateSalaryStructure(ctx context.Context, salaryID uint, updates map[string]interface{}) error {
	return uc.HrmsRepo.UpdateSalaryStructure(ctx, salaryID, updates)
}

// DeleteSalaryStructure 📌 Removes a salary structure.
func (uc *HrmsUsecase) DeleteSalaryStructure(ctx context.Context, salaryID uint) error {
	return uc.HrmsRepo.DeleteSalaryStructure(ctx, salaryID)
}

// convertSalaryDomainToDTO converts a SalaryStructure domain model to DTO.
func (uc *HrmsUsecase) convertSalaryDomainToDTO(salary *domain.SalaryStructure) *dto.SalaryStructureDTO {
	return &dto.SalaryStructureDTO{
		ID:             salary.ID,
		OrganizationID: salary.OrganizationID,
		DesignationID:  salary.DesignationID,
		BaseSalary:     salary.BaseSalary,
		Allowances:     salary.Allowances,
		TaxPercentage:  salary.TaxPercentage,
		Deductions:     salary.Deductions,
	}
}
