package grpc

import (
    "context"
    pb "organization-service/grpc/organizationpb"  // Import the generated gRPC package
    "organization-service/internal/models"
    "organization-service/internal/repository"
    "log"
)

// OrganizationServiceServer implements the gRPC service
type OrganizationServiceServer struct {
    pb.UnimplementedOrganizationServiceServer
}

// Implement CreateOrganization method
func (s *OrganizationServiceServer) CreateOrganization(ctx context.Context, req *pb.CreateOrganizationRequest) (*pb.CreateOrganizationResponse, error) {
    // Create an Organization model instance from the request
    organization := models.Organization{
        GSTIn:      req.GstIn,
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

    // Call the repository to save the organization in the database
    if err := repository.CreateOrganization(&organization); err != nil {
        log.Printf("Failed to create organization: %v", err)
        return nil, err
    }

    // Return the response with the organization ID and name
    return &pb.CreateOrganizationResponse{
        Id:   uint32(organization.ID),
        Name: organization.Name,
    }, nil
}
