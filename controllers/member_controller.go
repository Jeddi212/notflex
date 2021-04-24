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

	// vars := mux.Vars(r)
	// email := vars["email"]
	email := GetEmailFromToken(r)

	var user models.User
	db.Where("email=?", email).First(&user)

	if name := r.Form.Get("name"); name != "" {
		user.Name = name
	}
	if birthdate := r.Form.Get("birthdate"); birthdate != "" {
		user.BirthDate = birthdate
	}
	if gender := r.Form.Get("gender"); gender != "" {
		user.Gender = gender
	}

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
