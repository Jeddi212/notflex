package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

func EditFilm(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// Get id from path
	vars := mux.Vars(r)
	filmID := vars["film_id"]

	var film model.Film
	// Get film from database by id
	err = db.First(&film, filmID).Error

	// Get film data
	title := r.Form.Get("title")
	genre := r.Form.Get("genre")
	year := r.Form.Get("year")
	director := r.Form.Get("director")
	actor := r.Form.Get("actor")
	synopsis := r.Form.Get("synopsis")

	// Set inputted data to object
	if title != "" {
		film.Title = title
	}
	if genre != "" {
		film.Genre = genre
	}
	if year != "" {
		film.Year = year
	}
	if director != "" {
		film.Director = director
	}
	if actor != "" {
		film.Actor = actor
	}
	if synopsis != "" {
		film.Synopsis = synopsis
	}

	// Set response
	var response model.UserResponse

	// Save (update) film object to database
	if err == nil {
		result := db.Save(&film)
		if result.Error == nil {
			// Output to console
			//printUser(film)
			fmt.Println("Success Edit Film Data")
			fmt.Println()

			response.Status = 200
			response.Message = "Success Edit Film Data"
		} else {
			// Output to console
			fmt.Println("Edit Film Data Failed\n", result.Error.Error())

			response.Status = 400
			response.Message = "Edit Film Data Failed " + result.Error.Error()
		}
	} else {
		// Output to console
		fmt.Println("Edit Film Data Failed, ID Not Valid")
		fmt.Println()

		response.Status = 400
		response.Message = "Edit Film Data Failed, ID Not Valid"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
