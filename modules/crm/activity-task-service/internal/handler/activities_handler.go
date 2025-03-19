// grpc/handler/activities_handler.go

package handler

import (
	"activity-task-service/grpc/activitypb"
	"activity-task-service/internal/models"
	"activity-task-service/internal/services"
	"context"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ActivityHandler struct {
	activityService services.ActivityService
	activitypb.UnimplementedActivityServiceServer
}

func NewActivityHandler(service services.ActivityService) *ActivityHandler {
	return &ActivityHandler{activityService: service}
}

// CreateActivity handles the creation of a new activity.
func (h *ActivityHandler) CreateActivity(ctx context.Context, req *activitypb.CreateActivityRequest) (*activitypb.CreateActivityResponse, error) {
	log.Printf("Received CreateActivity request: %+v", req)

	// Convert Proto to Model
	activity := convertProtoToModel(req.Activity)

	// Create Activity via Service
	createdActivity, err := h.activityService.CreateActivity(ctx,activity)
	if err != nil {
		log.Printf("Error creating activity: %v", err)
		switch err {
		case services.ErrActivityExists:
			return nil, status.Error(codes.AlreadyExists, err.Error())
		case services.ErrInvalidActivityData:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Error(codes.Internal, "failed to create activity")
		}
	}

	// Convert Model to Proto
	return &activitypb.CreateActivityResponse{
		Activity: convertModelToProto(createdActivity),
	}, nil
}

// GetActivity handles retrieval of an activity by ID.
func (h *ActivityHandler) GetActivity(ctx context.Context, req *activitypb.GetActivityRequest) (*activitypb.GetActivityResponse, error) {
	activity, err := h.activityService.GetActivity(uint(req.Id))
	if err != nil {
		log.Printf("Error getting activity: %v", err)
		switch err {
		case services.ErrActivityNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "failed to get activity")
		}
	}

	return &activitypb.GetActivityResponse{
		Activity: convertModelToProto(activity),
	}, nil
}

// UpdateActivity handles updating an existing activity.
func (h *ActivityHandler) UpdateActivity(ctx context.Context, req *activitypb.UpdateActivityRequest) (*activitypb.UpdateActivityResponse, error) {
	log.Printf("Received UpdateActivity request: %+v", req)

	// Convert Proto to Model
	activity := convertProtoToModel(req.Activity)

	// Update Activity via Service
	updatedActivity, err := h.activityService.UpdateActivity(activity)
	if err != nil {
		log.Printf("Error updating activity: %v", err)
		switch err {
		case services.ErrActivityNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		case services.ErrActivityExists:
			return nil, status.Error(codes.AlreadyExists, err.Error())
		case services.ErrInvalidActivityData:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Error(codes.Internal, "failed to update activity")
		}
	}

	// Convert Model to Proto
	return &activitypb.UpdateActivityResponse{
		Activity: convertModelToProto(updatedActivity),
	}, nil
}

// DeleteActivity handles deletion of an activity by ID.
func (h *ActivityHandler) DeleteActivity(ctx context.Context, req *activitypb.DeleteActivityRequest) (*activitypb.DeleteActivityResponse, error) {
	log.Printf("Received DeleteActivity request: %+v", req)

	// Delete Activity via Service
	err := h.activityService.DeleteActivity(uint(req.Id))
	if err != nil {
		log.Printf("Error deleting activity: %v", err)
		switch err {
		case services.ErrActivityNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "failed to delete activity")
		}
	}

	// Return Success
	return &activitypb.DeleteActivityResponse{
		Success: true,
	}, nil
}

// ListActivities handles listing activities with pagination and optional filtering.
func (h *ActivityHandler) ListActivities(ctx context.Context, req *activitypb.ListActivitiesRequest) (*activitypb.ListActivitiesResponse, error) {
	activities, err := h.activityService.ListActivities(uint(req.PageNumber), uint(req.PageSize), req.SortBy, req.Ascending, uint(req.ContactId))
	if err != nil {
		log.Printf("Error listing activities: %v", err)
		switch err {
		case services.ErrInvalidActivityData:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Error(codes.Internal, "failed to list activities")
		}
	}

	// Convert Models to Proto
	var protoActivities []*activitypb.Activity
	for _, activity := range activities {
		protoActivities = append(protoActivities, convertModelToProto(&activity))
	}

	return &activitypb.ListActivitiesResponse{
		Activities: protoActivities,
	}, nil
}

// CreateTask handles the creation of a new task.
func (h *ActivityHandler) CreateTask(ctx context.Context, req *activitypb.CreateTaskRequest) (*activitypb.CreateTaskResponse, error) {
	log.Printf("Received CreateTask request: %+v", req)

	// Convert Proto to Model
	task := convertProtoToModelTask(req.Task)

	// Create Task via Service
	createdTask, err := h.activityService.CreateTask(task)
	if err != nil {
		log.Printf("Error creating task: %v", err)
		switch err {
		case services.ErrTaskExists:
			return nil, status.Error(codes.AlreadyExists, err.Error())
		case services.ErrInvalidTaskData:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case services.ErrActivityNotFound:
			return nil, status.Error(codes.NotFound, "associated activity not found")
		default:
			return nil, status.Error(codes.Internal, "failed to create task")
		}
	}

	// Convert Model to Proto
	return &activitypb.CreateTaskResponse{
		Task: convertTaskModelToProto(createdTask),
	}, nil
}

