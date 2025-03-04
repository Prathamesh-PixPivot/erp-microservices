package usecase

import (
	"context"

	"hrms/internal/domain"
	"hrms/internal/dto"
)


// CreatePerformanceKPI 📌 Adds a new KPI entry for a performance review.
func (uc *HrmsUsecase) CreatePerformanceKPI(ctx context.Context, kpiDTO dto.PerformanceKPIDTO) (*dto.PerformanceKPIDTO, error) {
	kpi := &domain.PerformanceKPI{
		ReviewID: kpiDTO.ReviewID,
		KPIName:  kpiDTO.KPIName,
		Score:    kpiDTO.Score, // Already float64
		Comments: kpiDTO.Comments,
	}

	createdKPI, err := uc.HrmsRepo.CreatePerformanceKPI(ctx, kpi)
	if err != nil {
		return nil, err
	}

	return uc.convertKPIDomainToDTO(createdKPI), nil
}

// GetPerformanceKPIByID 📌 Fetches a specific KPI by its ID.
func (uc *HrmsUsecase) GetPerformanceKPIByID(ctx context.Context, kpiID uint) (*dto.PerformanceKPIDTO, error) {
	kpi, err := uc.HrmsRepo.GetPerformanceKPIByID(ctx, kpiID)
	if err != nil {
		return nil, err
	}

	return uc.convertKPIDomainToDTO(kpi), nil
}

// UpdatePerformanceKPI 📌 Updates an existing KPI entry.
func (uc *HrmsUsecase) UpdatePerformanceKPI(ctx context.Context, kpiID uint, updates map[string]interface{}) error {
	return uc.HrmsRepo.UpdatePerformanceKPI(ctx, kpiID, updates)
}

// DeletePerformanceKPI 📌 Deletes a KPI entry.
func (uc *HrmsUsecase) DeletePerformanceKPI(ctx context.Context, kpiID uint) error {
	return uc.HrmsRepo.DeletePerformanceKPI(ctx, kpiID)
}

// ListPerformanceKPIs 📌 Retrieves all KPIs linked to a performance review.
func (uc *HrmsUsecase) ListPerformanceKPIs(ctx context.Context, reviewID uint, limit, offset int) ([]dto.PerformanceKPIDTO, int64, error) {
	kpis, totalCount, err := uc.HrmsRepo.ListPerformanceKPIs(ctx, reviewID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	var kpiDTOs []dto.PerformanceKPIDTO
	for _, kpi := range kpis {
		kpiDTOs = append(kpiDTOs, *uc.convertKPIDomainToDTO(&kpi))
	}

	return kpiDTOs, totalCount, nil
}

// convertKPIDomainToDTO converts a PerformanceKPI domain model to DTO.
func (uc *HrmsUsecase) convertKPIDomainToDTO(kpi *domain.PerformanceKPI) *dto.PerformanceKPIDTO {
	return &dto.PerformanceKPIDTO{
		ID:       kpi.ID,
		ReviewID: kpi.ReviewID,
		KPIName:  kpi.KPIName,
		Score:    kpi.Score, // Already float64
		Comments: kpi.Comments,
	}
}
