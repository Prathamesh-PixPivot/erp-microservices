package usecase

import (
	"context"

	"hrms/internal/domain"
	"hrms/internal/dto"
)

// CreatePerformanceReview 📌 Adds a new performance review entry.
func (uc *HrmsUsecase) CreatePerformanceReview(ctx context.Context, reviewDTO dto.PerformanceReviewDTO) (*dto.PerformanceReviewDTO, error) {
	review := &domain.PerformanceReview{
		EmployeeID:    reviewDTO.EmployeeID,
		ReviewerID:    reviewDTO.ReviewerID,
		ReviewDate:    reviewDTO.ReviewDate,
		ReviewPeriod:  reviewDTO.ReviewPeriod,
		OverallRating: reviewDTO.OverallRating,
		Feedback:      reviewDTO.Feedback,
		Promotion:     reviewDTO.Promotion,
	}

	createdReview, err := uc.HrmsRepo.CreatePerformanceReview(ctx, review)
	if err != nil {
		return nil, err
	}

	return uc.convertPerformanceReviewDomainToDTO(createdReview), nil
}

// GetPerformanceReviewByID 📌 Fetches a specific performance review entry by its ID.
func (uc *HrmsUsecase) GetPerformanceReviewByID(ctx context.Context, reviewID uint) (*dto.PerformanceReviewDTO, error) {
	review, err := uc.HrmsRepo.GetPerformanceReviewByID(ctx, reviewID)
	if err != nil {
		return nil, err
	}

	return uc.convertPerformanceReviewDomainToDTO(review), nil
}

// ListPerformanceReviews 📌 Retrieves performance review records for an employee.
func (uc *HrmsUsecase) ListPerformanceReviews(ctx context.Context, employeeID uint, limit, offset int) ([]dto.PerformanceReviewDTO, int64, error) {
	reviews, totalCount, err := uc.HrmsRepo.ListPerformanceReviews(ctx, employeeID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	var reviewDTOs []dto.PerformanceReviewDTO
	for _, review := range reviews {
		reviewDTOs = append(reviewDTOs, *uc.convertPerformanceReviewDomainToDTO(&review))
	}

	return reviewDTOs, totalCount, nil
}

// UpdatePerformanceReview 📌 Updates an existing performance review entry.
func (uc *HrmsUsecase) UpdatePerformanceReview(ctx context.Context, reviewID uint, updates map[string]interface{}) error {
	return uc.HrmsRepo.UpdatePerformanceReview(ctx, reviewID, updates)
}

// DeletePerformanceReview 📌 Soft deletes a performance review entry.
func (uc *HrmsUsecase) DeletePerformanceReview(ctx context.Context, reviewID uint) error {
	return uc.HrmsRepo.DeletePerformanceReview(ctx, reviewID)
}

// convertPerformanceReviewDomainToDTO converts a PerformanceReview domain model to DTO.
func (uc *HrmsUsecase) convertPerformanceReviewDomainToDTO(review *domain.PerformanceReview) *dto.PerformanceReviewDTO {
	kpiDTOs := []dto.PerformanceKPIDTO{}
	for _, kpi := range review.KPIs {
		kpiDTOs = append(kpiDTOs, dto.PerformanceKPIDTO{
			ID:       kpi.ID,
			ReviewID: kpi.ReviewID,
			KPIName:  kpi.KPIName,
			Score:    kpi.Score,
			Comments: kpi.Comments,
		})
	}

	skillDTOs := []dto.SkillDevelopmentDTO{}
	for _, skill := range review.SkillDevelopments {
		skillDTOs = append(skillDTOs, dto.SkillDevelopmentDTO{
			ID:       skill.ID,
			ReviewID: skill.ReviewID,
			Skill:    skill.Skill,
			Progress: skill.Progress,
		})
	}

	return &dto.PerformanceReviewDTO{
		ID:               review.ID,
		EmployeeID:       review.EmployeeID,
		ReviewerID:       review.ReviewerID,
		ReviewDate:       review.ReviewDate,
		ReviewPeriod:     review.ReviewPeriod,
		OverallRating:    review.OverallRating,
		Feedback:         review.Feedback,
		Promotion:        review.Promotion,
		KPIs:             kpiDTOs,
		SkillDevelopments: skillDTOs,
	}
}
