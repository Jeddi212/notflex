package models

// type User struct {
// 	Email    string `form:"email" json:"id" gorm:"primaryKey"`
// 	Password string `form:"password" json:"password"`
// 	Level    string `form:"level" json:"level"`
// }

// type UserResponse struct {
// 	Status  int    `form:"status" json:"status"`
// 	Message string `form:"message" json:"message"`
// 	Data    []User `form:"data" json:"data"`
// }

// func (user *User) SetUser(email string, password string) {
// 	user.email = email
// 	user.password = password
// }

// func (user *User) GetEmail() string {
// 	return user.email
// }

// func (user *User) GetPassword() string {
// 	return user.password
// }

// func (user *User) SetEmail(email string) {
// 	user.email = email
// }

// func (user *User) SetPassword(password string) {
// 	user.password = password
// }

type User struct {
	Email       string `form:"email" json:"id" gorm:"primaryKey"`
	Password    string `form:"password" json:"password"`
	Name        string `form:"name" json:"name"`
	BirthDate   string `form:"birthdate" json:"birthdate"`
	Gender      string `form:"gender" json:"gender"`
	Nationality string `form:"nationality" json:"nationality"`
	Status      string `form:"status" json:"status"`
	Subscribe   string `fomr:"subscribe" json:"subscribe"`
	Level       int    `form:"level" json:"level"`
}

type UserResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}
