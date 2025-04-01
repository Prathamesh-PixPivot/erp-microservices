//go:build wireinject
// +build wireinject

package main

import (
	"amaa/config"
	"amaa/infrastructure"
	"amaa/internal/repository"
	"amaa/internal/services"

	// Removed unused import

	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Dependencies holds the application-level dependencies.
type Dependencies struct {
	Config     *config.Config
	Logger     *zap.Logger
	Database   *infrastructure.Database
	GRPCServer *grpc.Server
	Services   struct {
		AssetService       services.AssetService
		AllocationService  services.AllocationService
		MaintenanceService services.MaintenanceService
		AuditService       services.AuditService
		DisposalService    services.DisposalService
		LicenseService     services.LicenseService
	}
}

// InitializeConfig loads the configuration.
func InitializeConfig() (*config.Config, error) {
	return config.LoadConfig()
}

// InitializeLogger sets up the logger using the full configuration.
func InitializeLogger(cfg *config.Config) *zap.Logger {
	return zap.NewExample()
}

// InitializeDatabase creates a new database connection.
func InitializeDatabase(cfg *config.Config, logger *zap.Logger) (*infrastructure.Database, error) {
	return infrastructure.NewDatabase(cfg, logger)
}

// ProvideAssetRepository initializes the AssetRepository.
func ProvideAssetRepository(db *infrastructure.Database, logger *zap.Logger) repository.AssetRepository {
	return repository.NewAssetRepository(db.DB)
}

// ProvideAssetUsecase initializes the AssetService with AssetRepository.
func ProvideAssetUsecase(assetRepo repository.AssetRepository, logger *zap.Logger) services.AssetService {
	return services.NewAssetService(assetRepo)
}

// ProvideAllocationRepository initializes the AllocationRepository.
func ProvideAllocationRepository(db *infrastructure.Database, logger *zap.Logger) repository.AllocationRepository {
	return repository.NewAllocationRepository()
}

// ProvideAllocationUsecase initializes the AllocationService with AllocationRepository.
func ProvideAllocationUsecase(allocRepo repository.AllocationRepository, logger *zap.Logger) services.AllocationService {
	return services.NewAllocationService(allocRepo)
}

// ProvideMaintenanceRepository initializes the MaintenanceRepository.
func ProvideMaintenanceRepository(db *infrastructure.Database, logger *zap.Logger) repository.MaintenanceRepository {
	return repository.NewMaintenanceRepository()
}

// ProvideMaintenanceService initializes the MaintenanceService.
func ProvideMaintenanceService(repo repository.MaintenanceRepository, logger *zap.Logger) services.MaintenanceService {
	return services.NewMaintenanceService(repo)
}

// ProvideAuditService initializes the AuditService.
func ProvideAuditService(repo repository.AuditRepository, logger *zap.Logger) services.AuditService {
	return services.NewAuditService(repo)
}

// ProvideAuditRepository initializes the AuditRepository.
func ProvideAuditRepository(db *infrastructure.Database, logger *zap.Logger) repository.AuditRepository {
	return repository.NewAuditRepository()
}

// ProvideAuditUsecase initializes the AuditService with AuditRepository.
func ProvideAuditUsecase(auditRepo repository.AuditRepository, logger *zap.Logger) services.AuditService {
	return services.NewAuditService(auditRepo)
}

// InitializeDispoalService initializes the DisposalService with DisposalRepository.
func ProvideDisposalRepository(db *infrastructure.Database, logger *zap.Logger) repository.DisposalRepository {
	return repository.NewDisposalRepository()
}

// ProvideDisposalUsecase initializes the DisposalService with DisposalRepository.
func ProvideDisposalUsecase(disposalRepo repository.DisposalRepository, logger *zap.Logger) services.DisposalService {
	return services.NewDisposalService(disposalRepo)
}

// InitializelicenseService initializes the LicenseService with LicenseRepository.
func ProvideLicenseRepository(db *infrastructure.Database, logger *zap.Logger) repository.LicenseRepository {
	return repository.NewLicenseRepository()
}

// ProvideLicenseUsecase initializes the LicenseService with LicenseRepository.
func ProvideLicenseUsecase(licenseRepo repository.LicenseRepository, logger *zap.Logger) services.LicenseService {
	return services.NewLicenseService(licenseRepo)
}

// ProvideLicenseService initializes the LicenseService with LicenseRepository.
func ProvideLicenseService(licenseRepo repository.LicenseRepository, logger *zap.Logger) services.LicenseService {
	return services.NewLicenseService(licenseRepo)
}

// InitializeGRPCServer initializes gRPC Service
func InitializeGRPCServer(allocService services.AllocationService, assetService services.AssetService, maintenanceSrv services.MaintenanceService, auditService services.AuditService, disposalService services.DisposalService, licenseService services.LicenseService, logger *zap.Logger) *grpc.Server {
	return grpc.NewServer()
}

func InitializeServices(assetService services.AssetService, allocService services.AllocationService, maintenanceService services.MaintenanceService, auditService services.AuditService, disposalService services.DisposalService, licenseService services.LicenseService) struct {
	AssetService       services.AssetService
	AllocationService  services.AllocationService
	MaintenanceService services.MaintenanceService
	AuditService       services.AuditService
	DisposalService    services.DisposalService
	LicenseService     services.LicenseService
} {
	return struct {
		AssetService       services.AssetService
		AllocationService  services.AllocationService
		MaintenanceService services.MaintenanceService
		AuditService       services.AuditService
		DisposalService    services.DisposalService
		LicenseService     services.LicenseService
	}{
		AssetService:       assetService,
		AllocationService:  allocService,
		MaintenanceService: maintenanceService,
		AuditService:       auditService,
		DisposalService:    disposalService,
		LicenseService:     licenseService,
	}
}

// WireSet groups all dependency providers.
var WireSet = wire.NewSet(
	InitializeConfig,             // Load configuration
	InitializeDatabase,           // Initialize database connection
	InitializeLogger,             // Initialize logger
	InitializeGRPCServer,         // Initialize gRPC server with handlers
	InitializeServices,           // Initialize all services
	ProvideMaintenanceService,    // Initialize MaintenanceService
	ProvideMaintenanceRepository, // Initialize MaintenanceRepository
	ProvideAssetRepository,       // Initialize AssetRepository
	ProvideAssetUsecase,          // Initialize AssetService
	ProvideAllocationRepository,  // Initialize AllocationRepository
	ProvideAllocationUsecase,     // Initialize AllocationService
	ProvideAuditService,          // Initialize AuditService
	ProvideAuditRepository,       // Initialize AuditRepository
	ProvideDisposalUsecase,       // Initialize DisposalService
	ProvideDisposalRepository,
	ProvideLicenseService,
	ProvideLicenseRepository,
)

// InitializeDependencies sets up the entire dependency graph.
func InitializeDI() (*Dependencies, error) {
	wire.Build(WireSet, wire.Struct(new(Dependencies), "Config", "Logger", "Database", "GRPCServer", "Services"))
	return &Dependencies{}, nil
}
