package repository

import (
	"errors"
	"sync"
	"time"

	"amaa/internal/models"
)

type AuditRepository interface {
	CreateAudit(audit *models.Audit) (*models.Audit, error)
	GetAuditHistory(assetID string) ([]*models.Audit, error)
}

type inMemoryAuditRepo struct {
	mu     sync.RWMutex
	audits map[string][]*models.Audit // keyed by AssetID
}

func NewAuditRepository() AuditRepository {
	return &inMemoryAuditRepo{
		audits: make(map[string][]*models.Audit),
	}
}

func (repo *inMemoryAuditRepo) CreateAudit(audit *models.Audit) (*models.Audit, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	audit.AuditDate = time.Now()
	repo.audits[audit.AssetID] = append(repo.audits[audit.AssetID], audit)
	return audit, nil
}

func (repo *inMemoryAuditRepo) GetAuditHistory(assetID string) ([]*models.Audit, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	history, exists := repo.audits[assetID]
	if !exists {
		return nil, errors.New("no audit records found")
	}
	return history, nil
}
