package grpc

import (
	"context"
	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto" // Assuming `public_holiday.proto` is compiled here
	"hrms/internal/usecase"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// PublicHolidayHandler handles gRPC requests for Public Holidays.
type PublicHolidayHandler struct {
	proto.UnimplementedPublicHolidayServiceServer
	Usecase *usecase.HrmsUsecase
}

// CreatePublicHoliday handles the gRPC request to create a new public holiday.
func (h *PublicHolidayHandler) CreatePublicHoliday(ctx context.Context, req *proto.CreatePublicHolidayRequest) (*proto.PublicHolidayResponse, error) {
	holidayDTO := dto.PublicHolidayDTO{
		OrganizationID: uint(req.OrganizationId),
		Name:           req.Name,
		Date:           req.Date.AsTime(),
	}

	createdHoliday, err := h.Usecase.CreatePublicHoliday(ctx, holidayDTO)
	if err != nil {
		return nil, err
	}

	return &proto.PublicHolidayResponse{
		Holiday: &proto.PublicHoliday{
			Id:             uint64(createdHoliday.ID),
			OrganizationId: uint64(createdHoliday.OrganizationID),
			Name:           createdHoliday.Name,
			Date:           timestamppb.New(createdHoliday.Date),
		},
	}, nil
}

// GetPublicHoliday retrieves a holiday by ID.
func (h *PublicHolidayHandler) GetPublicHoliday(ctx context.Context, req *proto.GetPublicHolidayRequest) (*proto.PublicHolidayResponse, error) {
	holiday, err := h.Usecase.GetPublicHoliday(ctx, uint(req.HolidayId))
	if err != nil {
		return nil, err
	}

	return &proto.PublicHolidayResponse{
		Holiday: &proto.PublicHoliday{
			Id:             uint64(holiday.ID),
			OrganizationId: uint64(holiday.OrganizationID),
			Name:           holiday.Name,
			Date:           timestamppb.New(holiday.Date),
		},
	}, nil
}

// UpdatePublicHoliday updates a public holiday.
func (h *PublicHolidayHandler) UpdatePublicHoliday(ctx context.Context, req *proto.UpdatePublicHolidayRequest) (*proto.UpdatePublicHolidayResponse, error) {
	updates := make(map[string]interface{})

	for key, value := range req.Updates {
		updates[key] = value
	}

	err := h.Usecase.UpdatePublicHoliday(ctx, uint(req.HolidayId), updates)
	if err != nil {
		return nil, err
	}

	return &proto.UpdatePublicHolidayResponse{Success: true}, nil
}

// DeletePublicHoliday deletes a public holiday.
func (h *PublicHolidayHandler) DeletePublicHoliday(ctx context.Context, req *proto.DeletePublicHolidayRequest) (*proto.DeletePublicHolidayResponse, error) {
	err := h.Usecase.DeletePublicHoliday(ctx, uint(req.HolidayId))
	if err != nil {
		return nil, err
	}

	return &proto.DeletePublicHolidayResponse{Success: true}, nil
}

// ListPublicHolidays retrieves all public holidays for an organization.
func (h *PublicHolidayHandler) ListPublicHolidays(ctx context.Context, req *proto.ListPublicHolidaysRequest) (*proto.ListPublicHolidaysResponse, error) {
	var year *int
	if req.Year != nil {
		yearInt := int(req.Year.Value)
		year = &yearInt
	}

	holidays, err := h.Usecase.ListPublicHolidays(ctx, uint(req.OrganizationId), year)
	if err != nil {
		return nil, err
	}

	var holidayResponses []*proto.PublicHoliday
	for _, holiday := range holidays {
		holidayResponses = append(holidayResponses, &proto.PublicHoliday{
			Id:             uint64(holiday.ID),
			OrganizationId: uint64(holiday.OrganizationID),
			Name:           holiday.Name,
			Date:           timestamppb.New(holiday.Date),
		})
	}

	return &proto.ListPublicHolidaysResponse{
		Holidays: holidayResponses,
	}, nil
}
