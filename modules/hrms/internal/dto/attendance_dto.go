package dto

import "time"

// CreateAttendanceRequest represents the request for employee check-in
type CreateAttendanceRequest struct {
	EmployeeID  uint      `json:"employee_id" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
	CheckIn     time.Time `json:"check_in" validate:"required"`
	Location    string    `json:"location" validate:"required"`
	IsRemote    bool      `json:"is_remote"`
	PunchMethod string    `json:"punch_method" validate:"required,oneof=Manual Biometric Geolocation"`
}

// CheckOutAttendanceRequest represents the request for employee check-out
type CheckOutAttendanceRequest struct {
	EmployeeID uint      `json:"employee_id" validate:"required"`
	CheckOut   time.Time `json:"check_out" validate:"required"`
}

// AttendanceDTO represents the response structure for attendance records
type AttendanceDTO struct {
	ID          uint       `json:"id"`
	EmployeeID  uint       `json:"employee_id"`
	Date        time.Time  `json:"date"`
	CheckIn     *time.Time `json:"check_in,omitempty"`
	CheckOut    *time.Time `json:"check_out,omitempty"`
	WorkHours   float64    `json:"work_hours"`
	Overtime    float64    `json:"overtime"`
	BreakTime   float64    `json:"break_time"`
	Location    string     `json:"location"`
	IsRemote    bool       `json:"is_remote"`
	PunchMethod string     `json:"punch_method"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
