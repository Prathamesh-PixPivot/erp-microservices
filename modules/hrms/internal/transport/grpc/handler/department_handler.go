package grpc

import (
	"context"

	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"

	"go.uber.org/zap"
)

// CreateDepartment handles the creation of a new department
func (h *HrmsHandler) CreateDepartment(ctx context.Context, req *proto.CreateDepartmentRequest) (*proto.CreateDepartmentResponse, error) {
	departmentDTO, err := h.HrmsUsecase.CreateDepartment(ctx, dto.CreateDepartmentRequest{
		Name:           req.Name,
		OrganizationID: uint(req.OrganizationId),
	})
	if err != nil {
		h.Logger.Error("Failed to create department", zap.Error(err))
		return nil, err
	}

	return &proto.CreateDepartmentResponse{Department: mapToProtoDepartment(departmentDTO)}, nil
}

// GetDepartmentByID retrieves a department by ID
func (h *HrmsHandler) GetDepartmentByID(ctx context.Context, req *proto.GetDepartmentByIDRequest) (*proto.GetDepartmentByIDResponse, error) {
	departmentDTO, err := h.HrmsUsecase.GetDepartmentByID(ctx, uint(req.Id))
	if err != nil {
		h.Logger.Error("Failed to fetch department", zap.Error(err))
		return nil, err
	}

	return &proto.GetDepartmentByIDResponse{Department: mapToProtoDepartment(departmentDTO)}, nil
}

// UpdateDepartment updates department details
func (h *HrmsHandler) UpdateDepartment(ctx context.Context, req *proto.UpdateDepartmentRequest) (*proto.UpdateDepartmentResponse, error) {
	updateReq := dto.UpdateDepartmentRequest{}
	if req.Name != "" {
		updateReq.Name = req.Name
	}

	err := h.HrmsUsecase.UpdateDepartment(ctx, uint(req.Id), updateReq)
	if err != nil {
		h.Logger.Error("Failed to update department", zap.Error(err))
		return nil, err
	}

	return &proto.UpdateDepartmentResponse{Success: true}, nil
}

// DeleteDepartment deletes a department by ID
func (h *HrmsHandler) DeleteDepartment(ctx context.Context, req *proto.DeleteDepartmentRequest) (*proto.DeleteDepartmentResponse, error) {
	err := h.HrmsUsecase.DeleteDepartment(ctx, uint(req.Id))
	if err != nil {
		h.Logger.Error("Failed to delete department", zap.Error(err))
		return nil, err
	}

	return &proto.DeleteDepartmentResponse{Success: true}, nil
}

// ListDepartments retrieves a list of departments with pagination and search filters
func (h *HrmsHandler) ListDepartments(ctx context.Context, req *proto.ListDepartmentsRequest) (*proto.ListDepartmentsResponse, error) {
	departments, totalCount, err := h.HrmsUsecase.ListDepartments(ctx, uint(req.OrganizationId), int(req.Limit), int(req.Offset), req.Search)
	if err != nil {
		h.Logger.Error("Failed to list departments", zap.Error(err))
		return nil, err
	}

	protoDepartments := make([]*proto.Department, len(departments))
	for i, dept := range departments {
		protoDepartments[i] = mapToProtoDepartment(&dept)
	}

	return &proto.ListDepartmentsResponse{
		Departments: protoDepartments,
		TotalCount:  totalCount,
	}, nil
}

// mapToProtoDepartment converts a DTO to a protobuf response
func mapToProtoDepartment(d *dto.DepartmentDTO) *proto.Department {
	return &proto.Department{
		Id:             uint64(d.ID),
		Name:           d.Name,
		OrganizationId: uint64(d.OrganizationID),
	}
}
