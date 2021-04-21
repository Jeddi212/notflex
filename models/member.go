package models

type Member struct {
	User
	Name        string `form:"name" json:"name"`
	BirthDate   string `form:"birthdate" json:"birthdate"`
	Gender      string `form:"gender" json:"gender"`
	Nationality string `form:"nationality" json:"nationality"`
	Status      string `form:"status" json:"status"`
	Subscribe   string `fomr:"subscribe" json:"subscribe"`
}

// func (member *Member) SetMember(email string, password string, name string, birthDate string, gender string, nationality string, status string, subscribe string) {
// 	member.SetUser(email, password)
// 	member.name = name
// 	member.birthDate = birthDate
// 	member.gender = gender
// 	member.nationality = nationality
// 	member.status = status
// 	member.subscribe = subscribe
// }

// func (member *Member) SetName(name string) {
// 	member.name = name
// }

// func (member *Member) SetBirthDate(birthDate string) {
// 	member.birthDate = birthDate
// }

// func (member *Member) SetGender(gender string) {
// 	member.gender = gender
// }

// func (member *Member) SetNationality(nationality string) {
// 	member.nationality = nationality
// }

// func (member *Member) SetStatus(status string) {
// 	member.status = status
// }

// func (member *Member) SetSubsribe(subscribe string) {
// 	member.subscribe = subscribe
// }

// func (member *Member) GetName() string {
// 	return member.name
// }

// func (member *Member) GetBirthDate() string {
// 	return member.birthDate
// }
// func (member *Member) GetGender() string {
// 	return member.gender
// }
// func (member *Member) GetNationality() string {
// 	return member.nationality
// }
// func (member *Member) GetStatus() string {
// 	return member.status
// }
