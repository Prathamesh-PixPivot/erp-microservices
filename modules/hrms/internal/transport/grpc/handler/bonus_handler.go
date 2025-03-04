package grpc

import (
	"context"
	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateBonus handles creating a new bonus record
func (h *HrmsHandler) CreateBonus(ctx context.Context, req *proto.CreateBonusRequest) (*proto.BonusResponse, error) {
	bonusReq := dto.CreateBonusRequest{
		EmployeeID:  uint(req.EmployeeId),
		Amount:      req.Amount,
		BonusType:   req.BonusType,
		Description: req.Description,
		ApprovedBy:  uint(req.ApprovedBy),
		IssueDate:   req.IssueDate.AsTime(),
		Status:      req.Status,
	}

	if req.ApprovalDate != nil {
		approvalDate := req.ApprovalDate.AsTime()
		bonusReq.ApprovalDate = &approvalDate
	}

	bonus, err := h.HrmsUsecase.CreateBonus(ctx, bonusReq)
	if err != nil {
		return nil, err
	}

	return &proto.BonusResponse{Bonus: mapToProtoBonus(bonus)}, nil
}

// GetBonusByID fetches a bonus record by ID
func (h *HrmsHandler) GetBonusByID(ctx context.Context, req *proto.GetBonusRequest) (*proto.BonusResponse, error) {
	bonus, err := h.HrmsUsecase.GetBonusByID(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &proto.BonusResponse{Bonus: mapToProtoBonus(bonus)}, nil
}

// ListBonuses fetches multiple bonus records with optional filters
func (h *HrmsHandler) ListBonuses(ctx context.Context, req *proto.ListBonusesRequest) (*proto.ListBonusesResponse, error) {
	var employeeID uint
	if req.EmployeeId != nil {
		employeeID = uint(*req.EmployeeId)
	}

	var status string
	if req.Status != nil {
		status = *req.Status
	}

	bonuses, err := h.HrmsUsecase.ListBonuses(ctx, employeeID, status)
	if err != nil {
		return nil, err
	}

	protoBonuses := make([]*proto.Bonus, len(bonuses))
	for i, bonus := range bonuses {
		protoBonuses[i] = mapToProtoBonus(&bonus)
	}

	return &proto.ListBonusesResponse{Bonuses: protoBonuses}, nil
}


// UpdateBonus updates a bonus record
func (h *HrmsHandler) UpdateBonus(ctx context.Context, req *proto.UpdateBonusRequest) (*proto.BonusResponse, error) {
	updateReq := dto.UpdateBonusRequest{}

	// Handle optional fields
	if req.Status != nil {
		updateReq.Status = *req.Status
	}
	if req.Description != nil {
		updateReq.Description = *req.Description
	}
	if req.ApprovalDate != nil {
		approvalDate := req.ApprovalDate.AsTime()
		updateReq.ApprovalDate = &approvalDate
	}

	err := h.HrmsUsecase.UpdateBonus(ctx, uint(req.Id), updateReq)
	if err != nil {
		return nil, err
	}

	bonus, err := h.HrmsUsecase.GetBonusByID(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &proto.BonusResponse{Bonus: mapToProtoBonus(bonus)}, nil
}

// DeleteBonus removes a bonus record
func (h *HrmsHandler) DeleteBonus(ctx context.Context, req *proto.DeleteBonusRequest) (*proto.BonusResponse, error) {
	err := h.HrmsUsecase.DeleteBonus(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &proto.BonusResponse{}, nil
}

// mapToProtoBonus converts domain model to gRPC response
func mapToProtoBonus(b *dto.BonusDTO) *proto.Bonus {
	return &proto.Bonus{
		Id:          uint64(b.ID),
		EmployeeId:  uint64(b.EmployeeID),
		Amount:      b.Amount,
		BonusType:   b.BonusType,
		Description: b.Description,
		ApprovedBy:  uint64(b.ApprovedBy),
		ApprovalDate: toProtoTimestamp(b.ApprovalDate),
		IssueDate:   timestamppb.New(b.IssueDate),
		Status:      b.Status,
		CreatedAt:   timestamppb.New(b.CreatedAt),
		UpdatedAt:   timestamppb.New(b.UpdatedAt),
	}
}
