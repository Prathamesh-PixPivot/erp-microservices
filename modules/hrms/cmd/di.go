//go:build wireinject
// +build wireinject

package main

import (
	"hrms/internal/repository"
	"hrms/internal/transport/grpc/handler"
	"hrms/internal/usecase"

	"github.com/google/wire"
)

// ProvideRepositories sets up the repository dependencies.
func ProvideRepositories() *repository.HrmsRepository {
	return repository.NewHrmsRepository()
}

// ProvideUsecases sets up the use case dependencies.
func ProvideUsecases(repo *repository.HrmsRepository) *usecase.HrmsUsecase {
	return usecase.NewHrmsUsecase(repo)
}

// ProvideHandlers initializes all gRPC handlers with dependencies.
func ProvideHandlers(usecase *usecase.HrmsUsecase) *Handlers {
	return &Handlers{
		AttendanceHandler:        &handler.AttendanceHandler{Usecase: usecase},
		BonusHandler:             &handler.BonusHandler{Usecase: usecase},
		DepartmentHandler:        &handler.DepartmentHandler{Usecase: usecase},
		DesignationHandler:       &handler.DesignationHandler{Usecase: usecase},
		EmpBenefitsHandler:       &handler.EmpBenefitsHandler{Usecase: usecase},
		EmpDocHandler:            &handler.EmpDocHandler{Usecase: usecase},
		EmpExitHandler:           &handler.EmpExitHandler{Usecase: usecase},
		EmpPerkHandler:           &handler.EmpPerkHandler{Usecase: usecase},
		EmployeeHandler:          &handler.EmployeeHandler{Usecase: usecase},
		ExpenseHandler:           &handler.ExpenseHandler{Usecase: usecase},
		LeaveBalanceHandler:      &handler.LeaveBalanceHandler{Usecase: usecase},
		LeaveHandler:             &handler.LeaveHandler{Usecase: usecase},
		LeavePolicyHandler:       &handler.LeavePolicyHandler{Usecase: usecase},
		LoanAdvanceHandler:       &handler.LoanAdvanceHandler{Usecase: usecase},
		OrganizationHandler:      &handler.OrganizationHandler{Usecase: usecase},
		PayrollHandler:           &handler.PayrollHandler{Usecase: usecase},
		PerformanceKPIHandler:    &handler.PerformanceKPIHandler{Usecase: usecase},
		PerformanceReviewHandler: &handler.PerformanceReviewHandler{Usecase: usecase},
		PublicHolidayHandler:     &handler.PublicHolidayHandler{Usecase: usecase},
		SalaryStructureHandler:   &handler.SalaryStructureHandler{Usecase: usecase},
		ShiftHandler:             &handler.ShiftHandler{Usecase: usecase},
		SkillDevelopmentHandler:  &handler.SkillDevelopmentHandler{Usecase: usecase},
		WorkHistoryHandler:       &handler.WorkHistoryHandler{Usecase: usecase},
	}
}

// Handlers struct aggregates all gRPC handlers.
type Handlers struct {
	AttendanceHandler        *handler.AttendanceHandler
	BonusHandler             *handler.BonusHandler
	DepartmentHandler        *handler.DepartmentHandler
	DesignationHandler       *handler.DesignationHandler
	EmpBenefitsHandler       *handler.EmpBenefitsHandler
	EmpDocHandler            *handler.EmpDocHandler
	EmpExitHandler           *handler.EmpExitHandler
	EmpPerkHandler           *handler.EmpPerkHandler
	EmployeeHandler          *handler.EmployeeHandler
	ExpenseHandler           *handler.ExpenseHandler
	LeaveBalanceHandler      *handler.LeaveBalanceHandler
	LeaveHandler             *handler.LeaveHandler
	LeavePolicyHandler       *handler.LeavePolicyHandler
	LoanAdvanceHandler       *handler.LoanAdvanceHandler
	OrganizationHandler      *handler.OrganizationHandler
	PayrollHandler           *handler.PayrollHandler
	PerformanceKPIHandler    *handler.PerformanceKPIHandler
	PerformanceReviewHandler *handler.PerformanceReviewHandler
	PublicHolidayHandler     *handler.PublicHolidayHandler
	SalaryStructureHandler   *handler.SalaryStructureHandler
	ShiftHandler             *handler.ShiftHandler
	SkillDevelopmentHandler  *handler.SkillDevelopmentHandler
	WorkHistoryHandler       *handler.WorkHistoryHandler
}

// InitializeHandlers wires everything together using Google Wire.
func InitializeDependencies() *Handlers {
	wire.Build(ProvideRepositories, ProvideUsecases, ProvideHandlers)
	return nil
}
