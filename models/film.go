package models

type Film struct {
	ID       int    `form:"id" json:"id" gorm:"primaryKey"`
	Title    string `form:"title" json:"title"`
	Genre    string `form:"genre" json:"genre"`
	Year     string `form:"year" json:"year"`
	Director string `form:"director" json:"director"`
	Actor    string `form:"actor" json:"actor"`
	Synopsis string `form:"synopsis" json:"synopsis"`
}

// func (film *Film) SetFilm(title string, genre string, year string, director string, actor string, synopsis string) {
// 	film.title = title
// 	film.genre = genre
// 	film.year = year
// 	film.director = director
// 	film.actor = actor
// 	film.synopsis = synopsis
// }

// func (film *Film) SetTitle(title string) {
// 	film.title = title
// }

// func (film *Film) SetGenre(genre string) {
// 	film.genre = genre
// }

// func (film *Film) SetYear(year string) {
// 	film.year = year
// }

// func (film *Film) SetDirector(director string) {
// 	film.director = director
// }

// func (film *Film) SetActor(actor string) {
// 	film.actor = actor
// }

// func (film *Film) SetSynopsis(synopsis string) {
// 	film.synopsis = synopsis
// }

// func (film *Film) GetTitle() string {
// 	return film.title
// }

// func (film *Film) GetGenre() string {
// 	return film.genre
// }

// func (film *Film) GetYear() string {
// 	return film.year
// }

// func (film *Film) GetDirector() string {
// 	return film.director
// }

// func (film *Film) GetActor() string {
// 	return film.actor
// }

// func (film *Film) GetSynopsis() string {
// 	return film.synopsis
// }
