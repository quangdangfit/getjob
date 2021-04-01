package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Model              `json:"inline"`
	Username           string `json:"username" gorm:"unique;not null;index"`
	Email              string `json:"email" gorm:"unique;not null;index"`
	Password           string `json:"password" gorm:"not null;index"`
	Title              string `json:"title"`
	CompanyID          string `json:"company_id" gorm:"index"`
	AvatarURL          string `json:"avatar_url"`
	BackgroundPhotoURL string `json:"background_photo_url"`
	CV                 string `json:"cv"`
	CountryID          string `json:"country_id" gorm:"index"`
	CityID             string `json:"city_id" gorm:"index"`
	ViewCount          int    `json:"view_count" gorm:"index"`
	ConnectionCount    int    `json:"connection_count" gorm:"index"`
	Description        string `json:"description"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	//err := u.Model.BeforeCreate(scope)
	//if err != nil {
	//	return err
	//}
	//
	//hashedPassword, err := utils.HashPassword([]byte(u.Password))
	//if err != nil {
	//	return errors.Wrap(err, "User.BeforeCreate")
	//}
	//u.Password = hashedPassword
	//
	return nil
}
