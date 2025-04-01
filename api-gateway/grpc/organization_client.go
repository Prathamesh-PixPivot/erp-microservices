package grpc

import (
	organizationpb "api-gateway/grpc/organizationpb"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var OrganizationServiceClient organizationpb.OrganizationServiceClient

func InitOrganizationClient() {
	conn, err := grpc.DialContext(context.Background(), "organization-service:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to organization-service: %v", err)
	}
	OrganizationServiceClient = organizationpb.NewOrganizationServiceClient(conn)
}
