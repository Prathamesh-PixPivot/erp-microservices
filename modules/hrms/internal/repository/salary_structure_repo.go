package repository

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreateSalaryStructure 📌 Adds a new salary structure (Prevents duplicates)
func (r *HrmsRepository) CreateSalaryStructure(ctx context.Context, salaryStructure *domain.SalaryStructure) (*domain.SalaryStructure, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Check if salary structure already exists for this Designation
	var existingStructure domain.SalaryStructure
	if err := tx.Where("designation_id = ?", salaryStructure.DesignationID).First(&existingStructure).Error; err == nil {
		tx.Rollback()
		return nil, hrmsErrors.ErrSalaryStructureExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 2: Validate salary structure data
	if salaryStructure.BaseSalary < 0 || salaryStructure.Allowances < 0 || salaryStructure.TaxPercentage < 0 || salaryStructure.Deductions < 0 {
		tx.Rollback()
		return nil, hrmsErrors.ErrInvalidSalaryData
	}

	// ✅ Step 3: Insert Salary Structure
	if err := tx.Create(salaryStructure).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create salary structure", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return salaryStructure, nil
}

// GetSalaryStructure 📌 Fetch salary structure by ID
func (r *HrmsRepository) GetSalaryStructure(ctx context.Context, salaryID uint) (*domain.SalaryStructure, error) {
	var salaryStructure domain.SalaryStructure
	err := r.DB.WithContext(ctx).Where("id = ?", salaryID).First(&salaryStructure).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrSalaryStructureNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetSalaryStructure", zap.Error(err))
		return nil, err
	}

	return &salaryStructure, nil
}

// ListSalaryStructures 📌 Fetch all salary structures (Optional: filter by Organization or Designation)
func (r *HrmsRepository) ListSalaryStructures(ctx context.Context, organizationID, designationID uint) ([]domain.SalaryStructure, error) {
	var salaryStructures []domain.SalaryStructure
	query := r.DB.WithContext(ctx).Model(&domain.SalaryStructure{})

	if organizationID > 0 {
		query = query.Where("organization_id = ?", organizationID)
	}

	if designationID > 0 {
		query = query.Where("designation_id = ?", designationID)
	}

	if err := query.Order("designation_id ASC").Find(&salaryStructures).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch salary structures", zap.Error(err))
		return nil, err
	}

	return salaryStructures, nil
}

// UpdateSalaryStructure 📌 Update salary structure details
func (r *HrmsRepository) UpdateSalaryStructure(ctx context.Context, salaryID uint, updates map[string]interface{}) error {
	tx := r.DB.WithContext(ctx).Begin()

	// Validate salary-related fields
	if baseSalary, ok := updates["base_salary"].(float64); ok && baseSalary < 0 {
		return hrmsErrors.ErrInvalidSalaryData
	}
	if allowances, ok := updates["allowances"].(float64); ok && allowances < 0 {
		return hrmsErrors.ErrInvalidSalaryData
	}
	if taxPercentage, ok := updates["tax_percentage"].(float64); ok && taxPercentage < 0 {
		return hrmsErrors.ErrInvalidSalaryData
	}
	if deductions, ok := updates["deductions"].(float64); ok && deductions < 0 {
		return hrmsErrors.ErrInvalidSalaryData
	}

	// ✅ Update salary structure details
	if err := tx.Model(&domain.SalaryStructure{}).Where("id = ?", salaryID).Updates(updates).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to update salary structure", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// DeleteSalaryStructure 📌 Remove a salary structure
func (r *HrmsRepository) DeleteSalaryStructure(ctx context.Context, salaryID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	var salaryStructure domain.SalaryStructure
	if err := tx.Where("id = ?", salaryID).First(&salaryStructure).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return hrmsErrors.ErrSalaryStructureNotFound
		}
		r.Logger.Error("❌ Error finding salary structure for deletion", zap.Error(err))
		return err
	}

	if err := tx.Delete(&domain.SalaryStructure{}, salaryID).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete salary structure", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}
