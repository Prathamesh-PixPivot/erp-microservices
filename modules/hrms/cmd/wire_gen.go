// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"hrms/config"
	"hrms/infrastructure"
	"hrms/internal/repository"
	"hrms/internal/transport/grpc/handler"
	"hrms/internal/usecase"
)

// Injectors from di.go:

// InitializeDependencies sets up the entire dependency graph
func InitializeDependencies() (*Dependencies, error) {
	config, err := InitializeConfig()
	if err != nil {
		return nil, err
	}
	logger := InitializeLogger(config)
	database, err := InitializeDatabase(config, logger)
	if err != nil {
		return nil, err
	}
	hrmsRepository := ProvideHrmsRepository(database, logger)
	hrmsUsecase := ProvideHrmsUsecase(hrmsRepository, logger)
	hrmsHandler := InitializeGRPCServer(hrmsUsecase, logger)
	dependencies := &Dependencies{
		Config:     config,
		Logger:     logger,
		Database:   database,
		GRPCServer: hrmsHandler,
	}
	return dependencies, nil
}

// di.go:

type Dependencies struct {
	Config     *config.Config
	Logger     *zap.Logger
	Database   *infrastructure.Database
	GRPCServer *grpc.HrmsHandler // ✅ Fix: Use gRPC service implementation
}

// InitializeConfig provides the app configuration
func InitializeConfig() (*config.Config, error) {
	return config.LoadConfig()
}

// InitializeLogger sets up the logger using the entire `cfg` object
func InitializeLogger(cfg *config.Config) *zap.Logger {
	infrastructure.InitLogger(cfg)
	return infrastructure.GetLogger()
}

// InitializeDatabase sets up the PostgreSQL database connection
func InitializeDatabase(cfg *config.Config, logger *zap.Logger) (*infrastructure.Database, error) {
	return infrastructure.NewDatabase(cfg, logger)
}

// ProvideHrmsUsecase initializes HrmsUsecase with HrmsRepository
func ProvideHrmsUsecase(hrmsRepo *repository.HrmsRepository, logger *zap.Logger) *usecase.HrmsUsecase {
	return usecase.NewHrmsUsecase(hrmsRepo, logger)
}

// ProvideHrmsRepository initializes HrmsRepository
func ProvideHrmsRepository(db *infrastructure.Database, logger *zap.Logger) *repository.HrmsRepository {
	return repository.NewHrmsRepository(db.DB, logger)
}

// InitializeGRPCServer initializes gRPC HrmsService
func InitializeGRPCServer(hrmsUsecase *usecase.HrmsUsecase, logger *zap.Logger) *grpc.HrmsHandler {
	return grpc.NewHrmsGrpcHandler(hrmsUsecase, logger)
}

// WireSet groups dependencies for Wire to generate them automatically
var WireSet = wire.NewSet(
	InitializeConfig,
	InitializeLogger,
	InitializeDatabase,
	InitializeGRPCServer,
	ProvideHrmsRepository,
	ProvideHrmsUsecase,
)
