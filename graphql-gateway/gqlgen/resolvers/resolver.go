package resolvers

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"
	"errors"
	"fmt"
	"graphql-gateway/gqlgen/generated"
	"graphql-gateway/gqlgen/model"
	"graphql-gateway/grpc/activitypb"
	"graphql-gateway/grpc/authpb"
	"graphql-gateway/grpc/contactpb"
	"graphql-gateway/grpc/finance_pb"
	"graphql-gateway/grpc/inventory_pb"
	"graphql-gateway/grpc/leadspb"
	"graphql-gateway/grpc/opportunitypb"
	"graphql-gateway/grpc/organizationpb"
	"graphql-gateway/grpc/userpb"
	"graphql-gateway/grpc/vms_pb"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Resolver struct {
	AuthClient            authpb.AuthServiceClient
	UserClient            userpb.UserServiceClient
	OrganizationClient    organizationpb.OrganizationServiceClient
	LeadClient            leadspb.LeadServiceClient
	OpportunityClient     opportunitypb.OpportunityServiceClient
	ContactClient         contactpb.ContactServiceClient
	ActivityClient        activitypb.ActivityServiceClient
	VendorClient          vms_pb.VendorServiceClient
	PerformanceClient     vms_pb.PerformanceServiceClient
	PurchaseOrderClient   vms_pb.PurchaseOrderServiceClient
	PaymentClient         vms_pb.PaymentServiceClient
	InvoiceClient         finance_pb.InvoiceServiceClient
	CreditDebitNoteClient finance_pb.CreditDebitNoteServiceClient
	PaymentDueClient      finance_pb.PaymentServiceClient
	LedgerClient          finance_pb.LedgerServiceClient
	InventoryClient       inventory_pb.InventoryServiceClient
}

// Signup is the resolver for the signup field.
func (r *mutationResolver) Signup(ctx context.Context, firstName string, lastName string, email string, password string, phone string, role string, organizationName string, gstIn string, address string, city string, state string, country string, zipcode string, website string) (*model.SignupResponse, error) {
	authResponse, err := r.AuthClient.Signup(ctx, &authpb.SignupRequest{
		FirstName:        firstName,
		LastName:         lastName,
		Email:            email,
		Password:         password,
		Phone:            phone,
		Role:             role,
		OrganizationName: organizationName,
		GstIn:            gstIn,
		Address:          address,
		City:             city,
		State:            state,
		Country:          country,
		Zipcode:          zipcode,
		Website:          website,
	})
	if err != nil {
		return nil, err
	}

	return &model.SignupResponse{
		Message:        authResponse.Message,
		UserID:         authResponse.UserId,
		OrganizationID: authResponse.OrganizationId,
	}, nil
}

// Signin is the resolver for the signin field.
func (r *mutationResolver) Signin(ctx context.Context, email string, password string) (*model.SigninResponse, error) {
	authResponse, err := r.AuthClient.Signin(ctx, &authpb.SigninRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return &model.SigninResponse{
		AccessToken:  authResponse.AccessToken,
		RefreshToken: authResponse.RefreshToken,
	}, nil
}

// CreateLead is the resolver for the createLead field.
func (r *mutationResolver) CreateLead(ctx context.Context, input model.CreateLeadInput) (*model.Lead, error) {
	// Log the input values for debugging purposes
	log.Printf("CreateLead called with: firstName=%s, lastName=%s, email=%s, phone=%v, status=%s, assignedTo=%s, organizationID=%s",
		input.FirstName, input.LastName, input.Email, input.Phone, input.Status, input.AssignedTo, input.OrganizationID)

	// Validate required fields before making the gRPC request
	if input.FirstName == "" {
		return nil, fmt.Errorf("firstName is required")
	}
	if input.LastName == "" {
		return nil, fmt.Errorf("lastName is required")
	}
	if input.Email == "" {
		return nil, fmt.Errorf("email is required")
	}
	if input.Status == "" {
		return nil, fmt.Errorf("status is required")
	}

	// Prepare the request for the gRPC service
	leadRequest := &leadspb.CreateLeadRequest{
		Lead: &leadspb.Lead{
			FirstName:  input.FirstName,
			LastName:   input.LastName,
			Email:      input.Email,
			Phone:      getValue(input.Phone),
			Status:     string(input.Status),
			AssignedTo: input.AssignedTo,
			OrganizationId: func() uint32 {
				orgID, err := strconv.ParseUint(input.OrganizationID, 10, 32)
				if err != nil {
					log.Printf("Invalid organization ID: %v", err)
					return 0
				}
				return uint32(orgID)
			}(),
		},
	}

	// Log the gRPC request to verify what is being sent
	log.Printf("Sending gRPC request: %+v", leadRequest)

	// Call the gRPC leads service to create a lead
	leadResponse, err := r.LeadClient.CreateLead(ctx, leadRequest)
	if err != nil {
		log.Printf("gRPC CreateLead error: %v", err)
		return nil, err
	}

	if leadResponse == nil || leadResponse.Lead == nil {
		log.Println("Received nil response or lead from CreateLead")
		return nil, fmt.Errorf("invalid response from leads service")
	}

	// Log the gRPC response for debugging
	log.Printf("Received gRPC response: %+v", leadResponse)

	// Convert gRPC response to GraphQL model
	return &model.Lead{
		ID:         strconv.Itoa(int(leadResponse.Lead.Id)),
		FirstName:  leadResponse.Lead.FirstName,
		LastName:   leadResponse.Lead.LastName,
		Email:      leadResponse.Lead.Email,
		Phone:      getPointer(leadResponse.Lead.Phone),
		Status:     model.LeadStatus(leadResponse.Lead.Status),
		AssignedTo: leadResponse.Lead.AssignedTo,
		Organization: &model.Organization{
			ID: strconv.Itoa(int(leadResponse.Lead.OrganizationId)),
		},
	}, nil
}

// UpdateLead is the resolver for the updateLead field.
func (r *mutationResolver) UpdateLead(ctx context.Context, input model.UpdateLeadInput) (*model.Lead, error) {
	leadID, err := strconv.ParseUint(input.ID, 10, 32)
	if err != nil {
		return nil, err
	}

	var orgID uint32
	if input.OrganizationID != nil {
		parsedOrgID, err := strconv.ParseUint(*input.OrganizationID, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid organization ID: %v", err)
		}
		orgID = uint32(parsedOrgID)
	}

	leadResponse, err := r.LeadClient.UpdateLead(ctx, &leadspb.UpdateLeadRequest{
		Lead: &leadspb.Lead{
			Id:             uint32(leadID),
			FirstName:      getValue(input.FirstName),
			LastName:       getValue(input.LastName),
			Email:          getValue(input.Email),
			Phone:          getValue(input.Phone),
			Status:         getValue((*string)(input.Status)),
			AssignedTo:     input.AssignedTo,
			OrganizationId: orgID,
		},
	})
	if err != nil {
		return nil, err
	}

	return &model.Lead{
		ID:         strconv.Itoa(int(leadResponse.Lead.Id)),
		FirstName:  leadResponse.Lead.FirstName,
		LastName:   leadResponse.Lead.LastName,
		Email:      leadResponse.Lead.Email,
		Phone:      getPointer(leadResponse.Lead.Phone),
		Status:     model.LeadStatus(leadResponse.Lead.Status),
		AssignedTo: leadResponse.Lead.AssignedTo,
		Organization: &model.Organization{
			ID: strconv.Itoa(int(leadResponse.Lead.OrganizationId)),
		},
	}, nil
}

// DeleteLead is the resolver for the deleteLead field.
func (r *mutationResolver) DeleteLead(ctx context.Context, id string) (*bool, error) {
	leadID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}

	_, err = r.LeadClient.DeleteLead(ctx, &leadspb.DeleteLeadRequest{Id: uint32(leadID)})
	if err != nil {
		return nil, err
	}

	success := true
	return &success, nil
}

// CreateOpportunity is the resolver for the createOpportunity field.
func (r *mutationResolver) CreateOpportunity(ctx context.Context, input model.CreateOpportunityInput) (*model.Opportunity, error) {
	// Parse and validate IDs
	leadId, err := parseUint32(input.LeadID)
	if err != nil {
		return nil, fmt.Errorf("invalid lead ID: %v", err)
	}

	ownerID, err := parseUint32(input.OwnerID)
	if err != nil {
		return nil, fmt.Errorf("invalid owner ID: %v", err)
	}

	accountId, err := parseOptionalUint32(input.AccountID)
	if err != nil {
		return nil, fmt.Errorf("invalid account ID: %v", err)
	}

	// Validate probability (if provided)
	if input.Probability != nil {
		if *input.Probability < 0 || *input.Probability > 1 {
			return nil, fmt.Errorf("probability must be between 0 and 1")
		}
	}

	// Validate amount
	if input.Amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than zero")
	}

	// Validate closeDate format
	_, err = time.Parse("2006-01-02", input.CloseDate)
	if err != nil {
		return nil, fmt.Errorf("invalid close date format (expected YYYY-MM-DD): %v", err)
	}

	// Prepare the opportunity data
	opportunityProto := &opportunitypb.Opportunity{
		Name:        wrapperspb.String(input.Name),
		Description: wrapperspb.String(getValue(input.Description)),
		Stage:       wrapperspb.String(input.Stage),
		Amount:      wrapperspb.Double(input.Amount),
		CloseDate:   wrapperspb.String(input.CloseDate),
		Probability: wrapperspb.Double(getFloatValue(input.Probability)),
		LeadId:      leadId,
		AccountId:   accountId,
		OwnerId:     ownerID,
	}

	// Log the opportunity data being sent
	log.Printf("Creating opportunity with data: %+v", opportunityProto)

	// Call the gRPC service
	res, err := r.OpportunityClient.CreateOpportunity(ctx, &opportunitypb.CreateOpportunityRequest{
		Opportunity: opportunityProto,
	})
	if err != nil {
		log.Printf("Error creating opportunity: %v", err)
		return nil, err
	}

	return convertProtoToGraphQLOpportunity(res.Opportunity), nil
}

