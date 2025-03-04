package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"go.uber.org/zap"

	proto "hrms/internal/transport/grpc/proto"
	"hrms/internal/domain"
)


// CreateExpense handles gRPC request to create an expense
func (h *HrmsHandler) CreateExpense(ctx context.Context, req *proto.CreateExpenseRequest) (*proto.ExpenseResponse, error) {
	h.Logger.Info("Processing CreateExpense request", zap.Uint64("employee_id", req.EmployeeId))

	expense := &domain.Expense{
		EmployeeID:  uint(req.EmployeeId),
		ExpenseType: req.ExpenseType,
		Amount:      req.Amount,
	}

	savedExpense, err := h.HrmsUsecase.CreateExpense(ctx, expense)
	if err != nil {
		return nil, err
	}

	return &proto.ExpenseResponse{Expense: mapExpenseToProto(savedExpense)}, nil
}

// GetExpense retrieves an expense by ID
func (h *HrmsHandler) GetExpense(ctx context.Context, req *proto.GetExpenseRequest) (*proto.ExpenseResponse, error) {
	expense, err := h.HrmsUsecase.GetExpense(ctx, uint(req.ExpenseId))
	if err != nil {
		return nil, err
	}

	return &proto.ExpenseResponse{Expense: mapExpenseToProto(expense)}, nil
}

// GetEmployeeExpenses retrieves all expenses for an employee
func (h *HrmsHandler) GetEmployeeExpenses(ctx context.Context, req *proto.GetEmployeeExpensesRequest) (*proto.EmployeeExpensesResponse, error) {
	expenses, err := h.HrmsUsecase.GetEmployeeExpenses(ctx, uint(req.EmployeeId))
	if err != nil {
		return nil, err
	}

	var protoExpenses []*proto.Expense
	for _, exp := range expenses {
		protoExpenses = append(protoExpenses, mapExpenseToProto(&exp))
	}

	return &proto.EmployeeExpensesResponse{Expenses: protoExpenses}, nil
}

// UpdateExpenseStatus updates the status of an expense
func (h *HrmsHandler) UpdateExpenseStatus(ctx context.Context, req *proto.UpdateExpenseStatusRequest) (*emptypb.Empty, error) {
	err := h.HrmsUsecase.UpdateExpenseStatus(ctx, uint(req.ExpenseId), uint(req.ApproverId), req.NewStatus)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeleteExpense deletes an expense if it's pending
func (h *HrmsHandler) DeleteExpense(ctx context.Context, req *proto.DeleteExpenseRequest) (*emptypb.Empty, error) {
	err := h.HrmsUsecase.DeleteExpense(ctx, uint(req.ExpenseId), uint(req.EmployeeId))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// mapExpenseToProto converts a domain Expense to a proto Expense
func mapExpenseToProto(expense *domain.Expense) *proto.Expense {
	var approverID *uint64
	if expense.ApproverID != 0 { // Ensure only non-zero values are assigned
		approverID = new(uint64)
		*approverID = uint64(expense.ApproverID)
	}

	return &proto.Expense{
		Id:          uint64(expense.ID),
		EmployeeId:  uint64(expense.EmployeeID),
		ExpenseType: expense.ExpenseType,
		Amount:      expense.Amount,
		Status:      expense.Status,
		ApproverId:  approverID, // Assigning *uint64
	}
}
