package repository

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreateSkillDevelopment 📌 Adds a new skill development entry
func (r *HrmsRepository) CreateSkillDevelopment(ctx context.Context, skill *domain.SkillDevelopment) (*domain.SkillDevelopment, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Check if skill entry already exists for the review
	var existingSkill domain.SkillDevelopment
	if err := tx.Where("review_id = ? AND skill = ?", skill.ReviewID, skill.Skill).First(&existingSkill).Error; err == nil {
		tx.Rollback()
		r.Logger.Warn("⚠️ Skill development entry already exists", zap.String("skill", skill.Skill))
		return nil, hrmsErrors.ErrSkillAlreadyExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 2: Insert SkillDevelopment record
	if err := tx.Create(skill).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create skill development entry", zap.Error(err))
		return nil, hrmsErrors.ErrSkillCreationFailed
	}

	tx.Commit() // ✅ Commit transaction
	return skill, nil
}

// GetSkillDevelopmentByID 📌 Fetches skill development entry by ID
func (r *HrmsRepository) GetSkillDevelopmentByID(ctx context.Context, skillID uint) (*domain.SkillDevelopment, error) {
	var skill domain.SkillDevelopment
	err := r.DB.WithContext(ctx).Where("id = ?", skillID).First(&skill).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrSkillNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetSkillDevelopmentByID", zap.Error(err))
		return nil, err
	}

	return &skill, nil
}

// UpdateSkillDevelopment 📌 Updates an existing skill development entry
func (r *HrmsRepository) UpdateSkillDevelopment(ctx context.Context, skillID uint, updates map[string]interface{}) error {
	if err := r.DB.WithContext(ctx).Model(&domain.SkillDevelopment{}).
		Where("id = ?", skillID).
		Updates(updates).Error; err != nil {
		r.Logger.Error("❌ Failed to update skill development entry", zap.Error(err))
		return hrmsErrors.ErrSkillUpdateFailed
	}

	return nil
}

// DeleteSkillDevelopment 📌 Soft deletes a skill development entry
func (r *HrmsRepository) DeleteSkillDevelopment(ctx context.Context, skillID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Where("id = ?", skillID).Delete(&domain.SkillDevelopment{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete skill development entry", zap.Error(err))
		return hrmsErrors.ErrSkillDeletionFailed
	}

	tx.Commit()
	return nil
}

// ListSkillDevelopments 📌 Fetches all skill developments for a given review ID
func (r *HrmsRepository) ListSkillDevelopments(ctx context.Context, reviewID uint, limit, offset int) ([]domain.SkillDevelopment, int64, error) {
	var skills []domain.SkillDevelopment
	var totalCount int64

	query := r.DB.WithContext(ctx).Model(&domain.SkillDevelopment{}).Where("review_id = ?", reviewID)

	query.Count(&totalCount)

	if err := query.Limit(limit).Offset(offset).Find(&skills).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch skill development records", zap.Error(err))
		return nil, 0, err
	}

	return skills, totalCount, nil
}
