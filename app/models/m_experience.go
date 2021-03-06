package models

import (
	"github.com/jinzhu/gorm"

	"github.com/quangdangfit/getjob/app/schema"
	"github.com/quangdangfit/getjob/pkg/utils"
)

// Experience model
type Experience struct {
	Model       `json:",inline"`
	UserID      string `json:"user_idr" gorm:"not null;index"`
	User        User
	CompanyID   string `json:"company_id" gorm:"index"`
	Company     Company
	Title       string `json:"title"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Description string `json:"description"`
}

// BeforeCreate handle before create to database
func (e *Experience) BeforeCreate(scope *gorm.Scope) error {
	err := e.Model.BeforeCreate(scope)
	if err != nil {
		return err
	}
	return nil
}

// ToSchema convert model to schema
func (e *Experience) ToSchema() *schema.Experience {
	var scExperience *schema.Experience
	err := utils.Copy(&scExperience, &e)
	if err != nil {
		return nil
	}

	return scExperience
}

// Experiences type
type Experiences []*Experience

// ToSchema convert model to schema
func (e Experiences) ToSchema() []*schema.Experience {
	scExperience := make([]*schema.Experience, len(e))
	for i, ex := range e {
		scExperience[i] = ex.ToSchema()
	}

	return scExperience
}
