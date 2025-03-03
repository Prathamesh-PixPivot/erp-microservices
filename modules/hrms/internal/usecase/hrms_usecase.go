package usecase

import (
	"go.uber.org/zap"
	"hrms/internal/repository"
)

// HrmsUsecase defines business logic for the HRMS system
type HrmsUsecase struct {
	HrmsRepo *repository.HrmsRepository
	Logger   *zap.Logger
}

// NewHrmsUsecase initializes the HRMS usecase
func NewHrmsUsecase(hrmsRepo *repository.HrmsRepository, logger *zap.Logger) *HrmsUsecase {
	return &HrmsUsecase{
		HrmsRepo: hrmsRepo,
		Logger:   logger,
	}
}
