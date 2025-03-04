package dto

import "time"

// WorkHistoryDTO represents a structured data transfer object for Work History.
type WorkHistoryDTO struct {
	ID            uint       `json:"id,omitempty"`
	EmployeeID    uint       `json:"employee_id"`
	Company       string     `json:"company"`
	Designation   string     `json:"designation"`
	StartDate     time.Time  `json:"start_date"`
	EndDate       *time.Time `json:"end_date,omitempty"`
	ReasonForExit string     `json:"reason_for_exit"`
}
