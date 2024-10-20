package handler

import (
	"leads-service/grpc/leadspb"
	"leads-service/internal/models"
	"log"
	"strconv"
)

// ConvertProtoToModelLead converts a protobuf lead to a models lead.
func ConvertProtoToModelLead(protoLead *leadspb.Lead) *models.Lead {
	return &models.Lead{
		ID:        uint(protoLead.Id),
		FirstName: protoLead.FirstName,
		LastName:  protoLead.LastName,
		Email:     protoLead.Email,
		Phone:     protoLead.Phone,
		Status:    protoLead.Status, // Added Status field
		AssignedTo: func() int {
			assignedTo, err := strconv.Atoi(protoLead.AssignedTo)
			if err != nil {
				return 0 // or handle the error appropriately
			}
			return assignedTo
		}(),
		OrganizationID: int(protoLead.OrganizationId),
	}
}

// ConvertModelToProtoLead converts a models lead to a protobuf lead.
func ConvertModelToProtoLead(modelLead *models.Lead) *leadspb.Lead {
	if modelLead == nil {
		log.Println("ConvertModelToProtoLead received nil modelLead")
		return nil
	}

	return &leadspb.Lead{
		Id:             uint32(modelLead.ID),
		FirstName:      modelLead.FirstName,
		LastName:       modelLead.LastName,
		Email:          modelLead.Email,
		Phone:          modelLead.Phone,
		Status:         modelLead.Status,
		AssignedTo:     strconv.Itoa(modelLead.AssignedTo),
		OrganizationId: uint32(modelLead.OrganizationID),
	}
}
