package models

import "time"

type Film struct {
	ID       int    `form:"id" json:"id" gorm:"primaryKey"`
	Title    string `form:"title" json:"title"`
	Genre    string `form:"genre" json:"genre"`
	Year     string `form:"year" json:"year"`
	Director string `form:"director" json:"director"`
	Actor    string `form:"actor" json:"actor"`
	Synopsis string `form:"synopsis" json:"synopsis"`
}

type FilmResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
	Data    []Film `form:"data" json:"data"`
}

type WatchResponse struct {
	Status  int       `form:"status" json:"status"`
	Message string    `form:"message" json:"message"`
	Movie   Film      `form:"film" json:"film"`
	Date    time.Time `form:"date" json:"date"`
}
