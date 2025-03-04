package repository

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreateEmployeeBenefits 📌 Adds employee benefits
func (r *HrmsRepository) CreateEmployeeBenefits(ctx context.Context, benefits *domain.EmployeeBenefits) (*domain.EmployeeBenefits, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Validate Employee ID
	if benefits.EmployeeID == 0 {
		tx.Rollback()
		return nil, hrmsErrors.ErrInvalidBenefits
	}

	// ✅ Check if Employee Benefits already exist
	var existing domain.EmployeeBenefits
	if err := tx.Where("employee_id = ?", benefits.EmployeeID).First(&existing).Error; err == nil {
		tx.Rollback()
		return nil, errors.New("employee benefits already exist")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err // Some other DB error
	}

	// ✅ Insert Employee Benefits
	if err := tx.Create(benefits).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create employee benefits record", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return benefits, nil
}

// GetEmployeeBenefits 📌 Fetch employee benefits by Employee ID
func (r *HrmsRepository) GetEmployeeBenefits(ctx context.Context, employeeID uint) (*domain.EmployeeBenefits, error) {
	var benefits domain.EmployeeBenefits
	err := r.DB.WithContext(ctx).Preload("Perks").Where("employee_id = ?", employeeID).First(&benefits).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrBenefitsNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetEmployeeBenefits", zap.Error(err))
		return nil, err
	}

	return &benefits, nil
}

// ListEmployeeBenefits 📌 Fetch all employee benefits with optional filtering
func (r *HrmsRepository) ListEmployeeBenefits(ctx context.Context, healthPlan, retirementPlan string) ([]domain.EmployeeBenefits, error) {
	var benefits []domain.EmployeeBenefits
	query := r.DB.WithContext(ctx).Model(&domain.EmployeeBenefits{}).Preload("Perks")

	if healthPlan != "" {
		query = query.Where("health_plan = ?", healthPlan)
	}
	if retirementPlan != "" {
		query = query.Where("retirement_plan = ?", retirementPlan)
	}

	if err := query.Find(&benefits).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch employee benefits", zap.Error(err))
		return nil, err
	}

	return benefits, nil
}

// UpdateEmployeeBenefits 📌 Update employee benefits details
func (r *HrmsRepository) UpdateEmployeeBenefits(ctx context.Context, employeeID uint, updates map[string]interface{}) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Validate if Employee Benefits exist
	var benefits domain.EmployeeBenefits
	if err := tx.Where("employee_id = ?", employeeID).First(&benefits).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return hrmsErrors.ErrBenefitsNotFound
		}
		r.Logger.Error("❌ Error finding employee benefits record for update", zap.Error(err))
		return err
	}

	// ✅ Update Employee Benefits
	if err := tx.Model(&benefits).Updates(updates).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to update employee benefits", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// DeleteEmployeeBenefits 📌 Soft delete employee benefits
func (r *HrmsRepository) DeleteEmployeeBenefits(ctx context.Context, employeeID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Where("employee_id = ?", employeeID).Delete(&domain.EmployeeBenefits{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete employee benefits", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}
