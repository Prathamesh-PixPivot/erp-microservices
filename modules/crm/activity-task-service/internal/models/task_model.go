// internal/models/task.go

package models

import "time"

// Task represents a specific actionable item associated with an activity.
type Task struct {
    ID          uint      `gorm:"primaryKey"`
    Title       string    `gorm:"size:255;not null;unique"`
    Description string    `gorm:"type:text"`
    Status      string    `gorm:"size:50;not null"` // e.g., Pending, In Progress, Completed
    Priority    string    `gorm:"size:50;not null"` // e.g., Low, Medium, High
    DueDate     time.Time
    CreatedAt   time.Time
    UpdatedAt   time.Time
    // Relationships
    ActivityID uint      `gorm:"not null;index"`
    Activity   Activity  `gorm:"foreignKey:ActivityID"`
}
