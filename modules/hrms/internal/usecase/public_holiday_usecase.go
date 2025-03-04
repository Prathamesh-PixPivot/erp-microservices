package usecase

import (
	"context"

	"hrms/internal/domain"
	"hrms/internal/dto"
)

// CreatePublicHoliday 📌 Adds a new public holiday.
func (uc *HrmsUsecase) CreatePublicHoliday(ctx context.Context, holidayDTO dto.PublicHolidayDTO) (*dto.PublicHolidayDTO, error) {
	holiday := &domain.PublicHoliday{
		OrganizationID: holidayDTO.OrganizationID,
		Name:           holidayDTO.Name,
		Date:           holidayDTO.Date,
	}

	createdHoliday, err := uc.HrmsRepo.CreatePublicHoliday(ctx, holiday)
	if err != nil {
		return nil, err
	}

	return uc.convertHolidayDomainToDTO(createdHoliday), nil
}

// GetPublicHoliday 📌 Fetch a specific public holiday by its ID.
func (uc *HrmsUsecase) GetPublicHoliday(ctx context.Context, holidayID uint) (*dto.PublicHolidayDTO, error) {
	holiday, err := uc.HrmsRepo.GetPublicHoliday(ctx, holidayID)
	if err != nil {
		return nil, err
	}

	return uc.convertHolidayDomainToDTO(holiday), nil
}

// ListPublicHolidays 📌 Retrieves all holidays for an organization (with optional year filter).
func (uc *HrmsUsecase) ListPublicHolidays(ctx context.Context, organizationID uint, year *int) ([]dto.PublicHolidayDTO, error) {
	holidays, err := uc.HrmsRepo.ListPublicHolidays(ctx, organizationID, year)
	if err != nil {
		return nil, err
	}

	var holidayDTOs []dto.PublicHolidayDTO
	for _, holiday := range holidays {
		holidayDTOs = append(holidayDTOs, *uc.convertHolidayDomainToDTO(&holiday))
	}

	return holidayDTOs, nil
}

// UpdatePublicHoliday 📌 Updates an existing public holiday.
func (uc *HrmsUsecase) UpdatePublicHoliday(ctx context.Context, holidayID uint, updates map[string]interface{}) error {
	return uc.HrmsRepo.UpdatePublicHoliday(ctx, holidayID, updates)
}

// DeletePublicHoliday 📌 Removes a public holiday.
func (uc *HrmsUsecase) DeletePublicHoliday(ctx context.Context, holidayID uint) error {
	return uc.HrmsRepo.DeletePublicHoliday(ctx, holidayID)
}

// convertHolidayDomainToDTO converts a PublicHoliday domain model to DTO.
func (uc *HrmsUsecase) convertHolidayDomainToDTO(holiday *domain.PublicHoliday) *dto.PublicHolidayDTO {
	return &dto.PublicHolidayDTO{
		ID:             holiday.ID,
		OrganizationID: holiday.OrganizationID,
		Name:           holiday.Name,
		Date:           holiday.Date,
	}
}
