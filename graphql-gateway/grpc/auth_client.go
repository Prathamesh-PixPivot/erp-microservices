package grpc

import (
	"context"
	authpb "graphql-gateway/grpc/authpb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var AuthServiceClient authpb.AuthServiceClient

func InitAuthClient() {
	conn, err := grpc.DialContext(context.Background(), "auth-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to auth-service: %v", err)
	}
	AuthServiceClient = authpb.NewAuthServiceClient(conn)
}
