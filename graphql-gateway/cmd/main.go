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
	"graphql-gateway/grpc/vms_pb"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

// ANSI color codes for green text and reset
const (
	greenColor = "\033[32m"
	resetColor = "\033[0m"
)

// CORS middleware to allow all origins, headers, and methods
var cors = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers to allow all origins, headers, and methods
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// If the request is an OPTIONS preflight request, return immediately
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Continue processing the request
		next.ServeHTTP(w, r)
	})
}

// attemptGrpcReconnect tries to establish a gRPC connection with retries and waits for READY state
func attemptGrpcReconnect(ctx context.Context, serviceHost string, servicePort int, serviceName string, clientSetter func(conn *grpc.ClientConn, client interface{}), wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			log.Printf("Shutdown signal received. Stopping connection attempts to %s.", serviceName)
			return
		default:
			conn, err := grpc.Dial(
				fmt.Sprintf("%s:%d", serviceHost, servicePort),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
				grpc.WithBlock(), // Waits for connection readiness
				grpc.WithTimeout(5*time.Second),
			)
			if err != nil {
				// log.Printf("Warning: Failed to connect to %s: %v. Retrying in 10 seconds...", serviceName, err)
				time.Sleep(10 * time.Second) // Retry delay
				continue
			}

			// Monitor connection state changes until it reaches READY
			for {
				state := conn.GetState()
				if state == connectivity.Ready {
					log.Printf("%sConnected to %s.%s", greenColor, serviceName, resetColor)
					break
				} else {
					log.Printf("Waiting for %s to be ready. Current state: %s", serviceName, state.String())
					conn.WaitForStateChange(ctx, state)
				}
			}

			// Set up the appropriate client based on serviceName
			var client interface{}
			switch serviceName {
			case "auth-service":
				client = authpb.NewAuthServiceClient(conn)
			case "user-service":
				client = userpb.NewUserServiceClient(conn)
			case "organization-service":
				client = organizationpb.NewOrganizationServiceClient(conn)
			case "leads-service":
				client = leadspb.NewLeadServiceClient(conn)
			case "opportunity-service":
				client = opportunitypb.NewOpportunityServiceClient(conn)
			case "contact-service":
				client = contactpb.NewContactServiceClient(conn)
			case "activity-service":
				client = activitypb.NewActivityServiceClient(conn)
			case "vendor-service":
				client = vms_pb.NewVendorServiceClient(conn)
			case "payment-service":
				client = vms_pb.NewPaymentServiceClient(conn)
			case "performance-service":
				client = vms_pb.NewPerformanceServiceClient(conn)
			case "purchase-order-service":
				client = vms_pb.NewPurchaseOrderServiceClient(conn)
			case "invoice-service":
				client = finance_pb.NewInvoiceServiceClient(conn)
			case "credit-debit-note-service":
				client = finance_pb.NewCreditDebitNoteServiceClient(conn)
			case "ledger-service":
				client = finance_pb.NewLedgerServiceClient(conn)
			default:
				log.Printf("No client found for service: %s", serviceName)
				conn.Close()
				return
			}

			// Assign the client and return to break the retry loop
			clientSetter(conn, client)
			return
		}
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: %s %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the resolver with client fields set to nil initially
	resolver := &resolvers.Resolver{}
	var wg sync.WaitGroup

	// Create a context to manage graceful shutdown of goroutines
	ctx, cancel := context.WithCancel(context.Background())

	// Start goroutines for each gRPC client with automatic reconnection logic
	wg.Add(1)
	go attemptGrpcReconnect(ctx, cfg.AuthServiceHost, cfg.AuthServicePort, "auth-service", func(conn *grpc.ClientConn, client interface{}) {
		resolver.AuthClient = client.(authpb.AuthServiceClient)
	}, &wg)

	wg.Add(1)
	go attemptGrpcReconnect(ctx, cfg.UserServiceHost, cfg.UserServicePort, "user-service", func(conn *grpc.ClientConn, client interface{}) {
		resolver.UserClient = client.(userpb.UserServiceClient)
	}, &wg)

	wg.Add(1)
	go attemptGrpcReconnect(ctx, cfg.OrganizationServiceHost, cfg.OrganizationServicePort, "organization-service", func(conn *grpc.ClientConn, client interface{}) {
		resolver.OrganizationClient = client.(organizationpb.OrganizationServiceClient)
	}, &wg)

	wg.Add(1)
	go attemptGrpcReconnect(ctx, cfg.LeadServiceHost, cfg.LeadServicePort, "leads-service", func(conn *grpc.ClientConn, client interface{}) {
		resolver.LeadClient = client.(leadspb.LeadServiceClient)
	}, &wg)

	wg.Add(1)
	go attemptGrpcReconnect(ctx, cfg.OpportunityServiceHost, cfg.OpportunityServicePort, "opportunity-service", func(conn *grpc.ClientConn, client interface{}) {
		resolver.OpportunityClient = client.(opportunitypb.OpportunityServiceClient)
	}, &wg)

	wg.Add(1)
	go attemptGrpcReconnect(ctx, cfg.ContactServiceHost, cfg.ContactServicePort, "contact-service", func(conn *grpc.ClientConn, client interface{}) {
		resolver.ContactClient = client.(contactpb.ContactServiceClient)
	}, &wg)

	wg.Add(1)
	go attemptGrpcReconnect(ctx, cfg.ActivityServiceHost, cfg.ActivityServicePort, "activity-service", func(conn *grpc.ClientConn, client interface{}) {
		resolver.ActivityClient = client.(activitypb.ActivityServiceClient)
	}, &wg)

	wg.Add(1)
	go attemptGrpcReconnect(ctx, cfg.VendorServiceHost, cfg.VendorServicePort, "vendor-service", func(conn *grpc.ClientConn, client interface{}) {
		resolver.VendorClient = client.(vms_pb.VendorServiceClient)
	}, &wg)

	wg.Add(1)
	go attemptGrpcReconnect(ctx, cfg.PaymentServiceHost, cfg.PaymentServicePort, "payment-service", func(conn *grpc.ClientConn, client interface{}) {
		resolver.PaymentClient = client.(vms_pb.PaymentServiceClient)
	}, &wg)

	wg.Add(1)
	go attemptGrpcReconnect(ctx, cfg.PerformanceServiceHost, cfg.PerformanceServicePort, "performance-service", func(conn *grpc.ClientConn, client interface{}) {
		resolver.PerformanceClient = client.(vms_pb.PerformanceServiceClient)
	}, &wg)

	wg.Add(1)
	go attemptGrpcReconnect(ctx, cfg.PurchaseOrderServiceHost, cfg.PurchaseOrderServicePort, "purchase-order-service", func(conn *grpc.ClientConn, client interface{}) {
		resolver.PurchaseOrderClient = client.(vms_pb.PurchaseOrderServiceClient)
	}, &wg)

	wg.Add(1)
	go attemptGrpcReconnect(ctx, cfg.InvoiceServiceHost, cfg.InvoiceServicePort, "invoice-service", func(conn *grpc.ClientConn, client interface{}) {
		resolver.InvoiceClient = client.(finance_pb.InvoiceServiceClient)
	}, &wg)

	wg.Add(1)
	go attemptGrpcReconnect(ctx, cfg.CreditDebitNoteServiceHost, cfg.CreditDebitNoteServicePort, "credit-debit-note-service", func(conn *grpc.ClientConn, client interface{}) {
		resolver.CreditDebitNoteClient = client.(finance_pb.CreditDebitNoteServiceClient)
	}, &wg)

	wg.Add(1)
	go attemptGrpcReconnect(ctx, cfg.LedgerServiceHost, cfg.LedgerServicePort, "ledger-service", func(conn *grpc.ClientConn, client interface{}) {
		resolver.LedgerClient = client.(finance_pb.LedgerServiceClient)
	}, &wg)

	// Set up GraphQL Server with the resolver that will dynamically receive clients
	gqlSrv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))

	router := mux.NewRouter()
	log.Println("Applying CORS middleware...") // Confirm CORS middleware application
	router.Use(loggingMiddleware)              // Log all incoming requests
	router.Use(cors)

	// Set up GraphQL and Playground routes with logging
	log.Println("Setting up routes...")
	// Set up GraphQL and Playground routes
	router.Handle("/", playground.Handler("GraphQL Playground", "/graphql"))
	router.Handle("/graphql", gqlSrv)
	router.Handle("/query", gqlSrv) // If "/query" is another endpoint being used
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Test endpoint accessed")
		w.Write([]byte("CORS Middleware Test"))
	}).Methods("GET")
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.GraphQLPort),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start HTTP server in a goroutine
	go func() {
		log.Printf("GraphQL Gateway running on :%d", cfg.GraphQLPort)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to listen and serve: %v", err)
		}
	}()

	// Wait for any shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Trigger context cancellation to stop goroutines
	cancel()

	// Gracefully shut down the HTTP server with a timeout
	ctxShutDown, cancelShutDown := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelShutDown()

	if err := httpServer.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// Wait for all reconnection goroutines to finish
	wg.Wait()
	log.Println("Server exited gracefully")
}
