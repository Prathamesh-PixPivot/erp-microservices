package dto

import "time"

// PublicHolidayDTO represents a structured data transfer object for Public Holidays.
type PublicHolidayDTO struct {
	ID             uint      `json:"id,omitempty"`
	OrganizationID uint      `json:"organization_id"`
	Name           string    `json:"name"`
	Date           time.Time `json:"date"`
}
