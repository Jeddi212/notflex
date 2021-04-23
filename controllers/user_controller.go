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

func sendSuspendResponse(w http.ResponseWriter) {
	var response models.ErrorResponse
	response.Status = 400
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
	ccNumber := r.Form.Get("creditCardNumber")
	exp := r.Form.Get("expire")
	cvc := r.Form.Get("cvc")

	var userCek models.User
	var response models.UserResponse
	err = db.Where("email = ?", email).First(&userCek).Error
	if errors.Is(err, gorm.ErrRecordNotFound) == true {
		cc := models.Credit{CardNumber: ccNumber, Exp: exp, Cvc: cvc}
		user := models.User{Email: email, Password: password, Name: name, BirthDate: birthDate, Gender: gender, Nationality: nationality, Status: "active", Level: 1, CreditID: ccNumber}

		if ccNumber != "" {
			var ccCek models.Credit
			err = db.Where("card_number = ?", ccNumber).First(&ccCek).Error
			if errors.Is(err, gorm.ErrRecordNotFound) == true {
				db.Create(&cc)
				if err := db.Create(&user).Error; err == nil {
					response.Status = 200
					response.Message = "Registration success"
				} else {
					db.Where("card_number = ?", ccNumber).Delete(&cc)
					response.Status = 400
					response.Message = "Registration failed"
				}
			} else {
				response.Status = 400
				response.Message = "Credit card number has been registrered"
			}
		} else {
			response.Status = 400
			response.Message = "Registration credit card failed"
		}
	} else {
		response.Status = 400
		response.Message = "Email or has been taken"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

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
