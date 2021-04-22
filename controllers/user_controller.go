package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/notflex/models"
)

func CheckUserLogin(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	email := r.URL.Query()["email"]
	var user models.User

	if err := db.Where("email = ?", email[0]).First(&user).Error; err == nil {
		if user.Level == 1 {
			generateToken(w, user.Email, 1)
		} else if user.Level == 0 {
			generateToken(w, user.Email, 0)
		}
		sendSuccessResponse(w)
	} else {
		sendErrorResponse(w)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)

	var response models.UserResponse
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

func sendUnauthorizedResponse(w http.ResponseWriter) {
	var response models.ErrorResponse
	response.Status = 401
	response.Message = "Unauthorized Access"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
