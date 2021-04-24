package models

import "time"

type History struct {
	ID        int       `form:"id" json:"id" gorm:"primaryKey"`
	UserEmail string    `form:"userEmail" json:"userEmail"`
	FilmId    int       `form:"filmId" json:"filmId"`
	Date      time.Time `form:"date" json:"date"`
}

type Result struct {
	ID       int
	Title    string
	Genre    string
	Year     string
	Director string
	Actor    string
	Synopsis string
	Date     time.Time
}

type HistoryResponse struct {
	Status  int      `form:"status" json:"status"`
	Message string   `form:"message" json:"message"`
	History []Result `form:"film" json:"film"`
}
