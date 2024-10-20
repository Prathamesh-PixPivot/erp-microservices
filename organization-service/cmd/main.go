package main

import (
	"log"
	"net"
	"organization-service/config"
	"organization-service/grpc/organizationpb"
	"organization-service/internal/handler"

	"google.golang.org/grpc"
)

func main() {
	// Load configuration
	config.InitConfig()

	// Initialize database connection
	config.InitDB()

	// Start gRPC server
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}

	grpcServer := grpc.NewServer()
	organizationServiceServer := handler.NewOrganizationServiceServer() // Correct usage here

	organizationpb.RegisterOrganizationServiceServer(grpcServer, organizationServiceServer)

	log.Println("Organization service is running on port 50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
