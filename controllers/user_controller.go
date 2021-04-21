package controllers

import (
	"fmt"
	"net/http"

	model "github.com/notflex/models"
)

func CheckUserLogin(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	email := r.URL.Query()["email"]
	var admin model.Admin
	var member model.Member
	var found bool

	if err := db.Where("email = ?", email[0]).First(&admin).Error; err == nil {
		found = true
	} else if err := db.Where("email = ?", email[0]).First(&member).Error; err == nil {
		found = true
	} else {
		found = false
	}

	if found {
		// response.Status = 200
		// response.Message = "Success"
		// response.Data = users
		fmt.Print("A")
	} else {
		// response.Status = 400
		// response.Message = "Failed to get a user"
		fmt.Print("B")
	}
}
