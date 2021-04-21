package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        int `form:"id" json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User
}

func (admin *Admin) SetAdmin(email string, password string) {
	admin.SetUser(email, password)
}
