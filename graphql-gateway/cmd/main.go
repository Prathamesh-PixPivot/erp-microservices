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

	// Initialize VMS Client (New Code)
	vmsConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.VMSServiceHost, cfg.VMSServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to VMS service: %v", err)
	}
	defer vmsConn.Close()
	vmsClient := vms_pb.NewVendorServiceClient(vmsConn) // This creates a new VMS client
	log.Println("Connected to VMS service.")

	// Step 3: Initialize Resolver with all gRPC Clients
	resolver := &resolvers.Resolver{
		AuthClient:         authClient,
		UserClient:         userClient,
		OrganizationClient: orgClient,
		LeadClient:         leadClient,
		OpportunityClient:  opportunityClient,
		ContactClient:      contactClient,
		ActivityClient:     activityClient,
		VendorClient:       vmsClient, // Add VMS client to the resolver
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
