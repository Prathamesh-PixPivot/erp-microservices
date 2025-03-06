package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"hrms/config"
	handlergrpc "hrms/internal/transport/grpc/handler"
	proto "hrms/internal/transport/grpc/proto"
)

// Dependencies struct to hold injected dependencies
type Dependencies struct {
	Config     *config.Config
	Logger     *zap.Logger
	GRPCServer *handlergrpc.HrmsHandler
}

// StartServer initializes and runs the gRPC server
func StartServer(deps *Dependencies) {
	// ✅ Load gRPC port from config
	grpcPort := strconv.Itoa(deps.Config.Server.GRPCPort)
	logger := deps.Logger
	logger.Info("🔁 Starting the gRPC server", zap.String("port", grpcPort))

	// Create a listener on the configured port
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		logger.Fatal("❌ Failed to listen on gRPC port", zap.Error(err))
	}

	// Initialize the gRPC server
	grpcServer := grpc.NewServer()

	// ✅ Register the single HrmsHandler instance for all services
	proto.RegisterAttendanceServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterBonusServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterDepartmentServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterDesignationServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterEmployeeBenefitServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterEmployeeDocumentServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterEmployeeExitServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterEmployeePerkServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterEmployeeServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterExpenseServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterLeaveBalanceServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterLeaveServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterLeavePolicyServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterLoanAdvanceServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterOrganizationServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterPayrollServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterPerformanceKPIServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterPerformanceReviewServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterPublicHolidayServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterSalaryStructureServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterShiftServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterSkillDevelopmentServiceServer(grpcServer, deps.GRPCServer)
	proto.RegisterWorkHistoryServiceServer(grpcServer, deps.GRPCServer)

	// ✅ Enable reflection for debugging with grpcurl
	reflection.Register(grpcServer)

	// Run the server in a goroutine to avoid blocking
	go func() {
		logger.Info("🚀 gRPC server started", zap.String("port", grpcPort))
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal("❌ Failed to serve gRPC server", zap.Error(err))
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

	logger.Info("🛑 Gracefully shutting down gRPC server...")

	// Allow 5 seconds for ongoing requests to complete before forcing shutdown
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Stop the gRPC server
	grpcServer.GracefulStop()

	logger.Info("✅ gRPC server stopped cleanly")
}
