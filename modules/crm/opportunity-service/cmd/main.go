package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"opportunity-service/config"
	"opportunity-service/grpc/opportunitypb"
	"opportunity-service/internal/handler"
	"opportunity-service/internal/models"
	"opportunity-service/internal/repository"
	"opportunity-service/internal/services"
)

func main() {
	// Load config
	cfg := config.Load()
	log.Printf("Opportunity service configuration: %+v", cfg)

	// Initialize the GORM database
	db, err := gorm.Open(postgres.Open(cfg.DB), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// AutoMigrate the database based on the Opportunity model
	err = db.AutoMigrate(&models.Opportunity{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize services and handlers
	opportunityRepo := repository.NewOpportunityRepository(db)
	opportunityService := services.NewOpportunityService(opportunityRepo)
	opportunityHandler := handler.NewOpportunityHandler(opportunityService)

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the handler with the gRPC server
	opportunitypb.RegisterOpportunityServiceServer(grpcServer, opportunityHandler)

	// Start listening
	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", cfg.GRPCPort, err)
	}

	// Graceful shutdown
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down opportunity service...")
	grpcServer.GracefulStop()
	log.Println("Service stopped.")
}
