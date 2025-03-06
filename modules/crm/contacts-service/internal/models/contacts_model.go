package models

import "time"

// TaxationDetail holds generic taxation information per country and tax type.
type TaxationDetail struct {
    ID          uint      `gorm:"primaryKey"`
    Country     string    `gorm:"size:100;not null"` // ISO code or country name
    TaxType     string    `gorm:"size:50;not null"`  // e.g., "GST", "VAT", etc.
    Rate        float64   `gorm:"not null"`          // Percentage value (e.g., 20 for 20%)
    Description string    `gorm:"type:text"`         // Additional details if necessary
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// Contact represents both individual contacts and company contacts.
// It supports an optional association with taxation details for companies.
type Contact struct {
    ID                  uint           `gorm:"primaryKey"`
    ContactType         string         `gorm:"size:20;not null"` // "individual" or "company"
    
    // Fields for individual contacts.
    FirstName           string         `gorm:"size:100"` // Required for individuals
    LastName            string         `gorm:"size:100"` // Required for individuals
    
    // Company-related field.
    // For company contacts, this holds the company name.
    CompanyName         string         `gorm:"size:100"`
    
    Email               string         `gorm:"size:100;unique;not null"`
    Phone               string         `gorm:"size:20"`
    Address             string         `gorm:"size:255"`
    City                string         `gorm:"size:100"`
    State               string         `gorm:"size:100"`
    Country             string         `gorm:"size:100"`
    ZipCode             string         `gorm:"size:20"`
    Position            string         `gorm:"size:100"` // Useful for individuals
    SocialMediaProfiles string         `gorm:"type:text"` // Consider using JSON or a separate table
    Notes               string         `gorm:"type:text"`
    
    // Optional taxation details association.
    // Populated only for company contacts that require taxation details.
    TaxationDetailID    *uint
    TaxationDetail      *TaxationDetail `gorm:"foreignKey:TaxationDetailID"`
    
    CreatedAt           time.Time
    UpdatedAt           time.Time
}
