package grpc

import (
	"context"

	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"

	"go.uber.org/zap"
)


func (h *HrmsHandler) CreateEmployeeBenefits(ctx context.Context, req *proto.CreateEmployeeBenefitsRequest) (*proto.EmployeeBenefitsDTO, error) {
	benefitsReq := dto.CreateEmployeeBenefitsRequest{
		EmployeeID:     uint(req.EmployeeId),
		HealthPlan:     req.HealthPlan,
		RetirementPlan: req.RetirementPlan,
	}

	benefits, err := h.HrmsUsecase.CreateEmployeeBenefits(ctx, benefitsReq)
	if err != nil {
		h.Logger.Error("Failed to create employee benefits", zap.Error(err))
		return nil, err
	}

	return mapToProtoEmployeeBenefits(benefits), nil
}

func (h *HrmsHandler) GetEmployeeBenefits(ctx context.Context, req *proto.GetEmployeeBenefitsRequest) (*proto.EmployeeBenefitsDTO, error) {
	benefits, err := h.HrmsUsecase.GetEmployeeBenefits(ctx, uint(req.EmployeeId))
	if err != nil {
		h.Logger.Error("Failed to fetch employee benefits", zap.Error(err))
		return nil, err
	}

	return mapToProtoEmployeeBenefits(benefits), nil
}

func (h *HrmsHandler) UpdateEmployeeBenefits(ctx context.Context, req *proto.UpdateEmployeeBenefitsRequest) (*proto.EmptyResponse, error) {
	updateReq := dto.UpdateEmployeeBenefitsRequest{
		HealthPlan:     req.HealthPlan,
		RetirementPlan: req.RetirementPlan,
	}

	err := h.HrmsUsecase.UpdateEmployeeBenefits(ctx, uint(req.EmployeeId), updateReq)
	if err != nil {
		h.Logger.Error("Failed to update employee benefits", zap.Error(err))
		return nil, err
	}

	return &proto.EmptyResponse{}, nil
}

func (h *HrmsHandler) DeleteEmployeeBenefits(ctx context.Context, req *proto.DeleteEmployeeBenefitsRequest) (*proto.EmptyResponse, error) {
	if err := h.HrmsUsecase.DeleteEmployeeBenefits(ctx, uint(req.EmployeeId)); err != nil {
		h.Logger.Error("Failed to delete employee benefits", zap.Error(err))
		return nil, err
	}

	return &proto.EmptyResponse{}, nil
}

func (h *HrmsHandler) ListEmployeeBenefits(ctx context.Context, req *proto.ListEmployeeBenefitsRequest) (*proto.ListEmployeeBenefitsResponse, error) {
	benefits, err := h.HrmsUsecase.ListEmployeeBenefits(ctx, req.HealthPlan, req.RetirementPlan)
	if err != nil {
		h.Logger.Error("Failed to list employee benefits", zap.Error(err))
		return nil, err
	}

	protoBenefits := make([]*proto.EmployeeBenefitsDTO, len(benefits))
	for i, b := range benefits {
		protoBenefits[i] = mapToProtoEmployeeBenefits(&b)
	}

	return &proto.ListEmployeeBenefitsResponse{Benefits: protoBenefits}, nil
}

func mapToProtoEmployeeBenefits(b *dto.EmployeeBenefitsDTO) *proto.EmployeeBenefitsDTO {
	return &proto.EmployeeBenefitsDTO{
		Id:             uint64(b.ID),
		EmployeeId:     uint64(b.EmployeeID),
		HealthPlan:     b.HealthPlan,
		RetirementPlan: b.RetirementPlan,
	}
}
