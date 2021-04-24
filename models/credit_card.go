package models

type Credit struct {
	CardNumber string `form:"cardnumber" json:"cardnumber" gorm:"primaryKey"`
	Exp        string `form:"exp" json:"exp"`
	Cvc        string `form:"cvc" json:"cvc"`
	UserID     string `form:"userId" json:"userId" gorm:"index"`
	User       User
}

type CreditResponse struct {
	CardNumber string `form:"cardnumber" json:"cardnumber"`
	Exp        string `form:"exp" json:"exp"`
	Cvc        string `form:"cvc" json:"cvc"`
}
