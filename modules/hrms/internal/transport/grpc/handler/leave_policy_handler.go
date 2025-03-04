package grpc

import (
	"context"
	"hrms/internal/domain"
	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

// CreateLeavePolicy handles the creation of a new leave policy
func (h *HrmsHandler) CreateLeavePolicy(ctx context.Context, req *proto.CreateLeavePolicyRequest) (*proto.LeavePolicyResponse, error) {
	leaveTypeStr := req.LeaveType.String()

	policyReq := dto.CreateLeavePolicyRequest{
		OrganizationID: uint(req.OrganizationId),
		LeaveType:      domain.LeaveType(leaveTypeStr), // Explicitly cast to domain.LeaveType
		MaxDays:        int(req.MaxDays),
		CarryForward:   req.CarryForward,
	}

	policy, err := h.HrmsUsecase.CreateLeavePolicy(ctx, policyReq)
	if err != nil {
		return nil, err
	}

	return &proto.LeavePolicyResponse{
		Policy: convertToProtoLeavePolicy(*policy),
	}, nil
}

// GetLeavePolicy retrieves a leave policy by ID
func (h *HrmsHandler) GetLeavePolicy(ctx context.Context, req *proto.GetLeavePolicyRequest) (*proto.LeavePolicyResponse, error) {
	policy, err := h.HrmsUsecase.GetLeavePolicy(ctx, uint(req.PolicyId))
	if err != nil {
		return nil, err
	}

	return &proto.LeavePolicyResponse{
		Policy: convertToProtoLeavePolicy(*policy),
	}, nil
}

// ListLeavePolicies lists leave policies for an organization
func (h *HrmsHandler) ListLeavePolicies(ctx context.Context, req *proto.ListLeavePoliciesRequest) (*proto.ListLeavePoliciesResponse, error) {
	policies, err := h.HrmsUsecase.ListLeavePolicies(ctx, uint(req.OrganizationId))
	if err != nil {
		return nil, err
	}

	response := &proto.ListLeavePoliciesResponse{}
	for _, policy := range policies {
		response.Policies = append(response.Policies, convertToProtoLeavePolicy(policy))
	}

	return response, nil
}

// UpdateLeavePolicy updates a leave policy
func (h *HrmsHandler) UpdateLeavePolicy(ctx context.Context, req *proto.UpdateLeavePolicyRequest) (*emptypb.Empty, error) {
	updateReq := dto.UpdateLeavePolicyRequest{}

	if req.MaxDays != nil {
		maxDays := int(req.GetMaxDays())
		updateReq.MaxDays = &maxDays
	}
	if req.CarryForward != nil {
		carryForward := req.GetCarryForward()
		updateReq.CarryForward = &carryForward
	}

	err := h.HrmsUsecase.UpdateLeavePolicy(ctx, uint(req.PolicyId), updateReq)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeleteLeavePolicy deletes a leave policy
func (h *HrmsHandler) DeleteLeavePolicy(ctx context.Context, req *proto.DeleteLeavePolicyRequest) (*emptypb.Empty, error) {
	err := h.HrmsUsecase.DeleteLeavePolicy(ctx, uint(req.PolicyId))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
func convertToProtoLeavePolicy(policy dto.LeavePolicyResponse) *proto.LeavePolicy {
	leaveTypeStr := string(policy.LeaveType) // Convert domain.LeaveType to string

	leaveTypeEnum, exists := proto.LeaveType_value[leaveTypeStr]
	if !exists {
		leaveTypeEnum = 0 // Default to an unknown leave type if not found
	}

	return &proto.LeavePolicy{
		Id:             uint64(policy.ID),
		OrganizationId: uint64(policy.OrganizationID),
		LeaveType:      proto.LeaveType(leaveTypeEnum),
		MaxDays:        int32(policy.MaxDays),
		CarryForward:   policy.CarryForward,
	}
}
