// activity-task-service/main.go

package main

import (
	"activity-task-service/config"
	"activity-task-service/grpc/activitypb"
	"activity-task-service/internal/handler"
	"activity-task-service/internal/models"
	"activity-task-service/internal/repository"
	"activity-task-service/internal/services"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Step 1: Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Step 2: Initialize Database Connection
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL database: %v", err)
	}

	// Step 3: Auto-Migrate Models
	if err := models.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}
	log.Println("Database connected and models auto-migrated.")

	// Step 4: Initialize Repository and Service Layers
	activityRepo := repository.NewActivityRepository(db)
	activityService := services.NewActivityService(activityRepo)

	// Step 5: Initialize gRPC Handler
	activityHandler := handler.NewActivityHandler(activityService)

	// Step 6: Initialize gRPC Server
	grpcServer := grpc.NewServer()

	// Step 7: Register ActivityService Server
	activitypb.RegisterActivityServiceServer(grpcServer, activityHandler)
	log.Println("ActivityService registered with gRPC server.")

	// Step 8: Start Listening on Specified Port
	address := fmt.Sprintf(":%d", cfg.GRPCPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", cfg.GRPCPort, err)
	}
	log.Printf("gRPC server listening on %s", address)

	// Step 9: Serve gRPC Server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
