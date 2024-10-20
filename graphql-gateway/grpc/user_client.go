package grpc

import (
	"context"
	userpb "graphql-gateway/grpc/userpb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var UserServiceClient userpb.UserServiceClient

func InitUserClient() {
	conn, err := grpc.DialContext(context.Background(), "user-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user-service: %v", err)
	}
	UserServiceClient = userpb.NewUserServiceClient(conn)
}
