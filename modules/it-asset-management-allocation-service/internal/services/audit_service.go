package services

import (
	"amaa/internal/models"
	"amaa/internal/repository"
)

type AuditService interface {
	AuditAsset(audit *models.Audit) (*models.Audit, error)
	GetAuditHistory(assetID string) ([]*models.Audit, error)
}

type auditService struct {
	repo repository.AuditRepository
}

func NewAuditService(repo repository.AuditRepository) AuditService {
	return &auditService{repo: repo}
}

func (s *auditService) AuditAsset(audit *models.Audit) (*models.Audit, error) {
	return s.repo.CreateAudit(audit)
}

func (s *auditService) GetAuditHistory(assetID string) ([]*models.Audit, error) {
	return s.repo.GetAuditHistory(assetID)
}
