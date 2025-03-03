package grpc

import (
	"context"
	errors "errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hrms/internal/dto"
	hrmsErrors "hrms/internal/errors"
	"hrms/internal/transport/grpc/proto"
	"hrms/internal/usecase"
	"hrms/utils"
)

// UserGrpcHandler implements the gRPC service for user operations
type UserGrpcHandler struct {
	proto.UnimplementedUserServiceServer
	UserUsecase *usecase.UserUsecase
	Logger      *zap.Logger
}

// NewUserGrpcHandler initializes a new gRPC UserGrpcHandler
func NewUserGrpcHandler(userUsecase *usecase.UserUsecase, logger *zap.Logger) *UserGrpcHandler {
	return &UserGrpcHandler{
		UserUsecase: userUsecase,
		Logger:      logger,
	}
}

// CreateUser 📌 (Handles user signup)
func (h *UserGrpcHandler) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	h.Logger.Info("📩 Received CreateUser request", zap.String("email", req.Email))

	// Convert gRPC request to DTO
	userDTO := dto.CreateUserDTO{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		OrgID:       req.OrgId,
	}

	// Call usecase
	createdUser, err := h.UserUsecase.CreateUser(ctx, userDTO)
	if err != nil {
		h.Logger.Error("❌ Failed to create admin user", zap.Error(err))

		// Handle specific errors
		switch {
		case errors.Is(err, hrmsErrors.ErrUserAlreadyExists):
			return nil, status.Errorf(codes.AlreadyExists, "user with this email already exists")
		case errors.Is(err, hrmsErrors.ErrMissingOrganization):
			return nil, status.Errorf(codes.InvalidArgument, "organization ID is required")
		default:
			return nil, status.Errorf(codes.Internal, "failed to create admin user: %v", err)
		}
	}

	return &proto.CreateUserResponse{
		UserId:    createdUser.ID,
		Email:     createdUser.Email,
		CreatedAt: utils.FormatTimestamp(createdUser.CreatedAt), // ✅ Fixes timestamp issue
	}, nil
}

// AdminCreateUser 📌 (Handles admin user creation)
func (h *UserGrpcHandler) AdminCreateUser(ctx context.Context, req *proto.AdminCreateUserRequest) (*proto.AdminCreateUserResponse, error) {
	h.Logger.Info("📩 Received AdminCreateUser request", zap.String("email", req.Email), zap.String("org_id", req.OrgId))

	// Convert gRPC request to DTO
	adminUserDTO := dto.AdminCreateUserDTO{
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		ProfileImage: req.ProfileImage,
		OrgID:        req.OrgId,
		Active:       req.GetActive(), // ✅ Handles `nil` values safely
	}

	// Call usecase
	createdUser, err := h.UserUsecase.AdminCreateUser(ctx, adminUserDTO)
	if err != nil {
		h.Logger.Error("❌ Failed to create admin user", zap.Error(err))

		// Handle specific errors
		switch {
		case errors.Is(err, hrmsErrors.ErrUserAlreadyExists):
			return nil, status.Errorf(codes.AlreadyExists, "user with this email already exists")
		case errors.Is(err, hrmsErrors.ErrMissingOrganization):
			return nil, status.Errorf(codes.InvalidArgument, "organization ID is required")
		default:
			return nil, status.Errorf(codes.Internal, "failed to create admin user: %v", err)
		}
	}

	return &proto.AdminCreateUserResponse{
		UserId:    createdUser.ID,
		Email:     createdUser.Email,
		Active:    createdUser.Active,
		CreatedAt: utils.FormatTimestamp(createdUser.CreatedAt),
	}, nil
}

