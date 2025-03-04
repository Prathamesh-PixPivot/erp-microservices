package usecase

import (
	"context"

	"hrms/internal/domain"
	"hrms/internal/dto"
)


// CreateSkillDevelopment 📌 Adds a new skill development entry.
func (uc *HrmsUsecase) CreateSkillDevelopment(ctx context.Context, skillDTO dto.SkillDevelopmentDTO) (*dto.SkillDevelopmentDTO, error) {
	skill := &domain.SkillDevelopment{
		ReviewID: skillDTO.ReviewID,
		Skill:    skillDTO.Skill,
		Progress: skillDTO.Progress,
	}

	createdSkill, err := uc.HrmsRepo.CreateSkillDevelopment(ctx, skill)
	if err != nil {
		return nil, err
	}

	return uc.convertSkillDevelopmentDomainToDTO(createdSkill), nil
}

// GetSkillDevelopmentByID 📌 Fetches a specific skill development entry by its ID.
func (uc *HrmsUsecase) GetSkillDevelopmentByID(ctx context.Context, skillID uint) (*dto.SkillDevelopmentDTO, error) {
	skill, err := uc.HrmsRepo.GetSkillDevelopmentByID(ctx, skillID)
	if err != nil {
		return nil, err
	}

	return uc.convertSkillDevelopmentDomainToDTO(skill), nil
}

// ListSkillDevelopments 📌 Retrieves skill development records for a given review.
func (uc *HrmsUsecase) ListSkillDevelopments(ctx context.Context, reviewID uint, limit, offset int) ([]dto.SkillDevelopmentDTO, int64, error) {
	skills, totalCount, err := uc.HrmsRepo.ListSkillDevelopments(ctx, reviewID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	var skillDTOs []dto.SkillDevelopmentDTO
	for _, skill := range skills {
		skillDTOs = append(skillDTOs, *uc.convertSkillDevelopmentDomainToDTO(&skill))
	}

	return skillDTOs, totalCount, nil
}

// UpdateSkillDevelopment 📌 Updates an existing skill development entry.
func (uc *HrmsUsecase) UpdateSkillDevelopment(ctx context.Context, skillID uint, updates map[string]interface{}) error {
	return uc.HrmsRepo.UpdateSkillDevelopment(ctx, skillID, updates)
}

// DeleteSkillDevelopment 📌 Soft deletes a skill development entry.
func (uc *HrmsUsecase) DeleteSkillDevelopment(ctx context.Context, skillID uint) error {
	return uc.HrmsRepo.DeleteSkillDevelopment(ctx, skillID)
}

// convertSkillDevelopmentDomainToDTO converts a SkillDevelopment domain model to DTO.
func (uc *HrmsUsecase) convertSkillDevelopmentDomainToDTO(skill *domain.SkillDevelopment) *dto.SkillDevelopmentDTO {
	return &dto.SkillDevelopmentDTO{
		ID:       skill.ID,
		ReviewID: skill.ReviewID,
		Skill:    skill.Skill,
		Progress: skill.Progress,
	}
}
