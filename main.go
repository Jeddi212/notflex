package main

import (
	"fmt"
	"log"
	"net/http"

	controller "github.com/notflex/controllers"
	model "github.com/notflex/models"

	"github.com/gorilla/mux"
)

func main() {
	db := controller.Connect()
	db.AutoMigrate(&model.User{}, &model.Film{}, &model.Credit{})
	// member level 1
	// admin level 0
	router := mux.NewRouter()

	// User Related Request
	router.HandleFunc("/login", controller.CheckUserLogin).Methods("GET")
	router.HandleFunc("/logout", controller.Logout).Methods("GET")
	router.HandleFunc("/registration", controller.InsertMember).Methods("POST")
	router.HandleFunc("/get-user", controller.Authenticate(controller.GetUser, 0)).Methods("GET")
	router.HandleFunc("/suspend", controller.Authenticate(controller.SuspendUser, 0)).Methods("PUT")
	router.HandleFunc("/updateProfile", controller.Authenticate(controller.UpdateMember, 1)).Methods("PUT")

	// Film Related Request

	router.HandleFunc("/search-film-by-id/{id_film}", controller.Authenticate(controller.GetFilmById, 0)).Methods("GET")
	router.HandleFunc("/search-film-by-title/{title_film}", controller.Authenticate(controller.GetFilmByTitle, 0)).Methods("GET")
	router.HandleFunc("/search-film", controller.Authenticate(controller.SearchFilm, 1)).Methods("GET")
	router.HandleFunc("/add-film", controller.Authenticate(controller.AddFilm, 0)).Methods("POST")
	router.HandleFunc("/edit-film/{film_id}", controller.Authenticate(controller.EditFilm, 0)).Methods("PUT")

	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
