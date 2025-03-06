package grpc

import (
	"context"

	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateEmployeeExit handles creating an employee exit record
func (h *HrmsHandler) CreateEmployeeExit(ctx context.Context, req *proto.CreateEmployeeExitRequest) (*proto.EmployeeExitResponse, error) {
	exitReq := dto.CreateEmployeeExitRequest{
		EmployeeID:      uint(req.EmployeeId),
		ExitType:        req.ExitType,
		ExitDate:        req.ExitDate.AsTime(),
		ClearanceStatus: req.ClearanceStatus,
	}

	exitDTO, err := h.HrmsUsecase.CreateEmployeeExit(ctx, exitReq)
	if err != nil {
		return nil, err
	}

	return &proto.EmployeeExitResponse{Exit: mapToProtoExit(exitDTO)}, nil
}

// GetEmployeeExitByID handles fetching an employee exit record by ID
func (h *HrmsHandler) GetEmployeeExitByID(ctx context.Context, req *proto.GetEmployeeExitByIDRequest) (*proto.EmployeeExitResponse, error) {
	exitDTO, err := h.HrmsUsecase.GetEmployeeExitByID(ctx, uint(req.ExitId))
	if err != nil {
		return nil, err
	}

	return &proto.EmployeeExitResponse{Exit: mapToProtoExit(exitDTO)}, nil
}

// GetExitRecordsByEmployee handles fetching all exit records for an employee
func (h *HrmsHandler) GetExitRecordsByEmployee(ctx context.Context, req *proto.GetExitRecordsByEmployeeRequest) (*proto.ListEmployeeExitsResponse, error) {
	exitsDTO, err := h.HrmsUsecase.GetExitRecordsByEmployee(ctx, uint(req.EmployeeId))
	if err != nil {
		return nil, err
	}

	var exits []*proto.EmployeeExit
	for _, exit := range exitsDTO {
		exits = append(exits, mapToProtoExit(&exit))
	}

	return &proto.ListEmployeeExitsResponse{Exits: exits}, nil
}

func (h *HrmsHandler) GetPendingClearances(ctx context.Context, _ *emptypb.Empty) (*proto.ListEmployeeExitsResponse, error) {
	exitsDTO, err := h.HrmsUsecase.GetPendingClearances(ctx)
	if err != nil {
		return nil, err
	}

	var exits []*proto.EmployeeExit
	for _, exit := range exitsDTO {
		exits = append(exits, mapToProtoExit(&exit))
	}

	return &proto.ListEmployeeExitsResponse{Exits: exits}, nil
}

// UpdateClearanceStatus handles updating clearance status
func (h *HrmsHandler) UpdateClearanceStatus(ctx context.Context, req *proto.UpdateClearanceStatusRequest) (*emptypb.Empty, error) {
	updateReq := dto.UpdateClearanceStatusRequest{
		ClearanceStatus: req.ClearanceStatus,
	}

	if err := h.HrmsUsecase.UpdateClearanceStatus(ctx, uint(req.ExitId), updateReq); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeleteEmployeeExit handles deleting an employee exit record
func (h *HrmsHandler) DeleteEmployeeExit(ctx context.Context, req *proto.DeleteEmployeeExitRequest) (*emptypb.Empty, error) {
	if err := h.HrmsUsecase.DeleteEmployeeExit(ctx, uint(req.ExitId)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// mapToProtoExit maps DTO to Protobuf message
func mapToProtoExit(e *dto.EmployeeExitDTO) *proto.EmployeeExit {
	return &proto.EmployeeExit{
		Id:              uint64(e.ID),
		EmployeeId:      uint64(e.EmployeeID),
		ExitType:        e.ExitType,
		ExitDate:        timestamppb.New(e.ExitDate),
		ClearanceStatus: e.ClearanceStatus,
	}
}