// UpdateOpportunity is the resolver for the updateOpportunity field.
func (r *mutationResolver) UpdateOpportunity(ctx context.Context, input model.UpdateOpportunityInput) (*model.Opportunity, error) {
	opportunityId, err := parseUint32(input.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid opportunity ID: %v", err)
	}

	opportunityProto := &opportunitypb.Opportunity{
		Id: opportunityId,
	}

	if input.Name != nil {
		opportunityProto.Name = wrapperspb.String(*input.Name)
	}
	if input.Description != nil {
		opportunityProto.Description = wrapperspb.String(*input.Description)
	}
	if input.Stage != nil {
		opportunityProto.Stage = wrapperspb.String(*input.Stage)
	}
	if input.Amount != nil {
		opportunityProto.Amount = wrapperspb.Double(*input.Amount)
	}
	if input.CloseDate != nil {
		// Validate closeDate format
		_, err = time.Parse("2006-01-02", *input.CloseDate)
		if err != nil {
			return nil, fmt.Errorf("invalid close date format (expected YYYY-MM-DD): %v", err)
		}
		opportunityProto.CloseDate = wrapperspb.String(*input.CloseDate)
	}
	if input.Probability != nil {
		if *input.Probability < 0 || *input.Probability > 1 {
			return nil, fmt.Errorf("probability must be between 0 and 1")
		}
		opportunityProto.Probability = wrapperspb.Double(*input.Probability)
	}
	if input.LeadID != nil {
		leadId, err := parseUint32(*input.LeadID)
		if err != nil {
			return nil, fmt.Errorf("invalid lead ID: %v", err)
		}
		opportunityProto.LeadId = leadId
	}
	if input.AccountID != nil {
		accountId, err := parseUint32(*input.AccountID)
		if err != nil {
			return nil, fmt.Errorf("invalid account ID: %v", err)
		}
		opportunityProto.AccountId = accountId
	}
	if input.OwnerID != nil {
		ownerID, err := parseUint32(*input.OwnerID)
		if err != nil {
			return nil, fmt.Errorf("invalid owner ID: %v", err)
		}
		opportunityProto.OwnerId = ownerID
	}

	res, err := r.OpportunityClient.UpdateOpportunity(ctx, &opportunitypb.UpdateOpportunityRequest{
		Opportunity: opportunityProto,
	})
	if err != nil {
		return nil, err
	}

	return convertProtoToGraphQLOpportunity(res.Opportunity), nil
}

// DeleteOpportunity is the resolver for the deleteOpportunity field.
func (r *mutationResolver) DeleteOpportunity(ctx context.Context, id string) (*bool, error) {
	opportunityId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid opportunity ID: %v", err)
	}

	res, err := r.OpportunityClient.DeleteOpportunity(ctx, &opportunitypb.DeleteOpportunityRequest{
		Id: uint32(opportunityId),
	})
	if err != nil {
		return nil, err
	}

	success := res.Success
	return &success, nil
}

// Implement createContact mutation
func (r *mutationResolver) CreateContact(ctx context.Context, input model.CreateContactInput) (*model.Contact, error) {
	// Validate input if necessary

	// Convert input to gRPC contact
	contactProto := &contactpb.Contact{
		FirstName:           input.FirstName,
		LastName:            input.LastName,
		Email:               input.Email,
		Phone:               getValue(input.Phone),
		Address:             getValue(input.Address),
		City:                getValue(input.City),
		State:               getValue(input.State),
		Country:             getValue(input.Country),
		ZipCode:             getValue(input.ZipCode),
		Company:             getValue(input.Company),
		Position:            getValue(input.Position),
		SocialMediaProfiles: getValue(input.SocialMediaProfiles),
		Notes:               getValue(input.Notes),
	}

	// Call gRPC service
	res, err := r.ContactClient.CreateContact(ctx, &contactpb.CreateContactRequest{Contact: contactProto})
	if err != nil {
		// Handle errors, possibly map gRPC errors to GraphQL errors
		return nil, fmt.Errorf("failed to create contact: %v", err)
	}

	// Convert gRPC response to GraphQL model
	return convertProtoToGraphQLContact(res.Contact), nil
}

// UpdateContact is the resolver for the updateContact field.
func (r *mutationResolver) UpdateContact(ctx context.Context, input model.UpdateContactInput) (*model.Contact, error) {
	contactID, err := strconv.ParseUint(input.ID, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid contact ID: %v", err)
	}

	contactProto := &contactpb.Contact{
		Id:                  uint32(contactID),
		FirstName:           getValue(input.FirstName),
		LastName:            getValue(input.LastName),
		Email:               getValue(input.Email),
		Phone:               getValue(input.Phone),
		Address:             getValue(input.Address),
		City:                getValue(input.City),
		State:               getValue(input.State),
		Country:             getValue(input.Country),
		ZipCode:             getValue(input.ZipCode),
		Company:             getValue(input.Company),
		Position:            getValue(input.Position),
		SocialMediaProfiles: getValue(input.SocialMediaProfiles),
		Notes:               getValue(input.Notes),
	}

	res, err := r.ContactClient.UpdateContact(ctx, &contactpb.UpdateContactRequest{Contact: contactProto})
	if err != nil {
		return nil, err
	}

	return convertProtoToGraphQLContact(res.Contact), nil
}

// DeleteContact is the resolver for the deleteContact mutation.
func (r *mutationResolver) DeleteContact(ctx context.Context, id string) (*bool, error) {
	// Parse the ID from string to uint32
	contactID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid contact ID format: %v", err)
	}

	// Call the gRPC DeleteContact method
	res, err := r.ContactClient.DeleteContact(ctx, &contactpb.DeleteContactRequest{Id: uint32(contactID)})
	if err != nil {
		// Handle gRPC errors and map them to meaningful GraphQL errors
		grpcErr, ok := status.FromError(err)
		if ok {
			switch grpcErr.Code() {
			case codes.NotFound:
				return nil, fmt.Errorf("contact not found")
			case codes.Internal:
				return nil, fmt.Errorf("internal server error")
			default:
				return nil, fmt.Errorf("failed to delete contact: %v", grpcErr.Message())
			}
		}
		return nil, fmt.Errorf("failed to delete contact: %v", err)
	}

	// Ensure res is not nil
	if res == nil {
		success := false
		return &success, nil
	}

	// Return the success status from the gRPC response
	return &res.Success, nil
}

// CreateActivity is the resolver for the createActivity mutation.
func (r *mutationResolver) CreateActivity(ctx context.Context, input model.CreateActivityInput) (*model.Activity, error) {
	// Convert GraphQL input to gRPC Activity
	activityProto := &activitypb.Activity{
		Title:       input.Title,
		Description: getValue(input.Description),
		Type:        input.Type,
		Status:      string(input.Status), // Convert enum to string
		DueDate:     getValue(input.DueDate),
		ContactId:   uint32(parseID(input.ContactID)),
	}

	// Call the gRPC CreateActivity method
	res, err := r.ActivityClient.CreateActivity(ctx, &activitypb.CreateActivityRequest{Activity: activityProto})
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok {
			switch grpcErr.Code() {
			case codes.AlreadyExists:
				return nil, fmt.Errorf("activity with this title already exists")
			case codes.InvalidArgument:
				return nil, fmt.Errorf("invalid activity data: %v", grpcErr.Message())
			case codes.NotFound:
				return nil, fmt.Errorf("associated contact not found")
			default:
				return nil, fmt.Errorf("failed to create activity: %v", grpcErr)
			}
		}
		return nil, fmt.Errorf("failed to create activity: %v", err)
	}

	// Fetch Contact details via ContactClient
	contactRes, err := r.ContactClient.GetContact(ctx, &contactpb.GetContactRequest{Id: res.Activity.ContactId})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch associated contact: %v", err)
	}

	// Convert the gRPC Activity and Contact to GraphQL Activity model
	activity := &model.Activity{
		ID:          strconv.Itoa(int(res.Activity.Id)),
		Title:       res.Activity.Title,
		Description: getPointer(res.Activity.Description),
		Type:        res.Activity.Type,
		Status:      model.ActivityStatus(res.Activity.Status), // Convert enum to model.ActivityStatus
		DueDate:     getPointer(res.Activity.DueDate),
		CreatedAt:   res.Activity.CreatedAt,
		UpdatedAt:   res.Activity.UpdatedAt,
		Contact: &model.Contact{
			ID:        strconv.Itoa(int(contactRes.Contact.Id)),
			FirstName: contactRes.Contact.FirstName,
			LastName:  contactRes.Contact.LastName,
			Email:     contactRes.Contact.Email,
			// Populate other fields as needed
		},
		Tasks: []*model.Task{}, // Populate tasks if needed
	}

	return activity, nil
}

// UpdateActivity is the resolver for the updateActivity field.
func (r *mutationResolver) UpdateActivity(ctx context.Context, input model.UpdateActivityInput) (*model.Activity, error) {
	// Convert GraphQL input to gRPC Activity
	activityProto := &activitypb.UpdateActivityRequest{
		Activity: &activitypb.Activity{
			Id:          uint32(parseID(input.ID)),
			Title:       getValue(input.Title),
			Description: getValue(input.Description),
			Type:        getValue(input.Type),
			Status:      getValue((*string)(input.Status)), // Ensure ActivityStatus is correctly converted
			DueDate:     getValue(input.DueDate),
			ContactId:   uint32(parseID(*input.ContactID)),
		},
	}

	// Call the gRPC UpdateActivity method
	res, err := r.ActivityClient.UpdateActivity(ctx, activityProto)
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok {
			switch grpcErr.Code() {
			case codes.NotFound:
				return nil, fmt.Errorf("activity not found")
			case codes.InvalidArgument:
				return nil, fmt.Errorf("invalid activity data: %v", grpcErr.Message())
			default:
				return nil, fmt.Errorf("failed to update activity: %v", grpcErr.Message())
			}
		}
		return nil, fmt.Errorf("failed to update activity: %v", err)
	}

	// Fetch Contact details via ContactClient
	contactRes, err := r.ContactClient.GetContact(ctx, &contactpb.GetContactRequest{Id: res.Activity.ContactId})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch associated contact: %v", err)
	}

	// Convert the gRPC Activity and Contact to GraphQL Activity model
	activity := &model.Activity{
		ID:          strconv.Itoa(int(res.Activity.Id)),
		Title:       res.Activity.Title,
		Description: getPointer(res.Activity.Description),
		Type:        res.Activity.Type,
		Status:      model.ActivityStatus(res.Activity.Status), // Use the correct enum type
		DueDate:     getPointer(res.Activity.DueDate),
		CreatedAt:   res.Activity.CreatedAt,
		UpdatedAt:   res.Activity.UpdatedAt,
		Contact: &model.Contact{
			ID:        strconv.Itoa(int(contactRes.Contact.Id)),
			FirstName: contactRes.Contact.FirstName,
			LastName:  contactRes.Contact.LastName,
			Email:     contactRes.Contact.Email,
			// Populate other fields as needed
		},
		Tasks: []*model.Task{}, // Populate tasks if needed
	}

	return activity, nil
}

// DeleteActivity is the resolver for the deleteActivity mutation.
func (r *mutationResolver) DeleteActivity(ctx context.Context, id string) (*bool, error) {
	// Parse the ID from string to uint32
	activityID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Printf("Invalid activity ID format: %v", err)
		return nil, fmt.Errorf("invalid activity ID format: %v", err)
	}

	// Call the gRPC DeleteActivity method
	res, err := r.ActivityClient.DeleteActivity(ctx, &activitypb.DeleteActivityRequest{Id: uint32(activityID)})
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok {
			switch grpcErr.Code() {
			case codes.NotFound:
				log.Printf("Activity not found: ID %d", activityID)
				return nil, fmt.Errorf("activity not found")
			case codes.Internal:
				log.Printf("Internal server error while deleting activity: %v", grpcErr.Message())
				return nil, fmt.Errorf("internal server error")
			default:
				log.Printf("Failed to delete activity: %v", grpcErr.Message())
				return nil, fmt.Errorf("failed to delete activity: %v", grpcErr.Message())
			}
		}
		log.Printf("Failed to delete activity: %v", err)
		return nil, fmt.Errorf("failed to delete activity: %v", err)
	}

	// Ensure res is not nil
	if res == nil {
		success := false
		return &success, nil
	}

	// Return the success status from the gRPC response
	return &res.Success, nil
}

