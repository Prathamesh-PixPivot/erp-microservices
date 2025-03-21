package grpc

import (
	"context"
	"log"
	userpb "organization-service/grpc/userpb" // Import the generated gRPC package for user-service

	"google.golang.org/grpc"
)

func GetUsersByOrganizationID(orgID uint) ([]*userpb.User, error) { // Returning []*pb.User instead of GetUserResponse
	// Connect to user-service gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user-service: %v", err)
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	// Make the gRPC request to fetch users for the organization
	req := &userpb.GetUsersByOrganizationRequest{
		OrganizationId: uint32(orgID),
	}

	res, err := client.GetUsersByOrganization(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res.Users, nil // Return the users from the response
}
