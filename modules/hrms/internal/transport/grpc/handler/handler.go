package grpc

import (
	"go.uber.org/zap"

	proto "hrms/internal/transport/grpc/proto"
	"hrms/internal/usecase"
)

// HrmsHandler implements all gRPC services for HRMS.
type HrmsHandler struct {
	proto.UnimplementedAttendanceServiceServer
	proto.UnimplementedBonusServiceServer
	proto.UnimplementedDepartmentServiceServer
	proto.UnimplementedDesignationServiceServer
	proto.UnimplementedEmployeeBenefitServiceServer
	proto.UnimplementedEmployeeDocumentServiceServer
	proto.UnimplementedEmployeeExitServiceServer
	proto.UnimplementedEmployeePerkServiceServer
	proto.UnimplementedEmployeeServiceServer
	proto.UnimplementedExpenseServiceServer
	proto.UnimplementedLeaveBalanceServiceServer
	proto.UnimplementedLeavePolicyServiceServer
	proto.UnimplementedLeaveServiceServer
	proto.UnimplementedLoanAdvanceServiceServer
	proto.UnimplementedOrganizationServiceServer
	proto.UnimplementedPayrollServiceServer
	proto.UnimplementedPerformanceKPIServiceServer
	proto.UnimplementedPerformanceReviewServiceServer
	proto.UnimplementedPublicHolidayServiceServer
	proto.UnimplementedSalaryStructureServiceServer
	proto.UnimplementedShiftServiceServer
	proto.UnimplementedSkillDevelopmentServiceServer
	proto.UnimplementedWorkHistoryServiceServer

	HrmsUsecase *usecase.HrmsUsecase
	Logger      *zap.Logger
}

// NewHrmsGrpcHandler initializes a new gRPC HrmsHandler
func NewHrmsGrpcHandler(usecase *usecase.HrmsUsecase, logger *zap.Logger) *HrmsHandler {
	return &HrmsHandler{
		HrmsUsecase: usecase,
		Logger:      logger,
	}
}
