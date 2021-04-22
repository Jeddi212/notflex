package controllers

import (
	"encoding/json"
	"fmt"
	model "github.com/notflex/models"
	"net/http"
)

func AddFilm(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// Get film data
	title := r.Form.Get("title")
	genre := r.Form.Get("genre")
	year := r.Form.Get("year")
	director := r.Form.Get("director")
	actor := r.Form.Get("actor")
	synopsis := r.Form.Get("synopsis")

	// Set inputted data to object
	film := model.Film{
		Title:    title,
		Genre:    genre,
		Year:     year,
		Director: director,
		Actor:    actor,
		Synopsis: synopsis,
	}

	// Insert object to database
	result := db.Create(&film)

	// Set response
	var response model.UserResponse
	if result.Error == nil {
		// Output to console
		//printUser(film)
		fmt.Println("Success Insert Film to Database")
		fmt.Println()

		response.Status = 200
		response.Message = "Success Insert Film to Database"
	} else {
		// Output to console
		fmt.Println("Insert Film Failed\n", result.Error.Error())

		response.Status = 400
		response.Message = "Insert Film Failed " + result.Error.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
