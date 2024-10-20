package handler

import (
	"context"
	"log"
	"user-service/grpc/userpb"
	"user-service/internal/models"
	"user-service/internal/repository"

	"github.com/segmentio/kafka-go"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	userpb.UnimplementedUserServiceServer
	repo *repository.UserRepository
}

// NewUserServiceServer initializes and returns an instance of UserServiceServer
func NewUserServiceServer(repo *repository.UserRepository) *UserServiceServer {
	return &UserServiceServer{repo: repo}
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

// hashPassword hashes the given password using bcrypt
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CreateUser handles creating a new user in the database via gRPC
func (s *UserServiceServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	// Hash the user's password before saving
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return nil, status.Errorf(codes.Internal, "Could not hash password: %v", err)
	}

	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Role:      req.Role,
		Password:  hashedPassword,
	}

	// Create user in the database
	if err := s.repo.CreateUser(user); err != nil {
		log.Printf("Failed to create user: %v", err)
		return nil, status.Errorf(codes.Internal, "Could not create user: %v", err)
	}

	// Produce event after successful user creation
	// eventMessage := fmt.Sprintf(`{"userId": "%d", "email": "%s"}`, user.ID, user.Email)
	// ProduceEvent("user-created", eventMessage)

	return &userpb.CreateUserResponse{Id: uint32(user.ID)}, nil
}

// GetUser retrieves a user from the database via gRPC
func (s *UserServiceServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := s.repo.GetUserByID(uint(req.Id))
	if err != nil {
		log.Printf("User not found: %v", err)
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	// Produce event after user is retrieved
	// eventMessage := fmt.Sprintf(`{"userId": "%d", "email": "%s"}`, user.ID, user.Email)
	// ProduceEvent("user-retrieved", eventMessage)

	return &userpb.GetUserResponse{
		Id:             uint32(user.ID),
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Phone:          user.Phone,
		Role:           user.Role,
		OrganizationId: uint32(user.OrganizationID),
		KeycloakId:     user.KeycloakID,
	}, nil
}

// GetUsersByOrganization retrieves all users belonging to an organization via gRPC
func (s *UserServiceServer) GetUsersByOrganization(ctx context.Context, req *userpb.GetUsersByOrganizationRequest) (*userpb.GetUsersByOrganizationResponse, error) {
	users, err := s.repo.GetUsersByOrganizationID(uint(req.OrganizationId))
	if err != nil {
		log.Printf("Error retrieving users: %v", err)
		return nil, status.Errorf(codes.Internal, "Error retrieving users")
	}

	var grpcUsers []*userpb.User
	for _, user := range users {
		grpcUsers = append(grpcUsers, &userpb.User{
			Id:             uint32(user.ID),
			FirstName:      user.FirstName,
			LastName:       user.LastName,
			Email:          user.Email,
			Phone:          user.Phone,
			Role:           user.Role,
			OrganizationId: uint32(user.OrganizationID),
			KeycloakId:     user.KeycloakID,
		})
	}

	// Produce event after users are retrieved
	// eventMessage := fmt.Sprintf(`{"organizationId": "%d", "numUsers": "%d"}`, req.OrganizationId, len(users))
	// ProduceEvent("users-retrieved", eventMessage)

	return &userpb.GetUsersByOrganizationResponse{Users: grpcUsers}, nil
}
