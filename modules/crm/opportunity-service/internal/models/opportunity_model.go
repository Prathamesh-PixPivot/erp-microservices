package models

import (
    "time"
)

type Opportunity struct {
    ID          uint        `gorm:"primaryKey"`
    Name        *string     `gorm:""`
    Description *string     `gorm:""`
    Stage       *string     `gorm:""`
    Amount      *float64    `gorm:""`
    CloseDate   *time.Time  `gorm:""`
    Probability *float64    `gorm:""`
    LeadID      *uint       `gorm:""`
    AccountID   *uint       `gorm:""`
    OwnerID     *uint       `gorm:""`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}