// CreateTask is the resolver for the createTask mutation.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
	// Convert GraphQL input to gRPC Task
	taskProto := &activitypb.Task{
		Title:       input.Title,
		Description: getValue(input.Description),
		Status:      string(input.Status),    // Convert enum to string
		Priority:    string(input.Priority),  // Convert enum to string
		DueDate:     getValue(input.DueDate), // Assuming ISO8601 format
		ActivityId:  uint32(parseID(input.ActivityID)),
	}

	// Call the gRPC CreateTask method
	res, err := r.ActivityClient.CreateTask(ctx, &activitypb.CreateTaskRequest{Task: taskProto})
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok {
			switch grpcErr.Code() {
			case codes.AlreadyExists:
				log.Printf("Task already exists: Title %s", input.Title)
				return nil, fmt.Errorf("task with this title already exists")
			case codes.InvalidArgument:
				log.Printf("Invalid argument for creating task: %v", grpcErr.Message())
				return nil, fmt.Errorf("invalid task data: %v", grpcErr.Message())
			case codes.NotFound:
				log.Printf("Associated activity not found: ID %s", input.ActivityID)
				return nil, fmt.Errorf("associated activity not found")
			case codes.Internal:
				log.Printf("Internal server error while creating task: %v", grpcErr.Message())
				return nil, fmt.Errorf("internal server error")
			default:
				log.Printf("Failed to create task: %v", grpcErr.Message())
				return nil, fmt.Errorf("failed to create task: %v", grpcErr.Message())
			}
		}
		log.Printf("Failed to create task: %v", err)
		return nil, fmt.Errorf("failed to create task: %v", err)
	}

	// Convert the gRPC Task to GraphQL Task model
	task := convertProtoToGraphQLTask(res.Task)
	return task, nil
}

// UpdateTask is the resolver for the updateTask mutation.
func (r *mutationResolver) UpdateTask(ctx context.Context, input model.UpdateTaskInput) (*model.Task, error) {
	// Parse the ID from string to uint32
	taskID, err := strconv.ParseUint(input.ID, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID: %v", err)
	}

	// Convert input to gRPC Task
	taskProto := &activitypb.Task{
		Id:          uint32(taskID),
		Title:       getValue(input.Title),
		Description: getValue(input.Description),
		Status:      getValue((*string)(input.Status)),   // Convert enum to string pointer
		Priority:    getValue((*string)(input.Priority)), // Convert enum to string pointer
		DueDate:     getValue(input.DueDate),
		ActivityId:  uint32(parseID(*input.ActivityID)),
	}

	// Call the gRPC UpdateTask method
	res, err := r.ActivityClient.UpdateTask(ctx, &activitypb.UpdateTaskRequest{Task: taskProto})
	if err != nil {
		return nil, err
	}

	// Convert the gRPC Task to GraphQL Task model
	task := convertProtoToGraphQLTask(res.Task)
	return task, nil
}

// DeleteTask is the resolver for the deleteTask mutation.
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (*bool, error) {
	// Parse the ID from string to uint32
	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID format: %v", err)
	}

	// Call the gRPC DeleteTask method
	res, err := r.ActivityClient.DeleteTask(ctx, &activitypb.DeleteTaskRequest{Id: uint32(taskID)})
	if err != nil {
		return nil, err
	}

	// Ensure res is not nil
	if res == nil {
		success := false
		return &success, nil
	}

	// Return the success status from the gRPC response
	return &res.Success, nil
}

// CreateVendor is the resolver for the createVendor field.
func (r *mutationResolver) CreateVendor(ctx context.Context, name string, category string, service string, industry string, gstin string, certifications *string, licenses *string) (*model.Vendor, error) {
	createdVendor, err := r.VendorClient.CreateVendor(ctx, &vms_pb.CreateVendorRequest{
		Vendor: &vms_pb.Vendor{
			Name:           name,
			Category:       category,
			Service:        service,
			Industry:       industry,
			Gstin:          gstin,
			Certifications: *certifications,
			Licenses:       *licenses,
		},
	})
	if err != nil {
		return nil, err
	}
	return &model.Vendor{
		ID:             createdVendor.Vendor.Id,
		Name:           createdVendor.Vendor.Name,
		Category:       createdVendor.Vendor.Category,
		Service:        createdVendor.Vendor.Service,
		Industry:       createdVendor.Vendor.Industry,
		Gstin:          createdVendor.Vendor.Gstin,
		Certifications: &createdVendor.Vendor.Certifications,
		Licenses:       &createdVendor.Vendor.Licenses,
	}, nil
}

// UpdateVendor is the resolver for the updateVendor field.
func (r *mutationResolver) UpdateVendor(ctx context.Context, id string, name string, category string, service string, industry string, gstin string, certifications *string, licenses *string) (*model.Vendor, error) {
	req := &vms_pb.UpdateVendorRequest{
		Vendor: &vms_pb.Vendor{
			Id:             id,
			Name:           name,
			Category:       category,
			Service:        service,
			Industry:       industry,
			Gstin:          gstin,
			Certifications: *certifications,
			Licenses:       *licenses,
		},
	}

	// Call UpdateVendor on VendorClient
	res, err := r.VendorClient.UpdateVendor(ctx, req)
	if err != nil {
		return nil, err
	}

	// Map response to GraphQL model.Vendor
	return &model.Vendor{
		ID:             res.Vendor.Id,
		Name:           res.Vendor.Name,
		Category:       res.Vendor.Category,
		Service:        res.Vendor.Service,
		Industry:       res.Vendor.Industry,
		Gstin:          res.Vendor.Gstin,
		Certifications: &res.Vendor.Certifications,
		Licenses:       &res.Vendor.Licenses,
	}, nil
}

// DeleteVendor is the resolver for the deleteVendor field.
func (r *mutationResolver) DeleteVendor(ctx context.Context, id string) (*string, error) {
	req := &vms_pb.DeleteVendorRequest{Id: id}

	// Call DeleteVendor on VendorClient
	res, err := r.VendorClient.DeleteVendor(ctx, req)
	if err != nil {
		return nil, err
	}

	return &res.Message, nil
}

// CreatePurchaseOrder is the resolver for the createPurchaseOrder field.
func (r *mutationResolver) CreatePurchaseOrder(ctx context.Context, vendorID string, orderDetails string, deliveryDate string) (*model.PurchaseOrder, error) {
	createdOrder, err := r.PurchaseOrderClient.CreatePurchaseOrder(ctx, &vms_pb.CreatePurchaseOrderRequest{
		VendorId:     vendorID,
		OrderDetails: orderDetails,
		DeliveryDate: func() *timestamppb.Timestamp {
			t, err := time.Parse("2006-01-02", deliveryDate)
			if err != nil {
				log.Printf("Invalid delivery date format: %v", err)
				return nil
			}
			return timestamppb.New(t)
		}(),
	})
	if err != nil {
		return nil, err
	}
	return &model.PurchaseOrder{
		ID:           createdOrder.Id,
		VendorID:     createdOrder.VendorId,
		OrderDetails: createdOrder.OrderDetails,
		DeliveryDate: &createdOrder.DeliveryDate,
	}, nil
}

// UpdatePurchaseOrder is the resolver for the updatePurchaseOrder field.
func (r *mutationResolver) UpdatePurchaseOrder(ctx context.Context, id string, orderDetails string, status string, deliveryDate *string, receivedDate *string) (*model.PurchaseOrder, error) {
	req := &vms_pb.UpdatePurchaseOrderRequest{
		PurchaseOrder: &vms_pb.PurchaseOrder{
			Id:           id,
			OrderDetails: orderDetails,
			Status:       status,
			DeliveryDate: func() *timestamppb.Timestamp {
				t, err := time.Parse("2006-01-02", *deliveryDate)
				if err != nil {
					log.Printf("Invalid delivery date format: %v", err)
					return nil
				}
				return timestamppb.New(t)
			}(),
			ReceivedDate: func() *timestamppb.Timestamp {
				t, err := time.Parse("2006-01-02", *receivedDate)
				if err != nil {
					log.Printf("Invalid received date format: %v", err)
					return nil
				}
				return timestamppb.New(t)
			}(),
		},
	}

	// Call UpdatePurchaseOrder on PurchaseOrderClient
	res, err := r.PurchaseOrderClient.UpdatePurchaseOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	return &model.PurchaseOrder{
		ID:           res.Id,
		OrderDetails: res.OrderDetails,
		Status:       res.Status,
		DeliveryDate: &res.DeliveryDate,
		ReceivedDate: &res.ReceivedDate,
	}, nil
}

// DeletePurchaseOrder is the resolver for the deletePurchaseOrder field.
func (r *mutationResolver) DeletePurchaseOrder(ctx context.Context, id string) (*string, error) {
	req := &vms_pb.DeletePurchaseOrderRequest{Id: id}

	// Call DeletePurchaseOrder on PurchaseOrderClient
	res, err := r.PurchaseOrderClient.DeletePurchaseOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	return &res.Message, nil
}

// RecordPerformance is the resolver for the recordPerformance field.
func (r *mutationResolver) RecordPerformance(ctx context.Context, vendorID string, score float64, riskLevel string, evaluatedAt string) (*model.VendorPerformance, error) {
	req := &vms_pb.RecordPerformanceRequest{
		VendorId:  vendorID,
		Score:     float32(score),
		RiskLevel: riskLevel,
		EvaluatedAt: func() *timestamppb.Timestamp {
			t, err := time.Parse("2006-01-02", evaluatedAt)
			if err != nil {
				log.Printf("Invalid evaluatedAt date format: %v", err)
				return nil
			}
			return timestamppb.New(t)
		}(),
	}

	// Call RecordPerformance on PerformanceClient
	res, err := r.PerformanceClient.RecordPerformance(ctx, req)
	if err != nil {
		return nil, err
	}

	return &model.VendorPerformance{
		VendorID:    res.VendorId,
		Score:       getPointerFloat64(float64(res.Score)),
		RiskLevel:   getPointer(res.RiskLevel),
		EvaluatedAt: getPointer(res.EvaluatedAt),
	}, nil
}

// ProcessInvoice is the resolver for the processInvoice field.
func (r *mutationResolver) ProcessInvoice(ctx context.Context, purchaseOrderID string, amount float64, paymentTerms string) (*model.Payment, error) {
	req := &vms_pb.ProcessInvoiceRequest{
		PurchaseOrderId: purchaseOrderID,
		Amount:          float32(amount),
		PaymentTerms:    paymentTerms,
	}

	// Call ProcessInvoice on PaymentClient
	res, err := r.PaymentClient.ProcessInvoice(ctx, req)
	if err != nil {
		return nil, err
	}

	return &model.Payment{
		ID:              res.Id,
		PurchaseOrderID: res.PurchaseOrderId,
		Amount:          float64(res.Amount),
		PaymentTerms:    &res.PaymentTerms,
		Status:          res.Status,
	}, nil
}

