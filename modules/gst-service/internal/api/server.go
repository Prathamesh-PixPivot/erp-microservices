package service

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"gst-service/internal/api/handler"
	pb "gst-service/internal/api/protobufs/gst-service"
	"gst-service/internal/infrastructure/config"
	"gst-service/internal/infrastructure/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// StartGRPCServer initializes and starts the gRPC server
func StartGRPCServer(cfg *config.Config, gstr1Handler *handler.GSTR1Handler, gstr1AHandler *handler.GSTR1AHandler, gstr2AHandler *handler.GSTR2AHandler, gstr3BHandler *handler.GSTR3BHandler, gstr9Handler *handler.GSTR9Handler, gstr9CHandler *handler.GSTR9CHandler, reconciliationHandler *handler.ReconciliationHandler) error {
	grpcAddress := fmt.Sprintf(":%d", cfg.Server.GRPCPort)
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return fmt.Errorf("failed to start gRPC listener: %w", err)
	}

	grpcServer := grpc.NewServer()

	// Register Services
	pb.RegisterGSTR1ServiceServer(grpcServer, gstr1Handler)
	pb.RegisterGSTR1AServiceServer(grpcServer, gstr1AHandler)
	pb.RegisterGSTR2AServiceServer(grpcServer, gstr2AHandler)
	pb.RegisterGSTR3BServiceServer(grpcServer, gstr3BHandler)
	pb.RegisterGSTR9ServiceServer(grpcServer, gstr9Handler)
	pb.RegisterGSTR9CServiceServer(grpcServer, gstr9CHandler)

	go func() {
		logger.AppLogger.Info("🚀 gRPC Server Started", zap.String("address", grpcAddress))
		if err := grpcServer.Serve(listener); err != nil {
			logger.AppLogger.Fatal("❌ gRPC Server failed", zap.Error(err))
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.AppLogger.Info("⚠️ Shutting down gRPC server...")
	grpcServer.GracefulStop()
	logger.AppLogger.Info("✅ gRPC Server Stopped Successfully")
	logger.CloseLoggers() // ✅ Ensures all logs are flushed before shutdown

	return nil
}
