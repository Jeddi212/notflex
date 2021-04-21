package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	ID        int `form:"id" json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// Email string `form:"email" json:"email"`
	// Password string `form:"password" json:"password"`
	User
	name        string `form:"name" json:"name"`
	birthDate   string `form:"birthdate" json:"birthdate"`
	gender      string `form:"gender" json:"gender"`
	nationality string `form:"nationality" json:"nationality"`
	status      string `form:"status" json:"status"`
}

func (member *Member) SetMember(email string, password string, name string, birthDate string, gender string, nationality string, status string) {
	member.SetUser(email, password)
	member.name = name
	member.birthDate = birthDate
	member.gender = gender
	member.nationality = nationality
	member.status = status
}

func (member *Member) SetName(name string) {
	member.name = name
}

func (member *Member) SetBirthDate(birthDate string) {
	member.birthDate = birthDate
}

func (member *Member) SetGender(gender string) {
	member.gender = gender
}

func (member *Member) SetNationality(nationality string) {
	member.nationality = nationality
}

func (member *Member) SetStatus(status string) {
	member.status = status
}

func (member *Member) GetName() string {
	return member.name
}

func (member *Member) GetBirthDate() string {
	return member.birthDate
}
func (member *Member) GetGender() string {
	return member.gender
}
func (member *Member) GetNationality() string {
	return member.nationality
}
func (member *Member) GetStatus() string {
	return member.status
}
