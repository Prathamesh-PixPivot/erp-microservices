package grpc

import (
	"context"
	"time"

	"hrms/internal/dto"
	"hrms/internal/usecase"
	proto "hrms/internal/transport/grpc/proto"
)

// WorkHistoryHandler handles gRPC requests for Work History management.
type WorkHistoryHandler struct {
	proto.UnimplementedWorkHistoryServiceServer
	Usecase *usecase.HrmsUsecase
}

// CreateWorkHistory handles the gRPC request to create a new work history entry.
func (h *WorkHistoryHandler) CreateWorkHistory(ctx context.Context, req *proto.CreateWorkHistoryRequest) (*proto.WorkHistoryResponse, error) {
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, err
	}

	var endDate *time.Time
	if req.EndDate != "" {
		parsedEndDate, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			return nil, err
		}
		endDate = &parsedEndDate
	}

	workHistoryDTO := dto.WorkHistoryDTO{
		EmployeeID:    uint(req.EmployeeId),
		Company:       req.Company,
		Designation:   req.Designation,
		StartDate:     startDate,
		EndDate:       endDate,
		ReasonForExit: req.ReasonForExit,
	}

	createdWorkHistory, err := h.Usecase.CreateWorkHistory(ctx, workHistoryDTO)
	if err != nil {
		return nil, err
	}

	return &proto.WorkHistoryResponse{
		WorkHistory: h.convertToProtoWorkHistory(createdWorkHistory),
	}, nil
}

// GetWorkHistoryByID retrieves a work history record by ID.
func (h *WorkHistoryHandler) GetWorkHistoryByID(ctx context.Context, req *proto.GetWorkHistoryByIDRequest) (*proto.WorkHistoryResponse, error) {
	workHistory, err := h.Usecase.GetWorkHistoryByID(ctx, uint(req.WorkHistoryId))
	if err != nil {
		return nil, err
	}

	return &proto.WorkHistoryResponse{
		WorkHistory: h.convertToProtoWorkHistory(workHistory),
	}, nil
}

// GetWorkHistoryByEmployee retrieves all work history records for an employee.
func (h *WorkHistoryHandler) GetWorkHistoryByEmployee(ctx context.Context, req *proto.GetWorkHistoryByEmployeeRequest) (*proto.GetWorkHistoryByEmployeeResponse, error) {
	workHistories, err := h.Usecase.GetWorkHistoryByEmployee(ctx, uint(req.EmployeeId))
	if err != nil {
		return nil, err
	}

	var workHistoryProtoList []*proto.WorkHistory
	for _, workHistory := range workHistories {
		workHistoryProtoList = append(workHistoryProtoList, h.convertToProtoWorkHistory(&workHistory))
	}

	return &proto.GetWorkHistoryByEmployeeResponse{
		WorkHistories: workHistoryProtoList,
	}, nil
}

// UpdateWorkHistory updates a work history record.
func (h *WorkHistoryHandler) UpdateWorkHistory(ctx context.Context, req *proto.UpdateWorkHistoryRequest) (*proto.UpdateWorkHistoryResponse, error) {
	updates := make(map[string]interface{})
	for key, value := range req.Updates {
		updates[key] = value
	}

	err := h.Usecase.UpdateWorkHistory(ctx, uint(req.WorkHistoryId), updates)
	if err != nil {
		return nil, err
	}

	return &proto.UpdateWorkHistoryResponse{Success: true}, nil
}

// DeleteWorkHistory deletes a work history record.
func (h *WorkHistoryHandler) DeleteWorkHistory(ctx context.Context, req *proto.DeleteWorkHistoryRequest) (*proto.DeleteWorkHistoryResponse, error) {
	err := h.Usecase.DeleteWorkHistory(ctx, uint(req.WorkHistoryId))
	if err != nil {
		return nil, err
	}

	return &proto.DeleteWorkHistoryResponse{Success: true}, nil
}

// convertToProtoWorkHistory converts a WorkHistoryDTO to its proto representation.
func (h *WorkHistoryHandler) convertToProtoWorkHistory(workHistory *dto.WorkHistoryDTO) *proto.WorkHistory {
	if workHistory == nil {
		return nil
	}

	var endDateStr string
	if workHistory.EndDate != nil {
		endDateStr = workHistory.EndDate.Format("2006-01-02")
	}

	return &proto.WorkHistory{
		Id:            uint64(workHistory.ID),
		EmployeeId:    uint64(workHistory.EmployeeID),
		Company:       workHistory.Company,
		Designation:   workHistory.Designation,
		StartDate:     workHistory.StartDate.Format("2006-01-02"),
		EndDate:       endDateStr,
		ReasonForExit: workHistory.ReasonForExit,
	}
}
