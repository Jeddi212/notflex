package controllers

import (
	"encoding/json"
	"fmt"
	model "github.com/notflex/models"
	"net/http"
)

// GetUser
func GetUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	var users []model.User
	var user model.User

	email := r.URL.Query()["email"]
	if email != nil {
		db.Where("email = ?", email[0]).First(&user)
	} else {
		db.Find(&user)
	}

	// Activate association mode
	db.Model(&user).Association("Credit")

	// Get data from many relations
	db.Model(&user).Association("Credit").Find(&user.Credit)

	// Add user to user list
	users = append(users, user)

	// Set response
	var response model.UserResponse
	if len(users) > 0 {
		// Output to console
		//printUsers(users)
		fmt.Println("Success get user data")

		response.Status = 200
		response.Message = "Success Get User Data"
		response.Data = users
	} else {
		// Output to console
		fmt.Println("Get Data Failed\n")

		response.Status = 204
		response.Message = "Not Found, No Content"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
