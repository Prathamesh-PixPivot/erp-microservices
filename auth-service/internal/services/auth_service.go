package services

import (
	"auth-service/grpc/authpb"
	"auth-service/grpc/organizationpb"
	"auth-service/grpc/userpb"
	"context"
	"fmt"
	"log"
	"net/smtp"
	"strconv"

	gocloak "github.com/Nerzal/gocloak/v12"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceServer struct {
	authpb.UnimplementedAuthServiceServer
	userClient         userpb.UserServiceClient
	organizationClient organizationpb.OrganizationServiceClient
	keycloakClient     gocloak.GoCloak
	adminToken         string
}

// NewAuthServiceServer initializes and returns an instance of AuthServiceServer
func NewAuthServiceServer(userClient userpb.UserServiceClient, organizationClient organizationpb.OrganizationServiceClient) *AuthServiceServer {
	keycloakClient := gocloak.NewClient(viper.GetString("KEYCLOAK_URL"))
	adminToken, err := getAdminToken(keycloakClient)
	if err != nil {
		log.Fatalf("Failed to get admin token from Keycloak: %v", err)
	}

	return &AuthServiceServer{
		userClient:         userClient,
		organizationClient: organizationClient,
		keycloakClient:     *keycloakClient,
		adminToken:         adminToken,
	}
}

// ProduceEvent is a helper function to send events to Kafka
func ProduceEvent(topic, message string) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:9092"}, // Adjust the Kafka broker address
		Topic:   topic,
	})

	defer writer.Close()

	err := writer.WriteMessages(context.Background(),
		kafka.Message{Value: []byte(message)},
	)
	if err != nil {
		log.Fatalf("Failed to produce event: %v", err)
	}
}

// Signup handles the user signup process via gRPC
func (s *AuthServiceServer) Signup(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupResponse, error) {
	// Step 1: Create user in Keycloak
	user := gocloak.User{
		Username:      gocloak.StringP(req.Email),
		Email:         gocloak.StringP(req.Email),
		FirstName:     gocloak.StringP(req.FirstName), // Added FirstName
		LastName:      gocloak.StringP(req.LastName),  // Added LastName
		Enabled:       gocloak.BoolP(true),
		EmailVerified: gocloak.BoolP(true), // Email not verified by default
		Credentials: &[]gocloak.CredentialRepresentation{
			{
				Type:      gocloak.StringP("password"),
				Value:     gocloak.StringP(req.Password),
				Temporary: gocloak.BoolP(false),
			},
		},
	}

	// Logging the user object for debugging
	log.Printf("[DEBUG] Creating user in Keycloak: %+v", user)

	_, err := s.keycloakClient.CreateUser(ctx, s.adminToken, viper.GetString("KEYCLOAK_REALM"), user)
	if err != nil {
		fmt.Println("admin token:", s.adminToken+"\n"+"Realm:", viper.GetString("KEYCLOAK_REALM")+"\n"+"User:", user)
		log.Printf("[DEBUG] Failed to create user in Keycloak: %v", err)
		return nil, fmt.Errorf("failed to create user in Keycloak")
	}

	// Step 2: Hash password before sending to user-service
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		log.Printf("[DEBUG] Failed to hash password: %v", err)
		return nil, fmt.Errorf("failed to hash password")
	}

	// Step 3: Create user in user-service via gRPC
	userReq := &userpb.CreateUserRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  hashedPassword,
		Role:      req.Role,
	}
	log.Printf("[DEBUG] Creating user in user-service with request: %+v", userReq)
	userRes, err := s.userClient.CreateUser(ctx, userReq)
	if err != nil {
		log.Printf("[DEBUG] Failed to create user in user-service: %v", err)
		return nil, fmt.Errorf("failed to create user in user-service")
	}

	// Step 4: Create organization in organization-service via gRPC
	orgReq := &organizationpb.CreateOrganizationRequest{
		Name:    req.OrganizationName,
		GstIn:   req.GstIn,
		Phone:   req.Phone,
		Email:   req.Email,
		Address: req.Address,
		City:    req.City,
		State:   req.State,
		Country: req.Country,
		Zipcode: req.Zipcode,
		Website: req.Website,
	}
	log.Printf("[DEBUG] Creating organization in organization-service with request: %+v", orgReq)
	orgRes, err := s.organizationClient.CreateOrganization(ctx, orgReq)
	if err != nil {
		log.Printf("[DEBUG] Failed to create organization in organization-service: %v", err)
		return nil, fmt.Errorf("failed to create organization in organization-service")
	}

	// Produce event after successful signup
	// eventMessage := fmt.Sprintf(`{"userId": "%d", "email": "%s"}`, userRes.Id, userReq.Email)
	// ProduceEvent("user-signup", eventMessage)

	// send custom verification email
	// Send custom verification email
	smtpHost := viper.GetString("SMTP_HOST")
	smtpPort := viper.GetString("SMTP_PORT")
	smtpUser := viper.GetString("SMTP_USER")
	smtpPass := viper.GetString("SMTP_PASSWORD")

	from := smtpUser
	to := []string{req.Email}
	subject := "Subject: Email Verification\r\n"
	body := fmt.Sprintf("Hello %s %s,\n\nPlease verify your email by clicking the link below:\n\nhttp://localhost:8080/verify?userID=%s\n\nThank you!", req.FirstName, req.LastName, strconv.Itoa(int(userRes.Id)))
	message := []byte(subject + "\r\n" + body)
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Printf("[DEBUG] Failed to send verification email: %v", err)
		return nil, err
	}

	log.Printf("[DEBUG] Verification email sent to user with ID: %s", strconv.Itoa(int(userRes.Id)))

	// Return success response
	return &authpb.SignupResponse{
		Message:        "Signup successful",
		UserId:         strconv.Itoa(int(userRes.Id)),
		OrganizationId: strconv.Itoa(int(orgRes.Id)),
	}, nil
}

// Signin handles user login and token generation via Keycloak via gRPC
func (s *AuthServiceServer) Signin(ctx context.Context, req *authpb.SigninRequest) (*authpb.SigninResponse, error) {
	// Authenticate user via Keycloak
	token, err := s.keycloakClient.Login(ctx, viper.GetString("KEYCLOAK_CLIENT_ID"), viper.GetString("KEYCLOAK_CLIENT_SECRET"), viper.GetString("KEYCLOAK_REALM"), req.Email, req.Password, "openid")
	if err != nil {
		log.Printf("Failed to authenticate user: %v", err)
		return nil, fmt.Errorf("invalid credentials")
	}

	// Produce event after successful signin
	// eventMessage := fmt.Sprintf(`{"email": "%s"}`, req.Email)
	// ProduceEvent("user-signin", eventMessage)

	// Return tokens (access token and refresh token)
	return &authpb.SigninResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

// hashPassword hashes the given password using bcrypt.
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// getAdminToken gets the Keycloak admin token
func getAdminToken(client *gocloak.GoCloak) (string, error) {
	token, err := client.LoginAdmin(context.Background(), viper.GetString("KEYCLOAK_ADMIN_USER"), viper.GetString("KEYCLOAK_ADMIN_PASSWORD"), viper.GetString("KEYCLOAK_ADMIN_REALM"))
	if err != nil {
		fmt.Println("admin user:", viper.GetString("KEYCLOAK_ADMIN_USER")+"\n"+"admin password:", viper.GetString("KEYCLOAK_ADMIN_PASSWORD")+"\n"+"admin realm:", viper.GetString("KEYCLOAK_ADMIN_REALM"))
		return "", fmt.Errorf("failed to get admin token from Keycloak: %v", err)
	}
	return token.AccessToken, nil
}
