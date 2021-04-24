package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	model "github.com/notflex/models"
	"github.com/gorilla/mux"

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
	var response model.FilmResponse
	if result.Error == nil {
		// Output to console
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
	var response model.FilmResponse
	if err == nil {
		result := db.Save(&film)
		if result.Error == nil {
			// Output to console
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
		response.Message = "Edit Film Data Failed, ID Not Valid " + err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SearchFilm(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	var films []model.Film

	// Get film data
	title := r.URL.Query()["title"]
	genre := r.URL.Query()["genre"]
	year := r.URL.Query()["year"]
	director := r.URL.Query()["director"]
	actor := r.URL.Query()["actor"]
	synopsis := r.URL.Query()["synopsis"]

	// Inialize db
	query := db

	// Set where condition
	if title != nil {
		titles := "%" + title[0] + "%"
		query = query.Where("title LIKE ?", titles)
	}
	if genre != nil {
		genres := "%" + genre[0] + "%"
		query = query.Where("genre LIKE ?", genres)
	}
	if year != nil {
		query = query.Where("year = ?", year[0])
	}
	if director != nil {
		directors := "%" + director[0] + "%"
		query = query.Where("director LIKE ?", directors)
	}
	if synopsis != nil {
		synopsises := "%" + synopsis[0] + "%"
		query = query.Where("synopsis LIKE ?", synopsises)
	}
	if actor != nil {
		actors := "%" + actor[0] + "%"
		query = query.Where("actor LIKE ?", actors)
	}

	// Finish the query
	result := query.Find(&films).Error

	// Set response
	var response model.FilmResponse
	if result == nil {
		if len(films) > 0 {
			// Output to console
			fmt.Println("Success search film")
			fmt.Println()

			response.Status = 200
			response.Message = "Success Search Film"
			response.Data = films
		} else {
			// Output to console
			fmt.Println("Film not found")
			fmt.Println()

			response.Status = 400
			response.Message = "Film not found"
		}
	} else {
		// Output to console
		fmt.Println("Search Film Failed\n" + result.Error())
		fmt.Println()

		response.Status = 400
		response.Message = "Search Film Failed " + result.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func WatchFilm(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	filmID, _ := strconv.Atoi(r.URL.Query()["film_id"][0])

	var film model.Film

	// Get email from logged member
	email := GetEmailFromToken(r)

	// Get user logged
	var user model.User
	db.Where("email = ?", email).First(&user)

	// Check expire date
	expDate := user.SubDate.AddDate(0, 0, 30)

	// Set response
	var response model.WatchResponse
	if expDate.After(time.Now()) {
		// Get film data
		err := db.First(&film, filmID).Error

		if err == nil {
			// Insert into user_histories
			var history model.History

			result := db.Raw("INSERT INTO histories (user_email, film_id, date) VALUES (?, ?, ?)", email, filmID, interface{}(time.Now())).Scan(&history)

			if result.Error == nil {
				// Output to console
				fmt.Println("Enjoy the movie")
				fmt.Println()

				response.Status = 200
				response.Message = "Enjoy the movie"
				response.Movie = film
				response.Date = time.Now()
			} else {
				// Output to console
				fmt.Println("Failed add to history\n" + result.Error.Error())

				response.Status = 400
				response.Message = "Failed add to history | " + result.Error.Error()
			}
		} else {
			// Output to console
			fmt.Println("Film not found\n" + err.Error())
			fmt.Println()

			response.Status = 400
			response.Message = "Film not found " + err.Error()
		}
	} else {
		response.Status = 400
		response.Message = "Subscribe was expired"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
