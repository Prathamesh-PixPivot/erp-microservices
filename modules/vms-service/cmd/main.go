package main

import (
	"log"
	"net"
	"vms-service/config"
	"vms-service/grpc/handler"
	"vms-service/internal/repository"
	"vms-service/internal/services"

	vms_pb "vms-service/grpc/vms_pb" // Generated package

	"google.golang.org/grpc"
)

func main() {
    config.InitConfig()
    db := config.InitDB()

    // Initialize repositories and services
    vendorRepo := repository.NewVendorRepository(db)
    poRepo := repository.NewPurchaseOrderRepository(db)
    performanceRepo := repository.NewPerformanceRepository(db)
    paymentRepo := repository.NewPaymentRepository(db)

    vendorService := services.NewVendorService(vendorRepo)
    poService := services.NewPurchaseOrderService(poRepo)
    performanceService := services.NewPerformanceService(performanceRepo)
    paymentService := services.NewPaymentService(paymentRepo)

    // Set up gRPC server
    grpcServer := grpc.NewServer()

    // Register the gRPC services
    vms_pb.RegisterVendorServiceServer(grpcServer, handler.NewVendorHandler(vendorService))
    vms_pb.RegisterPurchaseOrderServiceServer(grpcServer, handler.NewPurchaseOrderHandler(poService))
    vms_pb.RegisterPerformanceServiceServer(grpcServer, handler.NewPerformanceHandler(performanceService))
    vms_pb.RegisterPaymentServiceServer(grpcServer, handler.NewPaymentHandler(paymentService))

    // Start listening on the defined port
    lis, err := net.Listen("tcp", config.AppConfig.GRPCPort)
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    log.Printf("gRPC server is running on port %s", config.AppConfig.GRPCPort)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
