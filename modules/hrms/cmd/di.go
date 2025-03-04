//go:build wireinject
// +build wireinject

package main

import (
	"user-service/config"
	"user-service/infrastructure"
	"user-service/internal/repository"
	"user-service/internal/transport/grpc"
	"user-service/internal/usecase"

	"github.com/google/wire"
	"go.uber.org/zap"
)

type Dependencies struct {
	Config     *config.Config
	Logger     *zap.Logger
	Database   *infrastructure.Database
	GRPCServer *grpc.UserGrpcHandler // ✅ Fix: Use gRPC service implementation
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

// ProvideHrmsUsecase initializes UserUsecase with UserRepository
func ProvideHrmsUsecase(userRepo *repository.UserRepository, logger *zap.Logger) *usecase.UserUsecase {
	return usecase.NewUserUsecase(userRepo, logger)
}

// ProvideHrmsRepository initializes UserRepository
func ProvideHrmsRepository(db *infrastructure.Database, logger *zap.Logger) *repository.UserRepository {
	return repository.NewUserRepository(db.DB, logger)
}

// InitializeGRPCServer initializes gRPC UserService
func InitializeGRPCServer(userUseCase *usecase.UserUsecase, logger *zap.Logger) *grpc.UserGrpcHandler {
	return grpc.NewUserGrpcHandler(userUseCase, logger) // ✅ Fix: Correctly initialize gRPC server
}

// WireSet groups dependencies for Wire to generate them automatically
var WireSet = wire.NewSet(
	InitializeConfig,      // ✅ Load Config
	InitializeLogger,      // ✅ Initialize Logger
	InitializeDatabase,    // ✅ Initialize Database
	InitializeGRPCServer,  // ✅ Initialize gRPC server
	ProvideHrmsRepository, // ✅ Initialize UserRepository
	ProvideHrmsUsecase,    // ✅ Initialize UserUsecase
)

// InitializeDependencies sets up the entire dependency graph
func InitializeDependencies() (*Dependencies, error) {
	wire.Build(WireSet, wire.Struct(new(Dependencies), "Config", "Logger", "Database", "GRPCServer")) // ✅ Inject dependencies
	return &Dependencies{}, nil
}
