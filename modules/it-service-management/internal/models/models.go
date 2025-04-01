// internal/models/incident.go
package models

import "time"

// Incident represents an ITSM incident.
type Incident struct {
	ID          string    `json:"id" gorm:"primaryKey;size:36"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	HostID      string    `json:"hostid"`
	Status      string    `json:"status"` // e.g., "open", "in_progress", "resolved", "closed"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Change represents a planned change in the IT environment.
type Change struct {
	ID          string    `json:"id" gorm:"primaryKey;size:36"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Requester   string    `json:"requester"`
	Approver    string    `json:"approver,omitempty"`
	Status      string    `json:"status"` // e.g., "pending", "approved", "implemented", "rejected"
	PlannedTime time.Time `json:"planned_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ServiceRequest represents a routine IT service request.
type ServiceRequest struct {
	ID          string    `json:"id" gorm:"primaryKey;size:36"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Requester   string    `json:"requester"`
	Status      string    `json:"status"` // e.g., "open", "in_progress", "completed", "closed"
	SLA         string    `json:"sla"`    // e.g., "4h response", "8h resolution"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
