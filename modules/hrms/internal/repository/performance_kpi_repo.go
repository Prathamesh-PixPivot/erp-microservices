package repository

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreatePerformanceKPI 📌 Adds a new KPI entry for a performance review
func (r *HrmsRepository) CreatePerformanceKPI(ctx context.Context, kpi *domain.PerformanceKPI) (*domain.PerformanceKPI, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Ensure the related PerformanceReview exists
	var review domain.PerformanceReview
	if err := tx.Where("id = ?", kpi.ReviewID).First(&review).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, hrmsErrors.ErrPerformanceReviewNotFound
		}
		r.Logger.Error("❌ Database error in CreatePerformanceKPI", zap.Error(err))
		return nil, err
	}

	// ✅ Step 2: Check if KPI already exists for the review
	var existingKPI domain.PerformanceKPI
	if err := tx.Where("review_id = ? AND kpi_name = ?", kpi.ReviewID, kpi.KPIName).First(&existingKPI).Error; err == nil {
		tx.Rollback()
		r.Logger.Warn("⚠️ KPI already exists", zap.String("kpi_name", kpi.KPIName))
		return nil, hrmsErrors.ErrKPIAlreadyExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 3: Insert PerformanceKPI record
	if err := tx.Create(kpi).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create KPI entry", zap.Error(err))
		return nil, hrmsErrors.ErrKPICreationFailed
	}

	tx.Commit() // ✅ Commit transaction
	return kpi, nil
}

// GetPerformanceKPIByID 📌 Fetches KPI by ID (with optional Preloading)
func (r *HrmsRepository) GetPerformanceKPIByID(ctx context.Context, kpiID uint) (*domain.PerformanceKPI, error) {
	var kpi domain.PerformanceKPI
	err := r.DB.WithContext(ctx).Preload("Review").Where("id = ?", kpiID).First(&kpi).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrKPINotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetPerformanceKPIByID", zap.Error(err))
		return nil, err
	}

	return &kpi, nil
}
// UpdatePerformanceKPI 📌 Updates an existing KPI entry
func (r *HrmsRepository) UpdatePerformanceKPI(ctx context.Context, kpiID uint, updates map[string]interface{}) error {
	if err := r.DB.WithContext(ctx).Model(&domain.PerformanceKPI{}).
		Where("id = ?", kpiID).
		Updates(updates).Error; err != nil {
		r.Logger.Error("❌ Failed to update KPI entry", zap.Error(err))
		return hrmsErrors.ErrKPIUpdateFailed
	}

	return nil
}


// DeletePerformanceKPI 📌 Soft deletes a KPI entry
func (r *HrmsRepository) DeletePerformanceKPI(ctx context.Context, kpiID uint) error {
	// First, check if KPI exists before starting a transaction
	var kpi domain.PerformanceKPI
	if err := r.DB.WithContext(ctx).Where("id = ?", kpiID).First(&kpi).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return hrmsErrors.ErrKPINotFound
		}
		r.Logger.Error("❌ Database error in DeletePerformanceKPI", zap.Error(err))
		return err
	}

	// Begin transaction only when performing delete operation
	tx := r.DB.WithContext(ctx).Begin()
	if err := tx.Delete(&kpi).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete KPI entry", zap.Error(err))
		return hrmsErrors.ErrKPIDeletionFailed
	}

	tx.Commit()
	return nil
}

// ListPerformanceKPIs 📌 Fetches all KPIs for a given review ID
func (r *HrmsRepository) ListPerformanceKPIs(ctx context.Context, reviewID uint, limit, offset int) ([]domain.PerformanceKPI, int64, error) {
	var kpis []domain.PerformanceKPI
	var totalCount int64

	// Fetch KPI records directly without checking if the review exists
	query := r.DB.WithContext(ctx).Model(&domain.PerformanceKPI{}).Where("review_id = ?", reviewID)
	query.Count(&totalCount)

	if err := query.Limit(limit).Offset(offset).Find(&kpis).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch KPI records", zap.Error(err))
		return nil, 0, err
	}

	return kpis, totalCount, nil
}

