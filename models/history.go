package models

import "time"

type History struct {
	ID        int       `form:"id" json:"id" gorm:"primaryKey"`
	UserEmail string    `form:"userEmail" json:"userEmail"`
	FilmId    int       `form:"filmId" json:"filmId"`
	Date      time.Time `form:"date" json:"date"`
}

type HistoryResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
	History []Film `form:"film" json:"film"`
}
