package repository

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreateBonus 📌 Adds a new bonus record
func (r *HrmsRepository) CreateBonus(ctx context.Context, bonus *domain.Bonus) (*domain.Bonus, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Validate Amount
	if bonus.Amount <= 0 {
		tx.Rollback()
		return nil, hrmsErrors.ErrInvalidBonus
	}

	// ✅ Step 2: Ensure Approval Date is Valid
	if bonus.Status == "Approved" && bonus.ApprovalDate.IsZero() {
		tx.Rollback()
		return nil, errors.New("approval date must be set for approved bonuses")
	}

	// ✅ Step 3: Insert Bonus Record
	if err := tx.Create(bonus).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create bonus record", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return bonus, nil
}

// GetBonusByID 📌 Fetch a bonus record by ID
func (r *HrmsRepository) GetBonusByID(ctx context.Context, bonusID uint) (*domain.Bonus, error) {
	var bonus domain.Bonus
	err := r.DB.WithContext(ctx).Where("id = ?", bonusID).First(&bonus).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrBonusNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetBonusByID", zap.Error(err))
		return nil, err
	}

	return &bonus, nil
}

// ListBonuses 📌 Fetch all bonuses (Optional: filter by Employee ID & Status)
func (r *HrmsRepository) ListBonuses(ctx context.Context, employeeID uint, status string) ([]domain.Bonus, error) {
	var bonuses []domain.Bonus
	query := r.DB.WithContext(ctx).Model(&domain.Bonus{})

	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Order("issue_date DESC").Find(&bonuses).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch bonuses", zap.Error(err))
		return nil, err
	}

	return bonuses, nil
}

// UpdateBonus 📌 Update a bonus record (e.g., approval)
func (r *HrmsRepository) UpdateBonus(ctx context.Context, bonusID uint, updates map[string]interface{}) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Prevent unauthorized status updates
	if status, ok := updates["status"].(string); ok {
		if status == "Approved" && updates["approval_date"] == nil {
			return errors.New("approval date is required when approving bonus")
		}
	}

	// ✅ Update Bonus
	if err := tx.Model(&domain.Bonus{}).Where("id = ?", bonusID).Updates(updates).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to update bonus record", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// DeleteBonus 📌 Soft delete a bonus record
func (r *HrmsRepository) DeleteBonus(ctx context.Context, bonusID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	var bonus domain.Bonus
	if err := tx.Where("id = ?", bonusID).First(&bonus).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return hrmsErrors.ErrBonusNotFound
		}
		r.Logger.Error("❌ Error finding bonus record for deletion", zap.Error(err))
		return err
	}

	if err := tx.Delete(&domain.Bonus{}, bonusID).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete bonus record", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}
