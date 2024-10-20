package main

import (
    "log"
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"

    "google.golang.org/grpc"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "leads-service/config"
    "leads-service/grpc/leadspb"
    "leads-service/internal/handler"
    "leads-service/internal/models"
    "leads-service/internal/repository"
    "leads-service/internal/services"
    "leads-service/internal/websockets"
)

func main() {
    // Load config
    cfg := config.Load()

    // Initialize the GORM database
    db, err := gorm.Open(postgres.Open(cfg.DB), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    // AutoMigrate the database based on the Lead model
    err = db.AutoMigrate(&models.Lead{})
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }

    // Initialize WebSocket server (Single Instance)
    wsServer := websockets.NewServer()

    // Initialize services
    leadService := services.NewLeadService(repository.NewLeadRepository(db))

    // Create a new LeadHandler with the shared wsServer instance
    leadHandler := handler.NewLeadHandler(leadService, wsServer)

    // Create a new gRPC server
    grpcServer := grpc.NewServer()

    // Register the LeadHandler with the gRPC server
    leadspb.RegisterLeadServiceServer(grpcServer, leadHandler)

    // Start the gRPC server in a separate goroutine
    go func() {
        lis, err := net.Listen("tcp", cfg.GRPCPort)
        if err != nil {
            log.Fatalf("failed to listen: %v", err)
        }
        log.Printf("gRPC server listening on %s", cfg.GRPCPort)
        if err := grpcServer.Serve(lis); err != nil {
            log.Fatalf("failed to serve: %v", err)
        }
    }()

    // Handle WebSocket connections in a separate goroutine
    go func() {
        http.HandleFunc("/ws", wsServer.HandleConnections)
        log.Printf("WebSocket server listening on %s", cfg.WSAddr)
        if err := http.ListenAndServe(cfg.WSAddr, nil); err != nil {
            log.Fatalf("WebSocket server error: %v", err)
        }
    }()

    // Start broadcasting messages to WebSocket clients
    go wsServer.Start()

    // Graceful shutdown
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
    <-stop
    log.Println("Shutting down gracefully...")
    grpcServer.GracefulStop()
    wsServer.Shutdown()
    log.Println("Service stopped.")
}
