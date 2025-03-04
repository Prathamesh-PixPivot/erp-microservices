package grpc

import (
	"context"

	"errors"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"hrms/internal/domain"
	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"
)

// CreateLeaveBalance handles gRPC request to create a leave balance
func (h *HrmsHandler) CreateLeaveBalance(ctx context.Context, req *proto.CreateLeaveBalanceRequest) (*proto.LeaveBalanceResponse, error) {
	// Convert gRPC request to DTO request
	createRequest := dto.CreateLeaveBalanceRequest{
		EmployeeID:  uint(req.EmployeeId),
		LeaveType:   domain.LeaveType(req.LeaveType), // This remains the same
		TotalLeaves: float64(req.TotalLeaves),
	}

	// Call the use case
	balance, err := h.HrmsUsecase.CreateLeaveBalance(ctx, createRequest)
	if err != nil {
		h.Logger.Error("Failed to create leave balance", zap.Error(err))
		return nil, err
	}

	// Convert domain.LeaveType (string) to proto.LeaveType (enum)
	leaveTypeEnum, ok := proto.LeaveType_value[string(balance.LeaveType)]
	if !ok {
		h.Logger.Error("Invalid leave type mapping", zap.String("LeaveType", string(balance.LeaveType)))
		return nil, errors.New("invalid leave type mapping")
	}

	// Convert use case response to gRPC response
	return &proto.LeaveBalanceResponse{
		Balance: &proto.LeaveBalance{
			EmployeeId:  uint64(balance.EmployeeID),
			LeaveType:   proto.LeaveType(leaveTypeEnum), // Correct conversion
			TotalLeaves: int32(balance.TotalLeaves),
			UsedLeaves:  int32(balance.UsedLeaves),
			Remaining:   int32(balance.Remaining),
		},
	}, nil
}


// DeductLeaveBalance handles gRPC request to deduct leave balance
func (h *HrmsHandler) DeductLeaveBalance(ctx context.Context, req *proto.DeductLeaveBalanceRequest) (*emptypb.Empty, error) {
	// Convert gRPC request to DTO request
	deductRequest := dto.DeductLeaveBalanceRequest{
		EmployeeID: uint(req.EmployeeId),
		LeaveType:  domain.LeaveType(req.LeaveType),
		LeaveDays:  float64(req.LeaveDays),
	}

	// Call the use case with the DTO
	err := h.HrmsUsecase.DeductLeaveBalance(ctx, deductRequest)
	if err != nil {
		h.Logger.Error("Failed to deduct leave balance", zap.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// RestoreLeaveBalance handles gRPC request to restore leave balance
func (h *HrmsHandler) RestoreLeaveBalance(ctx context.Context, req *proto.RestoreLeaveBalanceRequest) (*emptypb.Empty, error) {
	// Convert gRPC request to DTO request
	restoreRequest := dto.RestoreLeaveBalanceRequest{
		EmployeeID: uint(req.EmployeeId),
		LeaveType:  domain.LeaveType(req.LeaveType),
		LeaveDays:  float64(req.LeaveDays),
	}

	// Call the use case with the DTO
	err := h.HrmsUsecase.RestoreLeaveBalance(ctx, restoreRequest)
	if err != nil {
		h.Logger.Error("Failed to restore leave balance", zap.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// GetLeaveBalance handles gRPC request to get a leave balance
func (h *HrmsHandler) GetLeaveBalance(ctx context.Context, req *proto.GetLeaveBalanceRequest) (*proto.LeaveBalanceResponse, error) {
	balance, err := h.HrmsUsecase.GetLeaveBalance(ctx, uint(req.EmployeeId), domain.LeaveType(req.LeaveType))
	if err != nil {
		h.Logger.Error("Failed to fetch leave balance", zap.Error(err))
		return nil, err
	}

	// Convert domain.LeaveType (string) to proto.LeaveType (enum)
	leaveTypeEnum, ok := proto.LeaveType_value[string(balance.LeaveType)]
	if !ok {
		h.Logger.Error("Invalid leave type mapping", zap.String("LeaveType", string(balance.LeaveType)))
		return nil, errors.New("invalid leave type mapping")
	}

	return &proto.LeaveBalanceResponse{
		Balance: &proto.LeaveBalance{
			EmployeeId:  uint64(balance.EmployeeID),
			LeaveType:   proto.LeaveType(leaveTypeEnum),
			TotalLeaves: int32(balance.TotalLeaves),
			UsedLeaves:  int32(balance.UsedLeaves),
			Remaining:   int32(balance.Remaining),
		},
	}, nil
}

// ListLeaveBalances handles gRPC request to list leave balances
func (h *HrmsHandler) ListLeaveBalances(ctx context.Context, req *proto.ListLeaveBalancesRequest) (*proto.ListLeaveBalancesResponse, error) {
	var employeeID *uint
	if req.EmployeeId != nil {
		empID := uint(*req.EmployeeId)
		employeeID = &empID
	}

	balances, total, err := h.HrmsUsecase.ListLeaveBalances(ctx, employeeID, int(req.Limit), int(req.Offset))
	if err != nil {
		h.Logger.Error("Failed to list leave balances", zap.Error(err))
		return nil, err
	}

	var protoBalances []*proto.LeaveBalance
	for _, balance := range balances {
		// Convert domain.LeaveType (string) to proto.LeaveType (enum)
		leaveTypeEnum, ok := proto.LeaveType_value[string(balance.LeaveType)]
		if !ok {
			h.Logger.Error("Invalid leave type mapping", zap.String("LeaveType", string(balance.LeaveType)))
			return nil, errors.New("invalid leave type mapping")
		}

		protoBalances = append(protoBalances, &proto.LeaveBalance{
			EmployeeId:  uint64(balance.EmployeeID),
			LeaveType:   proto.LeaveType(leaveTypeEnum),
			TotalLeaves: int32(balance.TotalLeaves),
			UsedLeaves:  int32(balance.UsedLeaves),
			Remaining:   int32(balance.Remaining),
		})
	}

	return &proto.ListLeaveBalancesResponse{
		Balances: protoBalances,
		Total:    total,
	}, nil
}

// DeleteLeaveBalance handles gRPC request to delete a leave balance
func (h *HrmsHandler) DeleteLeaveBalance(ctx context.Context, req *proto.DeleteLeaveBalanceRequest) (*emptypb.Empty, error) {
	err := h.HrmsUsecase.DeleteLeaveBalance(ctx, uint(req.EmployeeId), domain.LeaveType(req.LeaveType))
	if err != nil {
		h.Logger.Error("Failed to delete leave balance", zap.Error(err))
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
