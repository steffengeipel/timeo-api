package model

import "gorm.io/gorm"

type ProjectMember struct {
	gorm.Model

	UserID     string `gorm:"primaryKey;autoIncrement:false"`
	ProjectID  string `gorm:"primaryKey;autoIncrement:false"`
	Permission int
}

func (ProjectMember) BeforeCreate(db *gorm.DB) error {
	// ...
	return nil
}
