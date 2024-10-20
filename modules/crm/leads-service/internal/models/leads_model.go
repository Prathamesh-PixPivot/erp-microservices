package models

import "time"

type Lead struct {
    ID            uint      `gorm:"primaryKey"`
    FirstName     string    `gorm:"not null"`
    LastName      string    `gorm:"not null"`
    Email         string    `gorm:"uniqueIndex;not null"`
    Phone         string
    Status        string    `gorm:"not null"`
    AssignedTo    int       // This will reference a user ID from user-service
    OrganizationID int      // This will reference an organization ID from organization-service
    CreatedAt     time.Time
    UpdatedAt     time.Time
}
