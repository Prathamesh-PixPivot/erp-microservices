package dto

import "time"

// PerformanceReviewDTO represents a structured DTO for Performance Reviews.
type PerformanceReviewDTO struct {
	ID                uint                  `json:"id,omitempty"`
	EmployeeID        uint                  `json:"employee_id"`
	ReviewerID        uint                  `json:"reviewer_id"`
	ReviewDate        time.Time             `json:"review_date"`
	ReviewPeriod      string                `json:"review_period"`
	OverallRating     int                   `json:"overall_rating"`
	Feedback          string                `json:"feedback"`
	Promotion         bool                  `json:"promotion"`
	KPIs              []PerformanceKPIDTO   `json:"kpis,omitempty"`
	SkillDevelopments []SkillDevelopmentDTO `json:"skill_developments,omitempty"`
}

// PerformanceKPIDTO represents a structured DTO for Performance KPIs.
type PerformanceKPIDTO struct {
	ID       uint    `json:"id,omitempty"`
	ReviewID uint    `json:"review_id"`
	KPIName  string  `json:"kpi_name"`
	Score    float64 `json:"score"`
	Comments string  `json:"comments,omitempty"`
}
