package models

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/quangdangfit/getjob/app/schema"
	"github.com/quangdangfit/getjob/pkg/utils"
)

type Experience struct {
	Model       `json:",inline"`
	UserID      string    `json:"email" gorm:"unique;not null;index"`
	CompanyID   string    `json:"company_id" gorm:"index"`
	Title       string    `json:"title"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description"`
}

func (e *Experience) BeforeCreate(scope *gorm.Scope) error {
	err := e.Model.BeforeCreate(scope)
	if err != nil {
		return err
	}
	return nil
}

func (e *Experience) ToSchema() *schema.Experience {
	var scExperience *schema.Experience
	err := utils.Copy(&scExperience, &e)
	if err != nil {
		return nil
	}

	return scExperience
}
