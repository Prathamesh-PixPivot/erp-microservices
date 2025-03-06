package grpc

import (
	"context"

	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"
	"hrms/internal/usecase"

	"google.golang.org/protobuf/types/known/emptypb"
)

// SkillDevelopmentHandler handles gRPC requests for Skill Development management.
type SkillDevelopmentHandler struct {
	proto.UnimplementedSkillDevelopmentServiceServer
	Usecase *usecase.HrmsUsecase
}

// CreateSkillDevelopment handles the gRPC request to create a new skill development entry.
func (h *SkillDevelopmentHandler) CreateSkillDevelopment(ctx context.Context, req *proto.CreateSkillDevelopmentRequest) (*proto.SkillDevelopmentResponse, error) {
	skillDTO := dto.SkillDevelopmentDTO{
		ReviewID: uint(req.ReviewId),
		Skill:    req.Skill,
		Progress: req.Progress,
	}

	createdSkill, err := h.Usecase.CreateSkillDevelopment(ctx, skillDTO)
	if err != nil {
		return nil, err
	}

	return &proto.SkillDevelopmentResponse{
		SkillDev: h.convertToProtoSkillDevelopment(createdSkill),
	}, nil
}

// GetSkillDevelopment retrieves a skill development entry by ID.
func (h *SkillDevelopmentHandler) GetSkillDevelopment(ctx context.Context, req *proto.GetSkillDevelopmentRequest) (*proto.SkillDevelopmentResponse, error) {
	skill, err := h.Usecase.GetSkillDevelopmentByID(ctx, uint(req.SkillDevId))
	if err != nil {
		return nil, err
	}

	return &proto.SkillDevelopmentResponse{
		SkillDev: h.convertToProtoSkillDevelopment(skill),
	}, nil
}

// ListSkillDevelopments retrieves all skill development records for a given review.
func (h *SkillDevelopmentHandler) ListSkillDevelopments(ctx context.Context, req *proto.ListSkillDevelopmentsRequest) (*proto.ListSkillDevelopmentsResponse, error) {
	skills, total, err := h.Usecase.ListSkillDevelopments(ctx, uint(req.ReviewId), int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	var skillProtoList []*proto.SkillDevelopment
	for _, skill := range skills {
		skillProtoList = append(skillProtoList, h.convertToProtoSkillDevelopment(&skill))
	}

	return &proto.ListSkillDevelopmentsResponse{
		Total:     int32(total),
		Limit:     req.Limit,
		Offset:    req.Offset,
		SkillDevs: skillProtoList,
	}, nil
}

// UpdateSkillDevelopment updates a skill development entry.
func (h *SkillDevelopmentHandler) UpdateSkillDevelopment(ctx context.Context, req *proto.UpdateSkillDevelopmentRequest) (*emptypb.Empty, error) {
	updates := make(map[string]interface{})

	if req.Skill != nil {
		updates["skill"] = *req.Skill
	}
	if req.Progress != nil {
		updates["progress"] = *req.Progress
	}

	err := h.Usecase.UpdateSkillDevelopment(ctx, uint(req.SkillDevId), updates)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeleteSkillDevelopment deletes a skill development entry.
func (h *SkillDevelopmentHandler) DeleteSkillDevelopment(ctx context.Context, req *proto.DeleteSkillDevelopmentRequest) (*emptypb.Empty, error) {
	err := h.Usecase.DeleteSkillDevelopment(ctx, uint(req.SkillDevId))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// convertToProtoSkillDevelopment converts a SkillDevelopmentDTO to its proto representation.
func (h *SkillDevelopmentHandler) convertToProtoSkillDevelopment(skill *dto.SkillDevelopmentDTO) *proto.SkillDevelopment {
	if skill == nil {
		return nil
	}
	return &proto.SkillDevelopment{
		Id:       uint64(skill.ID),
		ReviewId: uint64(skill.ReviewID),
		Skill:    skill.Skill,
		Progress: skill.Progress,
	}
}
