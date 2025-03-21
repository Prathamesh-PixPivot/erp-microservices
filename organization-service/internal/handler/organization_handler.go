package handler

import (
	"context"
	"log"
	"organization-service/grpc/organizationpb"
	"organization-service/internal/models"
	"organization-service/internal/repository"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrganizationServiceServer struct {
	organizationpb.UnimplementedOrganizationServiceServer
}

// NewOrganizationServiceServer initializes and returns an instance of OrganizationServiceServer
func NewOrganizationServiceServer() *OrganizationServiceServer {
	return &OrganizationServiceServer{}
}

func ProduceEvent(topic, message string) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:9092"},
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

func (s *OrganizationServiceServer) CreateOrganization(ctx context.Context, req *organizationpb.CreateOrganizationRequest) (*organizationpb.CreateOrganizationResponse, error) {
	print("\n requested%v", req)
	org := &models.Organization{
		GstIn:      req.GstIn,
		Name:       req.Name,
		Phone:      req.Phone,
		Email:      req.Email,
		Address:    req.Address,
		TemplateID: req.TemplateId,
		Website:    req.Website,
		City:       req.City,
		Country:    req.Country,
		State:      req.State,
		ModuleID:   req.ModuleId,
		Zipcode:    req.Zipcode,
		Industry:   req.Industry,
	}

	// Create organization in the database
	if err := repository.CreateOrganization(org); err != nil {
		log.Printf("Failed to create organization in the database: %v", err)
		return nil, status.Errorf(codes.Internal, "Could not create organization: %v", err)
	}

	// Produce event after successful organization creation
	// eventMessage := fmt.Sprintf(`{"organizationId": "%d", "orgName": "%s"}`, org.ID, org.Name)
	// ProduceEvent("organization-created", eventMessage)

	return &organizationpb.CreateOrganizationResponse{Id: uint32(org.ID), Name: org.Name}, nil
}
