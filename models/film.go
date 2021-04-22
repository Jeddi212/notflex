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
