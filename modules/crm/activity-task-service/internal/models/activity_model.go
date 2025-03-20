// internal/models/activity.go

package models

import (
	"time"
)

// Activity represents a high-level action or event related to customer interactions.
type Activity struct {
	Id          uint `gorm:"primaryKey"`
	Title       string `gorm:"size:255;not null;unique"`
	Description string `gorm:"type:text"`
	Type        string `gorm:"size:50;not null"` // e.g., Call, Meeting, Email
	Status      string `gorm:"size:50;not null"` // e.g., Pending, Completed, Cancelled
	DueDate     time.Time 
	CreatedAt   time.Time `grom:"autoCreateTime"`
	UpdatedAt   time.Time
	// Relationships
	ContactID uint   `gorm:"not null;index"`
	Tasks     []Task `gorm:"foreignKey:ActivityID"`
}
