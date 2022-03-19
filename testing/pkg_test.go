package testing

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	c "github.com/notflex/controllers"
	"github.com/notflex/models"
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

func TestGetUser(t *testing.T) {
	router := mux.NewRouter()

	//success
	// request, _ := http.NewRequest("GET", "/get-user?email=jeddi@gmail.com", nil)

	//no content
	request, _ := http.NewRequest("GET", "/get-user?email=lalala@gmail.com", nil)

	response := httptest.NewRecorder()
	router.HandleFunc("/get-user", c.GetUser).Methods("GET").GetHandler().ServeHTTP(response, request)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var userResponse models.UserResponse
	json.Unmarshal(body, &userResponse)
	if userResponse.Status == 204 {
		t.Error(userResponse.Message)
	}
}

func TestSearchFilm(t *testing.T) {
	router := mux.NewRouter()

	//success
	request, _ := http.NewRequest("GET", "/search-film?year=2006", nil)

	//no content
	// request, _ := http.NewRequest("GET", "/search-film?year=2022", nil)

	response := httptest.NewRecorder()
	router.HandleFunc("/search-film", c.SearchFilm).Methods("GET").GetHandler().ServeHTTP(response, request)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var filmResponse models.FilmResponse
	json.Unmarshal(body, &filmResponse)
	if filmResponse.Status == 400 {
		t.Error(filmResponse.Message)
	}
}
