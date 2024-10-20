package grpc

import (
	"context"
	"fmt"
	"log"
	pb "user-service/grpc/userpb" // Import the generated gRPC package
	"user-service/internal/repository"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	repo *repository.UserRepository
}

// NewUserServiceServer initializes and returns an instance of UserServiceServer
func NewUserServiceServer(repo *repository.UserRepository) *UserServiceServer {
	if repo == nil {
		log.Fatal("Repository cannot be nil")
	}
	return &UserServiceServer{repo: repo}
}

// ProduceEvent is a helper function to send events to Kafka
func ProduceEvent(topic, message string) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"}, // Ensure this matches your Kafka broker address
		Topic:   topic,
	})

	defer func() {
		if err := writer.Close(); err != nil {
			log.Printf("Failed to close Kafka writer: %v", err)
		}
	}()

	err := writer.WriteMessages(context.Background(),
		kafka.Message{Value: []byte(message)},
	)
	if err != nil {
		log.Printf("Failed to produce event to Kafka: %v", err)
	}
}

// GetUsersByOrganization handles retrieving users by organization via gRPC
func (s *UserServiceServer) GetUsersByOrganization(ctx context.Context, req *pb.GetUsersByOrganizationRequest) (*pb.GetUsersByOrganizationResponse, error) {
	if s.repo == nil {
		return nil, status.Error(codes.Internal, "User repository not initialized")
	}

	// Fetch users from the repository by organization ID
	users, err := s.repo.GetUsersByOrganizationID(uint(req.OrganizationId))
	if err != nil {
		log.Printf("Error fetching users by organization ID %d: %v", req.OrganizationId, err)
		return nil, status.Errorf(codes.Internal, "Failed to retrieve users for organization ID %d: %v", req.OrganizationId, err)
	}

	// Prepare the gRPC response with a list of users
	var grpcUsers []*pb.User
	for _, user := range users {
		grpcUsers = append(grpcUsers, &pb.User{
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

	// Optionally produce an event after users are retrieved (if needed)
	eventMessage := fmt.Sprintf(`{"organizationId": "%d", "numUsers": "%d"}`, req.OrganizationId, len(users))
	ProduceEvent("users-retrieved", eventMessage)

	return &pb.GetUsersByOrganizationResponse{
		Users: grpcUsers,
	}, nil
}
