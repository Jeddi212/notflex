package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	model "github.com/notflex/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// GetUser | Admin Function - Get Specific User From Database By Email
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Connect to database
	db := Connect()

	// Set object model to hold user data
	var users []model.User
	var user model.User

	// Set variable for nil possibility
	checkNil := false

	// Get user email from query
	email := r.URL.Query()["email"]
	var err error
	if email != nil {
		err = db.Where("email = ?", email[0]).First(&user).Error
	} else {
		checkNil = true
	}

	// Initialize Credit Card Response
	var creditCard model.Credit
	var creditResponse model.CreditResponse

	// If a record found
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Find user data by email
		db.Where("user_id = ?", user.Email).Find(&creditCard)

		// Set Credit Card Response
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

// SuspendUser | Admin Function - Suspend Specific User From Database By Email
func SuspendUser(w http.ResponseWriter, r *http.Request) {
	// Connect to database
	db := Connect()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// Get user email from form
	email := r.Form.Get("email")

	// Update user status
	result := db.Model(&model.User{}).
		Where("email = ?", email).
		Update("status", "suspend")

	// Set response
	var response model.BasicResponse
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

// GetFilmById | Admin Function - Search a Film From Database By FilmId
func GetFilmById(w http.ResponseWriter, r *http.Request) {
	// Connect to database
	db := Connect()

	// Get FilmId from request
	vars := mux.Vars(r)
	idFilm := vars["id_film"]
	fmt.Println(idFilm)

	// Set object model to hold film data
	var films []model.Film

	// Get film from database by id
	db.First(&films, idFilm)

	// Set response
	var response model.FilmResponse

	// If record(s) found
	if len(films) > 0 {
		response.Status = 200
		response.Message = "Success Get Film By ID"
		response.Data = films
	} else {
		response.Status = 204
		response.Message = "Get film failed, data empty"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetFilmByTitle | Admin Function - Search Film(s) From Database By Film Tile
func GetFilmByTitle(w http.ResponseWriter, r *http.Request) {
	// Connect to database
	db := Connect()

	// Get Film Title from request
	vars := mux.Vars(r)
	titleFilm := vars["title_film"]

	// Set object model to hold film data
	var films []model.Film

	// Get film from database by title
	db.Where("title = ?", titleFilm).Find(&films)

	// Set response
	var response model.FilmResponse

	// If record(s) found
	if len(films) > 0 {
		response.Status = 200
		response.Message = "Success Get Film By Title"
		response.Data = films
	} else {
		response.Status = 204
		response.Message = "Get film failed, data empty"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