// CreateInvoice is the resolver for the createInvoice field.
func (r *mutationResolver) CreateInvoice(ctx context.Context, input model.CreateInvoiceInput) (*model.Invoice, error) {
	if input.OrganizationID == "" {
		log.Println("OrganizationID is missing or empty")
		return nil, errors.New("organization ID is required")
	}
	log.Printf("Creating invoice with input: %+v", input)

	req := &finance_pb.CreateInvoiceRequest{
		Invoice: &finance_pb.Invoice{
			Type:           input.Type,
			VendorId:       getValue(input.VendorID),
			CustomerId:     getValue(input.CustomerID),
			OrganizationId: input.OrganizationID, // Dereference the pointer
			InvoiceDate:    timestamppb.Now(),
			Items:          convertInvoiceItemsToProto(input.Items),
		},
	}

	res, err := r.InvoiceClient.CreateInvoice(ctx, req)
	if err != nil {
		log.Printf("Error creating invoice: %v", err)
		return nil, err
	}

	return convertProtoToGraphQLInvoice(res.Invoice), nil
}

// UpdateInvoice is the resolver for the updateInvoice field.
func (r *mutationResolver) UpdateInvoice(ctx context.Context, input model.UpdateInvoiceInput) (*model.Invoice, error) {
	req := &finance_pb.UpdateInvoiceRequest{
		Invoice: &finance_pb.Invoice{
			Id:          *&input.InvoiceID,
			InvoiceDate: timestamppb.Now(),
			Items:       convertInvoiceItemsToProto(input.Items),
		},
	}

	res, err := r.InvoiceClient.UpdateInvoice(ctx, req)
	if err != nil {
		log.Printf("Error updating invoice: %v", err)
		return nil, err
	}

	return convertProtoToGraphQLInvoice(res.Invoice), nil
}

// DeleteInvoice is the resolver for the deleteInvoice field.
func (r *mutationResolver) DeleteInvoice(ctx context.Context, id string) (string, error) {
	req := &finance_pb.DeleteInvoiceRequest{InvoiceId: id}
	res, err := r.InvoiceClient.DeleteInvoice(ctx, req)
	if err != nil {
		log.Printf("Error deleting invoice: %v", err)
		return "", err
	}

	return res.Message, nil
}

// CreateCreditDebitNote is the resolver for the createCreditDebitNote field.
func (r *mutationResolver) CreateCreditDebitNote(ctx context.Context, input model.CreateCreditDebitNoteInput) (*model.CreditDebitNote, error) {
	req := &finance_pb.CreateCreditDebitNoteRequest{
		InvoiceId: input.InvoiceID,
		Type:      input.Type,
		Amount:    input.Amount,
		Reason:    input.Reason,
	}

	res, err := r.CreditDebitNoteClient.CreateCreditDebitNote(ctx, req)
	if err != nil {
		log.Printf("Error creating credit/debit note: %v", err)
		return nil, err
	}

	return convertProtoToGraphQLCreditDebitNote(&finance_pb.CreditDebitNote{
		NoteId:    res.NoteId,
		InvoiceId: res.InvoiceId,
		Type:      res.Type,
		Amount:    res.Amount,
		Reason:    res.Reason,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}), nil
}

// DeleteCreditDebitNote is the resolver for the deleteCreditDebitNote field.
func (r *mutationResolver) DeleteCreditDebitNote(ctx context.Context, id string) (string, error) {
	req := &finance_pb.DeleteCreditDebitNoteRequest{NoteId: id}
	res, err := r.CreditDebitNoteClient.DeleteCreditDebitNote(ctx, req)
	if err != nil {
		log.Printf("Error deleting credit/debit note: %v", err)
		return "", err
	}

	return res.Message, nil
}

// AddLedgerEntry is the resolver for the addLedgerEntry field.
func (r *mutationResolver) AddLedgerEntry(ctx context.Context, input model.CreateLedgerEntryInput) (*model.LedgerEntry, error) {
	req := &finance_pb.AddLedgerEntryRequest{
		TransactionId:   input.TransactionID,
		Description:     input.Description,
		Debit:           *input.Debit,
		Credit:          *input.Credit,
		TransactionDate: timestamppb.Now(), // Convert input.TransactionDate if needed
	}

	res, err := r.LedgerClient.AddLedgerEntry(ctx, req)
	if err != nil {
		log.Printf("Error adding ledger entry: %v", err)
		return nil, err
	}

	return convertProtoToGraphQLLedgerEntry(res), nil
}

// DeleteLedgerEntry is the resolver for the deleteLedgerEntry field.
func (r *mutationResolver) DeleteLedgerEntry(ctx context.Context, id string) (string, error) {
	req := &finance_pb.DeleteLedgerEntryRequest{EntryId: id}
	res, err := r.LedgerClient.DeleteLedgerEntry(ctx, req)
	if err != nil {
		log.Printf("Error deleting ledger entry: %v", err)
		return "", err
	}

	return res.Message, nil
}

// AddPaymentDue is the resolver for the addPaymentDue field.
func (r *mutationResolver) AddPaymentDue(ctx context.Context, input model.AddPaymentDueInput) (*model.PaymentDue, error) {
	req := &finance_pb.AddPaymentDueRequest{
		InvoiceId: input.InvoiceID,
		AmountDue: input.AmountDue,
		DueDate:   timestamppb.Now(), // You may convert input.DueDate to `timestamppb` if provided.
	}

	res, err := r.PaymentDueClient.AddPaymentDue(ctx, req)
	if err != nil {
		log.Printf("Error adding payment due: %v", err)
		return nil, err
	}

	// Assuming res.PaymentDue has been updated to return a PaymentDueResponse
	return convertProtoToGraphQLPaymentDue(res), nil
}

// Payment Due-related resolvers
func (r *mutationResolver) MarkPaymentAsPaid(ctx context.Context, id string) (*model.PaymentDue, error) {
	req := &finance_pb.MarkPaymentAsPaidRequest{PaymentDueId: id}
	res, err := r.PaymentDueClient.MarkPaymentAsPaid(ctx, req)
	if err != nil {
		log.Printf("Error marking payment as paid: %v", err)
		return nil, err
	}

	return convertProtoToGraphQLPaymentDue(res), nil
}

// CreateInventoryItem is the resolver for the createInventoryItem field.
func (r *mutationResolver) CreateInventoryItem(ctx context.Context, productID string, productName string, productDescription *string, sku string, supplierID string, category *string, price float64, availableQuantity int, reorderPoint int) (*model.InventoryItem, error) {
	// Create item data structure
	newItem := &model.InventoryItem{
		ProductID:          productID,
		ProductName:        productName,
		ProductDescription: productDescription,
		Sku:                sku,
		SupplierID:         supplierID,
		Category:           category,
		Price:              float64(price),
		AvailableQuantity:  availableQuantity,
		ReorderPoint:       reorderPoint,
	}

	// Call service layer to create item
	// Create gRPC request
	req := &inventory_pb.CreateInventoryItemRequest{
		Item: &inventory_pb.InventoryItem{
			ProductId:          newItem.ProductID,
			ProductName:        newItem.ProductName,
			ProductDescription: getValue(newItem.ProductDescription),
			Sku:                newItem.Sku,
			SupplierId:         newItem.SupplierID,
			Category:           getValue(newItem.Category),
			Price:              float32(newItem.Price),
			AvailableQuantity:  int32(newItem.AvailableQuantity),
			ReorderPoint:       int32(newItem.ReorderPoint),
		},
	}

	// Call gRPC service
	res, err := r.InventoryClient.CreateInventoryItem(ctx, req)
	if err != nil {
		return nil, err
	}

	// Convert created item to GraphQL model
	return &model.InventoryItem{
		ProductID:          res.Item.ProductId,
		ProductName:        res.Item.ProductName,
		ProductDescription: &res.Item.ProductDescription,
		Sku:                res.Item.Sku,
		SupplierID:         res.Item.SupplierId,
		Category:           &res.Item.Category,
		Price:              float64(res.Item.Price),
		AvailableQuantity:  int(res.Item.AvailableQuantity),
		ReorderPoint:       int(res.Item.ReorderPoint),
	}, nil
}

// UpdateInventoryItem is the resolver for the updateInventoryItem field.
func (r *mutationResolver) UpdateInventoryItem(ctx context.Context, productID string, productName *string, productDescription *string, sku *string, supplierID *string, category *string, price *float64, availableQuantity *int, reorderPoint *int) (*model.InventoryItem, error) {
	updatedItem := &model.InventoryItem{
		ProductID:          productID,
		ProductName:        *productName,
		ProductDescription: productDescription,
		Sku:                *sku,
		SupplierID:         *supplierID,
		Category:           category,
		Price:              *price,
		AvailableQuantity:  *availableQuantity,
		ReorderPoint:       *reorderPoint,
	}
	req := &inventory_pb.UpdateInventoryItemRequest{
		Item: &inventory_pb.InventoryItem{
			ProductId:          updatedItem.ProductID,
			ProductName:        updatedItem.ProductName,
			ProductDescription: getValue(updatedItem.ProductDescription),
			Sku:                updatedItem.Sku,
			SupplierId:         updatedItem.SupplierID,
			Category:           getValue(updatedItem.Category),
			Price:              float32(updatedItem.Price),
			AvailableQuantity:  int32(updatedItem.AvailableQuantity),
			ReorderPoint:       int32(updatedItem.ReorderPoint),
		},
	}
	result, err := r.InventoryClient.UpdateInventoryItem(ctx, req)
	if err != nil {
		return nil, err
	}
	return &model.InventoryItem{
		ProductID:          result.Item.ProductId,
		ProductName:        result.Item.ProductName,
		ProductDescription: &result.Item.ProductDescription,
		Sku:                result.Item.Sku,
		SupplierID:         result.Item.SupplierId,
		Category:           &result.Item.Category,
		Price:              float64(result.Item.Price),
		AvailableQuantity:  int(result.Item.AvailableQuantity),
		ReorderPoint:       int(result.Item.ReorderPoint),
	}, nil
}

// DeleteInventoryItem is the resolver for the deleteInventoryItem field.
func (r *mutationResolver) DeleteInventoryItem(ctx context.Context, productID string) (*bool, error) {
	// Call service layer to delete item
	_, err := r.InventoryClient.DeleteInventoryItem(ctx, &inventory_pb.DeleteInventoryItemRequest{ProductId: productID})
	if err != nil {
		return nil, err
	}

	// Return success
	success := true
	return &success, nil
}

// SetReorderPoint is the resolver for the setReorderPoint field.
func (r *mutationResolver) SetReorderPoint(ctx context.Context, productID string, reorderPoint int) (*bool, error) {
	req := &inventory_pb.SetReorderPointRequest{
		ProductId:    productID,
		ReorderPoint: int32(reorderPoint),
	}
	_, err := r.InventoryClient.SetReorderPoint(ctx, req)
	success := err == nil
	return &success, err
}

