package grpc

import (
	"context"
	"fmt"
	"log"
	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateAttendance handles employee check-in
func (h *HrmsHandler) CreateAttendance(ctx context.Context, req *proto.CreateAttendanceRequest) (*proto.AttendanceResponse, error) {
	attendanceReq := dto.CreateAttendanceRequest{
		EmployeeID:  uint(req.EmployeeId),
		Date:        req.Date.AsTime(),
		CheckIn:     req.CheckIn.AsTime(),
		Location:    req.Location,
		IsRemote:    req.IsRemote,
		PunchMethod: req.PunchMethod,
	}

	attendance, err := h.HrmsUsecase.CreateAttendance(ctx, attendanceReq)
	if err != nil {
		log.Printf("Error in CreateAttendance: %v", err)
		return nil, fmt.Errorf("failed to create attendance: %w", err)
	}

	return &proto.AttendanceResponse{Attendance: mapToProtoAttendance(attendance)}, nil
}

// CheckOutAttendance handles employee check-out
func (h *HrmsHandler) CheckOutAttendance(ctx context.Context, req *proto.CheckOutAttendanceRequest) (*proto.CheckOutResponse, error) {
	checkOutReq := dto.CheckOutAttendanceRequest{
		EmployeeID: uint(req.EmployeeId),
		CheckOut:   req.CheckOut.AsTime(),
	}

	err := h.HrmsUsecase.CheckOutAttendance(ctx, checkOutReq)
	if err != nil {
		log.Printf("Error in CheckOutAttendance: %v", err)
		return nil, fmt.Errorf("failed to check out: %w", err)
	}

	return &proto.CheckOutResponse{Message: "Check-out successful"}, nil
}

// GetAttendanceByID fetches an attendance record by ID
func (h *HrmsHandler) GetAttendanceByID(ctx context.Context, req *proto.GetAttendanceByIDRequest) (*proto.AttendanceResponse, error) {
	attendance, err := h.HrmsUsecase.GetAttendanceByID(ctx, uint(req.Id))
	if err != nil {
		log.Printf("Error in GetAttendanceByID: %v", err)
		return nil, fmt.Errorf("failed to get attendance record: %w", err)
	}

	return &proto.AttendanceResponse{Attendance: mapToProtoAttendance(attendance)}, nil
}

// ListAttendances fetches multiple attendance records with optional filters
func (h *HrmsHandler) ListAttendances(ctx context.Context, req *proto.ListAttendancesRequest) (*proto.ListAttendancesResponse, error) {
	attendances, total, err := h.HrmsUsecase.ListAttendances(ctx, uint(req.EmployeeId), req.StartDate.AsTime(), req.EndDate.AsTime(), req.IsRemote, int(req.Limit), int(req.Offset))
	if err != nil {
		log.Printf("Error in ListAttendances: %v", err)
		return nil, fmt.Errorf("failed to list attendances: %w", err)
	}

	protoAttendances := make([]*proto.Attendance, len(attendances))
	for i, att := range attendances {
		protoAttendances[i] = mapToProtoAttendance(&att)
	}

	return &proto.ListAttendancesResponse{
		Attendances: protoAttendances,
		Total:       total,
	}, nil
}

// DeleteAttendance removes an attendance record
func (h *HrmsHandler) DeleteAttendance(ctx context.Context, req *proto.DeleteAttendanceRequest) (*proto.DeleteAttendanceResponse, error) {
	err := h.HrmsUsecase.DeleteAttendance(ctx, uint(req.Id))
	if err != nil {
		log.Printf("Error in DeleteAttendance: %v", err)
		return nil, fmt.Errorf("failed to delete attendance record: %w", err)
	}

	return &proto.DeleteAttendanceResponse{Message: "Attendance record deleted successfully"}, nil
}

// mapToProtoAttendance converts domain model to gRPC response
func mapToProtoAttendance(a *dto.AttendanceDTO) *proto.Attendance {
	return &proto.Attendance{
		Id:          uint64(a.ID),
		EmployeeId:  uint64(a.EmployeeID),
		Date:        timestamppb.New(a.Date),
		CheckIn:     toProtoTimestamp(a.CheckIn),
		CheckOut:    toProtoTimestamp(a.CheckOut),
		WorkHours:   a.WorkHours,
		Overtime:    a.Overtime,
		BreakTime:   a.BreakTime,
		Location:    a.Location,
		IsRemote:    a.IsRemote,
		PunchMethod: a.PunchMethod,
		CreatedAt:   timestamppb.New(a.CreatedAt),
		UpdatedAt:   timestamppb.New(a.UpdatedAt),
	}
}

