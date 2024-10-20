// main.go

package main

import (
	"contacts-service/config"
	"contacts-service/grpc/contactpb"
	"contacts-service/internal/handler"
	"contacts-service/internal/models"
	"contacts-service/internal/repository"
	"contacts-service/internal/services"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize Database Connection
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

	// Auto Migrate the Contact model
	if err := db.AutoMigrate(&models.Contact{}); err != nil {
		log.Fatalf("Failed to auto-migrate Contact model: %v", err)
	}
	log.Println("Database connection established and auto-migrated.")

	// Initialize Repository
	contactRepo := repository.NewContactRepository(db)

	// Initialize Service
	contactService := services.NewContactService(contactRepo)

	// Initialize gRPC Handler
	contactHandler := handler.NewContactHandler(contactService)

	// Initialize gRPC Server
	grpcServer := grpc.NewServer()

	// Register ContactService Server
	contactpb.RegisterContactServiceServer(grpcServer, contactHandler)
	log.Println("ContactService registered with gRPC server.")

	// Start Listening on the specified port
	address := fmt.Sprintf(":%d", cfg.GRPCPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", cfg.GRPCPort, err)
	}
	log.Printf("gRPC server listening on %s", address)

	// Start the gRPC server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
