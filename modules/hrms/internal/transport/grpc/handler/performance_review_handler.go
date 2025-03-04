package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/emptypb"

	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"

)

func (h *HrmsHandler) CreatePerformanceReview(ctx context.Context, req *proto.CreatePerformanceReviewRequest) (*proto.PerformanceReviewResponse, error) {
	reviewDTO := dto.PerformanceReviewDTO{
		EmployeeID:   uint(req.EmployeeId),
		ReviewerID:   uint(req.ReviewerId),
		ReviewDate:   req.ReviewDate.AsTime(),
		ReviewPeriod: req.ReviewPeriod,
		OverallRating: int(req.OverallRating),
		Feedback:     req.Feedback,
		Promotion:    req.Promotion,
	}

	createdReview, err := h.HrmsUsecase.CreatePerformanceReview(ctx, reviewDTO)
	if err != nil {
		return nil, err
	}

	return &proto.PerformanceReviewResponse{
		Review: mapPerformanceReviewToProto(createdReview),
	}, nil
}

func (h *HrmsHandler) GetPerformanceReviewByID(ctx context.Context, req *proto.GetPerformanceReviewRequest) (*proto.PerformanceReviewResponse, error) {
	review, err := h.HrmsUsecase.GetPerformanceReviewByID(ctx, uint(req.ReviewId))
	if err != nil {
		return nil, err
	}

	return &proto.PerformanceReviewResponse{
		Review: mapPerformanceReviewToProto(review),
	}, nil
}

func (h *HrmsHandler) ListPerformanceReviews(ctx context.Context, req *proto.ListPerformanceReviewsRequest) (*proto.ListPerformanceReviewsResponse, error) {
	reviews, totalCount, err := h.HrmsUsecase.ListPerformanceReviews(ctx, uint(req.EmployeeId), int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	var protoReviews []*proto.PerformanceReview
	for _, review := range reviews {
		protoReviews = append(protoReviews, mapPerformanceReviewToProto(&review))
	}

	return &proto.ListPerformanceReviewsResponse{
		Reviews:    protoReviews,
		TotalCount: totalCount,
	}, nil
}

func (h *HrmsHandler) UpdatePerformanceReview(ctx context.Context, req *proto.UpdatePerformanceReviewRequest) (*emptypb.Empty, error) {
	updates := make(map[string]interface{})
	for key, val := range req.Updates {
		updates[key] = val.AsInterface() // Convert *structpb.Value to native Go type
	}

	if err := h.HrmsUsecase.UpdatePerformanceReview(ctx, uint(req.ReviewId), updates); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}


func (h *HrmsHandler) DeletePerformanceReview(ctx context.Context, req *proto.DeletePerformanceReviewRequest) (*emptypb.Empty, error) {
	if err := h.HrmsUsecase.DeletePerformanceReview(ctx, uint(req.ReviewId)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func mapPerformanceReviewToProto(review *dto.PerformanceReviewDTO) *proto.PerformanceReview {
	return &proto.PerformanceReview{
		Id:           uint64(review.ID),
		EmployeeId:   uint64(review.EmployeeID),
		ReviewerId:   uint64(review.ReviewerID),
		ReviewDate:   timestamppb.New(review.ReviewDate),
		ReviewPeriod: review.ReviewPeriod,
		OverallRating: int32(review.OverallRating),
		Feedback:     review.Feedback,
		Promotion:    review.Promotion,
	}
}
