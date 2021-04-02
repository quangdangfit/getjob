package schema

import "time"

// Company schema
type Company struct {
	ID                 string     `json:"id"`
	Email              string     `json:"email" gorm:"unique;not null;index"`
	Fields             string     `json:"fields"`
	LogoURL            string     `json:"logo_url"`
	BackgroundPhotoURL string     `json:"background_photo_url"`
	CountryID          string     `json:"country_id" gorm:"index"`
	CityID             string     `json:"city_id" gorm:"index"`
	Description        string     `json:"description"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at"`
}

// Request
// =============================================================================

// CompanyCreateParams schema
type CompanyCreateParams struct {
	Email              string `json:"email" validate:"required"`
	Fields             string `json:"fields" validate:"required"`
	LogoURL            string `json:"logo_url"`
	BackgroundPhotoURL string `json:"background_photo_url"`
	CountryID          string `json:"country_id"`
	CityID             string `json:"city_id"`
	Description        string `json:"description"`
}
