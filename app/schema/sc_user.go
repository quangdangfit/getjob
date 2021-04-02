package schema

import "time"

// User schema
type User struct {
	ID                 string    `json:"id"`
	Email              string    `json:"email"`
	Title              string    `json:"title"`
	CompanyID          string    `json:"company_id"`
	AvatarURL          string    `json:"avatar_url"`
	BackgroundPhotoURL string    `json:"background_photo_url"`
	CV                 string    `json:"cv"`
	CountryID          string    `json:"country_id"`
	CityID             string    `json:"city_id"`
	ViewCount          int       `json:"view_count"`
	ConnectionCount    int       `json:"connection_count"`
	Description        string    `json:"description"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	DeletedAt          time.Time `json:"deleted_at,omitempty"`
}

// Request
// =============================================================================

type UserRegisterParams struct {
	Email              string `json:"email" validate:"required,email"`
	Password           string `json:"password" validate:"required,password"`
	ConfirmPassword    string `json:"confirm_password" validate:"required,eqfield=Password"`
	Title              string `json:"title"`
	CompanyID          string `json:"company_id"`
	AvatarURL          string `json:"avatar_url"`
	BackgroundPhotoURL string `json:"background_photo_url"`
	CV                 string `json:"cv"`
	CountryID          string `json:"country_id"`
	CityID             string `json:"city_id"`
	Description        string `json:"description"`
}

// Response
//==============================================================================
