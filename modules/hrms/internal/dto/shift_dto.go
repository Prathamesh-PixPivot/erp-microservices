package dto

// ShiftDTO represents a structured data transfer object for Shifts.
type ShiftDTO struct {
	ID         uint   `json:"id,omitempty"`
	Name       string `json:"name"`
	ShiftType  string `json:"shift_type"` // Stored as string (morning, evening, night)
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	DaysOfWeek string `json:"days_of_week"` // e.g., "Mon-Fri"
}