// AddOrUpdateInventoryItem is the resolver for the addOrUpdateInventoryItem field.
func (r *mutationResolver) AddOrUpdateInventoryItem(ctx context.Context, productID string, productName string, productDescription *string, sku string, supplierID string, category *string, price float64, availableQuantity int, reorderPoint int, warehouseStocks []*model.WarehouseStockInput) (*model.InventoryItem, error) {
	// Prepare the gRPC request for AddOrUpdateInventoryItem
	req := &inventory_pb.AddOrUpdateInventoryItemRequest{
		Item: &inventory_pb.InventoryItem{
			ProductId:          productID,
			ProductName:        productName,
			ProductDescription: getValue(productDescription),
			Sku:                sku,
			SupplierId:         supplierID,
			Category:           getValue(category),
			Price:              float32(price),
			AvailableQuantity:  int32(availableQuantity),
			ReorderPoint:       int32(reorderPoint),
		},
	}

	// Populate warehouse stock data if provided
	for _, ws := range warehouseStocks {
		req.Item.WarehouseStocks = append(req.Item.WarehouseStocks, &inventory_pb.WarehouseStock{
			WarehouseId: ws.WarehouseID,
			StockLevel:  int32(ws.StockLevel),
		})
	}

	// Call the AddOrUpdateInventoryItem gRPC method
	res, err := r.InventoryClient.AddOrUpdateInventoryItem(ctx, req)
	if err != nil {
		return nil, err
	}

	// Check if the operation was successful
	if !res.Success {
		return nil, errors.New("failed to add or update inventory item")
	}

	// Convert the updated item to GraphQL model.InventoryItem
	updatedItem := res.UpdatedItem
	return &model.InventoryItem{
		ProductID:          updatedItem.ProductId,
		ProductName:        updatedItem.ProductName,
		ProductDescription: &updatedItem.ProductDescription,
		Sku:                updatedItem.Sku,
		SupplierID:         updatedItem.SupplierId,
		Category:           &updatedItem.Category,
		Price:              float64(updatedItem.Price),
		AvailableQuantity:  int(updatedItem.AvailableQuantity),
		ReorderPoint:       int(updatedItem.ReorderPoint),
		WarehouseStocks: func() []*model.WarehouseStock {
			var stocks []*model.WarehouseStock
			for _, ws := range updatedItem.WarehouseStocks {
				stocks = append(stocks, &model.WarehouseStock{
					WarehouseID: ws.WarehouseId,
					StockLevel:  int(ws.StockLevel),
				})
			}
			return stocks
		}(),
	}, nil
}

// ProcessOrder is the resolver for the processOrder field.
func (r *mutationResolver) ProcessOrder(ctx context.Context, orderItems []*model.OrderItemInput) (*bool, error) {
	var serviceOrderItems []*inventory_pb.OrderItem
	for _, item := range orderItems {
		serviceOrderItems = append(serviceOrderItems, &inventory_pb.OrderItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity),
		})
	}

	// Call service layer to process order
	_, err := r.InventoryClient.ProcessOrder(ctx, &inventory_pb.ProcessOrderRequest{
		Items: serviceOrderItems,
	})
	if err != nil {
		return nil, err
	}

	// Return success
	success := true
	return &success, nil
}

// GeneratePickingList is the resolver for the generatePickingList field.
func (r *mutationResolver) GeneratePickingList(ctx context.Context, orderID string) ([]*model.PickingItem, error) {
	// Call service layer to generate picking list
	pickingItems, err := r.InventoryClient.GeneratePickingList(ctx, &inventory_pb.GeneratePickingListRequest{OrderId: orderID})
	if err != nil {
		return nil, err
	}

	// Convert service picking items to GraphQL model
	var gqlPickingItems []*model.PickingItem
	for _, item := range pickingItems.PickingList {
		gqlPickingItems = append(gqlPickingItems, &model.PickingItem{
			ProductID:   item.ProductId,
			Quantity:    int(item.Quantity),
			WarehouseID: item.WarehouseId,
		})
	}

	return gqlPickingItems, nil
}

// UpdateInventoryStock is the resolver for the updateInventoryStock field.
func (r *mutationResolver) UpdateInventoryStock(ctx context.Context, productID string, quantity int, warehouseID string) (*bool, error) {
	req := &inventory_pb.UpdateInventoryRequest{
		ProductId:   productID,
		Quantity:    int32(quantity),
		WarehouseId: warehouseID,
	}
	_, err := r.InventoryClient.UpdateInventory(ctx, req)
	success := err == nil
	return &success, err
}

// PlaceVendorOrder is the resolver for the placeVendorOrder field.
func (r *mutationResolver) PlaceVendorOrder(ctx context.Context, vendorID string, orderItems []*model.OrderItemInput) (*bool, error) {
	var items []model.OrderItemInput
	for _, i := range orderItems {
		items = append(items, model.OrderItemInput{
			ProductID: i.ProductID,
			Quantity:  i.Quantity,
		})
	}
	ctx = context.Background()
	res, err := r.InventoryClient.PlaceVendorOrder(ctx, &inventory_pb.PlaceVendorOrderRequest{
		VendorId: vendorID,
		Items: func() []*inventory_pb.VendorOrderItem {
			var vendorOrderItems []*inventory_pb.VendorOrderItem
			for _, item := range items {
				vendorOrderItems = append(vendorOrderItems, &inventory_pb.VendorOrderItem{
					ProductId: item.ProductID,
					Quantity:  int32(item.Quantity),
				})
			}
			return vendorOrderItems
		}(),
	})
	if err != nil {
		return nil, err
	}
	success := res.Success
	return &success, nil
}

// NotifyFinanceForOrder is the resolver for the notifyFinanceForOrder field.
func (r *mutationResolver) NotifyFinanceForOrder(ctx context.Context, orderID string, totalAmount float64) (*bool, error) {
	req := &inventory_pb.NotifyFinanceRequest{
		OrderId:     orderID,
		TotalAmount: float32(totalAmount),
	}
	_, err := r.InventoryClient.NotifyFinanceForOrder(ctx, req)
	success := err == nil
	return &success, err
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}
	userResponse, err := r.UserClient.GetUser(ctx, &userpb.GetUserRequest{Id: uint32(userID)})
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        strconv.Itoa(int(userResponse.Id)),
		FirstName: userResponse.FirstName,
		LastName:  userResponse.LastName,
		Email:     userResponse.Email,
		Phone:     userResponse.Phone,
		Role:      userResponse.Role,
		Organization: &model.Organization{
			ID: strconv.Itoa(int(userResponse.OrganizationId)),
		}, // Ensure the field exists in the struct
	}, nil
}

// GetOrganization is the resolver for the getOrganization field.
func (r *queryResolver) GetOrganization(ctx context.Context, id string) (*model.Organization, error) {
	orgID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}
	orgResponse, err := r.OrganizationClient.GetOrganization(ctx, &organizationpb.GetOrganizationRequest{Id: uint32(orgID)})
	if err != nil {
		return nil, err
	}

	return &model.Organization{
		ID:       strconv.Itoa(int(orgResponse.Id)),
		GstIn:    orgResponse.GstIn,
		Name:     orgResponse.Name,
		Phone:    orgResponse.Phone,
		Email:    orgResponse.Email,
		Address:  orgResponse.Address,
		City:     orgResponse.City,
		Country:  orgResponse.Country,
		State:    orgResponse.State,
		Zipcode:  orgResponse.Zipcode,
		Website:  orgResponse.Website,
		Industry: orgResponse.Industry,
	}, nil
}

// Leads is the resolver for the leads field.
func (r *queryResolver) Leads(ctx context.Context) ([]*model.Lead, error) {
	res, err := r.LeadClient.GetAllLeads(ctx, &leadspb.GetAllLeadsRequest{})
	if err != nil {
		return nil, err
	}

	var leads []*model.Lead
	for _, grpcLead := range res.Leads {
		leads = append(leads, &model.Lead{
			ID:         strconv.Itoa(int(grpcLead.Id)),
			FirstName:  grpcLead.FirstName,
			LastName:   grpcLead.LastName,
			Email:      grpcLead.Email,
			Phone:      getPointer(grpcLead.Phone),
			Status:     model.LeadStatus(grpcLead.Status), // Convert string to enum
			AssignedTo: grpcLead.AssignedTo,
			Organization: &model.Organization{ // Ensure the field exists in the struct
				ID: strconv.Itoa(int(grpcLead.OrganizationId)),
			},
		})
	}

	return leads, nil
}

// Lead is the resolver for the lead field.
func (r *queryResolver) Lead(ctx context.Context, id string) (*model.Lead, error) {
	leadID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}

	res, err := r.LeadClient.GetLead(ctx, &leadspb.GetLeadRequest{Id: uint32(leadID)})
	if err != nil {
		return nil, err
	}

	return &model.Lead{
		ID:        strconv.Itoa(int(res.Lead.Id)),
		FirstName: res.Lead.FirstName,
		LastName:  res.Lead.LastName,
		Status:    model.LeadStatus(res.Lead.Status), // Convert string to enum
		Phone:     getPointer(res.Lead.Phone),
		Organization: &model.Organization{ // Ensure the field exists in the struct
			ID: strconv.Itoa(int(res.Lead.OrganizationId)),
		},
		AssignedTo: res.Lead.AssignedTo,
	}, nil
}

// LeadByEmail is the resolver for the leadByEmail field.
func (r *queryResolver) LeadByEmail(ctx context.Context, email string) (*model.Lead, error) {
	res, err := r.LeadClient.GetLeadByEmail(ctx, &leadspb.GetLeadByEmailRequest{Email: email})
	if err != nil {
		return nil, err
	}

	return &model.Lead{
		ID:         strconv.Itoa(int(res.Lead.Id)),
		FirstName:  res.Lead.FirstName,
		LastName:   res.Lead.LastName,
		Email:      res.Lead.Email,
		Phone:      getPointer(res.Lead.Phone),
		Status:     model.LeadStatus(res.Lead.Status),
		AssignedTo: res.Lead.AssignedTo,
		Organization: &model.Organization{
			ID: strconv.Itoa(int(res.Lead.OrganizationId)),
		},
	}, nil
}

// GetOpportunity is the resolver for the getOpportunity field.
func (r *queryResolver) GetOpportunity(ctx context.Context, id string) (*model.Opportunity, error) {
	opportunityId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}

	res, err := r.OpportunityClient.GetOpportunity(ctx, &opportunitypb.GetOpportunityRequest{
		Id: uint32(opportunityId),
	})
	if err != nil {
		return nil, err
	}

	return convertProtoToGraphQLOpportunity(res.Opportunity), nil
}

// ListOpportunities is the resolver for the listOpportunities field.
func (r *queryResolver) ListOpportunities(ctx context.Context, ownerID *string) ([]*model.Opportunity, error) {
	var ownerIDUint uint32
	if ownerID != nil {
		id, err := strconv.ParseUint(*ownerID, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid owner ID: %v", err)
		}
		ownerIDUint = uint32(id)
	}

	res, err := r.OpportunityClient.ListOpportunities(ctx, &opportunitypb.ListOpportunitiesRequest{
		OwnerId: ownerIDUint,
	})
	if err != nil {
		return nil, err
	}

	var opportunities []*model.Opportunity
	for _, opp := range res.Opportunities {
		opportunities = append(opportunities, convertProtoToGraphQLOpportunity(opp))
	}

	return opportunities, nil
}

