package grpc

import (
	"context"

	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"

	"go.uber.org/zap"
)


func (h *HrmsHandler) CreateDesignation(ctx context.Context, req *proto.CreateDesignationRequest) (*proto.DesignationResponse, error) {
	designationReq := dto.CreateDesignationRequest{
		Title:          req.Title,
		Level:          req.Level,
		HierarchyLevel: int(req.HierarchyLevel),
		DepartmentID:   uint(req.DepartmentId),
	}

	designation, err := h.HrmsUsecase.CreateDesignation(ctx, designationReq)
	if err != nil {
		h.Logger.Error("Failed to create designation", zap.Error(err))
		return nil, err
	}

	return &proto.DesignationResponse{Designation: mapToProtoDesignation(designation)}, nil
}

func (h *HrmsHandler) GetDesignationByID(ctx context.Context, req *proto.GetDesignationRequest) (*proto.DesignationResponse, error) {
	designation, err := h.HrmsUsecase.GetDesignationByID(ctx, uint(req.Id))
	if err != nil {
		h.Logger.Error("Failed to fetch designation", zap.Error(err))
		return nil, err
	}

	return &proto.DesignationResponse{Designation: mapToProtoDesignation(designation)}, nil
}

func (h *HrmsHandler) UpdateDesignation(ctx context.Context, req *proto.UpdateDesignationRequest) (*proto.EmptyResponse, error) {
	updateReq := dto.UpdateDesignationRequest{}

	if req.Title != nil {
		updateReq.Title = *req.Title
	}
	if req.Level != nil {
		updateReq.Level = *req.Level
	}
	if req.HierarchyLevel != nil {
		updateReq.HierarchyLevel = int(*req.HierarchyLevel)
	}

	err := h.HrmsUsecase.UpdateDesignation(ctx, uint(req.Id), updateReq)
	if err != nil {
		h.Logger.Error("Failed to update designation", zap.Error(err))
		return nil, err
	}

	return &proto.EmptyResponse{}, nil
}

func (h *HrmsHandler) DeleteDesignation(ctx context.Context, req *proto.DeleteDesignationRequest) (*proto.EmptyResponse, error) {
	if err := h.HrmsUsecase.DeleteDesignation(ctx, uint(req.Id)); err != nil {
		h.Logger.Error("Failed to delete designation", zap.Error(err))
		return nil, err
	}
	return &proto.EmptyResponse{}, nil
}

func (h *HrmsHandler) ListDesignations(ctx context.Context, req *proto.ListDesignationsRequest) (*proto.ListDesignationsResponse, error) {
	designations, totalCount, err := h.HrmsUsecase.ListDesignations(ctx, uint(req.DepartmentId), int(req.Limit), int(req.Offset), req.Search)
	if err != nil {
		h.Logger.Error("Failed to list designations", zap.Error(err))
		return nil, err
	}

	protoDesignations := make([]*proto.Designation, len(designations))
	for i, d := range designations {
		protoDesignations[i] = mapToProtoDesignation(&d)
	}

	return &proto.ListDesignationsResponse{
		Designations: protoDesignations,
		TotalCount:   totalCount,
	}, nil
}

func mapToProtoDesignation(d *dto.DesignationDTO) *proto.Designation {
	return &proto.Designation{
		Id:             uint64(d.ID),
		Title:          d.Title,
		Level:          d.Level,
		HierarchyLevel: uint32(d.HierarchyLevel),
		DepartmentId:   uint64(d.DepartmentID),
	}
}
