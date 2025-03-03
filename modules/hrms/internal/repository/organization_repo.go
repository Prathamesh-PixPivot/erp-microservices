package repository

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"regexp"
	"time"

	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
	
)

// CreateOrganization 📌 Create a new organization
func (r *HrmsRepository) CreateOrganization(ctx context.Context, organization *domain.Organization) (*domain.Organization, error) {
	// ✅ Step 1: Validate email format
	if !isValidEmail(organization.Email) {
		return nil, hrmsErrors.ErrInvalidEmailFormat
	}

	// ✅ Step 2: Validate phone number format
	if !isValidPhone(organization.Phone) {
		return nil, hrmsErrors.ErrInvalidPhoneFormat
	}

	// ✅ Step 3: Check for existing organization by name and email
	var existingOrg domain.Organization
	if err := r.DB.WithContext(ctx).
		Where("email = ? OR name = ?", organization.Email, organization.Name).
		First(&existingOrg).Error; err == nil {
		r.Logger.Warn("⚠️ Organization already exists", zap.String("email", organization.Email))
		return nil, hrmsErrors.ErrOrganizationAlreadyExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// ✅ Step 4: Insert Organization
	if err := r.DB.WithContext(ctx).Create(organization).Error; err != nil {
		r.Logger.Error("❌ Failed to create organization", zap.Error(err))
		return nil, err
	}

	return organization, nil
}

// GetOrganizationByID 📌 Fetch organization by ID
func (r *HrmsRepository) GetOrganizationByID(ctx context.Context, orgID uint) (*domain.Organization, error) {
	var organization domain.Organization
	err := r.DB.WithContext(ctx).
		Preload("Employees").
		Preload("Departments").
		Where("id = ?", orgID).
		First(&organization).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrOrganizationNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetOrganizationByID", zap.Error(err))
		return nil, err
	}
	return &organization, nil
}

// UpdateOrganization 📌 Update organization details
func (r *HrmsRepository) UpdateOrganization(ctx context.Context, orgID uint, updates map[string]interface{}) error {
	// Prevent email modification
	delete(updates, "email")

	// Validate phone number if being updated
	if phone, exists := updates["phone"]; exists {
		if !isValidPhone(phone.(string)) {
			return hrmsErrors.ErrInvalidPhoneFormat
		}
	}

	if err := r.DB.WithContext(ctx).Model(&domain.Organization{}).
		Where("id = ?", orgID).
		Updates(updates).Error; err != nil {
		r.Logger.Error("❌ Failed to update organization", zap.Error(err))
		return err
	}
	return nil
}

// DeleteOrganization 📌 Soft delete organization
func (r *HrmsRepository) DeleteOrganization(ctx context.Context, orgID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	// ✅ Ensure no employees exist before deletion
	var employeeCount int64
	if err := tx.Model(&domain.Employee{}).Where("organization_id = ?", orgID).Count(&employeeCount).Error; err != nil {
		tx.Rollback()
		return err
	}
	if employeeCount > 0 {
		tx.Rollback()
		return hrmsErrors.ErrCannotDeleteOrgWithEmployees
	}

	// ✅ Soft Delete Organization
	if err := tx.Model(&domain.Organization{}).
		Where("id = ?", orgID).
		Update("deleted_at", gorm.DeletedAt{Time: time.Now(), Valid: true}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to soft delete organization", zap.Error(err))
		return err
	}

	tx.Commit()
	r.Logger.Info("✅ Organization soft deleted", zap.Uint("organization_id", orgID))
	return nil
}

// ListOrganizations 📌 Fetch organizations with pagination & search
func (r *HrmsRepository) ListOrganizations(ctx context.Context, limit, offset int, search string) ([]domain.Organization, int64, error) {
	var organizations []domain.Organization
	var totalCount int64

	query := r.DB.WithContext(ctx).Model(&domain.Organization{}).
		Preload("Departments").
		Order("created_at DESC") // Default sorting

	if search != "" {
		query = query.Where("name ILIKE ? OR email ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	query.Count(&totalCount)

	if err := query.Limit(limit).Offset(offset).Find(&organizations).Error; err != nil {
		r.Logger.Error("❌ Failed to fetch organizations", zap.Error(err))
		return nil, 0, err
	}

	return organizations, totalCount, nil
}

// Validate Email Format
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// Validate Phone Format
func isValidPhone(phone string) bool {
	re := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	return re.MatchString(phone)
}
