package usecase

import (
	"context"
	"hrms/internal/domain"
	"hrms/internal/dto"
	"hrms/internal/errors"

	"go.uber.org/zap"
)

// CreateLeavePolicy handles leave policy creation
func (u *HrmsUsecase) CreateLeavePolicy(ctx context.Context, req dto.CreateLeavePolicyRequest) (*dto.LeavePolicyResponse, error) {
	// Validate input
	if req.OrganizationID == 0 || req.MaxDays < 0 {
		return nil, errors.ErrInvalidLeaveType
	}

	// Create leave policy entity
	policy := &domain.LeavePolicy{
		OrganizationID: req.OrganizationID,
		LeaveType:      req.LeaveType,
		MaxDays:        req.MaxDays,
		CarryForward:   req.CarryForward,
	}

	// Persist data
	createdPolicy, err := u.HrmsRepo.CreateLeavePolicy(ctx, policy)
	if err != nil {
		u.Logger.Error("❌ Failed to create leave policy", zap.Error(err))
		return nil, err
	}

	// Response mapping
	return &dto.LeavePolicyResponse{
		ID:             createdPolicy.ID,
		OrganizationID: createdPolicy.OrganizationID,
		LeaveType:      createdPolicy.LeaveType,
		MaxDays:        createdPolicy.MaxDays,
		CarryForward:   createdPolicy.CarryForward,
	}, nil
}

// GetLeavePolicy fetches a leave policy by ID
func (u *HrmsUsecase) GetLeavePolicy(ctx context.Context, policyID uint) (*dto.LeavePolicyResponse, error) {
	// Fetch policy from repository
	policy, err := u.HrmsRepo.GetLeavePolicy(ctx, policyID)
	if err != nil {
		u.Logger.Error("❌ Failed to fetch leave policy", zap.Error(err))
		return nil, err
	}

	// Response mapping
	return &dto.LeavePolicyResponse{
		ID:             policy.ID,
		OrganizationID: policy.OrganizationID,
		LeaveType:      policy.LeaveType,
		MaxDays:        policy.MaxDays,
		CarryForward:   policy.CarryForward,
	}, nil
}

// ListLeavePolicies fetches all leave policies for an organization
func (u *HrmsUsecase) ListLeavePolicies(ctx context.Context, organizationID uint) ([]dto.LeavePolicyResponse, error) {
	// Fetch policies from repository
	policies, err := u.HrmsRepo.ListLeavePolicies(ctx, organizationID)
	if err != nil {
		u.Logger.Error("❌ Failed to list leave policies", zap.Error(err))
		return nil, err
	}

	// Convert domain models to DTOs
	var response []dto.LeavePolicyResponse
	for _, policy := range policies {
		response = append(response, dto.LeavePolicyResponse{
			ID:             policy.ID,
			OrganizationID: policy.OrganizationID,
			LeaveType:      policy.LeaveType,
			MaxDays:        policy.MaxDays,
			CarryForward:   policy.CarryForward,
		})
	}

	return response, nil
}

// UpdateLeavePolicy handles updating leave policy fields
func (u *HrmsUsecase) UpdateLeavePolicy(ctx context.Context, policyID uint, req dto.UpdateLeavePolicyRequest) error {
	// Validate input
	updates := make(map[string]interface{})
	if req.MaxDays != nil {
		if *req.MaxDays < 0 {
			return errors.ErrInvalidLeaveType
		}
		updates["max_days"] = *req.MaxDays
	}
	if req.CarryForward != nil {
		updates["carry_forward"] = *req.CarryForward
	}

	// Update policy in repository
	err := u.HrmsRepo.UpdateLeavePolicy(ctx, policyID, updates)
	if err != nil {
		u.Logger.Error("❌ Failed to update leave policy", zap.Error(err))
		return err
	}

	return nil
}

// DeleteLeavePolicy deletes a leave policy
func (u *HrmsUsecase) DeleteLeavePolicy(ctx context.Context, policyID uint) error {
	// Validate input
	if policyID == 0 {
		return errors.ErrInvalidLeaveType
	}

	// Delete leave policy from repository
	err := u.HrmsRepo.DeleteLeavePolicy(ctx, policyID)
	if err != nil {
		u.Logger.Error("❌ Failed to delete leave policy", zap.Error(err))
		return err
	}

	return nil
}
