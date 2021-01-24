package model

type ProjectMember struct {
	UserID    int `gorm:"primaryKey;autoIncrement:false"`
	ProjectID int `gorm:"primaryKey;autoIncrement:false"`
}