// Similar implementations for getContact, updateContact, deleteContact, listContacts
func (r *queryResolver) GetContact(ctx context.Context, id string) (*model.Contact, error) {
	contactID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid contact ID: %v", err)
	}

	res, err := r.ContactClient.GetContact(ctx, &contactpb.GetContactRequest{Id: uint32(contactID)})
	if err != nil {
		return nil, err
	}

	return convertProtoToGraphQLContact(res.Contact), nil
}

// ListContacts is the resolver for the listContacts field.
func (r *queryResolver) ListContacts(ctx context.Context, pageNumber *int, pageSize *int, sortBy *model.ContactSortField, ascending *bool) ([]*model.Contact, error) {
	res, err := r.ContactClient.ListContacts(ctx, &contactpb.ListContactsRequest{})
	if err != nil {
		return nil, err
	}

	var contacts []*model.Contact
	for _, contact := range res.Contacts {
		contacts = append(contacts, convertProtoToGraphQLContact(contact))
	}

	return contacts, nil
}

// GetActivity is the resolver for the getActivity query.
func (r *queryResolver) GetActivity(ctx context.Context, id string) (*model.Activity, error) {
	// Parse the ID from string to uint32
	activityID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Printf("Invalid activity ID format: %v", err)
		return nil, fmt.Errorf("invalid activity ID format: %v", err)
	}

	// Call the gRPC GetActivity method
	res, err := r.ActivityClient.GetActivity(ctx, &activitypb.GetActivityRequest{Id: uint32(activityID)})
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok {
			switch grpcErr.Code() {
			case codes.NotFound:
				log.Printf("Activity not found: ID %d", activityID)
				return nil, fmt.Errorf("activity not found")
			case codes.Internal:
				log.Printf("Internal server error while retrieving activity: %v", grpcErr.Message())
				return nil, fmt.Errorf("internal server error")
			default:
				log.Printf("Failed to get activity: %v", grpcErr.Message())
				return nil, fmt.Errorf("failed to get activity: %v", grpcErr.Message())
			}
		}
		log.Printf("Failed to get activity: %v", err)
		return nil, fmt.Errorf("failed to get activity: %v", err)
	}

	// Convert the gRPC Activity to GraphQL Activity model
	activity := convertProtoToGraphQLActivity(res.Activity)
	return activity, nil
}

// ListActivities is the resolver for the listActivities field.
func (r *queryResolver) ListActivities(ctx context.Context, pageNumber *int, pageSize *int, sortBy model.ActivitySortField, ascending bool, contactID *string) ([]*model.Activity, error) {
	log.Println("Resolver: ListActivities called")

	// Defensive checks for clients
	if r.ActivityClient == nil {
		log.Println("Error: ActivityClient is nil")
		return nil, fmt.Errorf("internal server error: ActivityClient not initialized")
	}
	if r.ContactClient == nil {
		log.Println("Error: ContactClient is nil")
		return nil, fmt.Errorf("internal server error: ContactClient not initialized")
	}

	// Mapping GraphQL enums to repository sort fields
	sortByMap := map[model.ActivitySortField]string{
		model.ActivitySortFieldTitle:     "title",
		model.ActivitySortFieldCreatedat: "created_at",
		model.ActivitySortFieldUpdatedat: "updated_at",
		model.ActivitySortFieldDuedate:   "due_date",
	}

	// Validate and map sortBy
	sortByStr, ok := sortByMap[sortBy]
	if !ok {
		log.Printf("Error: Invalid sort field '%s'", sortBy)
		return nil, fmt.Errorf("invalid sort field '%s'", sortBy)
	}

	log.Printf("SortBy mapped to '%s'", sortByStr)

	// Set default pagination values if not provided
	pn := uint32(1) // Default page number
	if pageNumber != nil && *pageNumber > 0 {
		pn = uint32(*pageNumber)
	}
	log.Printf("PageNumber: %d", pn)

	ps := uint32(10) // Default page size
	if pageSize != nil && *pageSize > 0 {
		ps = uint32(*pageSize)
	}
	log.Printf("PageSize: %d", ps)

	// Parse contactID if provided
	var contactIDUint uint32 = 0
	if contactID != nil {
		parsedID, err := strconv.ParseUint(*contactID, 10, 32)
		if err != nil {
			log.Printf("Error: Invalid contact ID format '%s'", *contactID)
			return nil, fmt.Errorf("invalid contact ID format")
		}
		contactIDUint = uint32(parsedID)
		log.Printf("ContactID: %d", contactIDUint)
	} else {
		log.Println("No ContactID provided")
	}

	// Call the gRPC ListActivities method
	log.Println("Calling gRPC ActivityClient.ListActivities")
	res, err := r.ActivityClient.ListActivities(ctx, &activitypb.ListActivitiesRequest{
		PageNumber: pn,
		PageSize:   ps,
		SortBy:     sortByStr,
		Ascending:  ascending,
		ContactId:  contactIDUint,
	})
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok {
			switch grpcErr.Code() {
			case codes.InvalidArgument:
				log.Printf("Invalid argument for listing activities: %v", grpcErr.Message())
				return nil, fmt.Errorf("invalid request: %v", grpcErr.Message())
			case codes.Internal:
				log.Printf("Internal server error while listing activities: %v", grpcErr.Message())
				return nil, fmt.Errorf("internal server error")
			default:
				log.Printf("Failed to list activities: %v", grpcErr.Message())
				return nil, fmt.Errorf("failed to list activities: %v", grpcErr.Message())
			}
		}
		log.Printf("Failed to list activities: %v", err)
		return nil, fmt.Errorf("failed to list activities: %v", err)
	}

	log.Printf("Fetched %d activities from gRPC service", len(res.Activities))

	// Convert the list of gRPC Activities to GraphQL Activity models
	activities := make([]*model.Activity, 0, len(res.Activities))
	for _, protoActivity := range res.Activities {
		log.Printf("Processing activity ID: %d", protoActivity.Id)

		// Fetch Contact details via ContactClient
		log.Printf("Fetching contact for activity ID: %d", protoActivity.Id)
		contactRes, err := r.ContactClient.GetContact(ctx, &contactpb.GetContactRequest{Id: protoActivity.ContactId})
		if err != nil {
			log.Printf("Failed to fetch contact for activity ID %d: %v", protoActivity.Id, err)
			return nil, fmt.Errorf("failed to fetch contact for activity ID %d", protoActivity.Id)
		}

		if contactRes.Contact == nil {
			log.Printf("Contact is nil for activity ID %d", protoActivity.Id)
			return nil, fmt.Errorf("contact not found for activity ID %d", protoActivity.Id)
		}

		// Convert gRPC Activity to GraphQL Activity
		activity := &model.Activity{
			ID:          strconv.Itoa(int(protoActivity.Id)),
			Title:       protoActivity.Title,
			Description: nullableString(protoActivity.Description),
			Type:        protoActivity.Type,
			Status:      model.ActivityStatus(protoActivity.Status), // Convert string to enum
			DueDate:     nullableString(protoActivity.DueDate),
			CreatedAt:   protoActivity.CreatedAt,
			UpdatedAt:   protoActivity.UpdatedAt,
			Contact: &model.Contact{
				ID:        strconv.Itoa(int(contactRes.Contact.Id)),
				FirstName: contactRes.Contact.FirstName,
				LastName:  contactRes.Contact.LastName,
				Email:     contactRes.Contact.Email,
				// Populate other fields as needed
			},
			Tasks: []*model.Task{}, // Optionally populate tasks
		}

		activities = append(activities, activity)
		log.Printf("Added activity ID: %d to response", protoActivity.Id)
	}

	log.Println("Resolver: ListActivities completed successfully")
	return activities, nil
}

// GetTask is the resolver for the getTask query.
func (r *queryResolver) GetTask(ctx context.Context, id string) (*model.Task, error) {
	// Parse the ID from string to uint32
	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Printf("Invalid task ID format: %v", err)
		return nil, fmt.Errorf("invalid task ID format: %v", err)
	}

	// Call the gRPC GetTask method
	res, err := r.ActivityClient.GetTask(ctx, &activitypb.GetTaskRequest{Id: uint32(taskID)})
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok {
			switch grpcErr.Code() {
			case codes.NotFound:
				log.Printf("Task not found: ID %d", taskID)
				return nil, fmt.Errorf("task not found")
			case codes.Internal:
				log.Printf("Internal server error while retrieving task: %v", grpcErr.Message())
				return nil, fmt.Errorf("internal server error")
			default:
				log.Printf("Failed to get task: %v", grpcErr.Message())
				return nil, fmt.Errorf("failed to get task: %v", grpcErr.Message())
			}
		}
		log.Printf("Failed to get task: %v", err)
		return nil, fmt.Errorf("failed to get task: %v", err)
	}

	// Convert the gRPC Task to GraphQL Task model
	task := convertProtoToGraphQLTask(res.Task)
	return task, nil
}

// ListTasks is the resolver for the listTasks query.
func (r *queryResolver) ListTasks(ctx context.Context, pageNumber *int, pageSize *int, sortBy model.TaskSortField, ascending bool, activityID *string) ([]*model.Task, error) {
	// Set default pagination parameters if not provided
	pn := uint(1)
	ps := uint(10)
	if pageNumber != nil && *pageNumber > 0 {
		pn = uint(*pageNumber)
	}
	if pageSize != nil && *pageSize > 0 {
		ps = uint(*pageSize)
	}
	sb := ""
	if sortBy != "" {
		sb = string(sortBy)
	}

	asc := true
	if ascending {
		asc = ascending
	}

	// Parse activity ID if provided
	aid := uint(0)
	if activityID != nil && *activityID != "" {
		aidParsed, err := strconv.ParseUint(*activityID, 10, 32)
		if err != nil {
			log.Printf("Invalid activity ID format: %v", err)
			return nil, fmt.Errorf("invalid activity ID format: %v", err)
		}
		aid = uint(aidParsed)
	}

	// Call the gRPC ListTasks method
	res, err := r.ActivityClient.ListTasks(ctx, &activitypb.ListTasksRequest{
		PageNumber: uint32(pn),
		PageSize:   uint32(ps),
		SortBy:     sb,
		Ascending:  asc,
		ActivityId: uint32(aid),
	})
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok {
			switch grpcErr.Code() {
			case codes.InvalidArgument:
				log.Printf("Invalid argument for listing tasks: %v", grpcErr.Message())
				return nil, fmt.Errorf("invalid request: %v", grpcErr.Message())
			case codes.Internal:
				log.Printf("Internal server error while listing tasks: %v", grpcErr.Message())
				return nil, fmt.Errorf("internal server error")
			default:
				log.Printf("Failed to list tasks: %v", grpcErr.Message())
				return nil, fmt.Errorf("failed to list tasks: %v", grpcErr.Message())
			}
		}
		log.Printf("Failed to list tasks: %v", err)
		return nil, fmt.Errorf("failed to list tasks: %v", err)
	}

	// Convert the list of gRPC Tasks to GraphQL Task models
	tasks := make([]*model.Task, len(res.Tasks))
	for i, protoTask := range res.Tasks {
		tasks[i] = convertProtoToGraphQLTask(protoTask)
	}

	return tasks, nil
}

