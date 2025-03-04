package repository

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreatePerformanceReview 📌 Adds a new review entry
func (r *HrmsRepository) CreatePerformanceReview(ctx context.Context, review *domain.PerformanceReview) (*domain.PerformanceReview, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Prevent duplicate review for the same period
	var existingReview domain.PerformanceReview
	if err := tx.Where("employee_id = ? AND review_period = ?", review.EmployeeID, review.ReviewPeriod).
		First(&existingReview).Error; err == nil {
		tx.Rollback()
		r.Logger.Warn("⚠️ Review already exists", zap.Uint("employee_id", review.EmployeeID))
		return nil, hrmsErrors.ErrReviewAlreadyExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 2: Insert PerformanceReview record
	if err := tx.Create(review).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create PerformanceReview", zap.Error(err))
		return nil, hrmsErrors.ErrReviewCreationFailed
	}

	tx.Commit() // ✅ Commit transaction
	return review, nil
}

// GetPerformanceReviewByID 📌 Fetch a review by ID (Preloads KPIs & Skill Developments)
func (r *HrmsRepository) GetPerformanceReviewByID(ctx context.Context, reviewID uint) (*domain.PerformanceReview, error) {
	var review domain.PerformanceReview

	// Fetch review with related KPIs & skill developments
	err := r.DB.WithContext(ctx).Preload("KPIs").Preload("SkillDevelopments").
		Where("id = ?", reviewID).First(&review).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrPerformanceReviewNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetPerformanceReviewByID", zap.Error(err))
		return nil, err
	}

	return &review, nil
}

// UpdatePerformanceReview 📌 Updates an existing review
func (r *HrmsRepository) UpdatePerformanceReview(ctx context.Context, reviewID uint, updates map[string]interface{}) error {
	if err := r.DB.WithContext(ctx).Model(&domain.PerformanceReview{}).
		Where("id = ?", reviewID).
		Updates(updates).Error; err != nil {
		r.Logger.Error("❌ Failed to update review", zap.Error(err))
		return hrmsErrors.ErrReviewUpdateFailed
	}

	return nil
}

// DeletePerformanceReview 📌 Soft deletes a review and associated KPIs & Skill Developments
func (r *HrmsRepository) DeletePerformanceReview(ctx context.Context, reviewID uint) error {
	// Step 1: Check if review exists before starting a transaction
	var review domain.PerformanceReview
	if err := r.DB.WithContext(ctx).Where("id = ?", reviewID).First(&review).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return hrmsErrors.ErrPerformanceReviewNotFound
		}
		r.Logger.Error("❌ Database error in DeletePerformanceReview", zap.Error(err))
		return err
	}

	// Step 2: Begin transaction for cascading deletion
	tx := r.DB.WithContext(ctx).Begin()

	// Delete associated KPIs (soft delete)
	if err := tx.Where("review_id = ?", reviewID).Delete(&domain.PerformanceKPI{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete associated KPIs", zap.Error(err))
		return err
	}

	// Delete associated Skill Developments (soft delete)
	if err := tx.Where("review_id = ?", reviewID).Delete(&domain.SkillDevelopment{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete associated Skill Developments", zap.Error(err))
		return err
	}

	// Delete the review itself
	if err := tx.Delete(&review).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete performance review", zap.Error(err))
		return hrmsErrors.ErrReviewDeletionFailed
	}

	tx.Commit()
	return nil
}

// ListPerformanceReviews 📌 Fetch all reviews for an employee with pagination
func (r *HrmsRepository) ListPerformanceReviews(ctx context.Context, employeeID uint, limit, offset int) ([]domain.PerformanceReview, int64, error) {
	var reviews []domain.PerformanceReview
	var totalCount int64

	// Fetch total count in a single query
	query := r.DB.WithContext(ctx).Model(&domain.PerformanceReview{}).Where("employee_id = ?", employeeID)

	// Use subquery to count total records before applying pagination
	if err := query.Count(&totalCount).Error; err != nil {
		r.Logger.Error("❌ Failed to count performance reviews", zap.Error(err))
		return nil, 0, err
	}

	// Fetch paginated records with KPIs & Skill Developments
	if err := query.Limit(limit).Offset(offset).Preload("KPIs").Preload("SkillDevelopments").Find(&reviews).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch performance reviews", zap.Error(err))
		return nil, 0, err
	}

	return reviews, totalCount, nil
}
