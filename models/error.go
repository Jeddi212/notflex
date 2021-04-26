package models

type BasicResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
}
