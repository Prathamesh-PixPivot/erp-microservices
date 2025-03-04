package grpc

import (
	"context"
	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

// CreateEmployee handles gRPC request to create an employee
func (h *HrmsHandler) CreateEmployee(ctx context.Context, req *proto.CreateEmployeeRequest) (*proto.EmployeeResponse, error) {
	hiredDate := time.Unix(req.HiredDate.GetSeconds(), 0)
	employeeReq := dto.CreateEmployeeRequest{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          req.Phone,
		DateOfBirth:    time.Unix(req.DateOfBirth.GetSeconds(), 0),
		EmploymentType: req.EmploymentType,
		HiredDate:      &hiredDate,
		OrganizationID: uint(req.OrganizationId),
		DepartmentID:   uint(req.DepartmentId),
		DesignationID:  uint(req.DesignationId),
		ReportsTo:      (*uint)(req.ReportsTo),
	}

	employee, err := h.HrmsUsecase.CreateEmployee(ctx, employeeReq)
	if err != nil {
		return nil, err
	}

	return &proto.EmployeeResponse{Employee: mapToProtoEmployee(employee)}, nil
}

// GetEmployee handles gRPC request to fetch an employee by ID
func (h *HrmsHandler) GetEmployee(ctx context.Context, req *proto.GetEmployeeRequest) (*proto.EmployeeResponse, error) {
	employee, err := h.HrmsUsecase.GetEmployee(ctx, uint(req.EmployeeId))
	if err != nil {
		return nil, err
	}
	return &proto.EmployeeResponse{Employee: mapToProtoEmployee(employee)}, nil
}

// UpdateEmployee handles gRPC request to update an employee
func (h *HrmsHandler) UpdateEmployee(ctx context.Context, req *proto.UpdateEmployeeRequest) (*emptypb.Empty, error) {
	updateReq := dto.UpdateEmployeeRequest{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Phone:          req.Phone,
		EmploymentType: req.EmploymentType,
		Status:         req.Status,
	}

	err := h.HrmsUsecase.UpdateEmployee(ctx, uint(req.EmployeeId), updateReq)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// DeleteEmployee handles gRPC request to delete an employee
func (h *HrmsHandler) DeleteEmployee(ctx context.Context, req *proto.DeleteEmployeeRequest) (*emptypb.Empty, error) {
	err := h.HrmsUsecase.DeleteEmployee(ctx, uint(req.EmployeeId), req.Reason)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// mapToProtoEmployee converts EmployeeDTO to gRPC Employee
func mapToProtoEmployee(emp *dto.EmployeeDTO) *proto.Employee {
	return &proto.Employee{
		Id:             uint64(emp.ID),
		FirstName:      emp.FirstName,
		LastName:       emp.LastName,
		Email:          emp.Email,
		Phone:          emp.Phone,
		DateOfBirth:    &proto.Timestamp{Seconds: emp.DateOfBirth.Unix()},
		EmploymentType: emp.EmploymentType,
		Status:         emp.Status,
		HiredDate:      &proto.Timestamp{Seconds: emp.HiredDate.Unix()},
		OrganizationId: uint64(emp.OrganizationID),
		DepartmentId:   uint64(emp.DepartmentID),
		DesignationId:  uint64(emp.DesignationID),
		ReportsTo:      uint64(*emp.ReportsTo),
	}
}
