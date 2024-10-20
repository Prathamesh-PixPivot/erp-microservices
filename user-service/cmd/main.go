package main

import (
	"log"
	"net"
	"user-service/config"
	"user-service/grpc/userpb"
	"user-service/internal/handler"
	"user-service/internal/repository"

	"google.golang.org/grpc"
)

func main() {
	// Load configuration
	config.InitConfig()

	// Initialize the database connection
	config.InitDB()

	// Create repository
	userRepo := repository.NewUserRepository(config.DB)

	// Create gRPC server and register UserServiceServer
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	userServiceServer := handler.NewUserServiceServer(userRepo)

	userpb.RegisterUserServiceServer(grpcServer, userServiceServer)

	log.Println("User service is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
