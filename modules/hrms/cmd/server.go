package main

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
	"hrms/internal/transport/grpc/proto"
)

func StartServer(deps *Dependencies) {
	// ✅ Get gRPC port from config
	grpcPort := strconv.Itoa(deps.Config.Server.GRPCPort)
	logger := deps.Logger
	logger.Info("🔁 Starting the GRPC server", zap.String("port", grpcPort))

	// Create a listener on the configured port
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		logger.Fatal("Failed to listen on gRPC port", zap.Error(err))
	}

	// Initialize the gRPC server
	grpcServer := grpc.NewServer()

	// Register gRPC services
	proto.RegisterHrmsServiceServer(grpcServer, deps.GRPCServer)

	// Enable reflection for debugging tools like grpcurl
	reflection.Register(grpcServer)

	// Use a goroutine to start the server with logger
	go func() {
		logger.Info("🚀 gRPC server started", zap.String("port", grpcPort))
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal("Failed to serve gRPC server", zap.Error(err))
		}
	}()

	// Handle graceful shutdown
	gracefulShutdown(grpcServer, logger)
}

// gracefulShutdown ensures that the gRPC server stops cleanly
func gracefulShutdown(grpcServer *grpc.Server, logger *zap.Logger) {
	// Capture termination signals (Ctrl+C, kill command, SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit // Wait here until an OS signal is received

	logger.Info("🔄 Gracefully shutting down gRPC server...")

	// Allow 5 seconds for ongoing requests to complete before forcing shutdown
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Stop the gRPC server
	grpcServer.GracefulStop()

	logger.Info("✅ gRPC server stopped")
}