// GetVendorByID is the resolver for the getVendorByID field.
func (r *queryResolver) GetVendorByID(ctx context.Context, id string) (*model.Vendor, error) {
	vendorID := id
	req := &vms_pb.GetVendorByIDRequest{Id: vendorID}
	res, err := r.VendorClient.GetVendorByID(ctx, req)
	if err != nil {
		log.Printf("Error fetching vendor: %v", err)
		return nil, err
	}
	return convertProtoToGraphQLVendor(res.Vendor), nil
}

// SearchVendors is the resolver for the searchVendors field.
func (r *queryResolver) SearchVendors(ctx context.Context, query string) ([]*model.Vendor, error) {
	req := &vms_pb.SearchVendorsRequest{Query: query}
	res, err := r.VendorClient.SearchVendors(ctx, req)
	if err != nil {
		log.Printf("Error searching vendors: %v", err)
		return nil, err
	}

	var vendors []*model.Vendor
	for _, vendor := range res.Vendors {
		vendors = append(vendors, convertProtoToGraphQLVendor(vendor))
	}

	return vendors, nil
}

// GetPurchaseOrderByID is the resolver for the getPurchaseOrderByID field.
func (r *queryResolver) GetPurchaseOrderByID(ctx context.Context, id string) (*model.PurchaseOrder, error) {
	// Prepare the gRPC request
	req := &vms_pb.GetPurchaseOrderByIDRequest{Id: id}

	// Call the gRPC service to get the purchase order by ID
	res, err := r.PurchaseOrderClient.GetPurchaseOrderByID(ctx, req)
	if err != nil {
		log.Printf("Error fetching purchase order: %v", err)
		return nil, err
	}

	// Convert the gRPC response to the GraphQL model
	return &model.PurchaseOrder{
		ID:           res.Id,
		VendorID:     res.VendorId,
		OrderDetails: res.OrderDetails,
		Status:       res.Status,
		DeliveryDate: &res.DeliveryDate, // Assuming this is a string in GraphQL
		ReceivedDate: &res.ReceivedDate, // Assuming this is a string in GraphQL
	}, nil
}

// GetPerformanceByID is the resolver for the getPerformanceByID field.
func (r *queryResolver) GetPerformanceByID(ctx context.Context, id string) (*model.VendorPerformance, error) {
	// Prepare the gRPC request
	req := &vms_pb.GetPerformanceByIDRequest{Id: id}

	// Call the gRPC service to get performance data by ID
	res, err := r.PerformanceClient.GetPerformanceByID(ctx, req)
	if err != nil {
		log.Printf("Error fetching performance data: %v", err)
		return nil, err
	}

	// Convert the gRPC response to the GraphQL model
	return &model.VendorPerformance{
		ID:          res.VendorId,
		Score:       getPointerFloat64(float64(res.Score)),
		RiskLevel:   &res.RiskLevel,
		EvaluatedAt: &res.EvaluatedAt, // Assuming this is a string in GraphQL
	}, nil
}

// GetPaymentByID is the resolver for the getPaymentByID field.
func (r *queryResolver) GetPaymentByID(ctx context.Context, id string) (*model.Payment, error) {
	// Prepare the gRPC request
	req := &vms_pb.GetPaymentByIDRequest{Id: id}

	// Call the gRPC service to get payment information by ID
	res, err := r.PaymentClient.GetPaymentByID(ctx, req)
	if err != nil {
		log.Printf("Error fetching payment: %v", err)
		return nil, err
	}

	// Convert the gRPC response to the GraphQL model
	return &model.Payment{
		ID:              res.Id,
		PurchaseOrderID: res.PurchaseOrderId,
		Amount:          float64(res.Amount),
		Status:          res.Status,
		PaymentTerms:    &res.PaymentTerms,
		PaidAt:          &res.PaidAt, // Assuming this is a string in GraphQL
	}, nil
}

// GetInvoiceByID is the resolver for the getInvoiceById field.
func (r *queryResolver) GetInvoiceByID(ctx context.Context, id string) (*model.Invoice, error) {
	req := &finance_pb.GetInvoiceByIDRequest{InvoiceId: id}
	res, err := r.InvoiceClient.GetInvoiceByID(ctx, req)
	if err != nil {
		log.Printf("Error fetching invoice by ID: %v", err)
		return nil, err
	}

	return convertProtoToGraphQLInvoice(res.Invoice), nil
}

// ListInvoices is the resolver for the listInvoices field.
func (r *queryResolver) ListInvoices(ctx context.Context, page *int, pageSize *int) ([]*model.Invoice, error) {
	req := &finance_pb.ListInvoicesRequest{
		Page:     int32(*page),
		PageSize: int32(*pageSize),
	}
	res, err := r.InvoiceClient.ListInvoices(ctx, req)
	if err != nil {
		log.Printf("Error listing invoices: %v", err)
		return nil, err
	}

	var invoices []*model.Invoice
	for _, protoInvoice := range res.Invoices {
		invoices = append(invoices, convertProtoToGraphQLInvoice(protoInvoice))
	}

	return invoices, nil
}

// Credit/Debit Note-related resolvers
func (r *queryResolver) GetCreditDebitNoteByID(ctx context.Context, id string) (*model.CreditDebitNote, error) {
	req := &finance_pb.GetCreditDebitNoteByIDRequest{NoteId: id}
	res, err := r.CreditDebitNoteClient.GetCreditDebitNoteByID(ctx, req)
	if err != nil {
		log.Printf("Error fetching credit/debit note by ID: %v", err)
		return nil, err
	}

	// Extract the CreditDebitNote from the response and pass it to the conversion function
	return convertProtoToGraphQLCreditDebitNote(&finance_pb.CreditDebitNote{
		NoteId:    res.NoteId,
		InvoiceId: res.InvoiceId,
		Type:      res.Type,
		Amount:    res.Amount,
		Reason:    res.Reason,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}), nil
}

// ListCreditDebitNotes is the resolver for the listCreditDebitNotes field.
func (r *queryResolver) ListCreditDebitNotes(ctx context.Context) ([]*model.CreditDebitNote, error) {
	res, err := r.CreditDebitNoteClient.ListAllCreditDebitNotes(ctx, &finance_pb.ListCreditDebitNotesRequest{})
	if err != nil {
		log.Printf("Error listing credit/debit notes: %v", err)
		return nil, err
	}

	var notes []*model.CreditDebitNote
	for _, protoNote := range res.Notes {
		notes = append(notes, convertProtoToGraphQLCreditDebitNote(protoNote))
	}

	return notes, nil
}

// GetLedgerEntryByID implements generated.QueryResolver.
func (r *queryResolver) GetLedgerEntryByID(ctx context.Context, id string) (*model.LedgerEntry, error) {
	req := &finance_pb.GetLedgerEntryByIDRequest{EntryId: id}
	res, err := r.LedgerClient.GetLedgerEntryByID(ctx, req)
	if err != nil {
		log.Printf("Error fetching ledger entry by ID: %v", err)
		return nil, err
	}

	// Convert the protobuf LedgerEntry to the GraphQL model
	return convertProtoToGraphQLLedgerEntry(&finance_pb.LedgerResponse{
		EntryId:         res.EntryId,
		TransactionId:   res.TransactionId,
		Description:     res.Description,
		Debit:           res.Debit,
		Credit:          res.Credit,
		Balance:         res.Balance,
		TransactionDate: res.TransactionDate,
	}), nil
}

// ListLedgerEntries is the resolver for the listLedgerEntries field.
func (r *queryResolver) ListLedgerEntries(ctx context.Context) ([]*model.LedgerEntry, error) {
	res, err := r.LedgerClient.ListLedgerEntries(ctx, &finance_pb.ListLedgerEntriesRequest{})
	if err != nil {
		log.Printf("Error listing ledger entries: %v", err)
		return nil, err
	}

	var entries []*model.LedgerEntry
	for _, protoEntry := range res.LedgerEntries {
		entries = append(entries, convertProtoToGraphQLLedgerEntry(protoEntry))
	}

	return entries, nil
}

// GetPaymentDueByID is the resolver for the getPaymentDueById field.
func (r *queryResolver) GetPaymentDueByID(ctx context.Context, id string) (*model.PaymentDue, error) {
	req := &finance_pb.GetPaymentDueByIDRequest{PaymentDueId: id}
	res, err := r.PaymentDueClient.GetPaymentDueByID(ctx, req)
	if err != nil {
		log.Printf("Error fetching payment due by ID: %v", err)
		return nil, err
	}

	// Convert the protobuf PaymentDue to the GraphQL model
	return convertProtoToGraphQLPaymentDue(&finance_pb.PaymentDueResponse{
		PaymentDueId: res.PaymentDueId,
		InvoiceId:    res.InvoiceId,
		AmountDue:    res.AmountDue,
		DueDate:      res.DueDate,
		Status:       res.Status,
	}), nil
}

// ListPaymentDues implements generated.QueryResolver.
func (r *queryResolver) ListPaymentDues(ctx context.Context) ([]*model.PaymentDue, error) {
	req := &finance_pb.ListPaymentDueRequest{}
	res, err := r.PaymentDueClient.ListPaymentDues(ctx, req)
	if err != nil {
		log.Printf("Error listing payment dues: %v", err)
		return nil, err
	}

	// Convert the list of protobuf PaymentDue to the GraphQL model
	var paymentDues []*model.PaymentDue
	for _, protoDue := range res.Payments {
		paymentDues = append(paymentDues, convertProtoToGraphQLPaymentDue(protoDue))
	}

	return paymentDues, nil
}

// GetInventoryItem is the resolver for the getInventoryItem field.
func (r *queryResolver) GetInventoryItem(ctx context.Context, productID string) (*model.InventoryItem, error) {
	// Call service layer to fetch inventory item by productID
	item, err := r.InventoryClient.GetInventoryItem(ctx, &inventory_pb.GetInventoryItemRequest{ProductId: productID})
	if err != nil {
		return nil, err
	}

	// Convert service layer item to GraphQL model (if necessary)
	return &model.InventoryItem{
		ProductID:          item.Item.ProductId,
		ProductName:        item.Item.ProductName,
		ProductDescription: &item.Item.ProductDescription,
		Sku:                item.Item.Sku,
		SupplierID:         item.Item.SupplierId,
		Category:           &item.Item.Category,
		Price:              float64(item.Item.Price),
		AvailableQuantity:  int(item.Item.AvailableQuantity),
		ReorderPoint:       int(item.Item.ReorderPoint),
	}, nil
}

