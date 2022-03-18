package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/notflex/models"

	"gorm.io/gorm"
)

// CheckUserLogin | User Function - Login as Admin or Member
func CheckUserLogin(w http.ResponseWriter, r *http.Request) {
	// Connect to database
	db := Connect()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// Get email and password given from form
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	//email := r.URL.Query()["email"]
	//password := r.URL.Query()["password"]

	// Set object model to hold user data
	var user models.User

	// Variable to check if the user has logged in
	cekEmail := GetEmailFromToken(r)

	// If there's no email logged yet from token | continue
	if cekEmail == "" {
		// If user credentials is correct | continue
		if err := db.Where("email = ? AND password = ?", email, password).First(&user).Error; err == nil {
			// If user account is in active status | continue
			if user.Status == "active" {
				// Set user token | Use user email - user type - current time -> Generate Unique Token
				if user.Level == 1 {
					generateToken(w, user.Email, 1)
				} else if user.Level == 0 {
					generateToken(w, user.Email, 0)
				}
				// Send response
				sendSuccessResponse(w)
			} else if user.Status == "suspend" {
				sendSuspendResponse(w)
			}
		} else {
			sendErrorResponse(w)
		}
	} else {
		var response models.BasicResponse

		// If there already a user logged in | Return Error response
		response.Status = 406
		response.Message = "Already logged in as " + cekEmail

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// Logout | User Function - Admin or Member Can Logout
func Logout(w http.ResponseWriter, r *http.Request) {
	// Reset Current Token
	resetUserToken(w)

	// Set response
	var response models.BasicResponse
	response.Status = 200
	response.Message = "Success Logout |  Bye - Bye"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponse(w http.ResponseWriter) {
	var response models.BasicResponse
	response.Status = 200
	response.Message = "Success Login"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter) {
	var response models.BasicResponse
	response.Status = 400
	response.Message = "Failed Login"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuspendResponse(w http.ResponseWriter) {
	var response models.BasicResponse
	response.Status = 400
	response.Message = "Account Suspended"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendUnauthorizedResponse(w http.ResponseWriter) {
	var response models.BasicResponse
	response.Status = 401
	response.Message = "Unauthorized Access"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Insertmember | System Function - A.k.a Registration - Register New Member to Database
func InsertMember(w http.ResponseWriter, r *http.Request) {
	// Connect to database
	db := Connect()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// Get provided data from form
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	name := r.Form.Get("name")
	birthDate := r.Form.Get("birthdate")
	gender := r.Form.Get("gender")
	nationality := r.Form.Get("nationality")

	// Set object model to hold user data
	var userCek models.User

	// Set reponse
	var response models.BasicResponse

	// Check if the email has been taken
	err = db.Where("email = ?", email).First(&userCek).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Set user data to object
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

		// Insert user data to database
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
