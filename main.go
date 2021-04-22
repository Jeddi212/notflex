package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/notflex/controllers"
	model "github.com/notflex/models"
)

func main() {
	db := controller.Connect()
	db.AutoMigrate(&model.User{}, &model.Film{})
	// member level 1
	// admin level 0
	router := mux.NewRouter()
	router.HandleFunc("/login", controller.CheckUserLogin).Methods("GET")
	router.HandleFunc("/logout", controller.Logout).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
