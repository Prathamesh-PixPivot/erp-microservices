package services

import (
	"amaa/internal/models"
	"amaa/internal/repository"
)

type DisposalService interface {
	DecommissionAsset(disposal *models.Disposal) (*models.Disposal, error)
}

type disposalService struct {
	repo repository.DisposalRepository
}

func NewDisposalService(repo repository.DisposalRepository) DisposalService {
	return &disposalService{repo: repo}
}

func (s *disposalService) DecommissionAsset(disposal *models.Disposal) (*models.Disposal, error) {
	return s.repo.DecommissionAsset(disposal)
}
