package database

import (
	"fmt"
	"gst-service/internal/domain"
	"gst-service/internal/infrastructure/config"
	"log"

	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB global database instance
var DB *gorm.DB

// ConnectDatabase initializes the PostgreSQL connection
func ConnectDatabase(cfg *config.Config) error {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.SSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to configure database connection: %w", err)
	}
	sqlDB.SetMaxOpenConns(cfg.Database.MaxConnections)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConnections)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ Connected to the database successfully")
	return nil
}

// MigrateDatabase runs the schema migrations for GSTR tables
func MigrateDatabase() error {
	err := DB.AutoMigrate(
		&domain.Invoice{},       // Migrate Invoice table first ✅
		&domain.GSTR1Request{},  // Then GSTR1 ✅
		&domain.GSTR1ARequest{}, // Then GSTR1A ✅
		&domain.GSTR2ARequest{}, // Then GSTR2A ✅
		&domain.GSTR3BRequest{}, // Then GSTR3B ✅
		&domain.GSTR9Request{},  // Then GSTR9 ✅
		&domain.GSTR9CRequest{}, // Then GSTR9C ✅
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	log.Println("✅ Database migrations completed successfully")
	return nil
}
