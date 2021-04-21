package main

import (
	"fmt"

	model "github.com/notflex/models"
)

func main() {
	fmt.Println("Notflex Asoy")
	var admin model.Admin
	// admin.SetAdmin("@gmail.com", "123")
	// admin.
	// fmt.Println(admin.GetPassword())
	// fmt.Println(admin.GetEmail())
	admin.User.SetUser("AA", "ASDF")
	fmt.Println(admin.GetEmail())
	fmt.Println(admin.GetPassword())
}
