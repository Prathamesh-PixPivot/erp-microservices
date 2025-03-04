package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// HrmsRepository provides database operations for users
type HrmsRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

// NewHrmsRepository initializes HrmsRepository
func NewHrmsRepository(db *gorm.DB, logger *zap.Logger) *HrmsRepository {
	return &HrmsRepository{DB: db, Logger: logger}
}
