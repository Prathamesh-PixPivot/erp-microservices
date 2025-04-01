package main

import (
	"amaa/internal/transport/grpc/handler"
	pb "amaa/internal/transport/grpc/proto"
	"net"
	"strconv"

	"google.golang.org/grpc"

	"go.uber.org/zap"
)

// Server encapsulates the gRPC server and its startup logic.
type Server struct {
	GRPCServer *grpc.Server
	Port       string
	Logger     *zap.Logger
}

// NewServer creates a new Server instance.
func NewServer(grpcServer *grpc.Server, port string, log *zap.Logger) *Server {
	return &Server{
		GRPCServer: grpcServer,
		Port:       port,
		Logger:     log,
	}
}

// Start begins serving gRPC requests.
func StartServer(deps *Dependencies) error {
	// Register the gRPC services
	pb.RegisterAllocationServiceServer(deps.GRPCServer, handler.NewAllocationHandler(deps.Services.AllocationService, deps.Logger))
	pb.RegisterAssetServiceServer(deps.GRPCServer, handler.NewAssetHandler(deps.Services.AssetService, deps.Logger))
	pb.RegisterMaintenanceServiceServer(deps.GRPCServer, handler.NewMaintenanceHandler(deps.Services.MaintenanceService, deps.Logger))
	pb.RegisterAuditServiceServer(deps.GRPCServer, handler.NewAuditHandler(deps.Services.AuditService, deps.Logger))
	pb.RegisterDisposalServiceServer(deps.GRPCServer, handler.NewDisposalHandler(deps.Services.DisposalService, deps.Logger))
	pb.RegisterLicenseServiceServer(deps.GRPCServer, handler.NewLicenseHandler(deps.Services.LicenseService, deps.Logger))

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(deps.Config.Server.GRPCPort))
	if err != nil {
		deps.Logger.Error("Failed to listen", zap.Int("port", deps.Config.Server.GRPCPort), zap.Error(err))
		return err
	}

	deps.Logger.Info("gRPC server starting", zap.Int("port", deps.Config.Server.GRPCPort))
	return deps.GRPCServer.Serve(lis)
}
