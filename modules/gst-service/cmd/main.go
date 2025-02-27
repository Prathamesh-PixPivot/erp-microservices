package main

import (
	"fmt"
	apiService "gst-service/internal/api"
	"gst-service/internal/api/handler"
	"gst-service/internal/infrastructure/config"
	"gst-service/internal/infrastructure/database"
	"gst-service/internal/infrastructure/external"
	"gst-service/internal/infrastructure/logger"
	"gst-service/internal/repository"
	"gst-service/internal/service"
	"log"

	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize Logger
	if err := logger.InitializeLoggers(cfg); err != nil {
		log.Fatalf("Logger initialization failed: %v", err)
	}
	defer logger.CloseLoggers() // ✅ Ensure logs are flushed before exit

	// Connect to Database
	if err := database.ConnectDatabase(cfg); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Run Database Migrations
	if err := database.MigrateDatabase(); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	// Initialize Repositories
	gstr1Repo := repository.NewGSTRRepository(database.DB)
	gstr1ARepo := repository.NewGSTR1ARepository(database.DB)
	gstr2ARepo := repository.NewGSTR2ARepository(database.DB)
	gstr3BRepo := repository.NewGSTR3BRepository(database.DB)
	gstr9Repo := repository.NewGSTR9Repository(database.DB)
	gstr9CRepo := repository.NewGSTR9CRepository(database.DB)

	// Initialize Services
	gstr1Service := service.NewGSTR1Service(gstr1Repo)
	gstr1AService := service.NewGSTR1AService(gstr1ARepo)
	gstr2AService := service.NewGSTR2AService(gstr2ARepo)
	gstr3BService := service.NewGSTR3BService(gstr3BRepo)
	gstr9Service := service.NewGSTR9Service(gstr9Repo)
	gstr9CService := service.NewGSTR9CService(gstr9CRepo)

	// Initialize Handlers with their respective services
	gstr1Handler := handler.NewGSTR1Handler(gstr1Service)
	gstr1AHandler := handler.NewGSTR1AHandler(gstr1AService)
	gstr2AHandler := handler.NewGSTR2AHandler(gstr2AService)
	gstr3BHandler := handler.NewGSTR3BHandler(gstr3BService)
	gstr9Handler := handler.NewGSTR9Handler(gstr9Service)
	gstr9CHandler := handler.NewGSTR9CHandler(gstr9CService)

	// Initialize External Services
	gstClient := external.NewGSTClient(cfg)

	// Initialize Repositories
	gstr1Repo = repository.NewGSTRRepository(database.DB)
	gstr2ARepo = repository.NewGSTR2ARepository(database.DB)

	// Initialize Services
	reconciliationService := service.NewReconciliationService(gstr1Repo, gstr2ARepo, gstClient)

	// Initialize Handlers
	reconciliationHandler := handler.NewReconciliationHandler(reconciliationService)

	// Start gRPC Server
	if err := apiService.StartGRPCServer(cfg, gstr1Handler, gstr1AHandler, gstr2AHandler, gstr3BHandler, gstr9Handler, gstr9CHandler, reconciliationHandler); err != nil {
		logger.AppLogger.Fatal("❌ gRPC Server startup failed", zap.Error(err))
	}

	fmt.Printf("gRPC Server running on port: %d\n", cfg.Server.GRPCPort)
	fmt.Printf("Database host: %s\n", cfg.Database.Host)
	logger.AppLogger.Info("GST Microservice Started Successfully")
}
