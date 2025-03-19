// internal/services/activities_service.go

package services

import (
	"activity-task-service/internal/models"
	"activity-task-service/internal/repository"
	"context"
	"errors"
	"regexp"
	"time"
)

var (
	ErrActivityNotFound    = errors.New("activity not found")
	ErrInvalidActivityData = errors.New("invalid activity data")
	ErrActivityExists      = errors.New("activity with this title already exists")
	ErrTaskNotFound        = errors.New("task not found")
	ErrInvalidTaskData     = errors.New("invalid task data")
	ErrTaskExists          = errors.New("task with this title already exists")
)

// ActivityService defines the methods for activity and task management.
type ActivityService interface {
	CreateActivity(ctx context.Context, activity *models.Activity) (*models.Activity, error)
	GetActivity(id uint) (*models.Activity, error)
	UpdateActivity(activity *models.Activity) (*models.Activity, error)
	DeleteActivity(id uint) error
	ListActivities(pageNumber, pageSize uint, sortBy string, ascending bool, contactID uint) ([]models.Activity, error)
	GetActivityByID(id uint) (*models.Activity, error)
	CreateTask(task *models.Task) (*models.Task, error)
	GetTask(id uint) (*models.Task, error)
	UpdateTask(task *models.Task) (*models.Task, error)
	DeleteTask(id uint) error
	ListTasks(pageNumber uint, pageSize uint, sortBy string, ascending bool, activityID uint) ([]models.Task, error)
}

type activityService struct {
	repo repository.ActivityRepository
}

// GetActivityByID implements ActivityService.
func (s *activityService) GetActivityByID(id uint) (*models.Activity, error) {
	activity, err := s.repo.GetActivityByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrActivityNotFound) {
			return nil, ErrActivityNotFound
		}
		return nil, err
	}
	return activity, nil
}

func NewActivityService(repo repository.ActivityRepository) ActivityService {
	return &activityService{repo: repo}
}

// CreateActivity validates and creates a new activity.
func (s *activityService) CreateActivity(ctx context.Context, activity *models.Activity) (*models.Activity, error) {
	// Validate required fields
	if activity.Title == "" || activity.Type == "" || activity.Status == "" || activity.ContactID == 0 {
		return nil, ErrInvalidActivityData
	}

	// Validate ActivityStatus using constants or enums
	validStatuses := map[string]bool{
		"Pending":    true,
		"InProgress": true,
		"Completed":  true,
		"Canceled":   true,
		"Scheduled":  true, // Include if 'Scheduled' is a valid status
	}

	if !validStatuses[activity.Status] {
		return nil, errors.New("invalid activity status")
	}

	// Set timestamps
	now := time.Now()
	activity.CreatedAt = now
	activity.UpdatedAt = now

	// Attempt to create the activity
	createdActivity, err := s.repo.CreateActivity(ctx, activity)
	if err != nil {
		if errors.Is(err, repository.ErrActivityExists) {
			return nil, ErrActivityExists
		}
		return nil, err
	}

	return createdActivity, nil
}

// GetActivity retrieves an activity by ID.
func (s *activityService) GetActivity(id uint) (*models.Activity, error) {
	activity, err := s.repo.GetActivityByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrActivityNotFound) {
			return nil, ErrActivityNotFound
		}
		return nil, err
	}
	return activity, nil
}

// UpdateActivity validates and updates an existing activity.
func (s *activityService) UpdateActivity(activity *models.Activity) (*models.Activity, error) {
	// Validate activity ID
	if activity.ID == 0 {
		return nil, ErrInvalidActivityData
	}

	// Validate Type and Status if provided
	if activity.Type != "" {
		validTypes := map[string]bool{
			"Call":    true,
			"Meeting": true,
			"Email":   true,
		}
		if !validTypes[activity.Type] {
			return nil, errors.New("invalid activity type")
		}
	}

	if activity.Status != "" {
		validStatuses := map[string]bool{
			"Pending":   true,
			"Completed": true,
			"Canceled":  true,
		}
		if !validStatuses[activity.Status] {
			return nil, errors.New("invalid activity status")
		}
	}

	// Validate DueDate if provided
	if !activity.DueDate.IsZero() && activity.DueDate.Before(time.Now()) {
		return nil, errors.New("due date cannot be in the past")
	}

	// Set the UpdatedAt timestamp
	activity.UpdatedAt = time.Now()

	// Update the activity
	updatedActivity, err := s.repo.UpdateActivity(activity)
	if err != nil {
		if errors.Is(err, repository.ErrActivityNotFound) {
			return nil, ErrActivityNotFound
		}
		if errors.Is(err, repository.ErrActivityExists) {
			return nil, ErrActivityExists
		}
		return nil, err
	}

	return updatedActivity, nil
}

