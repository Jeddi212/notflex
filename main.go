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
	db.AutoMigrate(&model.Admin{}, &model.Member{}, &model.Film{})

	router := mux.NewRouter()
	router.HandleFunc("/login", controller.CheckUserLogin).Methods("GET")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
