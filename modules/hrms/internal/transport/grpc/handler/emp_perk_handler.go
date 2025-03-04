package grpc

import (
	"context"

	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"


	"google.golang.org/protobuf/types/known/emptypb"
)

// CreateEmployeePerk handles gRPC request to add a perk for an employee
func (h *HrmsHandler) CreateEmployeePerk(ctx context.Context, req *proto.CreateEmployeePerkRequest) (*proto.EmployeePerkResponse, error) {
	perkDTO, err := h.HrmsUsecase.CreateEmployeePerk(ctx, dto.CreateEmployeePerkRequest{
		EmployeeID: uint(req.EmployeeId),
		Perk:       req.Perk,
	})
	if err != nil {
		return nil, err
	}

	return &proto.EmployeePerkResponse{
		Perk: &proto.EmployeePerk{
			Id:          uint64(perkDTO.ID),
			EmployeeId: uint64(perkDTO.EmployeeID),
			Perk:       perkDTO.Perk,
		},
	}, nil
}

// GetEmployeePerks fetches all perks assigned to an employee
func (h *HrmsHandler) GetEmployeePerks(ctx context.Context, req *proto.GetEmployeePerksRequest) (*proto.ListEmployeePerksResponse, error) {
	perks, err := h.HrmsUsecase.GetEmployeePerks(ctx, uint(req.EmployeeId))
	if err != nil {
		return nil, err
	}

	var perkList []*proto.EmployeePerk
	for _, p := range perks {
		perkList = append(perkList, &proto.EmployeePerk{
			Id:          uint64(p.ID),
			EmployeeId: uint64(p.EmployeeID),
			Perk:       p.Perk,
		})
	}

	return &proto.ListEmployeePerksResponse{Perks: perkList}, nil
}

// UpdateEmployeePerk updates an existing employee perk
func (h *HrmsHandler) UpdateEmployeePerk(ctx context.Context, req *proto.UpdateEmployeePerkRequest) (*emptypb.Empty, error) {
	err := h.HrmsUsecase.UpdateEmployeePerk(ctx, uint(req.PerkId), dto.UpdateEmployeePerkRequest{
		Perk: req.Perk,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeleteEmployeePerk removes an employee's perk
func (h *HrmsHandler) DeleteEmployeePerk(ctx context.Context, req *proto.DeleteEmployeePerkRequest) (*emptypb.Empty, error) {
	if err := h.HrmsUsecase.DeleteEmployeePerk(ctx, uint(req.PerkId)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
