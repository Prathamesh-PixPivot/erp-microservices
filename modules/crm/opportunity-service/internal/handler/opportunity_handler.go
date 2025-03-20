package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"opportunity-service/grpc/opportunitypb"
	"opportunity-service/internal/models"
	"opportunity-service/internal/services"
	"time"
	// "google.golang.org/protobuf/types/known/wrapperspb"
)

type OpportunityHandler struct {
	opportunityService services.OpportunityService
	opportunitypb.UnimplementedOpportunityServiceServer
}

func NewOpportunityHandler(service services.OpportunityService) *OpportunityHandler {
	return &OpportunityHandler{opportunityService: service}
}

func (h *OpportunityHandler) CreateOpportunity(ctx context.Context, req *opportunitypb.CreateOpportunityRequest) (*opportunitypb.CreateOpportunityResponse, error) {
	log.Printf("Received CreateOpportunity request: %+v", req)

	opportunity, err := convertProtoToModel(req.Opportunity)

	if err != nil {
		fmt.Print("failed to convert proto to model %v", err)
		return nil, err
	}

	createdOpportunity, err := h.opportunityService.CreateOpportunity(opportunity)
	if err != nil {
		log.Printf("Error creating opportunity: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &opportunitypb.CreateOpportunityResponse{
		Opportunity: convertModelToProto(createdOpportunity),
	}, nil
}

func (h *OpportunityHandler) GetOpportunity(ctx context.Context, req *opportunitypb.GetOpportunityRequest) (*opportunitypb.GetOpportunityResponse, error) {
	opportunity, err := h.opportunityService.GetOpportunity(uint(req.Id))
	if err != nil {
		if err == services.ErrOpportunityNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &opportunitypb.GetOpportunityResponse{
		Opportunity: convertModelToProto(opportunity),
	}, nil
}

// Implement UpdateOpportunity, DeleteOpportunity, ListOpportunities similarly...
func (h *OpportunityHandler) UpdateOpportunity(ctx context.Context, req *opportunitypb.UpdateOpportunityRequest) (*opportunitypb.UpdateOpportunityResponse, error) {
	log.Printf("Received UpdateOpportunity request: %+v", req)

	// Retrieve the existing opportunity from the database
	existingOpportunity, err := h.opportunityService.GetOpportunity(uint(req.Opportunity.Id))
	if err != nil {
		if err == services.ErrOpportunityNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Update fields only if they are provided (non-zero values)
	if req.Opportunity.Name != "" {
		existingOpportunity.Name = req.Opportunity.Name
	}
	if req.Opportunity.Description != "" {
		existingOpportunity.Description = req.Opportunity.Description
	}
	if req.Opportunity.Stage != "" {
		existingOpportunity.Stage = req.Opportunity.Stage
	}
	if req.Opportunity.Amount != 0 {
		existingOpportunity.Amount = req.Opportunity.Amount
	}
	if req.Opportunity.CloseDate != "" {
		parsedDate, _ := parseDate(req.Opportunity.CloseDate)
		existingOpportunity.CloseDate = parsedDate
	}
	if req.Opportunity.Probability != 0 {
		existingOpportunity.Probability = req.Opportunity.Probability
	}
	if req.Opportunity.LeadId != 0 {
		tempLeadID := uint(req.Opportunity.LeadId)
		existingOpportunity.LeadID = tempLeadID
	}
	if req.Opportunity.AccountId != 0 {
		tempAccountID := uint(req.Opportunity.AccountId)
		existingOpportunity.AccountID = tempAccountID
	}
	if req.Opportunity.OwnerId != 0 {
		tempOwnerID := uint(req.Opportunity.OwnerId)
		existingOpportunity.OwnerID = tempOwnerID
	}

	// Save the updated opportunity
	updatedOpportunity, err := h.opportunityService.UpdateOpportunity(existingOpportunity)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &opportunitypb.UpdateOpportunityResponse{
		Opportunity: convertModelToProto(updatedOpportunity),
	}, nil
}

func (h *OpportunityHandler) DeleteOpportunity(ctx context.Context, req *opportunitypb.DeleteOpportunityRequest) (*opportunitypb.DeleteOpportunityResponse, error) {
	log.Printf("Received DeleteOpportunity request: %+v", req)

	// Call the service layer to delete the opportunity
	err := h.opportunityService.DeleteOpportunity(uint(req.Id))
	if err != nil {
		if err == services.ErrOpportunityNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Return a successful response
	return &opportunitypb.DeleteOpportunityResponse{
		Success: true,
	}, nil
}

func (h *OpportunityHandler) ListOpportunities(ctx context.Context, req *opportunitypb.ListOpportunitiesRequest) (*opportunitypb.ListOpportunitiesResponse, error) {
	log.Printf("Received ListOpportunities request: %+v", req)

	ownerID := uint(req.OwnerId)

	// Call the service layer to list opportunities
	opportunities, err := h.opportunityService.ListOpportunities(ownerID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Convert to protobuf opportunities
	var protoOpps []*opportunitypb.Opportunity
	for _, opp := range opportunities {
		protoOpps = append(protoOpps, convertModelToProto(&opp))
	}

	return &opportunitypb.ListOpportunitiesResponse{
		Opportunities: protoOpps,
	}, nil
}

// Conversion functions
func convertProtoToModel(protoOpp *opportunitypb.Opportunity) (*models.Opportunity, error) {
	parsedDate, err := parseDate(protoOpp.CloseDate)
	if err != nil {
		return nil, fmt.Errorf("invalid close date format: %s", protoOpp.CloseDate)
	}

	// Convert protobuf Opportunity to models.Opportunity
	return &models.Opportunity{
		Id:          uint(protoOpp.Id), // Fixed casing to match struct
		Name:        protoOpp.Name,
		Description: protoOpp.Description,
		Stage:       protoOpp.Stage,
		Amount:      protoOpp.Amount,
		CloseDate:   parsedDate, // Using the parsed date
		Probability: protoOpp.Probability,
		LeadID:      uint(protoOpp.LeadId),
		AccountID:   uint(protoOpp.AccountId),
		OwnerID:     uint(protoOpp.OwnerId),
	}, nil
}

func convertModelToProto(modelOpp *models.Opportunity) *opportunitypb.Opportunity {
	// Convert models.Opportunity to protobuf Opportunity
	return &opportunitypb.Opportunity{
		Id:          uint32(modelOpp.Id),
		Name:        modelOpp.Name,
		Description: modelOpp.Description,
		Stage:       modelOpp.Stage,
		Amount:      modelOpp.Amount,
		CloseDate:   modelOpp.CloseDate.Format(time.RFC3339),
		Probability: modelOpp.Probability,
		LeadId:      uint32(modelOpp.LeadID),
		AccountId:   uint32(modelOpp.AccountID),
		OwnerId:     uint32(modelOpp.OwnerID),
		CreatedAt:   modelOpp.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   modelOpp.UpdatedAt.Format(time.RFC3339),
	}
}

// Helper functions

// func getFloatPointer(f *wrapperspb.DoubleValue) *float64 {
// 	if f != nil {
// 		val := f.GetValue()
// 		return &val
// 	}
// 	return nil
// }
// func getStringPointer(s *wrapperspb.StringValue) *string {
// 	if s != nil {
// 		str := s.GetValue()
// 		return &str
// 	}
// 	return nil
// }

// func getStringValue(s string) string {
// 	if s != nil {
// 		return *s
// 	}
// 	return ""
// }

// func getFloatValue(f float64) float64 {
// 	if f != nil {
// 		return *f
// 	}
// 	return 0
// }

// func getTimeValue(t time.Time) string {
// 	if t != nil {
// 		return t.Format("2006-01-02")
// 	}
// 	return ""
// }

// func parseDatePointer(dateStr string) *time.Time {
// 	if dateStr == "" {
// 		return nil
// 	}
// 	t, err := time.Parse("2006-01-02", dateStr)
// 	if err != nil {
// 		return nil
// 	}
// 	return &t
// }

// func uintPointer(u uint32) *uint {
// 	if u != 0 {
// 		temp := uint(u)
// 		return &temp
// 	}
// 	return nil
// }

// func getUint32Value(u *uint) uint32 {
// 	if u != nil {
// 		return uint32(*u)
// 	}
// 	return 0
// }

func parseDate(dateStr string) (time.Time, error) {
	// Try parsing in YYYY-MM-DD format
	date, err := time.Parse("2006-01-02", dateStr)
	if err == nil {
		return date, nil
	}

	// Try parsing RFC3339 (ISO 8601 format)
	date, err = time.Parse(time.RFC3339, dateStr)
	if err == nil {
		return date, nil
	}

	// Return error if both formats fail
	return time.Time{}, fmt.Errorf("invalid date format: %s", dateStr)
}
