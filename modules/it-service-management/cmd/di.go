//go:build wireinject
// +build wireinject

package main

import (
	"fmt"

	"github.com/google/wire"
	"go.uber.org/zap"

	"itsm/config"
	"itsm/internal/models"
	"itsm/internal/repository"
	"itsm/internal/services"
	grpcTransport "itsm/internal/transport/grpc/handler"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Dependencies holds all the dependencies for the ITSM services.
type Dependencies struct {
	Config      *config.Config
	DB          *gorm.DB
	Logger      *zap.Logger
	ITSMService services.ITSMService
	ITSMServer  *grpcTransport.ITSMServer
}

// NewZapLogger creates a new production Zap logger.
func NewZapLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

// NewDB initializes the PostgreSQL database connection using GORM.
// It constructs the DSN from the config and auto-migrates the ITSM models.
func NewDB(cfg *config.Config, logger *zap.Logger) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}
	// AutoMigrate the ITSM models: Incident, Change, and ServiceRequest.
	db.AutoMigrate(
		&models.Incident{},
		&models.Change{},
		&models.ServiceRequest{},
	)
	return db
}

// NewITSMService creates the ITSM service using GORM-based repositories.
func NewITSMService(db *gorm.DB) services.ITSMService {
	incidentRepo := repository.NewIncidentRepository(db)
	changeRepo := repository.NewChangeRepository(db)
	srRepo := repository.NewServiceRequestRepository()
	return services.NewITSMService(incidentRepo, changeRepo, srRepo)
}

// NewITSMServer creates the gRPC ITSM server handler.
func NewITSMServer(cfg *config.Config, itsmSrv services.ITSMService, logger *zap.Logger) *grpcTransport.ITSMServer {
	return grpcTransport.NewITSMServer(":"+cfg.Port, itsmSrv, logger)
}

var wireSet = wire.NewSet(
	config.LoadConfig, // loads config from file/env
	NewZapLogger,
	NewDB,
	NewITSMService,
	NewITSMServer,
)

// InitializeDependencies is the Wire injector function.
func InitializeDependencies() (*Dependencies, error) {
	wire.Build(
		wireSet,
		wire.Struct(new(Dependencies), "Config", "DB", "Logger", "ITSMService", "ITSMServer"),
	)
	return &Dependencies{}, nil // This line is never reached.
}