// ListInventoryItems is the resolver for the listInventoryItems field.
func (r *queryResolver) ListInventoryItems(ctx context.Context, limit *int, offset *int) ([]*model.InventoryItem, error) {
	req := &inventory_pb.ListInventoryItemsRequest{
		Limit:  int32(*limit),
		Offset: int32(*offset),
	}
	items, err := r.InventoryClient.ListInventoryItems(ctx, req)
	if err != nil {
		return nil, err
	}

	var result []*model.InventoryItem
	for _, item := range items.Items {
		result = append(result, &model.InventoryItem{
			ProductID:          item.ProductId,
			ProductName:        item.ProductName,
			ProductDescription: &item.ProductDescription,
			Sku:                item.Sku,
			SupplierID:         item.SupplierId,
			Category:           &item.Category,
			Price:              float64(item.Price),
			AvailableQuantity:  int(item.AvailableQuantity),
			ReorderPoint:       int(item.ReorderPoint),
		})
	}
	return result, nil
}

// TrackInventory is the resolver for the trackInventory field.
func (r *queryResolver) TrackInventory(ctx context.Context, productID string) (*model.InventoryItem, error) {
	item, err := r.InventoryClient.TrackInventory(ctx, &inventory_pb.TrackInventoryRequest{ProductId: productID})
	if err != nil {
		return nil, err
	}
	return &model.InventoryItem{
		ProductID:          item.Item.ProductId,
		ProductName:        item.Item.ProductName,
		ProductDescription: &item.Item.ProductDescription,
		Sku:                item.Item.Sku,
		SupplierID:         item.Item.SupplierId,
		Category:           &item.Item.Category,
		Price:              float64(item.Item.Price),
		AvailableQuantity:  int(item.Item.AvailableQuantity),
		ReorderPoint:       int(item.Item.ReorderPoint),
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.

func (r *mutationResolver) ManageWarehouses(ctx context.Context, warehouses []*model.WarehouseStockInput) (*bool, error) {
	var whs []model.WarehouseStockInput
	for _, w := range warehouses {
		whs = append(whs, model.WarehouseStockInput{
			WarehouseID: w.WarehouseID,
			StockLevel:  w.StockLevel,
		})
	}
	_, err := r.InventoryClient.ManageWarehouses(ctx, &inventory_pb.ManageWarehousesRequest{
		Warehouses: func() []*inventory_pb.Warehouse {
			var protoWarehouses []*inventory_pb.Warehouse
			for _, w := range whs {
				protoWarehouses = append(protoWarehouses, &inventory_pb.Warehouse{
					WarehouseId: w.WarehouseID,
				})
			}
			return protoWarehouses
		}(),
	})
	success := err == nil
	return &success, err
}
func convertInvoiceItemsToProto(items []*model.InvoiceItemInput) []*finance_pb.InvoiceItem {
	var protoItems []*finance_pb.InvoiceItem
	for _, item := range items {
		protoItems = append(protoItems, &finance_pb.InvoiceItem{
			ItemId:   item.ItemID,
			Name:     item.Name,
			Price:    item.Price,
			Quantity: int32(item.Quantity),
		})
	}
	return protoItems
}
func convertProtoToGraphQLPurchaseOrder(po *vms_pb.PurchaseOrder) *model.PurchaseOrder {
	return &model.PurchaseOrder{
		ID:           po.Id,
		VendorID:     po.VendorId,
		OrderDetails: po.OrderDetails,
		Status:       po.Status,
		DeliveryDate: getPointer(po.DeliveryDate.AsTime().String()),
		ReceivedDate: getPointer(po.ReceivedDate.AsTime().String()),
	}
}
func parseID(idStr string) uint {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0
	}
	return uint(id)
}
func convertProtoToGraphQLVendor(protoVendor *vms_pb.Vendor) *model.Vendor {
	return &model.Vendor{
		ID:             protoVendor.Id,
		Name:           protoVendor.Name,
		Category:       protoVendor.Category,
		Service:        protoVendor.Service,
		Industry:       protoVendor.Industry,
		Gstin:          protoVendor.Gstin,
		Certifications: &protoVendor.Certifications,
		Licenses:       &protoVendor.Licenses,
	}
}
func convertProtoToGraphQLPerformance(perf *vms_pb.VendorPerformance) *model.VendorPerformance {
	return &model.VendorPerformance{
		ID:          perf.Id,
		VendorID:    perf.VendorId,
		Score:       getPointerFloat64(float64(perf.Score)),
		RiskLevel:   getPointer(perf.RiskLevel),
		EvaluatedAt: getPointer(perf.EvaluatedAt.AsTime().String()),
	}
}
func getPointerFloat64(f float64) *float64 {
	return &f
}
func convertProtoToGraphQLActivity(protoActivity *activitypb.Activity) *model.Activity {
	return &model.Activity{
		Title:       protoActivity.Title,
		Description: getPointer(protoActivity.Description),
		Type:        protoActivity.Type,
		Status:      model.ActivityStatus(protoActivity.Status), // Convert string to enum
		DueDate:     getPointer(protoActivity.DueDate),
		CreatedAt:   protoActivity.CreatedAt,
		UpdatedAt:   protoActivity.UpdatedAt,
		Contact: &model.Contact{
			ID: strconv.Itoa(int(protoActivity.ContactId)),
			// Populate other Contact fields as needed
		},
		Tasks: []*model.Task{}, // Populate tasks if needed
	}
}
func convertProtoToGraphQLTask(protoTask *activitypb.Task) *model.Task {
	return &model.Task{
		Status:      model.TaskStatus(protoTask.Status),     // Convert string to enum
		Priority:    model.TaskPriority(protoTask.Priority), // Convert string to enum
		Title:       protoTask.Title,
		Description: &protoTask.Description,
		DueDate:     &protoTask.DueDate,
		CreatedAt:   protoTask.CreatedAt,
		UpdatedAt:   protoTask.UpdatedAt,
		Activity: &model.Activity{
			ID: strconv.Itoa(int(protoTask.ActivityId)),
			// Populate other Activity fields as needed
		},
	}
}
func convertProtoToGraphQLPayment(payment *vms_pb.Payment) *model.Payment {
	return &model.Payment{
		ID:              payment.Id,
		PurchaseOrderID: payment.PurchaseOrderId,
		Amount:          float64(payment.Amount),
		Status:          payment.Status,
		PaymentTerms:    &payment.PaymentTerms,
		PaidAt:          getPointer(payment.PaidAt.AsTime().String()),
	}
}
func convertProtoToGraphQLContact(protoContact *contactpb.Contact) *model.Contact {
	return &model.Contact{
		ID:                  strconv.Itoa(int(protoContact.Id)),
		FirstName:           protoContact.FirstName,
		LastName:            protoContact.LastName,
		Email:               protoContact.Email,
		Phone:               getPointer(protoContact.Phone),
		Address:             getPointer(protoContact.Address),
		City:                getPointer(protoContact.City),
		State:               getPointer(protoContact.State),
		Country:             getPointer(protoContact.Country),
		ZipCode:             getPointer(protoContact.ZipCode),
		Company:             getPointer(protoContact.Company),
		Position:            getPointer(protoContact.Position),
		SocialMediaProfiles: getPointer(protoContact.SocialMediaProfiles),
		Notes:               getPointer(protoContact.Notes),
		CreatedAt:           getPointer(protoContact.CreatedAt),
		UpdatedAt:           getPointer(protoContact.UpdatedAt),
	}
}
func convertProtoToGraphQLOpportunity(protoOpp *opportunitypb.Opportunity) *model.Opportunity {
	if protoOpp == nil {
		return nil
	}

	return &model.Opportunity{
		ID:          strconv.FormatUint(uint64(protoOpp.Id), 10),
		Name:        getValue(getProtoStringValue(protoOpp.Name)),
		Description: getProtoStringValue(protoOpp.Description),
		Lead: &model.Lead{ // Ensure the field exists in the struct
			ID: strconv.FormatUint(uint64(protoOpp.LeadId), 10),
		},
		Amount:      getProtoFloat64Value(protoOpp.Amount),
		CloseDate:   getValue(getProtoStringValue(protoOpp.CloseDate)),
		Probability: getProtoFloat64Pointer(protoOpp.Probability),
		Owner: &model.User{ // Ensure the field exists in the struct
			ID: strconv.FormatUint(uint64(protoOpp.OwnerId), 10),
		},
	}
}
func getProtoStringValue(sv *wrapperspb.StringValue) *string {
	if sv != nil {
		return &sv.Value
	}
	return nil
}
func getProtoFloat64Pointer(fv *wrapperspb.DoubleValue) *float64 {
	if fv != nil {
		return &fv.Value
	}
	return nil
}
func getProtoFloat64Value(fv *wrapperspb.DoubleValue) float64 {
	if fv != nil {
		return fv.Value
	}
	return 0
}
func parseUint32(idStr string) (uint32, error) {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid ID '%s': %v", idStr, err)
	}
	return uint32(id), nil
}
func parseOptionalUint32(idStr *string) (uint32, error) {
	if idStr != nil {
		return parseUint32(*idStr)
	}
	return 0, nil
}
func getValue(str *string) string {
	if str != nil {
		return *str
	}
	return ""
}
func getFloatValue(f *float64) float64 {
	if f != nil {
		return *f
	}
	return 0
}
func getPointer(str string) *string {
	return &str
}
func nullableString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
func convertProtoToGraphQLInvoice(protoInvoice *finance_pb.Invoice) *model.Invoice {
	var items []*model.InvoiceItem
	for _, protoItem := range protoInvoice.Items {
		items = append(items, &model.InvoiceItem{
			ID:       protoItem.Id,
			ItemID:   protoItem.ItemId,
			Name:     protoItem.Name,
			Price:    protoItem.Price,
			Quantity: int(protoItem.Quantity),
			Total:    protoItem.Total,
		})
	}

	return &model.Invoice{
		ID:          protoInvoice.Id,
		Type:        protoInvoice.Type,
		VendorID:    getPointer(protoInvoice.VendorId),
		CustomerID:  getPointer(protoInvoice.CustomerId),
		TotalAmount: protoInvoice.TotalAmount,
		Status:      protoInvoice.Status,
		InvoiceDate: protoInvoice.InvoiceDate.String(),
		Items:       items,
	}
}
func convertProtoToGraphQLCreditDebitNote(protoNote *finance_pb.CreditDebitNote) *model.CreditDebitNote {
	return &model.CreditDebitNote{
		ID:        protoNote.NoteId,
		Type:      protoNote.Type,
		InvoiceID: protoNote.InvoiceId,
		Amount:    protoNote.Amount,
		Reason:    protoNote.Reason,
	}
}
func convertProtoToGraphQLLedgerEntry(protoEntry *finance_pb.LedgerResponse) *model.LedgerEntry {
	return &model.LedgerEntry{
		TransactionID:   protoEntry.TransactionId,
		Description:     protoEntry.Description,
		Debit:           &protoEntry.Debit,
		Credit:          &protoEntry.Credit,
		Balance:         &protoEntry.Balance,
		TransactionDate: protoEntry.TransactionDate.String(),
	}
}
func convertProtoToGraphQLPaymentDue(protoDue *finance_pb.PaymentDueResponse) *model.PaymentDue {
	return &model.PaymentDue{
		ID:        protoDue.InvoiceId,
		InvoiceID: protoDue.InvoiceId,
		AmountDue: protoDue.AmountDue,
		DueDate:   protoDue.DueDate.String(),
		Status:    protoDue.Status,
	}
}