// GetTask handles retrieval of a task by ID.
func (h *ActivityHandler) GetTask(ctx context.Context, req *activitypb.GetTaskRequest) (*activitypb.GetTaskResponse, error) {
	task, err := h.activityService.GetTask(uint(req.Id))
	if err != nil {
		log.Printf("Error getting task: %v", err)
		switch err {
		case services.ErrTaskNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "failed to get task")
		}
	}

	return &activitypb.GetTaskResponse{
		Task: convertTaskModelToProto(task),
	}, nil
}

// UpdateTask handles updating an existing task.
func (h *ActivityHandler) UpdateTask(ctx context.Context, req *activitypb.UpdateTaskRequest) (*activitypb.UpdateTaskResponse, error) {
	log.Printf("Received UpdateTask request: %+v", req)

	// Convert Proto to Model
	task := convertProtoToModelTask(req.Task)

	// Update Task via Service
	updatedTask, err := h.activityService.UpdateTask(task)
	if err != nil {
		log.Printf("Error updating task: %v", err)
		switch err {
		case services.ErrTaskNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		case services.ErrTaskExists:
			return nil, status.Error(codes.AlreadyExists, err.Error())
		case services.ErrInvalidTaskData:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Error(codes.Internal, "failed to update task")
		}
	}

	// Convert Model to Proto
	return &activitypb.UpdateTaskResponse{
		Task: convertTaskModelToProto(updatedTask),
	}, nil
}

// DeleteTask handles deletion of a task by ID.
func (h *ActivityHandler) DeleteTask(ctx context.Context, req *activitypb.DeleteTaskRequest) (*activitypb.DeleteTaskResponse, error) {
	log.Printf("Received DeleteTask request: %+v", req)

	// Delete Task via Service
	err := h.activityService.DeleteTask(uint(req.Id))
	if err != nil {
		log.Printf("Error deleting task: %v", err)
		switch err {
		case services.ErrTaskNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "failed to delete task")
		}
	}

	// Return Success
	return &activitypb.DeleteTaskResponse{
		Success: true,
	}, nil
}

// ListTasks handles listing tasks with pagination and optional filtering.
func (h *ActivityHandler) ListTasks(ctx context.Context, req *activitypb.ListTasksRequest) (*activitypb.ListTasksResponse, error) {
	tasks, err := h.activityService.ListTasks(uint(req.PageNumber), uint(req.PageSize), req.SortBy, req.Ascending, uint(req.ActivityId))
	if err != nil {
		log.Printf("Error listing tasks: %v", err)
		switch err {
		case services.ErrInvalidTaskData:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Error(codes.Internal, "failed to list tasks")
		}
	}

	// Convert Models to Proto
	var protoTasks []*activitypb.Task
	for _, task := range tasks {
		protoTasks = append(protoTasks, convertTaskModelToProto(&task))
	}

	return &activitypb.ListTasksResponse{
		Tasks: protoTasks,
	}, nil
}

// Conversion Functions

func convertProtoToModel(protoActivity *activitypb.Activity) *models.Activity {
	dueDate, _ := time.Parse(time.RFC3339, protoActivity.DueDate)
	return &models.Activity{
		Title:       protoActivity.Title,
		Description: protoActivity.Description,
		Type:        protoActivity.Type,
		Status:      protoActivity.Status,
		DueDate:     dueDate,
		ContactID:   uint(protoActivity.ContactId),
	}
}

func convertModelToProto(modelActivity *models.Activity) *activitypb.Activity {
	dueDate := ""
	if !modelActivity.DueDate.IsZero() {
		dueDate = modelActivity.DueDate.Format(time.RFC3339)
	}
	return &activitypb.Activity{
		Id:          uint32(modelActivity.ID),
		Title:       modelActivity.Title,
		Description: modelActivity.Description,
		Type:        modelActivity.Type,
		Status:      modelActivity.Status,
		DueDate:     dueDate,
		CreatedAt:   modelActivity.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   modelActivity.UpdatedAt.Format(time.RFC3339),
		ContactId:   uint32(modelActivity.ContactID),
	}
}

func convertProtoToModelTask(protoTask *activitypb.Task) *models.Task {
	dueDate, _ := time.Parse(time.RFC3339, protoTask.DueDate)
	return &models.Task{
		Title:       protoTask.Title,
		Description: protoTask.Description,
		Status:      protoTask.Status,
		Priority:    protoTask.Priority,
		DueDate:     dueDate,
		ActivityID:  uint(protoTask.ActivityId),
	}
}

func convertTaskModelToProto(modelTask *models.Task) *activitypb.Task {
	dueDate := ""
	if !modelTask.DueDate.IsZero() {
		dueDate = modelTask.DueDate.Format(time.RFC3339)
	}
	return &activitypb.Task{
		Id:          uint32(modelTask.ID),
		Title:       modelTask.Title,
		Description: modelTask.Description,
		Status:      modelTask.Status,
		Priority:    modelTask.Priority,
		DueDate:     dueDate,
		CreatedAt:   modelTask.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   modelTask.UpdatedAt.Format(time.RFC3339),
		ActivityId:  uint32(modelTask.ActivityID),
	}
}
