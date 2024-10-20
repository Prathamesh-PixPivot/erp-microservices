// internal/repository/activities_repository.go

package repository

import (
	"activity-task-service/internal/models"
	"errors"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

var (
	ErrActivityExists   = errors.New("activity with this title already exists")
	ErrActivityNotFound = errors.New("activity not found")
	ErrTaskExists       = errors.New("task with this title already exists")
	ErrTaskNotFound     = errors.New("task not found")
)

// ActivityRepository defines the methods for activity-related database operations.
type ActivityRepository interface {
	CreateActivity(activity *models.Activity) (*models.Activity, error)
	GetActivity(id uint) (*models.Activity, error)
	UpdateActivity(activity *models.Activity) (*models.Activity, error)
	DeleteActivity(id uint) error
	ListActivities(pageNumber, pageSize uint, sortBy string, ascending bool, contactID uint) ([]models.Activity, error)
	GetActivityByID(id uint) (*models.Activity, error)
	CreateTask(task *models.Task) (*models.Task, error)
	GetTaskByID(id uint) (*models.Task, error)
	UpdateTask(task *models.Task) (*models.Task, error)
	DeleteTask(id uint) error
	ListTasks(pageNumber uint, pageSize uint, sortBy string, ascending bool, activityID uint) ([]models.Task, error)
}

type activityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &activityRepository{db: db}
}

// CreateActivity inserts a new activity into the database.
func (r *activityRepository) CreateActivity(activity *models.Activity) (*models.Activity, error) {
	if err := r.db.Create(activity).Error; err != nil {
		// Check for unique constraint violation on Title
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" { // unique_violation
				return nil, ErrActivityExists
			}
		}
		return nil, err
	}
	return activity, nil
}

func (r *activityRepository) GetActivity(id uint) (*models.Activity, error) {
	var activity models.Activity
	if err := r.db.Preload("Tasks").First(&activity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrActivityNotFound
		}
		return nil, err
	}
	return &activity, nil
}

// UpdateActivity modifies an existing activity.
func (r *activityRepository) UpdateActivity(activity *models.Activity) (*models.Activity, error) {
	result := r.db.Model(&models.Activity{}).Where("id = ?", activity.ID).Updates(activity)
	if result.Error != nil {
		if isUniqueConstraintError(result.Error, "activities_title_key") {
			return nil, ErrActivityExists
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, ErrActivityNotFound
	}
	return activity, nil
}

// DeleteActivity removes an activity by its ID.
func (r *activityRepository) DeleteActivity(id uint) error {
	result := r.db.Delete(&models.Activity{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrActivityNotFound
	}
	return nil
}

// ListActivities retrieves activities with pagination, sorting, and optional filtering by contact.
func (r *activityRepository) ListActivities(pageNumber uint, pageSize uint, sortBy string, ascending bool, contactID uint) ([]models.Activity, error) {
	var activities []models.Activity

	query := r.db.Model(&models.Activity{})

	// Apply filter by contact ID if provided
	if contactID != 0 {
		query = query.Where("contact_id = ?", contactID)
	}

	// Apply sorting
	if sortBy != "" {
		order := sortBy
		if ascending {
			order += " ASC"
		} else {
			order += " DESC"
		}
		query = query.Order(order)
	}

	// Apply pagination
	offset := (pageNumber - 1) * pageSize
	query = query.Offset(int(offset)).Limit(int(pageSize))

	if err := query.Preload("Tasks").Find(&activities).Error; err != nil {
		return nil, err
	}

	return activities, nil
}

// GetActivityByID retrieves an activity by its ID.
func (r *activityRepository) GetActivityByID(id uint) (*models.Activity, error) {
	var activity models.Activity
	if err := r.db.First(&activity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrActivityNotFound
		}
		return nil, err
	}
	return &activity, nil
}

// CreateTask inserts a new task into the database.
func (r *activityRepository) CreateTask(task *models.Task) (*models.Task, error) {
	if err := r.db.Create(task).Error; err != nil {
		if isUniqueConstraintError(err, "tasks_title_key") {
			return nil, ErrTaskExists
		}
		return nil, err
	}
	return task, nil
}

// GetTaskByID retrieves a task by its ID.
func (r *activityRepository) GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}
	return &task, nil
}

// UpdateTask modifies an existing task.
func (r *activityRepository) UpdateTask(task *models.Task) (*models.Task, error) {
	result := r.db.Model(&models.Task{}).Where("id = ?", task.ID).Updates(task)
	if result.Error != nil {
		if isUniqueConstraintError(result.Error, "tasks_title_key") {
			return nil, ErrTaskExists
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

// DeleteTask removes a task by its ID.
func (r *activityRepository) DeleteTask(id uint) error {
	result := r.db.Delete(&models.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrTaskNotFound
	}
	return nil
}

// ListTasks retrieves tasks with pagination, sorting, and optional filtering by activity.
func (r *activityRepository) ListTasks(pageNumber uint, pageSize uint, sortBy string, ascending bool, activityID uint) ([]models.Task, error) {
	var tasks []models.Task

	query := r.db.Model(&models.Task{})

	// Apply filter by activity ID if provided
	if activityID != 0 {
		query = query.Where("activity_id = ?", activityID)
	}

	// Apply sorting
	if sortBy != "" {
		order := sortBy
		if ascending {
			order += " ASC"
		} else {
			order += " DESC"
		}
		query = query.Order(order)
	}

	// Apply pagination
	offset := (pageNumber - 1) * pageSize
	query = query.Offset(int(offset)).Limit(int(pageSize))

	if err := query.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

// Helper function to check for unique constraint violations.
func isUniqueConstraintError(err error, constraintName string) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" && pgErr.ConstraintName == constraintName {
			return true
		}
	}
	return false
}
