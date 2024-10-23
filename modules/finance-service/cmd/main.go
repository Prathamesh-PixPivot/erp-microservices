package main

import (
	"finance-service/config"
	"finance-service/grpc/finance_pb"
	"finance-service/grpc/handler"
	"finance-service/internal/models"
	"finance-service/internal/repository"
	"finance-service/internal/services"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Step 1: Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Step 2: Initialize Database
	db, err := gorm.Open(postgres.Open(cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Step 3: Auto-migrate models
	err = models.AutoMigrate(db)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Step 4: Initialize Repositories
	invoiceRepo := repository.NewInvoiceRepository(db)
	paymentDueRepo := repository.NewPaymentDueRepository(db)
	ledgerRepo := repository.NewLedgerRepository(db)
	creditDebitNoteRepo := repository.NewCreditDebitNoteRepository(db)

	// Step 5: Initialize Services
	invoiceService := services.NewInvoiceService(invoiceRepo)
	paymentDueService := services.NewPaymentDueService(paymentDueRepo)
	ledgerService := services.NewLedgerService(ledgerRepo)
	creditDebitNoteService := services.NewCreditDebitNoteService(creditDebitNoteRepo)

	// Step 6: Initialize gRPC server
	grpcServer := grpc.NewServer()

	// Step 7: Register gRPC Services
	finance_pb.RegisterInvoiceServiceServer(grpcServer, handler.NewInvoiceHandler(invoiceService))
	finance_pb.RegisterPaymentServiceServer(grpcServer, handler.NewPaymentDueHandler(paymentDueService))
	finance_pb.RegisterLedgerServiceServer(grpcServer, handler.NewLedgerHandler(ledgerService))
	finance_pb.RegisterCreditDebitNoteServiceServer(grpcServer, handler.NewCreditDebitNoteHandler(creditDebitNoteService))

	// Step 8: Start Listening on the Specified Port
	listener, err := net.Listen("tcp", cfg.Server.Port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", cfg.Server.Port, err)
	}

	// Graceful shutdown handling
	go func() {
		log.Printf("Finance gRPC service running on port %s...", cfg.Server.Port)
		if err := grpcServer.Serve(listener); err != nil && err != grpc.ErrServerStopped {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Step 9: Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Step 10: Graceful shutdown
	log.Println("Shutting down gracefully...")
	grpcServer.GracefulStop()
	log.Println("Server stopped")
}
