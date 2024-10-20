package handler

import (
	"contacts-service/grpc/contactpb"
	"contacts-service/internal/models"
	"contacts-service/internal/services"
	"context"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ContactHandler struct {
	contactService services.ContactService
	contactpb.UnimplementedContactServiceServer
}

func NewContactHandler(service services.ContactService) *ContactHandler {
	return &ContactHandler{contactService: service}
}

func (h *ContactHandler) CreateContact(ctx context.Context, req *contactpb.CreateContactRequest) (*contactpb.CreateContactResponse, error) {
	log.Printf("Received CreateContact request: %+v", req)

	// Convert Proto to Model
	contact := convertProtoToModel(req.Contact)

	// Validate and Create Contact
	createdContact, err := h.contactService.CreateContact(contact)
	if err != nil {
		log.Printf("Error creating contact: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Convert Model to Proto
	return &contactpb.CreateContactResponse{
		Contact: convertModelToProto(createdContact),
	}, nil
}

// Implement other methods: GetContact, UpdateContact, DeleteContact, ListContacts
func (h *ContactHandler) GetContact(ctx context.Context, req *contactpb.GetContactRequest) (*contactpb.GetContactResponse, error) {
	log.Printf("Received GetContact request: %+v", req)

	// Get Contact
	contact, err := h.contactService.GetContact(uint(req.Id))
	if err != nil {
		log.Printf("Error getting contact: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	// Convert Model to Proto
	return &contactpb.GetContactResponse{
		Contact: convertModelToProto(contact),
	}, nil
}

func (h *ContactHandler) UpdateContact(ctx context.Context, req *contactpb.UpdateContactRequest) (*contactpb.UpdateContactResponse, error) {
	log.Printf("Received UpdateContact request: %+v", req)

	// Convert Proto to Model
	contact := convertProtoToModel(req.Contact)

	// Validate and Update Contact
	updatedContact, err := h.contactService.UpdateContact(contact)
	if err != nil {
		log.Printf("Error updating contact: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Convert Model to Proto
	return &contactpb.UpdateContactResponse{
		Contact: convertModelToProto(updatedContact),
	}, nil
}

func (h *ContactHandler) DeleteContact(ctx context.Context, req *contactpb.DeleteContactRequest) (*contactpb.DeleteContactResponse, error) {
    log.Printf("Received DeleteContact request: %+v", req)

    // Call the service layer to delete the contact
    err := h.contactService.DeleteContact(uint(req.Id))
    if err != nil {
        if err == services.ErrContactNotFound {
            return nil, status.Error(codes.NotFound, err.Error())
        }
        return nil, status.Error(codes.Internal, err.Error())
    }

    // Return a successful response
    return &contactpb.DeleteContactResponse{
        Success: true,
    }, nil
}

func (h *ContactHandler) ListContacts(ctx context.Context, req *contactpb.ListContactsRequest) (*contactpb.ListContactsResponse, error) {
	log.Printf("Received ListContacts request: %+v", req)

	// List Contacts
	// Example arguments: page, pageSize, sortBy, ascending
	contacts, err := h.contactService.ListContacts(1, 10, "created_at", true)
	if err != nil {
		log.Printf("Error listing contacts: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Convert Model to Proto
	var protoContacts []*contactpb.Contact
	for _, contact := range contacts {
		protoContacts = append(protoContacts, convertModelToProto(&contact))
	}

	return &contactpb.ListContactsResponse{
		Contacts: protoContacts,
	}, nil
}

// Helper functions to convert between Proto and Model
func convertProtoToModel(protoContact *contactpb.Contact) *models.Contact {
	return &models.Contact{
		ID:                  uint(protoContact.Id),
		FirstName:           protoContact.FirstName,
		LastName:            protoContact.LastName,
		Email:               protoContact.Email,
		Phone:               protoContact.Phone,
		Address:             protoContact.Address,
		City:                protoContact.City,
		State:               protoContact.State,
		Country:             protoContact.Country,
		ZipCode:             protoContact.ZipCode,
		Company:             protoContact.Company,
		Position:            protoContact.Position,
		SocialMediaProfiles: protoContact.SocialMediaProfiles,
		Notes:               protoContact.Notes,
		CreatedAt:           parseTime(protoContact.CreatedAt),
		UpdatedAt:           parseTime(protoContact.UpdatedAt),
	}
}

func convertModelToProto(modelContact *models.Contact) *contactpb.Contact {
	return &contactpb.Contact{
		Id:                  uint32(modelContact.ID),
		FirstName:           modelContact.FirstName,
		LastName:            modelContact.LastName,
		Email:               modelContact.Email,
		Phone:               modelContact.Phone,
		Address:             modelContact.Address,
		City:                modelContact.City,
		State:               modelContact.State,
		Country:             modelContact.Country,
		ZipCode:             modelContact.ZipCode,
		Company:             modelContact.Company,
		Position:            modelContact.Position,
		SocialMediaProfiles: modelContact.SocialMediaProfiles,
		Notes:               modelContact.Notes,
		CreatedAt:           modelContact.CreatedAt.Format(time.RFC3339),
		UpdatedAt:           modelContact.UpdatedAt.Format(time.RFC3339),
	}
}

func parseTime(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}
	}
	return t
}
