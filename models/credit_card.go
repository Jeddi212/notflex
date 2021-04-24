package models

type Credit struct {
	CardNumber string `form:"cardnumber" json:"cardnumber" gorm:"primaryKey"`
	Exp        string `form:"exp" json:"exp"`
	Cvc        string `form:"cvc" json:"cvc"`
	UserID     string `form:"userId" json:"userId" gorm:"foreignKey"`
	User       User
}
