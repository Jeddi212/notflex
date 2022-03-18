package testing

import (
	"testing"

	"github.com/gorilla/mux"
	c "github.com/notflex/controllers"
)

// test function
func TestReturnGeeks(t *testing.T) {
	actualString := c.ReturnGeeks()
	expectedString := "geeks"
	if actualString != expectedString {
		t.Errorf(
			"Expected String(%s) is not same as"+
				" actual string (%s)",
			expectedString,
			actualString,
		)
	}
}

// func TestGetUser(t *testing.T) {
// 	c.Authenticate(c.GetUser, 0)
// }

func TestLogin(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/login", c.CheckUserLogin).Methods("POST")
}
