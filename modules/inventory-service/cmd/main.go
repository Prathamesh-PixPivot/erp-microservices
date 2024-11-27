package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"inventory-service/config"
	invHandler "inventory-service/grpc/handler" // Alias to avoid conflict with grpc package
	pb "inventory-service/grpc/inventory_pb"    // Alias to `pb` for clarity
	"inventory-service/internal/models"
	"inventory-service/internal/repository"
	invService "inventory-service/internal/services" // Now correctly importing services
)

func main() {
	// Load configurations
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database connection
	db, err := initDatabase(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize repository and services
	inventoryRepo := repository.NewInventoryRepository(db)
	inventoryService := invService.NewInventoryService(inventoryRepo) // Correct usage of services package

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	inventoryServer := invHandler.NewInventoryServiceServer(inventoryService)
	pb.RegisterInventoryServiceServer(grpcServer, inventoryServer) // Registering using pb alias

	// Start gRPC server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Server.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", cfg.Server.GRPCPort, err)
	}
	log.Printf("Starting Inventory Service on port %s...", cfg.Server.GRPCPort)

	// Run gRPC server in a separate goroutine
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// Graceful shutdown
	waitForShutdown(grpcServer)
}

// initDatabase initializes the database connection with GORM
func initDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Auto-migrate database tables
	if err := db.AutoMigrate(&models.InventoryItem{}, &models.WarehouseStock{}); err != nil {
		return nil, fmt.Errorf("failed to auto-migrate database tables: %v", err)
	}
	return db, nil
}

// waitForShutdown waits for interrupt signals to gracefully shutdown the server
func waitForShutdown(grpcServer *grpc.Server) {
	// Listen for termination signals
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	<-ch
	log.Println("Shutting down gracefully...")
	grpcServer.GracefulStop()
	log.Println("Inventory Service stopped.")
}
