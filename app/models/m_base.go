package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Model base model
type Model struct {
	ID        string     `json:"id" gorm:"unique;not null;index;primary_key"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null;index"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"not null;index"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

// BeforeCreate handle before create from database
func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	if model.ID == "" {
		model.ID = uuid.New().String()
	}
	if model.CreatedAt.IsZero() {
		model.CreatedAt = time.Now()
	}
	if model.UpdatedAt.IsZero() {
		model.UpdatedAt = time.Now()
	}
	return nil
}

// BeforeUpdate handle before update to database
func (model *Model) BeforeUpdate(scope *gorm.Scope) error {
	model.UpdatedAt = time.Now().UTC()
	return nil
}

// BeforeDelete handle before delete from database
func (model *Model) BeforeDelete(scope *gorm.Scope) error {
	model.UpdatedAt = time.Now().UTC()
	return nil
}
