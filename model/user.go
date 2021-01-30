package model

import (
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	ID       string `gorm:"unique_index;not null;primaryKey;autoIncrement:false"`
	Username string `gorm:"unique_index;not null" json:"username"`
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Names    string `json:"names"`
	Times    []Time
	Projects []Project `gorm:"many2many:member_of_projects"`
}
