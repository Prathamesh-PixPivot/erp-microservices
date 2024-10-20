package models

import "time"

type Contact struct {
    ID                 uint       `gorm:"primaryKey"`
    FirstName          string     `gorm:"size:100;not null"`
    LastName           string     `gorm:"size:100;not null"`
    Email              string     `gorm:"size:100;unique;not null"`
    Phone              string     `gorm:"size:20"`
    Address            string     `gorm:"size:255"`
    City               string     `gorm:"size:100"`
    State              string     `gorm:"size:100"`
    Country            string     `gorm:"size:100"`
    ZipCode            string     `gorm:"size:20"`
    Company            string     `gorm:"size:100"`
    Position           string     `gorm:"size:100"`
    SocialMediaProfiles string     `gorm:"type:text"` // JSON string or consider separate table
    Notes              string     `gorm:"type:text"`
    CreatedAt          time.Time
    UpdatedAt          time.Time
}

