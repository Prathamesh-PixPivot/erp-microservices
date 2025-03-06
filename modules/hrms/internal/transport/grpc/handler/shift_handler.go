package grpc

import (
	"context"
	"strings"

	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"
	"hrms/internal/usecase"
)

// ShiftHandler handles gRPC requests for Shift management.
type ShiftHandler struct {
	proto.UnimplementedShiftServiceServer
	Usecase *usecase.HrmsUsecase
}

// CreateShift handles the gRPC request to create a new shift.
func (h *ShiftHandler) CreateShift(ctx context.Context, req *proto.CreateShiftRequest) (*proto.ShiftResponse, error) {
	shiftDTO := dto.ShiftDTO{
		Name:       req.Name,
		ShiftType:  req.ShiftType,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		DaysOfWeek: strings.Join(req.DaysOfWeek, ","),
	}

	createdShift, err := h.Usecase.CreateShift(ctx, shiftDTO)
	if err != nil {
		return nil, err
	}

	return &proto.ShiftResponse{
		Shift: h.convertToProtoShift(createdShift),
	}, nil
}

// GetShiftByID retrieves a shift by ID.
func (h *ShiftHandler) GetShiftByID(ctx context.Context, req *proto.GetShiftByIDRequest) (*proto.ShiftResponse, error) {
	shift, err := h.Usecase.GetShiftByID(ctx, uint(req.ShiftId))
	if err != nil {
		return nil, err
	}

	return &proto.ShiftResponse{
		Shift: h.convertToProtoShift(shift),
	}, nil
}

// ListShifts retrieves all shifts with optional filters.
func (h *ShiftHandler) ListShifts(ctx context.Context, req *proto.ListShiftsRequest) (*proto.ListShiftsResponse, error) {
	var limit, offset int
	var search string

	if req.Limit != nil {
		limit = int(req.Limit.Value)
	}
	if req.Offset != nil {
		offset = int(req.Offset.Value)
	}
	if req.Search != nil {
		search = req.Search.Value
	}

	shifts, total, err := h.Usecase.ListShifts(ctx, limit, offset, search)
	if err != nil {
		return nil, err
	}

	var shiftProtoList []*proto.Shift
	for _, shift := range shifts {
		shiftProtoList = append(shiftProtoList, h.convertToProtoShift(&shift))
	}

	return &proto.ListShiftsResponse{
		Shifts:     shiftProtoList,
		TotalCount: total,
	}, nil
}

// UpdateShift updates a shift.
func (h *ShiftHandler) UpdateShift(ctx context.Context, req *proto.UpdateShiftRequest) (*proto.UpdateShiftResponse, error) {
	updates := make(map[string]interface{})
	for key, value := range req.Updates {
		updates[key] = value
	}

	err := h.Usecase.UpdateShift(ctx, uint(req.ShiftId), updates)
	if err != nil {
		return nil, err
	}

	return &proto.UpdateShiftResponse{Success: true}, nil
}

// DeleteShift deletes a shift.
func (h *ShiftHandler) DeleteShift(ctx context.Context, req *proto.DeleteShiftRequest) (*proto.DeleteShiftResponse, error) {
	err := h.Usecase.DeleteShift(ctx, uint(req.ShiftId))
	if err != nil {
		return nil, err
	}

	return &proto.DeleteShiftResponse{Success: true}, nil
}

// convertToProtoShift converts a ShiftDTO to its proto representation.
func (h *ShiftHandler) convertToProtoShift(shift *dto.ShiftDTO) *proto.Shift {
	if shift == nil {
		return nil
	}
	return &proto.Shift{
		Id:         uint64(shift.ID),
		Name:       shift.Name,
		ShiftType:  shift.ShiftType,
		StartTime:  shift.StartTime,
		EndTime:    shift.EndTime,
		DaysOfWeek: []string{shift.DaysOfWeek}, // Ensuring correct conversion
	}
}
