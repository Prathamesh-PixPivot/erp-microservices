package main

import (
	"auth-service/config"
	"auth-service/grpc/authpb"
	"auth-service/grpc/organizationpb"
	"auth-service/grpc/userpb"
	"auth-service/internal/handler"
	"auth-service/internal/repository"
	"auth-service/internal/services"
	"context"
	"log"
	"net"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/Nerzal/gocloak/v12"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Load configuration
	config.InitConfig()

	// Connect to gRPC services (User and Organization)
	userConn, err := grpc.Dial("host.docker.internal:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer userConn.Close()

	organizationConn, err := grpc.Dial("host.docker.internal:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to organization service: %v", err)
	}
	defer organizationConn.Close()

	userClient := userpb.NewUserServiceClient(userConn)
	organizationClient := organizationpb.NewOrganizationServiceClient(organizationConn)

	// Initialize the GoCloak client
	goCloakClient := gocloak.NewClient(viper.GetString("KEYCLOAK_URL"))
	realm := viper.GetString("KEYCLOAK_REALM")
	token, err := goCloakClient.LoginAdmin(context.Background(), viper.GetString("KEYCLOAK_ADMIN_USER"), viper.GetString("KEYCLOAK_ADMIN_PASSWORD"), viper.GetString("KEYCLOAK_ADMIN_REALM"))
	if err != nil {
		log.Printf("Admin Username: %s\nAdmin Password: %s\nAdmin Realm: %s\n", viper.GetString("KEYCLOAK_ADMIN_USERNAME"), viper.GetString("KEYCLOAK_ADMIN_PASSWORD"), viper.GetString("KEYCLOAK_ADMIN_REALM"))
		log.Fatalf("Failed to login to Keycloak: %v", err)
	}

	// Start Prometheus metrics server in a separate goroutine
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":9090", nil); err != nil {
			log.Fatalf("Failed to start Prometheus metrics server: %v", err)
		}
	}()

	// Start listening on port 50053 for the gRPC server
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen on port 50053: %v", err)
	}

	// Create new gRPC server
	grpcServer := grpc.NewServer()

	// Initialize the AuthServiceServer
	authServiceServer := services.NewAuthServiceServer(userClient, organizationClient)

	// Register the AuthServiceServer with the gRPC server
	authpb.RegisterAuthServiceServer(grpcServer, authServiceServer)

	// Start the gRPC server in a separate goroutine
	go func() {
		log.Println("Auth gRPC server is running on port 50053")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// Setup Fiber and routes for HTTP server
	app := fiber.New()
	// Use CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://host.docker.internal:3000",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
	if _, err := repository.NewKeycloakRepository(); err != nil {
		log.Fatalf("Failed to initialize Keycloak repository: %v", err)
	}

	authHandler := handler.NewAuthHandler(userClient, organizationClient, goCloakClient, token.AccessToken, realm)

	// Define REST routes
	app.Post("/signup", authHandler.SignupUser)
	app.Post("/signin", authHandler.SigninUser)
	app.Post("/refresh", authHandler.RefreshToken)

	// Define the port for the HTTP server
	port := ":3001"
	log.Printf("Auth service running on port %s", port)

	// Start HTTP server
	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
