package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	model "github.com/notflex/models"
	"gorm.io/gorm"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	var users []model.User
	var user model.User

	checkNil := false

	email := r.URL.Query()["email"]
	var err error
	if email != nil {
		err = db.Where("email = ?", email[0]).First(&user).Error
	} else {
		checkNil = true
	}

	var creditCard model.Credit
	var creditResponse model.CreditResponse
	if errors.Is(err, gorm.ErrRecordNotFound) == false {
		// Activate association mode
		db.Where("user_id = ?", user.Email).Find(&creditCard)

		// Set Credit Response
		creditResponse.CardNumber = creditCard.CardNumber
		creditResponse.Exp = creditCard.Exp
		creditResponse.Cvc = creditCard.Cvc

		// Add user to user list
		users = append(users, user)
	}

	// Set response
	var response model.UserResponse
	if len(users) > 0 && !checkNil {
		// Output to console
		fmt.Println("Success get user data", email)

		response.Status = 200
		response.Message = "Success Get User Data"
		response.Data = users
		response.Credit = creditResponse
	} else {
		// Output to console
		fmt.Println("Get Data Failed")
		fmt.Println()

		response.Status = 204
		response.Message = "Not Found, No Content"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SuspendUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// Get data from request
	email := r.Form.Get("email")

	// Update status
	result := db.Model(&model.User{}).Where("email = ?", email).Update("status", "suspend")

	// Set response
	var response model.UserResponse
	if result.Error == nil {
		// Output to console
		fmt.Println("Suspend User Success", email)

		response.Status = 200
		response.Message = "Suspend User Success"
	} else {
		// Output to console
		fmt.Println("Suspend Failed\n" + result.Error.Error())

		response.Status = 400
		response.Message = "Suspend Failed " + result.Error.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
