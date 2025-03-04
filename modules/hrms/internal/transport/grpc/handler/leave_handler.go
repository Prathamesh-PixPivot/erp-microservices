package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"
	"time"
)

// CreateLeave handles the gRPC request for creating a leave request
func (h *HrmsHandler) CreateLeave(ctx context.Context, req *proto.CreateLeaveRequest) (*proto.LeaveResponse, error) {
	leaveReq := dto.CreateLeaveRequest{
		EmployeeID: uint(req.EmployeeId),
		LeaveType:  req.LeaveType.String(),
		StartDate:  time.Unix(req.StartDate.Seconds, 0),
		EndDate:    time.Unix(req.EndDate.Seconds, 0),
	}

	leave, err := h.HrmsUsecase.CreateLeave(ctx, leaveReq)
	if err != nil {
		return nil, err
	}

	// Ensure optional fields are properly handled
	var approverID *uint64
	if leave.ApproverID != 0 {
		tempApproverID := uint64(leave.ApproverID)
		approverID = &tempApproverID
	}

	var comments *string
	if leave.Comments != "" {
		comments = &leave.Comments
	}

	return &proto.LeaveResponse{
		Leave: &proto.Leave{
			Id:         uint64(leave.ID),
			EmployeeId: uint64(leave.EmployeeID),
			LeaveType:  proto.LeaveType(proto.LeaveType_value[leave.LeaveType]),
			StartDate:  timestamppb.New(leave.StartDate),
			EndDate:    timestamppb.New(leave.EndDate),
			Status:     proto.LeaveStatus(proto.LeaveStatus_value[leave.Status]),
			ApproverId: approverID,
			Comments:   comments,
		},
	}, nil
}

// GetLeave retrieves a leave request by ID
func (h *HrmsHandler) GetLeave(ctx context.Context, req *proto.GetLeaveRequest) (*proto.LeaveResponse, error) {
	leave, err := h.HrmsUsecase.GetLeaveByID(ctx, uint(req.LeaveId))
	if err != nil {
		return nil, err
	}

	// Ensure optional fields are properly handled
	var approverID *uint64
	if leave.ApproverID != 0 {
		tempApproverID := uint64(leave.ApproverID)
		approverID = &tempApproverID
	}
	var comments *string
	if leave.Comments != "" {
		comments = &leave.Comments
	}

	return &proto.LeaveResponse{
		Leave: &proto.Leave{
			Id:         uint64(leave.ID),
			EmployeeId: uint64(leave.EmployeeID),
			LeaveType:  proto.LeaveType(proto.LeaveType_value[leave.LeaveType]),
			StartDate:  timestamppb.New(leave.StartDate),
			EndDate:    timestamppb.New(leave.EndDate),
			Status:     proto.LeaveStatus(proto.LeaveStatus_value[leave.Status]),
			ApproverId: approverID,
			Comments:   comments,
		},
	}, nil
}

// UpdateLeaveStatus updates the status of a leave request
func (h *HrmsHandler) UpdateLeaveStatus(ctx context.Context, req *proto.UpdateLeaveStatusRequest) (*emptypb.Empty, error) {
	updateReq := dto.UpdateLeaveStatusRequest{
		ApproverID: uint(req.ApproverId),
		Status:     req.Status.String(),
		Comments:   "", // Default empty string
	}

	if req.Comments != nil {
		updateReq.Comments = *req.Comments
	}

	err := h.HrmsUsecase.UpdateLeaveStatus(ctx, uint(req.LeaveId), updateReq)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// ListLeaves fetches leave records with optional filters
func (h *HrmsHandler) ListLeaves(ctx context.Context, req *proto.ListLeavesRequest) (*proto.ListLeavesResponse, error) {
	var employeeID *uint
	if req.EmployeeId != nil {
		id := uint(req.GetEmployeeId())
		employeeID = &id
	}

	var status *string
	if req.Status != nil {
		st := req.Status.String()
		status = &st
	}

	leaves, err := h.HrmsUsecase.ListLeaves(ctx, employeeID, status, time.Unix(req.StartDate.Seconds, 0), time.Unix(req.EndDate.Seconds, 0), int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	response := &proto.ListLeavesResponse{}
	for _, leave := range leaves {

		// Ensure optional fields are properly handled
		var approverID *uint64
		if leave.ApproverID != 0 {
			tempApproverID := uint64(leave.ApproverID)
			approverID = &tempApproverID
		}
		var comments *string
		if leave.Comments != "" {
			comments = &leave.Comments
		}

		response.Leaves = append(response.Leaves, &proto.Leave{
			Id:         uint64(leave.ID),
			EmployeeId: uint64(leave.EmployeeID),
			LeaveType:  proto.LeaveType(proto.LeaveType_value[leave.LeaveType]),
			StartDate:  timestamppb.New(leave.StartDate),
			EndDate:    timestamppb.New(leave.EndDate),
			Status:     proto.LeaveStatus(proto.LeaveStatus_value[leave.Status]),
			ApproverId: approverID,
			Comments:   comments,
		})
	}

	return response, nil
}

// DeleteLeave handles soft deletion of a leave request
func (h *HrmsHandler) DeleteLeave(ctx context.Context, req *proto.DeleteLeaveRequest) (*emptypb.Empty, error) {
	err := h.HrmsUsecase.DeleteLeave(ctx, uint(req.LeaveId))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
