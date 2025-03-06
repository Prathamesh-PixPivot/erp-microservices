package grpc

import (
	"context"

	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"
	"hrms/internal/usecase"
)

// SalaryStructureHandler handles gRPC requests for Salary Structures.
type SalaryStructureHandler struct {
	proto.UnimplementedSalaryStructureServiceServer
	Usecase *usecase.HrmsUsecase
}

// CreateSalaryStructure handles the gRPC request to create a new salary structure.
func (h *SalaryStructureHandler) CreateSalaryStructure(ctx context.Context, req *proto.CreateSalaryStructureRequest) (*proto.SalaryStructureResponse, error) {
	salaryDTO := dto.SalaryStructureDTO{
		OrganizationID: uint(req.OrganizationId),
		DesignationID:  uint(req.DesignationId),
		BaseSalary:     req.BaseSalary,
		Allowances:     req.Allowances,
		TaxPercentage:  req.TaxPercentage,
		Deductions:     req.Deductions,
	}

	createdSalary, err := h.Usecase.CreateSalaryStructure(ctx, salaryDTO)
	if err != nil {
		return nil, err
	}

	return &proto.SalaryStructureResponse{
		SalaryStructure: h.convertToProtoSalary(createdSalary),
	}, nil
}

// GetSalaryStructure retrieves a salary structure by ID.
func (h *SalaryStructureHandler) GetSalaryStructure(ctx context.Context, req *proto.GetSalaryStructureRequest) (*proto.SalaryStructureResponse, error) {
	salary, err := h.Usecase.GetSalaryStructure(ctx, uint(req.SalaryId))
	if err != nil {
		return nil, err
	}

	return &proto.SalaryStructureResponse{
		SalaryStructure: h.convertToProtoSalary(salary),
	}, nil
}

// ListSalaryStructures retrieves all salary structures with optional filters.
func (h *SalaryStructureHandler) ListSalaryStructures(ctx context.Context, req *proto.ListSalaryStructuresRequest) (*proto.ListSalaryStructuresResponse, error) {
	var organizationID, designationID uint

	if req.OrganizationId != nil {
		organizationID = uint(req.OrganizationId.Value)
	}
	if req.DesignationId != nil {
		designationID = uint(req.DesignationId.Value)
	}

	salaryStructures, err := h.Usecase.ListSalaryStructures(ctx, organizationID, designationID)
	if err != nil {
		return nil, err
	}

	var salaryProtoList []*proto.SalaryStructure
	for _, salary := range salaryStructures {
		salaryProtoList = append(salaryProtoList, h.convertToProtoSalary(&salary))
	}

	return &proto.ListSalaryStructuresResponse{
		SalaryStructures: salaryProtoList,
	}, nil
}

// UpdateSalaryStructure updates a salary structure.
func (h *SalaryStructureHandler) UpdateSalaryStructure(ctx context.Context, req *proto.UpdateSalaryStructureRequest) (*proto.UpdateSalaryStructureResponse, error) {
	updates := make(map[string]interface{})
	for key, value := range req.Updates {
		updates[key] = value
	}

	err := h.Usecase.UpdateSalaryStructure(ctx, uint(req.SalaryId), updates)
	if err != nil {
		return nil, err
	}

	return &proto.UpdateSalaryStructureResponse{Success: true}, nil
}

// DeleteSalaryStructure deletes a salary structure.
func (h *SalaryStructureHandler) DeleteSalaryStructure(ctx context.Context, req *proto.DeleteSalaryStructureRequest) (*proto.DeleteSalaryStructureResponse, error) {
	err := h.Usecase.DeleteSalaryStructure(ctx, uint(req.SalaryId))
	if err != nil {
		return nil, err
	}

	return &proto.DeleteSalaryStructureResponse{Success: true}, nil
}

// convertToProtoSalary converts a SalaryStructureDTO to its proto representation.
func (h *SalaryStructureHandler) convertToProtoSalary(salary *dto.SalaryStructureDTO) *proto.SalaryStructure {
	if salary == nil {
		return nil
	}
	return &proto.SalaryStructure{
		Id:             uint64(salary.ID),
		OrganizationId: uint64(salary.OrganizationID),
		DesignationId:  uint64(salary.DesignationID),
		BaseSalary:     salary.BaseSalary,
		Allowances:     salary.Allowances,
		TaxPercentage:  salary.TaxPercentage,
		Deductions:     salary.Deductions,
	}
}
