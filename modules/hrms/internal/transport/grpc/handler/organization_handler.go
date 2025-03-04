package grpc

import (
	"context"

	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateOrganization handles gRPC request to create an organization
func (h *HrmsHandler) CreateOrganization(ctx context.Context, req *proto.CreateOrganizationRequest) (*proto.OrganizationResponse, error) {
	orgDTO := dto.CreateOrganizationDTO{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
		Email:   req.Email,
	}

	createdOrg, err := h.HrmsUsecase.CreateOrganization(ctx, orgDTO)
	if err != nil {
		return nil, err
	}

	return &proto.OrganizationResponse{
		Organization: &proto.Organization{
			Id:        uint64(createdOrg.ID),
			Name:      createdOrg.Name,
			Address:   createdOrg.Address,
			Phone:     createdOrg.Phone,
			Email:     createdOrg.Email,
			CreatedAt: timestamppb.New(createdOrg.CreatedAt),
			UpdatedAt: timestamppb.New(createdOrg.UpdatedAt),
		},
	}, nil
}

// GetOrganization handles gRPC request to fetch an organization by ID
func (h *HrmsHandler) GetOrganization(ctx context.Context, req *proto.GetOrganizationRequest) (*proto.OrganizationResponse, error) {
	org, err := h.HrmsUsecase.GetOrganization(ctx, uint(req.OrgId))
	if err != nil {
		return nil, err
	}

	return &proto.OrganizationResponse{
		Organization: &proto.Organization{
			Id:        uint64(org.ID),
			Name:      org.Name,
			Address:   org.Address,
			Phone:     org.Phone,
			Email:     org.Email,
			CreatedAt: timestamppb.New(org.CreatedAt),
			UpdatedAt: timestamppb.New(org.UpdatedAt),
		},
	}, nil
}

// UpdateOrganization handles gRPC request to update an organization's details
func (h *HrmsHandler) UpdateOrganization(ctx context.Context, req *proto.UpdateOrganizationRequest) (*emptypb.Empty, error) {
	updateDTO := dto.UpdateOrganizationDTO{}

	if req.Name != nil {
		updateDTO.Name = req.Name
	}
	if req.Address != nil {
		updateDTO.Address = req.Address
	}
	if req.Phone != nil {
		updateDTO.Phone = req.Phone
	}

	err := h.HrmsUsecase.UpdateOrganization(ctx, uint(req.OrgId), updateDTO)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeleteOrganization handles gRPC request to delete an organization
func (h *HrmsHandler) DeleteOrganization(ctx context.Context, req *proto.DeleteOrganizationRequest) (*emptypb.Empty, error) {
	err := h.HrmsUsecase.DeleteOrganization(ctx, uint(req.OrgId))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// ListOrganizations handles gRPC request to list organizations with pagination
func (h *HrmsHandler) ListOrganizations(ctx context.Context, req *proto.ListOrganizationsRequest) (*proto.ListOrganizationsResponse, error) {
	orgsResponse, err := h.HrmsUsecase.ListOrganizations(ctx, int(req.Limit), int(req.Offset), getStringValue(req.Search))
	if err != nil {
		return nil, err
	}

	var orgList []*proto.Organization
	for _, org := range orgsResponse.Organizations { // Use orgsResponse instead of separate total
		orgList = append(orgList, &proto.Organization{
			Id:        uint64(org.ID),
			Name:      org.Name,
			Address:   org.Address,
			Phone:     org.Phone,
			Email:     org.Email,
			CreatedAt: timestamppb.New(org.CreatedAt),
			UpdatedAt: timestamppb.New(org.UpdatedAt),
		})
	}

	return &proto.ListOrganizationsResponse{
		Total:         int32(orgsResponse.Total),
		Limit:         req.Limit,
		Offset:        req.Offset,
		Organizations: orgList,
	}, nil
}

func getStringValue(str *string) string {
	if str == nil {
		return "" // or return some default value if needed
	}
	return *str
}
