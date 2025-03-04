package dto

// SkillDevelopmentDTO represents a structured data transfer object for Skill Development.
type SkillDevelopmentDTO struct {
	ID       uint   `json:"id,omitempty"`
	ReviewID uint   `json:"review_id"`
	Skill    string `json:"skill"`
	Progress string `json:"progress"`
}
