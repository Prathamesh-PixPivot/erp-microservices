package usecase

import (
	"context"

	"go.uber.org/zap"
	"hrms/internal/domain"
	"hrms/internal/dto"
)

// RequestLoanAdvance 📌 Employee submits a loan/advance request
func (u *HrmsUsecase) RequestLoanAdvance(ctx context.Context, req dto.LoanAdvanceRequestDTO) (*dto.LoanAdvanceResponseDTO, error) {
	loan := &domain.LoanAdvance{
		EmployeeID:      req.EmployeeID,
		Amount:          req.Amount,
		Purpose:         req.Purpose,
		Status:          "Pending",
		RepaymentMonths: req.RepaymentMonths,
	}

	createdLoan, err := u.HrmsRepo.CreateLoanAdvance(ctx, loan)
	if err != nil {
		u.Logger.Error("Failed to create loan advance", zap.Error(err))
		return nil, err
	}

	response := &dto.LoanAdvanceResponseDTO{
		ID:              createdLoan.ID,
		EmployeeID:      createdLoan.EmployeeID,
		Amount:          createdLoan.Amount,
		Purpose:         createdLoan.Purpose,
		Status:          createdLoan.Status,
		RepaymentMonths: createdLoan.RepaymentMonths,
		CreatedAt:       createdLoan.CreatedAt,
	}

	return response, nil
}

// ApproveLoanAdvance 📌 Approves a loan/advance request
func (u *HrmsUsecase) ApproveLoanAdvance(ctx context.Context, req dto.ApproveLoanAdvanceDTO) error {
	err := u.HrmsRepo.ApproveLoanAdvance(ctx, req.LoanID, req.ApproverID, req.ApprovalDate, req.RepaymentStart)
	if err != nil {
		u.Logger.Error("Failed to approve loan advance", zap.Error(err))
		return err
	}
	return nil
}

// RejectLoanAdvance 📌 Rejects a loan/advance request
func (u *HrmsUsecase) RejectLoanAdvance(ctx context.Context, req dto.RejectLoanAdvanceDTO) error {
	err := u.HrmsRepo.RejectLoanAdvance(ctx, req.LoanID, req.ApproverID)
	if err != nil {
		u.Logger.Error("Failed to reject loan advance", zap.Error(err))
		return err
	}
	return nil
}

// GetLoanAdvance 📌 Fetches a loan/advance request by ID
func (u *HrmsUsecase) GetLoanAdvance(ctx context.Context, loanID uint) (*dto.LoanAdvanceResponseDTO, error) {
	loan, err := u.HrmsRepo.GetLoanAdvanceByID(ctx, loanID)
	if err != nil {
		u.Logger.Error("Failed to fetch loan advance", zap.Error(err))
		return nil, err
	}

	response := &dto.LoanAdvanceResponseDTO{
		ID:              loan.ID,
		EmployeeID:      loan.EmployeeID,
		Amount:          loan.Amount,
		Purpose:         loan.Purpose,
		Status:          loan.Status,
		ApprovedBy:      loan.ApprovedBy,
		ApprovalDate:    loan.ApprovalDate,
		RepaymentStart:  loan.RepaymentStart,
		RepaymentMonths: loan.RepaymentMonths,
		CreatedAt:       loan.CreatedAt,
	}

	return response, nil
}

// ListLoanAdvances 📌 Fetches all loan/advance requests with optional filters
func (u *HrmsUsecase) ListLoanAdvances(ctx context.Context, status string, employeeID *uint) ([]dto.LoanAdvanceResponseDTO, error) {
	loans, err := u.HrmsRepo.ListLoanAdvances(ctx, status, employeeID)
	if err != nil {
		u.Logger.Error("Failed to list loan advances", zap.Error(err))
		return nil, err
	}

	var response []dto.LoanAdvanceResponseDTO
	for _, loan := range loans {
		response = append(response, dto.LoanAdvanceResponseDTO{
			ID:              loan.ID,
			EmployeeID:      loan.EmployeeID,
			Amount:          loan.Amount,
			Purpose:         loan.Purpose,
			Status:          loan.Status,
			ApprovedBy:      loan.ApprovedBy,
			ApprovalDate:    loan.ApprovalDate,
			RepaymentStart:  loan.RepaymentStart,
			RepaymentMonths: loan.RepaymentMonths,
			CreatedAt:       loan.CreatedAt,
		})
	}

	return response, nil
}

// DeleteLoanAdvance 📌 Deletes a loan/advance request
func (u *HrmsUsecase) DeleteLoanAdvance(ctx context.Context, loanID uint) error {
	err := u.HrmsRepo.DeleteLoanAdvance(ctx, loanID)
	if err != nil {
		u.Logger.Error("Failed to delete loan advance", zap.Error(err))
		return err
	}
	return nil
}
