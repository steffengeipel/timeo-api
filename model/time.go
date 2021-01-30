package model

import (
	"time"

	"gorm.io/gorm"
)

// Time Model zum Speichern der Zeiten
type Time struct {
	gorm.Model

	ID     string    `gorm:"primaryKey;autoIncrement:false"`
	UserID string    `json:"userId"`
	Time   time.Time `json:"time"`
	Status int       `json:"status"`
}