// DeleteActivity removes an activity by ID.
func (s *activityService) DeleteActivity(id uint) error {
	// Check if the activity exists
	_, err := s.repo.GetActivityByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrActivityNotFound) {
			return ErrActivityNotFound
		}
		return err
	}

	// Delete the activity
	if err := s.repo.DeleteActivity(id); err != nil {
		return err
	}
	return nil
}

// ListActivities retrieves activities with pagination, sorting, and optional filtering by contact.
func (s *activityService) ListActivities(pageNumber uint, pageSize uint, sortBy string, ascending bool, contactID uint) ([]models.Activity, error) {
	// Validate pagination parameters
	if pageNumber == 0 {
		pageNumber = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}

	// Validate sortBy field
	validSortFields := map[string]bool{
		"title":      true,
		"due_date":   true,
		"created_at": true,
		"updated_at": true,
		"type":       true,
		"status":     true,
	}
	if sortBy != "" && !validSortFields[sortBy] {
		return nil, errors.New("invalid sort field")
	}

	activities, err := s.repo.ListActivities(pageNumber, pageSize, sortBy, ascending, contactID)
	if err != nil {
		return nil, err
	}
	return activities, nil
}

// CreateTask validates and creates a new task.
func (s *activityService) CreateTask(task *models.Task) (*models.Task, error) {
	// Validate required fields
	if task.Title == "" || task.Status == "" || task.Priority == "" || task.ActivityID == 0 {
		return nil, ErrInvalidTaskData
	}

	// Validate Status and Priority against predefined sets
	validStatuses := map[string]bool{
		"Pending":     true,
		"In Progress": true,
		"Completed":   true,
	}
	if !validStatuses[task.Status] {
		return nil, errors.New("invalid task status")
	}

	validPriorities := map[string]bool{
		"Low":    true,
		"Medium": true,
		"High":   true,
	}
	if !validPriorities[task.Priority] {
		return nil, errors.New("invalid task priority")
	}

	// Validate DueDate
	if !task.DueDate.IsZero() && task.DueDate.Before(time.Now()) {
		return nil, errors.New("due date cannot be in the past")
	}

	// Set timestamps
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	// Attempt to create the task
	createdTask, err := s.repo.CreateTask(task)
	if err != nil {
		if errors.Is(err, repository.ErrTaskExists) {
			return nil, ErrTaskExists
		}
		return nil, err
	}

	return createdTask, nil
}

// GetTask retrieves a task by ID.
func (s *activityService) GetTask(id uint) (*models.Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}
	return task, nil
}

// UpdateTask validates and updates an existing task.
func (s *activityService) UpdateTask(task *models.Task) (*models.Task, error) {
	// Validate task ID
	if task.ID == 0 {
		return nil, ErrInvalidTaskData
	}

	// Validate Status and Priority if provided
	if task.Status != "" {
		validStatuses := map[string]bool{
			"Pending":     true,
			"In Progress": true,
			"Completed":   true,
		}
		if !validStatuses[task.Status] {
			return nil, errors.New("invalid task status")
		}
	}

	if task.Priority != "" {
		validPriorities := map[string]bool{
			"Low":    true,
			"Medium": true,
			"High":   true,
		}
		if !validPriorities[task.Priority] {
			return nil, errors.New("invalid task priority")
		}
	}

	// Validate DueDate if provided
	if !task.DueDate.IsZero() && task.DueDate.Before(time.Now()) {
		return nil, errors.New("due date cannot be in the past")
	}

	// Set the UpdatedAt timestamp
	task.UpdatedAt = time.Now()

	// Update the task
	updatedTask, err := s.repo.UpdateTask(task)
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			return nil, ErrTaskNotFound
		}
		if errors.Is(err, repository.ErrTaskExists) {
			return nil, ErrTaskExists
		}
		return nil, err
	}

	return updatedTask, nil
}

// DeleteTask removes a task by ID.
func (s *activityService) DeleteTask(id uint) error {
	// Check if the task exists
	_, err := s.repo.GetTaskByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			return ErrTaskNotFound
		}
		return err
	}

	// Delete the task
	if err := s.repo.DeleteTask(id); err != nil {
		return err
	}
	return nil
}

// ListTasks retrieves tasks with pagination, sorting, and optional filtering by activity.
func (s *activityService) ListTasks(pageNumber uint, pageSize uint, sortBy string, ascending bool, activityID uint) ([]models.Task, error) {
	// Validate pagination parameters
	if pageNumber == 0 {
		pageNumber = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}

	// Validate sortBy field
	validSortFields := map[string]bool{
		"title":      true,
		"due_date":   true,
		"created_at": true,
		"updated_at": true,
		"status":     true,
		"priority":   true,
	}
	if sortBy != "" && !validSortFields[sortBy] {
		return nil, errors.New("invalid sort field")
	}

	tasks, err := s.repo.ListTasks(pageNumber, pageSize, sortBy, ascending, activityID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// Helper function to validate email format using regex (if needed for tasks).
func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
