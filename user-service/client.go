package main

import (
	"context"
	"log"
	"time"
	pb "user-service/grpc/userpb" // Import the generated gRPC package

	"google.golang.org/grpc"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Prepare the request to get a user by ID
	req := &pb.GetUserRequest{Id: 1}

	// Make the request with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the GetUser method from the gRPC server
	res, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("Error when calling GetUser: %v", err)
	}

	// Print the response from the gRPC server
	log.Printf("Response from server: %v", res)
}
