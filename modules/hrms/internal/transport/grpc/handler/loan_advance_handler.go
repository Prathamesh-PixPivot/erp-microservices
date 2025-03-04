package grpc

import (
	"context"

	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// RequestLoanAdvance handles gRPC request for loan advance
func (h *HrmsHandler) RequestLoanAdvance(ctx context.Context, req *proto.RequestLoanAdvanceRequest) (*proto.LoanAdvanceResponse, error) {
	loanDTO := dto.LoanAdvanceRequestDTO{
		EmployeeID:      uint(req.EmployeeId),
		Amount:          req.Amount,
		Purpose:         req.Purpose,
		RepaymentMonths: int(req.RepaymentMonths),
	}
	loan, err := h.HrmsUsecase.RequestLoanAdvance(ctx, loanDTO)
	if err != nil {
		return nil, err
	}

	createdAt := timestamppb.New(loan.CreatedAt)

	return &proto.LoanAdvanceResponse{
		Loan: &proto.LoanAdvance{
			Id:              uint64(loan.ID),
			EmployeeId:      uint64(loan.EmployeeID),
			Amount:          loan.Amount,
			Purpose:         loan.Purpose,
			Status:          proto.LoanStatus(proto.LoanStatus_value[loan.Status]),
			RepaymentMonths: int32(loan.RepaymentMonths),
			CreatedAt:       createdAt,
		},
	}, nil
}

// ApproveLoanAdvance handles gRPC approval request
func (h *HrmsHandler) ApproveLoanAdvance(ctx context.Context, req *proto.ApproveLoanAdvanceRequest) (*emptypb.Empty, error) {
	approvalDate := timestamppb.New(req.ApprovalDate.AsTime())
	repaymentStart := timestamppb.New(req.RepaymentStart.AsTime())
// 
	approveDTO := dto.ApproveLoanAdvanceDTO{
		LoanID:         uint(req.LoanId),
		ApproverID:     uint(req.ApproverId),
		ApprovalDate:   approvalDate.AsTime(),
		RepaymentStart: repaymentStart.AsTime(),
	}

	if err := h.HrmsUsecase.ApproveLoanAdvance(ctx, approveDTO); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// RejectLoanAdvance handles gRPC rejection request
func (h *HrmsHandler) RejectLoanAdvance(ctx context.Context, req *proto.RejectLoanAdvanceRequest) (*emptypb.Empty, error) {
	rejectDTO := dto.RejectLoanAdvanceDTO{
		LoanID:     uint(req.LoanId),
		ApproverID: uint(req.ApproverId),
	}

	if err := h.HrmsUsecase.RejectLoanAdvance(ctx, rejectDTO); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// GetLoanAdvance retrieves a loan by ID
func (h *HrmsHandler) GetLoanAdvance(ctx context.Context, req *proto.GetLoanAdvanceRequest) (*proto.LoanAdvanceResponse, error) {
	loan, err := h.HrmsUsecase.GetLoanAdvance(ctx, uint(req.LoanId))
	if err != nil {
		return nil, err
	}

	createdAt := timestamppb.New(loan.CreatedAt)

	var approvalDate *timestamppb.Timestamp
	if loan.ApprovalDate != nil {
		approvalDate = timestamppb.New(*loan.ApprovalDate)
	}

	var repaymentStart *timestamppb.Timestamp
	if loan.RepaymentStart != nil {
		repaymentStart = timestamppb.New(*loan.RepaymentStart)
	}

	var approvedBy *uint64
	if loan.ApprovedBy != 0 {
		temp := uint64(loan.ApprovedBy) // Convert to uint64
		approvedBy = &temp              // Take the address to create a pointer
	}

	return &proto.LoanAdvanceResponse{
		Loan: &proto.LoanAdvance{
			Id:              uint64(loan.ID),
			EmployeeId:      uint64(loan.EmployeeID),
			Amount:          loan.Amount,
			Purpose:         loan.Purpose,
			Status:          proto.LoanStatus(proto.LoanStatus_value[loan.Status]),
			ApprovedBy:      approvedBy,
			ApprovalDate:    approvalDate,
			RepaymentStart:  repaymentStart,
			RepaymentMonths: int32(loan.RepaymentMonths),
			CreatedAt:       createdAt,
		},
	}, nil
}

// ListLoanAdvances lists loan requests
func (h *HrmsHandler) ListLoanAdvances(ctx context.Context, req *proto.ListLoanAdvancesRequest) (*proto.ListLoanAdvancesResponse, error) {
	var employeeID *uint
	if req.EmployeeId != nil {
		empID := uint(*req.EmployeeId)
		employeeID = &empID
	}

	loans, err := h.HrmsUsecase.ListLoanAdvances(ctx, req.Status.String(), employeeID)
	if err != nil {
		return nil, err
	}

	var loanList []*proto.LoanAdvance
	for _, loan := range loans {
		createdAt := timestamppb.New(loan.CreatedAt)
		loanList = append(loanList, &proto.LoanAdvance{
			Id:              uint64(loan.ID),
			EmployeeId:      uint64(loan.EmployeeID),
			Amount:          loan.Amount,
			Purpose:         loan.Purpose,
			Status:          proto.LoanStatus(proto.LoanStatus_value[loan.Status]),
			RepaymentMonths: int32(loan.RepaymentMonths),
			CreatedAt:       createdAt,
		})
	}

	return &proto.ListLoanAdvancesResponse{Loans: loanList}, nil
}

// DeleteLoanAdvance handles loan deletion
func (h *HrmsHandler) DeleteLoanAdvance(ctx context.Context, req *proto.DeleteLoanAdvanceRequest) (*emptypb.Empty, error) {
	if err := h.HrmsUsecase.DeleteLoanAdvance(ctx, uint(req.LoanId)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
