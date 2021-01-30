package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model

	ID   string `gorm:"primaryKey;autoIncrement:false;"`
	Name string `json:"name"`
}
