package migrations

import (
	"itsm/internal/models"

	"go.uber.org/zap"

	"gorm.io/gorm"
)

// Migrate runs database migrations
func Migrate(db *gorm.DB, logger *zap.Logger) {
	err := db.AutoMigrate(
		// Core entities
		&models.Incident{},
		&models.ServiceRequest{},
		&models.Change{},
	)

	if err != nil {
		logger.Fatal("❌ Migration failed", zap.Error(err))
	}
	logger.Info("✅ Database migration completed successfully")
}
