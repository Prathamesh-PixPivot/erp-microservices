package usecase

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	"hrms/internal/domain"
	"hrms/internal/dto"
)

// CreatePayroll 📌 Creates a payroll entry
func (u *HrmsUsecase) CreatePayroll(ctx context.Context, req dto.CreatePayrollDTO) (*dto.PayrollResponseDTO, error) {
	payroll := &domain.Payroll{
		EmployeeID:        req.EmployeeID,
		Salary:            req.Salary,
		Tax:               req.Tax,
		Allowances:        req.Allowances,
		Deductions:        req.Deductions,
		NetSalary:         req.NetSalary,
		PaymentDate:       req.PaymentDate,
		Status:            req.Status,
		PayslipURL:        req.PayslipURL,
		BankName:          req.BankName,
		BankAccountNumber: req.BankAccountNumber,
		BranchCode:        req.BranchCode,
	}

	createdPayroll, err := u.HrmsRepo.CreatePayroll(ctx, payroll)
	if err != nil {
		u.Logger.Error("❌ Failed to create payroll", zap.Error(err))
		return nil, err
	}

	return mapPayrollToDTO(createdPayroll), nil
}

// GetPayroll 📌 Fetch a payroll record by ID
func (u *HrmsUsecase) GetPayroll(ctx context.Context, payrollID uint) (*dto.PayrollResponseDTO, error) {
	payroll, err := u.HrmsRepo.GetPayroll(ctx, payrollID)
	if err != nil {
		u.Logger.Error("❌ Failed to fetch payroll", zap.Error(err))
		return nil, err
	}

	return mapPayrollToDTO(payroll), nil
}

// ListPayrolls 📌 Fetch payrolls with optional employee & date filters
func (u *HrmsUsecase) ListPayrolls(ctx context.Context, employeeID uint, month time.Time, limit, offset int) (*dto.PaginatedPayrollResponse, error) {
	payrolls, err := u.HrmsRepo.ListPayrolls(ctx, employeeID, month)
	if err != nil {
		u.Logger.Error("❌ Failed to fetch payrolls", zap.Error(err))
		return nil, err
	}

	var payrollDTOs []dto.PayrollResponseDTO
	for _, payroll := range payrolls {
		payrollDTOs = append(payrollDTOs, *mapPayrollToDTO(&payroll))
	}

	return &dto.PaginatedPayrollResponse{
		Total:    int64(len(payrolls)),
		Limit:    limit,
		Offset:   offset,
		Payrolls: payrollDTOs,
	}, nil
}

// UpdatePayroll 📌 Updates payroll details
func (u *HrmsUsecase) UpdatePayroll(ctx context.Context, payrollID uint, req dto.UpdatePayrollDTO) error {
	updates := make(map[string]interface{})

	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.PayslipURL != nil {
		updates["payslip_url"] = *req.PayslipURL
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	err := u.HrmsRepo.UpdatePayroll(ctx, payrollID, updates)
	if err != nil {
		u.Logger.Error("❌ Failed to update payroll", zap.Error(err))
		return err
	}

	return nil
}

// DeletePayroll 📌 Deletes a payroll record (only if not processed)
func (u *HrmsUsecase) DeletePayroll(ctx context.Context, payrollID uint) error {
	err := u.HrmsRepo.DeletePayroll(ctx, payrollID)
	if err != nil {
		u.Logger.Error("❌ Failed to delete payroll", zap.Error(err))
		return err
	}

	return nil
}

// Helper function to map domain payroll to DTO
func mapPayrollToDTO(payroll *domain.Payroll) *dto.PayrollResponseDTO {
	return &dto.PayrollResponseDTO{
		ID:               payroll.ID,
		EmployeeID:       payroll.EmployeeID,
		Salary:           payroll.Salary,
		Tax:              payroll.Tax,
		Allowances:       payroll.Allowances,
		Deductions:       payroll.Deductions,
		NetSalary:        payroll.NetSalary,
		PaymentDate:      payroll.PaymentDate,
		Status:           payroll.Status,
		PayslipURL:       payroll.PayslipURL,
		BankName:         payroll.BankName,
		BankAccountNumber: payroll.BankAccountNumber,
		BranchCode:       payroll.BranchCode,
		CreatedAt:        payroll.CreatedAt,
		UpdatedAt:        payroll.UpdatedAt,
	}
}
