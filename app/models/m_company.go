package models

import (
	"github.com/jinzhu/gorm"

	"github.com/quangdangfit/getjob/app/schema"
	"github.com/quangdangfit/getjob/pkg/utils"
)

type Company struct {
	Model              `json:",inline"`
	Name               string `json:"name" gorm:"not null;index"`
	Email              string `json:"email" gorm:"unique;not null;index"`
	Fields             string `json:"fields"`
	LogoURL            string `json:"logo_url"`
	BackgroundPhotoURL string `json:"background_photo_url"`
	CountryID          string `json:"country_id" gorm:"index"`
	CityID             string `json:"city_id" gorm:"index"`
	Description        string `json:"description"`
}

func (c *Company) BeforeCreate(scope *gorm.Scope) error {
	err := c.Model.BeforeCreate(scope)
	if err != nil {
		return err
	}

	return nil
}

func (c *Company) ToSchema() *schema.Company {
	var scCompany *schema.Company
	err := utils.Copy(&scCompany, &c)
	if err != nil {
		return nil
	}

	return scCompany
}
