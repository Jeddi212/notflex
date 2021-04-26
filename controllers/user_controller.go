package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/notflex/models"

	"gorm.io/gorm"
)

func CheckUserLogin(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	email := r.URL.Query()["email"]
	password := r.URL.Query()["password"]
	var user models.User
	cekEmail := GetEmailFromToken(r)

	if cekEmail == "" {
		if err := db.Where("email = ? AND password = ?", email[0], password[0]).First(&user).Error; err == nil {
			if user.Status == "active" {
				if user.Level == 1 {
					generateToken(w, user.Email, 1)
				} else if user.Level == 0 {
					generateToken(w, user.Email, 0)
				}
				sendSuccessResponse(w)
			} else if user.Status == "suspend" {
				sendSuspendResponse(w)
			}
		} else {
			sendErrorResponse(w)
		}
	} else {
		var response models.ErrorResponse
		response.Status = 406
		response.Message = "Already logged in as " + email[0]
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)

	var response models.ErrorResponse
	response.Status = 200
	response.Message = "Success"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponse(w http.ResponseWriter) {
	var response models.ErrorResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter) {
	var response models.ErrorResponse
	response.Status = 400
	response.Message = "Failed"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuspendResponse(w http.ResponseWriter) {
	var response models.ErrorResponse
	response.Status = 403
	response.Message = "Suspended"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendUnauthorizedResponse(w http.ResponseWriter) {
	var response models.ErrorResponse
	response.Status = 401
	response.Message = "Unauthorized Access"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertMember(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	err := r.ParseForm()
	if err != nil {
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")
	name := r.Form.Get("name")
	birthDate := r.Form.Get("birthdate")
	gender := r.Form.Get("gender")
	nationality := r.Form.Get("nationality")

	var userCek models.User
	var response models.UserResponse
	err = db.Where("email = ?", email).First(&userCek).Error
	if errors.Is(err, gorm.ErrRecordNotFound) == true {
		user := models.User{
			Email:       email,
			Password:    password,
			Name:        name,
			BirthDate:   birthDate,
			Gender:      gender,
			Nationality: nationality,
			Status:      "active",
			Level:       1,
		}

		if err := db.Create(&user).Error; err == nil {
			response.Status = 200
			response.Message = "Registration success"
		} else {
			response.Status = 400
			response.Message = "Registration failed" + err.Error()
		}
	} else {
		response.Status = 400
		response.Message = "Email or has been taken"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
