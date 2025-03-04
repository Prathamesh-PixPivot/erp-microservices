package repository

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreateLeavePolicy 📌 Adds a new leave policy for an organization
func (r *HrmsRepository) CreateLeavePolicy(ctx context.Context, policy *domain.LeavePolicy) (*domain.LeavePolicy, error) {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Step 1: Check if a policy already exists for the given leave type in the organization
	var existingPolicy domain.LeavePolicy
	if err := tx.Where("organization_id = ? AND leave_type = ?", policy.OrganizationID, policy.LeaveType).
		First(&existingPolicy).Error; err == nil {
		tx.Rollback()
		return nil, hrmsErrors.ErrLeavePolicyExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	// ✅ Step 2: Insert the new policy
	if err := tx.Create(policy).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create leave policy", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return policy, nil
}

// GetLeavePolicy 📌 Fetch a leave policy by ID
func (r *HrmsRepository) GetLeavePolicy(ctx context.Context, policyID uint) (*domain.LeavePolicy, error) {
	var policy domain.LeavePolicy
	err := r.DB.WithContext(ctx).Where("id = ?", policyID).First(&policy).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrLeavePolicyNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetLeavePolicy", zap.Error(err))
		return nil, err
	}

	return &policy, nil
}

// ListLeavePolicies 📌 Fetch all leave policies for an organization
func (r *HrmsRepository) ListLeavePolicies(ctx context.Context, organizationID uint) ([]domain.LeavePolicy, error) {
	var policies []domain.LeavePolicy

	if err := r.DB.WithContext(ctx).Where("organization_id = ?", organizationID).
		Order("leave_type ASC").Find(&policies).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch leave policies", zap.Error(err))
		return nil, err
	}

	return policies, nil
}

// UpdateLeavePolicy 📌 Update leave policy (e.g., max days, carry-forward rules)
func (r *HrmsRepository) UpdateLeavePolicy(ctx context.Context, policyID uint, updates map[string]interface{}) error {
	tx := r.DB.WithContext(ctx).Begin()

	// Prevent duplication when changing leave type
	if leaveType, ok := updates["leave_type"].(domain.LeaveType); ok {
		var policy domain.LeavePolicy
		if err := tx.Where("id != ? AND leave_type = ?", policyID, leaveType).First(&policy).Error; err == nil {
			tx.Rollback()
			return hrmsErrors.ErrLeavePolicyExists
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return err
		}
	}

	// ✅ Update policy
	if err := tx.Model(&domain.LeavePolicy{}).Where("id = ?", policyID).Updates(updates).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to update leave policy", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// DeleteLeavePolicy 📌 Remove a leave policy
func (r *HrmsRepository) DeleteLeavePolicy(ctx context.Context, policyID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Where("id = ?", policyID).Delete(&domain.LeavePolicy{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete leave policy", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}
