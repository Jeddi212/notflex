package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/notflex/models"
)

func UpdateMember(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	err := r.ParseForm()
	if err != nil {
		return
	}

	name := r.Form.Get("name")
	birthDate := r.Form.Get("birthdate")
	gender := r.Form.Get("gender")

	email := r.Form.Get("email")

	var user models.User
	db.Where("email=?", email).First(&user)

	user.Name = name
	user.BirthDate = birthDate
	user.Gender = gender

	result := db.Save(&user)

	var response models.UserResponse
	if result.Error == nil {
		response.Status = 200
		response.Message = "Success Update User Data"
	} else {
		response.Status = 400
		response.Message = "Update Failed"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
