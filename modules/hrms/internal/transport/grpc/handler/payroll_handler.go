package grpc

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	proto "hrms/internal/transport/grpc/proto"

	"hrms/internal/dto"
)

// CreatePayroll handles gRPC request to create a payroll record
func (h *HrmsHandler) CreatePayroll(ctx context.Context, req *proto.CreatePayrollRequest) (*proto.PayrollResponse, error) {
	payrollDTO := dto.CreatePayrollDTO{
		EmployeeID:        uint(req.EmployeeId),
		Salary:            req.Salary,
		Tax:               req.Tax,
		Allowances:        req.Allowances,
		Deductions:        req.Deductions,
		NetSalary:         req.NetSalary,
		PaymentDate:       req.PaymentDate.AsTime(),
		Status:            req.Status,
		PayslipURL:        req.PayslipUrl,
		BankName:          req.BankName,
		BankAccountNumber: req.BankAccountNumber,
		BranchCode:        req.BranchCode,
	}

	payroll, err := h.HrmsUsecase.CreatePayroll(ctx, payrollDTO)
	if err != nil {
		h.Logger.Error("Failed to create payroll", zap.Error(err))
		return nil, err
	}

	return &proto.PayrollResponse{Payroll: mapPayrollToProto(payroll)}, nil
}

// GetPayroll handles gRPC request to fetch a payroll record by ID
func (h *HrmsHandler) GetPayroll(ctx context.Context, req *proto.GetPayrollRequest) (*proto.PayrollResponse, error) {
	payroll, err := h.HrmsUsecase.GetPayroll(ctx, uint(req.PayrollId))
	if err != nil {
		h.Logger.Error("Failed to get payroll", zap.Error(err))
		return nil, err
	}

	return &proto.PayrollResponse{Payroll: mapPayrollToProto(payroll)}, nil
}

// ListPayrolls handles gRPC request to fetch payroll records
func (h *HrmsHandler) ListPayrolls(ctx context.Context, req *proto.ListPayrollsRequest) (*proto.ListPayrollsResponse, error) {
	payrolls, err := h.HrmsUsecase.ListPayrolls(ctx, uint(req.EmployeeId), req.Month.AsTime(), int(req.Limit), int(req.Offset))
	if err != nil {
		h.Logger.Error("Failed to list payrolls", zap.Error(err))
		return nil, err
	}

	var payrollProtos []*proto.Payroll
	for _, payroll := range payrolls.Payrolls {
		payrollProtos = append(payrollProtos, mapPayrollToProto(&payroll))
	}

	return &proto.ListPayrollsResponse{
		Total:    int32(payrolls.Total),
		Limit:    int32(payrolls.Limit),
		Offset:   int32(payrolls.Offset),
		Payrolls: payrollProtos,
	}, nil
}

// UpdatePayroll handles gRPC request to update payroll details
func (h *HrmsHandler) UpdatePayroll(ctx context.Context, req *proto.UpdatePayrollRequest) (*emptypb.Empty, error) {
	updateDTO := dto.UpdatePayrollDTO{}

	if req.Status != nil {
		updateDTO.Status = req.Status
	}
	if req.PayslipUrl != nil {
		updateDTO.PayslipURL = req.PayslipUrl
	}

	err := h.HrmsUsecase.UpdatePayroll(ctx, uint(req.PayrollId), updateDTO)
	if err != nil {
		h.Logger.Error("Failed to update payroll", zap.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeletePayroll handles gRPC request to delete a payroll record
func (h *HrmsHandler) DeletePayroll(ctx context.Context, req *proto.DeletePayrollRequest) (*emptypb.Empty, error) {
	err := h.HrmsUsecase.DeletePayroll(ctx, uint(req.PayrollId))
	if err != nil {
		h.Logger.Error("Failed to delete payroll", zap.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func mapPayrollToProto(payroll *dto.PayrollResponseDTO) *proto.Payroll {
	return &proto.Payroll{
		Id:                uint64(payroll.ID),
		EmployeeId:        uint64(payroll.EmployeeID),
		Salary:            payroll.Salary,
		Tax:               payroll.Tax,
		Allowances:        payroll.Allowances,
		Deductions:        payroll.Deductions,
		NetSalary:         payroll.NetSalary,
		PaymentDate:       timestamppb.New(payroll.PaymentDate),
		Status:            payroll.Status,
		PayslipUrl:        payroll.PayslipURL,
		BankName:          payroll.BankName,
		BankAccountNumber: payroll.BankAccountNumber,
		BranchCode:        payroll.BranchCode,
		CreatedAt:         timestamppb.New(payroll.CreatedAt),
		UpdatedAt:         timestamppb.New(payroll.UpdatedAt),
	}
}
