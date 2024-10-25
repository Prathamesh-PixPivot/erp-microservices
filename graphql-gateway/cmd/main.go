package main

import (
	"context"
	"fmt"
	"graphql-gateway/config"
	"graphql-gateway/gqlgen/generated"
	"graphql-gateway/gqlgen/resolvers"
	"graphql-gateway/grpc/activitypb"
	"graphql-gateway/grpc/authpb"
	"graphql-gateway/grpc/contactpb"
	"graphql-gateway/grpc/finance_pb"
	"graphql-gateway/grpc/leadspb"
	"graphql-gateway/grpc/opportunitypb"
	"graphql-gateway/grpc/organizationpb"
	"graphql-gateway/grpc/userpb"
	"graphql-gateway/grpc/vms_pb" // Import the VMS gRPC package
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Step 1: Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Step 2: Initialize gRPC Clients for all required services

	// Initialize AuthService Client
	authConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.AuthServiceHost, cfg.AuthServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to auth-service: %v", err)
	}
	defer authConn.Close()
	authClient := authpb.NewAuthServiceClient(authConn)
	log.Println("Connected to auth-service.")

	// Initialize UserService Client
	userConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to user-service: %v", err)
	}
	defer userConn.Close()
	userClient := userpb.NewUserServiceClient(userConn)
	log.Println("Connected to user-service.")

	// Initialize OrganizationService Client
	orgConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.OrganizationServiceHost, cfg.OrganizationServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to organization-service: %v", err)
	}
	defer orgConn.Close()
	orgClient := organizationpb.NewOrganizationServiceClient(orgConn)
	log.Println("Connected to organization-service.")

	// Initialize LeadService Client
	leadsConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.LeadServiceHost, cfg.LeadServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to leads-service: %v", err)
	}
	defer leadsConn.Close()
	leadClient := leadspb.NewLeadServiceClient(leadsConn)
	log.Println("Connected to leads-service.")

	// Initialize OpportunityService Client
	opportunityConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.OpportunityServiceHost, cfg.OpportunityServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to opportunity-service: %v", err)
	}
	defer opportunityConn.Close()
	opportunityClient := opportunitypb.NewOpportunityServiceClient(opportunityConn)
	log.Println("Connected to opportunity-service.")

	// Initialize ContactService Client
	contactConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.ContactServiceHost, cfg.ContactServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to contact-service: %v", err)
	}
	defer contactConn.Close()
	contactClient := contactpb.NewContactServiceClient(contactConn)
	log.Println("Connected to contact-service.")

	// Initialize ActivityService Client
	activityConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.ActivityServiceHost, cfg.ActivityServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to activity-service: %v", err)
	}
	defer activityConn.Close()
	activityClient := activitypb.NewActivityServiceClient(activityConn)
	log.Println("Connected to activity-service.")

	// Initialize vendor Client (New Code)
	vendorConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.VendorServiceHost, cfg.VendorServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to vendor service: %v", err)
	}
	defer vendorConn.Close()
	vendorClient := vms_pb.NewVendorServiceClient(vendorConn) // This creates a new VMS client
	log.Println("Connected to VMS service.")

	// Initialize payment Client (New Code)
	paymentConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PaymentServiceHost, cfg.PaymentServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to payment service: %v", err)
	}
	defer paymentConn.Close()
	paymentClient := vms_pb.NewPaymentServiceClient(paymentConn) // This creates a new VMS client

	// Initialize performance Client (New Code)
	performanceConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PerformanceServiceHost, cfg.PerformanceServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to performance service: %v", err)
	}
	defer performanceConn.Close()
	performanceClient := vms_pb.NewPerformanceServiceClient(performanceConn) // This creates a new VMS client

	// Initialize purchase order Client (New Code)
	purchaseOrderConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PurchaseOrderServiceHost, cfg.PurchaseOrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to purchase order service: %v", err)
	}
	defer purchaseOrderConn.Close()
	purchaseOrderClient := vms_pb.NewPurchaseOrderServiceClient(purchaseOrderConn) // This creates a new VMS client

	// Initialize Invoice Client (New Code)
	invoiceConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.InvoiceServiceHost, cfg.InvoiceServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to invoice service: %v", err)
	}
	defer invoiceConn.Close()
	invoiceClient := finance_pb.NewInvoiceServiceClient(invoiceConn) // This creates a new VMS client

	// Initialize CreditNote Client (New Code)
	creditNoteConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CreditDebitNoteServiceHost, cfg.CreditDebitNoteServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to credit note service: %v", err)
	}
	defer creditNoteConn.Close()
	creditDebitNoteClient := finance_pb.NewCreditDebitNoteServiceClient(creditNoteConn) // This creates a new VMS client

	// Initialize PaymentDue Client (New Code)
	paymentDueConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PaymentDueServiceHost, cfg.PaymentDueServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to payment due service: %v", err)
	}
	defer paymentDueConn.Close()
	paymentDueClient := finance_pb.NewPaymentServiceClient(paymentDueConn) // This creates a new VMS client

	// Initialize Ledger Client (New Code)
	ledgerConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.LedgerServiceHost, cfg.LedgerServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to ledger service: %v", err)
	}
	defer ledgerConn.Close()
	ledgerClient := finance_pb.NewLedgerServiceClient(ledgerConn) // This creates a new VMS client

	// Step 3: Initialize Resolver with all gRPC Clients
	resolver := &resolvers.Resolver{
		AuthClient:            authClient,
		UserClient:            userClient,
		OrganizationClient:    orgClient,
		LeadClient:            leadClient,
		OpportunityClient:     opportunityClient,
		ContactClient:         contactClient,
		ActivityClient:        activityClient,
		VendorClient:          vendorClient,
		PaymentClient:         paymentClient,
		PerformanceClient:     performanceClient,
		PurchaseOrderClient:   purchaseOrderClient,
		InvoiceClient:         invoiceClient,
		CreditDebitNoteClient: creditDebitNoteClient,
		PaymentDueClient:      paymentDueClient,
		LedgerClient:          ledgerClient,
	}

	// Step 4: Setup GraphQL Server
	gqlSrv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))

	// Step 5: Setup HTTP Router
	router := mux.NewRouter()

	// Add GraphQL Playground handler
	router.Handle("/", playground.Handler("GraphQL Playground", "/graphql"))

	// Add GraphQL endpoint handler
	router.Handle("/graphql", gqlSrv)

	// Step 6: Create the HTTP Server
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.GraphQLPort),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Step 7: Start Listening on Specified Port
	go func() {
		log.Printf("GraphQL Gateway running on :%d", cfg.GraphQLPort)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to listen and serve: %v", err)
		}
	}()

	// Step 8: Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline to wait for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
