package usecase

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"hrms/internal/domain"
	"hrms/internal/dto"
)

// CreateOrganization 📌 Handles the creation of a new organization
func (u *HrmsUsecase) CreateOrganization(ctx context.Context, req dto.CreateOrganizationDTO) (*dto.OrganizationResponseDTO, error) {
	org := &domain.Organization{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
		Email:   req.Email,
	}

	createdOrg, err := u.HrmsRepo.CreateOrganization(ctx, org)
	if err != nil {
		u.Logger.Error("❌ Failed to create organization", zap.Error(err))
		return nil, err
	}

	return &dto.OrganizationResponseDTO{
		ID:        createdOrg.ID,
		Name:      createdOrg.Name,
		Address:   createdOrg.Address,
		Phone:     createdOrg.Phone,
		Email:     createdOrg.Email,
		CreatedAt: createdOrg.CreatedAt,
		UpdatedAt: createdOrg.UpdatedAt,
	}, nil
}

// GetOrganization 📌 Fetch an organization by ID
func (u *HrmsUsecase) GetOrganization(ctx context.Context, orgID uint) (*dto.OrganizationResponseDTO, error) {
	org, err := u.HrmsRepo.GetOrganizationByID(ctx, orgID)
	if err != nil {
		u.Logger.Error("❌ Failed to fetch organization", zap.Error(err))
		return nil, err
	}

	return &dto.OrganizationResponseDTO{
		ID:        org.ID,
		Name:      org.Name,
		Address:   org.Address,
		Phone:     org.Phone,
		Email:     org.Email,
		CreatedAt: org.CreatedAt,
		UpdatedAt: org.UpdatedAt,
	}, nil
}

// UpdateOrganization 📌 Updates organization details
func (u *HrmsUsecase) UpdateOrganization(ctx context.Context, orgID uint, req dto.UpdateOrganizationDTO) error {
	updates := make(map[string]interface{})

	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Address != nil {
		updates["address"] = *req.Address
	}
	if req.Phone != nil {
		updates["phone"] = *req.Phone
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	err := u.HrmsRepo.UpdateOrganization(ctx, orgID, updates)
	if err != nil {
		u.Logger.Error("❌ Failed to update organization", zap.Error(err))
		return err
	}

	return nil
}

// DeleteOrganization 📌 Soft delete organization
func (u *HrmsUsecase) DeleteOrganization(ctx context.Context, orgID uint) error {
	err := u.HrmsRepo.DeleteOrganization(ctx, orgID)
	if err != nil {
		u.Logger.Error("❌ Failed to delete organization", zap.Error(err))
		return err
	}

	return nil
}

// ListOrganizations 📌 Fetch a paginated list of organizations
func (u *HrmsUsecase) ListOrganizations(ctx context.Context, limit, offset int, search string) (*dto.PaginatedOrganizationsResponse, error) {
	orgs, total, err := u.HrmsRepo.ListOrganizations(ctx, limit, offset, search)
	if err != nil {
		u.Logger.Error("❌ Failed to fetch organizations", zap.Error(err))
		return nil, err
	}

	// Map domain objects to DTO response
	var orgDTOs []dto.OrganizationResponseDTO
	for _, org := range orgs {
		orgDTOs = append(orgDTOs, dto.OrganizationResponseDTO{
			ID:        org.ID,
			Name:      org.Name,
			Address:   org.Address,
			Phone:     org.Phone,
			Email:     org.Email,
			CreatedAt: org.CreatedAt,
			UpdatedAt: org.UpdatedAt,
		})
	}

	return &dto.PaginatedOrganizationsResponse{
		Total: total,
		Limit: limit,
		Offset: offset,
		Organizations: orgDTOs,
	}, nil
}
