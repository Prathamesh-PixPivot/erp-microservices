package repository

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"hrms/internal/domain"
	hrmsErrors "hrms/internal/errors"
)

// CreateEmployeeDocument 📌 Add a new employee document
func (r *HrmsRepository) CreateEmployeeDocument(ctx context.Context, document *domain.EmployeeDocument) (*domain.EmployeeDocument, error) {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Create(document).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to create employee document", zap.Error(err))
		return nil, err
	}

	tx.Commit()
	return document, nil
}

// GetEmployeeDocumentByID 📌 Fetch a specific employee document by ID
func (r *HrmsRepository) GetEmployeeDocumentByID(ctx context.Context, documentID uint) (*domain.EmployeeDocument, error) {
	var document domain.EmployeeDocument
	err := r.DB.WithContext(ctx).Where("id = ?", documentID).First(&document).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hrmsErrors.ErrEmployeeDocumentNotFound
	} else if err != nil {
		r.Logger.Error("❌ Database error in GetEmployeeDocumentByID", zap.Error(err))
		return nil, err
	}
	return &document, nil
}

// GetDocumentsByEmployee 📌 Fetch all documents for a specific employee
func (r *HrmsRepository) GetDocumentsByEmployee(ctx context.Context, employeeID uint) ([]domain.EmployeeDocument, error) {
	var documents []domain.EmployeeDocument
	err := r.DB.WithContext(ctx).Where("employee_id = ?", employeeID).Find(&documents).Error
	if err != nil {
		r.Logger.Error("❌ Failed to fetch employee documents", zap.Uint("employee_id", employeeID), zap.Error(err))
		return nil, err
	}
	return documents, nil
}

// GetExpiredDocuments 📌 Fetch all expired employee documents
func (r *HrmsRepository) GetExpiredDocuments(ctx context.Context) ([]domain.EmployeeDocument, error) {
	var expiredDocs []domain.EmployeeDocument
	currentTime := time.Now()

	err := r.DB.WithContext(ctx).Where("expiry_date IS NOT NULL AND expiry_date < ?", currentTime).Find(&expiredDocs).Error
	if err != nil {
		r.Logger.Error("❌ Failed to fetch expired employee documents", zap.Error(err))
		return nil, err
	}
	return expiredDocs, nil
}

// UpdateEmployeeDocument 📌 Update a specific employee document
func (r *HrmsRepository) UpdateEmployeeDocument(ctx context.Context, documentID uint, updates map[string]interface{}) error {
	if err := r.DB.WithContext(ctx).Model(&domain.EmployeeDocument{}).
		Where("id = ?", documentID).
		Updates(updates).Error; err != nil {
		r.Logger.Error("❌ Failed to update employee document", zap.Error(err))
		return err
	}
	return nil
}

// DeleteEmployeeDocument 📌 Soft delete an employee document
func (r *HrmsRepository) DeleteEmployeeDocument(ctx context.Context, documentID uint) error {
	tx := r.DB.WithContext(ctx).Begin()

	if err := tx.Where("id = ?", documentID).Delete(&domain.EmployeeDocument{}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error("❌ Failed to delete employee document", zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}
