package schema

import "time"

// Experience schema
type Experience struct {
	ID          string    `json:"id"`
	UserID      string    `json:"email"`
	CompanyID   string    `json:"company_id"`
	Title       string    `json:"title"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description"`
}

// Request
// =============================================================================

// ExperienceCreateParams schema
type ExperienceCreateParams struct {
	CompanyID   string `json:"company_id"`
	Title       string `json:"title"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Description string `json:"description"`
}
