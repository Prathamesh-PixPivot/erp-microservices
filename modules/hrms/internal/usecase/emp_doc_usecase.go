package usecase

import (
	"context"
	"hrms/internal/dto"
	"hrms/internal/domain"

	"go.uber.org/zap"
)

// CreateEmployeeDocument adds a new employee document
func (u *HrmsUsecase) CreateEmployeeDocument(ctx context.Context, req dto.CreateEmployeeDocumentRequest) (*dto.EmployeeDocumentDTO, error) {
	document := domain.EmployeeDocument{
		EmployeeID:   req.EmployeeID,
		DocumentName: req.DocumentName,
		DocumentURL:  req.DocumentURL,
		ExpiryDate:   req.ExpiryDate,
	}

	result, err := u.HrmsRepo.CreateEmployeeDocument(ctx, &document)
	if err != nil {
		u.Logger.Error("Failed to create employee document", zap.Error(err))
		return nil, err
	}

	return mapToEmployeeDocumentDTO(result), nil
}

// GetEmployeeDocumentByID fetches a specific document by ID
func (u *HrmsUsecase) GetEmployeeDocumentByID(ctx context.Context, documentID uint) (*dto.EmployeeDocumentDTO, error) {
	document, err := u.HrmsRepo.GetEmployeeDocumentByID(ctx, documentID)
	if err != nil {
		u.Logger.Error("Failed to fetch employee document", zap.Error(err))
		return nil, err
	}

	return mapToEmployeeDocumentDTO(document), nil
}

// GetDocumentsByEmployee fetches all documents of an employee
func (u *HrmsUsecase) GetDocumentsByEmployee(ctx context.Context, employeeID uint) ([]dto.EmployeeDocumentDTO, error) {
	documents, err := u.HrmsRepo.GetDocumentsByEmployee(ctx, employeeID)
	if err != nil {
		u.Logger.Error("Failed to fetch documents for employee", zap.Error(err))
		return nil, err
	}

	dtoDocuments := make([]dto.EmployeeDocumentDTO, len(documents))
	for i, doc := range documents {
		dtoDocuments[i] = *mapToEmployeeDocumentDTO(&doc)
	}

	return dtoDocuments, nil
}

// GetExpiredDocuments fetches all expired documents
func (u *HrmsUsecase) GetExpiredDocuments(ctx context.Context) ([]dto.EmployeeDocumentDTO, error) {
	documents, err := u.HrmsRepo.GetExpiredDocuments(ctx)
	if err != nil {
		u.Logger.Error("Failed to fetch expired documents", zap.Error(err))
		return nil, err
	}

	dtoDocuments := make([]dto.EmployeeDocumentDTO, len(documents))
	for i, doc := range documents {
		dtoDocuments[i] = *mapToEmployeeDocumentDTO(&doc)
	}

	return dtoDocuments, nil
}

// UpdateEmployeeDocument updates a specific employee document
func (u *HrmsUsecase) UpdateEmployeeDocument(ctx context.Context, documentID uint, req dto.UpdateEmployeeDocumentRequest) error {
	updates := make(map[string]interface{})

	if req.DocumentName != "" {
		updates["document_name"] = req.DocumentName
	}
	if req.DocumentURL != "" {
		updates["document_url"] = req.DocumentURL
	}
	if req.ExpiryDate != nil {
		updates["expiry_date"] = req.ExpiryDate
	}

	err := u.HrmsRepo.UpdateEmployeeDocument(ctx, documentID, updates)
	if err != nil {
		u.Logger.Error("Failed to update employee document", zap.Error(err))
		return err
	}
	return nil
}

// DeleteEmployeeDocument removes a specific employee document
func (u *HrmsUsecase) DeleteEmployeeDocument(ctx context.Context, documentID uint) error {
	if err := u.HrmsRepo.DeleteEmployeeDocument(ctx, documentID); err != nil {
		u.Logger.Error("Failed to delete employee document", zap.Error(err))
		return err
	}
	return nil
}

// mapToEmployeeDocumentDTO converts domain model to DTO
func mapToEmployeeDocumentDTO(d *domain.EmployeeDocument) *dto.EmployeeDocumentDTO {
	return &dto.EmployeeDocumentDTO{
		ID:           d.ID,
		EmployeeID:   d.EmployeeID,
		DocumentName: d.DocumentName,
		DocumentURL:  d.DocumentURL,
		ExpiryDate:   d.ExpiryDate,
	}
}
