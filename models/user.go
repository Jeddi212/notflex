package models

import "time"

type User struct {
	Email       string    `form:"email" json:"email" gorm:"primaryKey"`
	Password    string    `form:"password" json:"password"`
	Name        string    `form:"name" json:"name"`
	BirthDate   string    `form:"birthdate" json:"birthdate"`
	Gender      string    `form:"gender" json:"gender"`
	Nationality string    `form:"nationality" json:"nationality"`
	Status      string    `form:"status" json:"status"`
	Subscribe   string    `form:"subscribe" json:"subscribe"`
	SubDate     time.Time `form:"subDate" json:"subDate"`
	Level       int       `form:"level" json:"level"`
}

type UserResponse struct {
	Status  int            `form:"status" json:"status"`
	Message string         `form:"message" json:"message"`
	Data    []User         `form:"data" json:"data"`
	Credit  CreditResponse `form:"credit" json:"credit"`
}

type SubscribeResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
	Type    string `form:"type" json:"type"`
}

type UnsubscribeResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
}
