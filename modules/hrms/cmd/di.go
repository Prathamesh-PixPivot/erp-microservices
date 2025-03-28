//go:build wireinject
// +build wireinject

package main

import (
	"hrms/config"
	"hrms/infrastructure"
	"hrms/internal/repository"
	grpc "hrms/internal/transport/grpc/handler"
	"hrms/internal/usecase"

	"github.com/google/wire"
	"go.uber.org/zap"
)

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
	infrastructure.InitLogger(cfg) // ✅ Pass full configuration
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
	return grpc.NewHrmsGrpcHandler(hrmsUsecase, logger) // ✅ Fix: Correctly initialize gRPC server
}

// WireSet groups dependencies for Wire to generate them automatically
var WireSet = wire.NewSet(
	InitializeConfig,      // ✅ Load Config
	InitializeLogger,      // ✅ Initialize Logger
	InitializeDatabase,    // ✅ Initialize Database
	InitializeGRPCServer,  // ✅ Initialize gRPC server
	ProvideHrmsRepository, // ✅ Initialize HrmsRepository
	ProvideHrmsUsecase,    // ✅ Initialize HrmsUsecase
)

// InitializeDependencies sets up the entire dependency graph
func InitializeDependencies() (*Dependencies, error) {
	wire.Build(WireSet, wire.Struct(new(Dependencies), "Config", "Logger", "Database", "GRPCServer")) // ✅ Inject dependencies
	return &Dependencies{}, nil
}