// 📌 GetUserByID (Fetch user by ID)
func (h *UserGrpcHandler) GetUserProfile(ctx context.Context, req *proto.GetUserProfileRequest) (*proto.GetUserProfileResponse, error) {
	h.Logger.Info("📩 Received GetUserByID request", zap.String("user_id", req.UserId))

	user, err := h.UserUsecase.GetUserByID(ctx, req.UserId)
	if err != nil {
		if errors.Is(err, hrmsErrors.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		h.Logger.Error("❌ Failed to fetch user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to fetch user: %v", err)
	}

	return &proto.GetUserProfileResponse{
		UserId:    user.ID,
		Email:     user.Email,
		CreatedAt: utils.FormatTimestamp(user.CreatedAt),
	}, nil
}

// GetUserProfileByUserID 📌 (Fetch user profile by user ID)
func (h *UserGrpcHandler) GetUserProfileByUserID(ctx context.Context, req *proto.GetUserProfileRequest) (*proto.GetUserProfileResponse, error) {
	h.Logger.Info("📩 Received GetUserProfileByUserID request", zap.String("user_id", req.UserId))

	profile, err := h.UserUsecase.GetUserProfileByUserID(ctx, req.UserId)
	if err != nil {
		if errors.Is(err, hrmsErrors.ErrProfileNotFound) {
			return nil, status.Errorf(codes.NotFound, "user profile not found")
		}
		h.Logger.Error("❌ Failed to fetch user profile", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to fetch user profile: %v", err)
	}

	return &proto.GetUserProfileResponse{
		UserId:       profile.UserID,
		FirstName:    utils.GetString(profile.FirstName),
		LastName:     utils.GetString(profile.LastName),
		ProfileImage: profile.ProfileImage,
	}, nil
}

// GetUserPreferences 📌 (Fetch user preferences)
func (h *UserGrpcHandler) GetUserPreferences(ctx context.Context, req *proto.GetUserPreferencesRequest) (*proto.GetUserPreferencesResponse, error) {
	h.Logger.Info("📩 Received GetUserPreferences request", zap.String("user_id", req.UserId))

	preferences, err := h.UserUsecase.GetUserPreferences(ctx, req.UserId)
	if err != nil {
		if errors.Is(err, hrmsErrors.ErrPreferencesNotFound) {
			return nil, status.Errorf(codes.NotFound, "user preferences not found")
		}
		h.Logger.Error("❌ Failed to fetch user preferences", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to fetch user preferences: %v", err)
	}

	return &proto.GetUserPreferencesResponse{
		Preferences: utils.ConvertToProtoPreferences(preferences),
	}, nil
}

// 📌 UpdateUserProfile (Update user profile)
func (h *UserGrpcHandler) UpdateUserProfile(ctx context.Context, req *proto.UpdateUserProfileRequest) (*proto.UpdateUserProfileResponse, error) {
	h.Logger.Info("📩 Received UpdateUserProfile request", zap.String("user_id", req.UserId))

	updates := make(map[string]interface{})
	if req.FirstName != nil {
		updates["first_name"] = req.FirstName
	}
	if req.LastName != nil {
		updates["last_name"] = req.LastName
	}
	if req.ProfileImage != nil {
		updates["profile_image"] = req.ProfileImage
	}

	err := h.UserUsecase.UpdateUserProfile(ctx, req.UserId, updates)
	if err != nil {
		h.Logger.Error("❌ Failed to update user profile", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to update user profile: %v", err)
	}

	return &proto.UpdateUserProfileResponse{Success: true}, nil
}

// 📌 DeleteUser (Soft Delete or Permanent)
func (h *UserGrpcHandler) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	h.Logger.Info("📩 Received DeleteUser request", zap.String("user_id", req.UserId))

	err := h.UserUsecase.DeleteUser(ctx, req.UserId, req.Permanent, "User requested deletion")
	if err != nil {
		h.Logger.Error("❌ Failed to delete user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}

	return &proto.DeleteUserResponse{Success: true}, nil
}

// ListUsers 📌 (Paginated user listing)
func (h *UserGrpcHandler) ListUsers(ctx context.Context, req *proto.ListUsersRequest) (*proto.ListUsersResponse, error) {
	h.Logger.Info("📩 Received ListUsers request")

	users, total, err := h.UserUsecase.ListUsers(ctx, int(req.Limit), int(req.Offset), utils.GetString(req.Search))
	if err != nil {
		h.Logger.Error("❌ Failed to list users", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to list users: %v", err)
	}

	var userResponses []*proto.GetUserProfileResponse
	for _, user := range users {
		userResponses = append(userResponses, &proto.GetUserProfileResponse{
			UserId:    user.ID,
			Email:     user.Email,
			CreatedAt: utils.FormatTimestamp(user.CreatedAt),
		})
	}

	return &proto.ListUsersResponse{
		Users:      userResponses,
		TotalCount: int32(total),
	}, nil
}

// ListUsersByOrganization 📌 (Lists users within a specific organization)
func (h *UserGrpcHandler) ListUsersByOrganization(ctx context.Context, req *proto.ListUsersByOrganizationRequest) (*proto.ListUsersByOrganizationResponse, error) {
	h.Logger.Info("📩 Received ListUsersByOrganization request", zap.String("org_id", req.OrgId))

	users, total, err := h.UserUsecase.ListUsersByOrganization(ctx, req.OrgId, int(req.Limit), int(req.Offset), utils.GetString(req.Search))
	if err != nil {
		h.Logger.Error("❌ Failed to list users by organization", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to list users by organization: %v", err)
	}

	var userResponses []*proto.GetUserProfileResponse
	for _, user := range users {
		userResponses = append(userResponses, &proto.GetUserProfileResponse{
			UserId:    user.ID,
			Email:     user.Email,
			CreatedAt: utils.FormatTimestamp(user.CreatedAt),
		})
	}

	return &proto.ListUsersByOrganizationResponse{
		Users:      userResponses,
		TotalCount: int32(total),
	}, nil
}
