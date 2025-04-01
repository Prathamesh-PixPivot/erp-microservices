package infrastructure

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"itsm/config"
	"itsm/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database holds the DB connection instance.
type Database struct {
	DB *gorm.DB
}

// NewDatabase initializes the PostgreSQL database connection and configures the connection pool.
func NewDatabase(cfg *config.Config, logger *zap.Logger) (*Database, error) {
	// Construct DSN (Database Source Name)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	// Connect to PostgreSQL with GORM.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("failed to connect to database", zap.Error(err))
		return nil, err
	}

	// Configure Connection Pooling
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("failed to get underlying DB", zap.Error(err))
		return nil, err
	}
	sqlDB.SetMaxOpenConns(50)                 // Maximum open connections.
	sqlDB.SetMaxIdleConns(10)                 // Maximum idle connections.
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Maximum connection lifetime.

	// AutoMigrate models: Incident, Change, and ServiceRequest.
	err = db.AutoMigrate(&models.Incident{}, &models.Change{}, &models.ServiceRequest{})
	if err != nil {
		logger.Error("failed to auto-migrate models", zap.Error(err))
		return nil, err
	}

	logger.Info("✅ Database connection successfully established!")
	return &Database{DB: db}, nil
}
