package models

import (
	"time"

	"gorm.io/gorm"
)

// Allocation model
type Allocation struct {
	ID             string     `gorm:"primaryKey" json:"id"`
	AssetID        string     `gorm:"index" json:"asset_id"`
	AssignedTo     string     `json:"assigned_to"`
	AssignmentDate time.Time  `json:"assignment_date"`
	ReleaseDate    *time.Time `json:"release_date"` // nil if still allocated
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

// Audit model
type Audit struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	AssetID   string     `gorm:"index" json:"asset_id"`
	AuditedBy string     `json:"audited_by"`
	Condition string     `json:"condition"`
	Remarks   string     `json:"remarks"`
	AuditDate time.Time  `json:"audit_date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// Asset represents an organizational asset.
type Asset struct {
	ID               string     `gorm:"primaryKey" json:"id"` // UUID string
	Name             string     `json:"name"`                 // Asset name or description
	Description      string     `json:"description"`          // Detailed description
	Category         string     `gorm:"index" json:"category"`// Category (e.g., IT, furniture)
	PurchaseDate     time.Time  `json:"purchase_date"`        // Date of purchase
	PurchasePrice    float64    `json:"purchase_price"`       // Original cost
	CurrentValue     float64    `json:"current_value"`        // Value after depreciation
	Location         string     `json:"location"`             // Physical location
	Status           string     `gorm:"index" json:"status"`  // e.g., active, allocated, under_maintenance, disposed
	DepreciationRate float64    `json:"depreciation_rate"`    // Annual depreciation rate (e.g., 0.1 for 10%)
	Guidelines       string     `json:"guidelines"`           // Management guidelines or notes
	Allocations      []Allocation `gorm:"foreignKey:AssetID" json:"-"`
	Audits           []Audit      `gorm:"foreignKey:AssetID" json:"-"`
	Maintenances     []Maintenance `gorm:"foreignKey:AssetID" json:"-"`
	Licenses         []License     `gorm:"foreignKey:AssetID" json:"-"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *gorm.DeletedAt `gorm:"index" json:"-"`
}

// License model
type License struct {
	ID              string     `gorm:"primaryKey" json:"id"`
	AssetID         string     `gorm:"index" json:"asset_id"`
	LicenseKey      string     `json:"license_key"`
	ExpiryDate      time.Time  `json:"expiry_date"`
	Vendor          string     `json:"vendor"`
	ContractDetails string     `json:"contract_details"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

// Disposal model
type Disposal struct {
	ID               uint       `gorm:"primaryKey" json:"id"`
	AssetID          string     `gorm:"uniqueIndex" json:"asset_id"`
	Reason           string     `json:"reason"`
	DecommissionDate time.Time  `json:"decommission_date"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

// Maintenance model
type Maintenance struct {
	ID              string     `gorm:"primaryKey" json:"id"`
	AssetID         string     `gorm:"index" json:"asset_id"`
	MaintenanceDate time.Time  `json:"maintenance_date"`
	Description     string     `json:"description"`
	Cost            float64    `json:"cost"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
