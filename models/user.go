package models

type User struct {
	Email    string `form:"email" json:"id" gorm:"primaryKey"`
	Password string `form:"password" json:"password"`
	Level    string `form:"level" json:"password"`
}

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
