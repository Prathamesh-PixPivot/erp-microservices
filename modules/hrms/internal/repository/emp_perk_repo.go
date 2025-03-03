package repository

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreateEmployeePerk 📌 Adds a perk for an employee
func (r *HrmsRepository) CreateEmployeePerk(ctx context.Context, perk *domain.EmployeePerk) (*domain.EmployeePerk, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Validate Employee ID & Perk
	if perk.EmployeeID == 0 || perk.Perk == "" {
		tx.Rollback()
		return nil, hrmsErrors.ErrInvalidPerk
	}

	// ✅ Check if the Perk Already Exists
	var existing domain.EmployeePerk
	if err := tx.Where("employee_id = ? AND perk = ?", perk.EmployeeID, perk.Perk).First(&existing).Error; err == nil {
		tx.Rollback()
		return nil, hrmsErrors.ErrDuplicatePerk
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err // Some other DB error
	}

	// ✅ Insert Employee Perk
	if err := tx.Create(perk).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create employee perk", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return perk, nil
}

// GetEmployeePerks 📌 Fetch all perks for an employee
func (r *HrmsRepository) GetEmployeePerks(ctx context.Context, employeeID uint) ([]domain.EmployeePerk, error) {
	var perks []domain.EmployeePerk
	err := r.DB.WithContext(ctx).Where("employee_id = ?", employeeID).Find(&perks).Error

	if err != nil {
		r.Logger.Error("❌ Failed to fetch employee perks", zap.Error(err))
		return nil, err
	}

	return perks, nil
}

// UpdateEmployeePerk 📌 Update an employee's perk
func (r *HrmsRepository) UpdateEmployeePerk(ctx context.Context, perkID uint, updatedPerk string) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Validate if Perk Exists
	var perk domain.EmployeePerk
	if err := tx.Where("id = ?", perkID).First(&perk).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return hrmsErrors.ErrPerkNotFound
		}
		r.Logger.Error("❌ Error finding employee perk for update", zap.Error(err))
		return err
	}

	// ✅ Update Perk
	perk.Perk = updatedPerk
	if err := tx.Save(&perk).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to update employee perk", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// DeleteEmployeePerk 📌 Remove an employee's perk
func (r *HrmsRepository) DeleteEmployeePerk(ctx context.Context, perkID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Where("id = ?", perkID).Delete(&domain.EmployeePerk{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete employee perk", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}
