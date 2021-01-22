package model

import (
	"gorm.io/gorm"
	"time"
)

// Time Model zum Speichern der Zeiten
type Time struct {
	gorm.Model

	UserId int       `json:"userId"`
	Time   time.Time `json:"time"`
	Status int       `json:"status"`
}
