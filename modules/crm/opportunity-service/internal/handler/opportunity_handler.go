package handler

import (
	"context"
	"log"
	"opportunity-service/grpc/opportunitypb"
	"opportunity-service/internal/models"
	"opportunity-service/internal/services"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
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

	opportunity := convertProtoToModel(req.Opportunity)

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
	if req.Opportunity.Name != nil {
		existingOpportunity.Name = &req.Opportunity.Name.Value
	}
	if req.Opportunity.Description != nil {
		existingOpportunity.Description = &req.Opportunity.Description.Value
	}
	if req.Opportunity.Stage != nil {
		existingOpportunity.Stage = &req.Opportunity.Stage.Value
	}
	if req.Opportunity.Amount.GetValue() != 0 {
		existingOpportunity.Amount = &req.Opportunity.Amount.Value
	}
	if req.Opportunity.CloseDate != nil {
		parsedDate := parseDate(req.Opportunity.CloseDate.Value)
		existingOpportunity.CloseDate = &parsedDate
	}
	if req.Opportunity.Probability != nil && req.Opportunity.Probability.Value != 0 {
		existingOpportunity.Probability = &req.Opportunity.Probability.Value
	}
	if req.Opportunity.LeadId != 0 {
		tempLeadID := uint(req.Opportunity.LeadId)
		existingOpportunity.LeadID = &tempLeadID
	}
	if req.Opportunity.AccountId != 0 {
		tempAccountID := uint(req.Opportunity.AccountId)
		existingOpportunity.AccountID = &tempAccountID
	}
	if req.Opportunity.OwnerId != 0 {
		tempOwnerID := uint(req.Opportunity.OwnerId)
		existingOpportunity.OwnerID = &tempOwnerID
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
func convertProtoToModel(protoOpp *opportunitypb.Opportunity) *models.Opportunity {
	// Convert protobuf Opportunity to models.Opportunity
	return &models.Opportunity{
		ID:          uint(protoOpp.Id),
		Name:        getStringPointer(protoOpp.Name),
		Description: getStringPointer(protoOpp.Description),
		Stage:       getStringPointer(protoOpp.Stage),
		Amount:      getFloatPointer(protoOpp.Amount),
		CloseDate:   parseDatePointer(protoOpp.CloseDate.GetValue()),
		Probability: getFloatPointer(protoOpp.Probability),
		LeadID:      uintPointer(protoOpp.LeadId),
		AccountID:   uintPointer(protoOpp.AccountId),
		OwnerID:     uintPointer(protoOpp.OwnerId),
	}
}

func convertModelToProto(modelOpp *models.Opportunity) *opportunitypb.Opportunity {
	// Convert models.Opportunity to protobuf Opportunity
	return &opportunitypb.Opportunity{
		Id:          uint32(modelOpp.ID),
		Name:        wrapperspb.String(getStringValue(modelOpp.Name)),
		Description: wrapperspb.String(getStringValue(modelOpp.Description)),
		Stage:       wrapperspb.String(getStringValue(modelOpp.Stage)),
		Amount:      wrapperspb.Double(getFloatValue(modelOpp.Amount)),
		CloseDate:   wrapperspb.String(getTimeValue(modelOpp.CloseDate)),
		Probability: wrapperspb.Double(getFloatValue(modelOpp.Probability)),
		LeadId:      getUint32Value(modelOpp.LeadID),
		AccountId:   getUint32Value(modelOpp.AccountID),
		OwnerId:     getUint32Value(modelOpp.OwnerID),
	}
}

// Helper functions

func getFloatPointer(f *wrapperspb.DoubleValue) *float64 {
	if f != nil {
		val := f.GetValue()
		return &val
	}
	return nil
}
func getStringPointer(s *wrapperspb.StringValue) *string {
	if s != nil {
		str := s.GetValue()
		return &str
	}
	return nil
}

func getStringValue(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func getFloatValue(f *float64) float64 {
	if f != nil {
		return *f
	}
	return 0
}

func getTimeValue(t *time.Time) string {
	if t != nil {
		return t.Format("2006-01-02")
	}
	return ""
}

func parseDatePointer(dateStr string) *time.Time {
	if dateStr == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil
	}
	return &t
}

func uintPointer(u uint32) *uint {
	if u != 0 {
		temp := uint(u)
		return &temp
	}
	return nil
}

func getUint32Value(u *uint) uint32 {
	if u != nil {
		return uint32(*u)
	}
	return 0
}

func parseDate(dateStr string) time.Time {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}
	}
	return date
}
