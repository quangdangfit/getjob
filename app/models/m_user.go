package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"github.com/quangdangfit/getjob/app/schema"
	"github.com/quangdangfit/getjob/pkg/errors"
	"github.com/quangdangfit/getjob/pkg/utils"
)

type User struct {
	Model              `json:",inline"`
	Email              string `json:"email" gorm:"unique;not null;index"`
	Password           string `json:"password" gorm:"not null;index"`
	Title              string `json:"title"`
	Company            *Company
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
	err := u.Model.BeforeCreate(scope)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword([]byte(u.Password))
	if err != nil {
		return errors.New(errors.HashPasswordFail, err.Error())
	}
	u.Password = hashedPassword

	return nil
}

func (u *User) ToSchema() *schema.User {
	var scUser *schema.User
	err := utils.Copy(&scUser, &u)
	if err != nil {
		return nil
	}

	return scUser
}

func (u *User) CheckPassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false, errors.New(errors.InvalidPassword, err.Error())
	}

	return true, nil
}
