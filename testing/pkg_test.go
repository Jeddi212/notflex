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

func TestLogin(t *testing.T) {
	router := mux.NewRouter()
	request, _ := http.NewRequest("POST", "/", nil)
	response := httptest.NewRecorder()
	router.HandleFunc("/login", c.CheckUserLogin).Methods("POST").GetHandler().ServeHTTP(response, request)
	t.Error(response.Body)
}

func TestGetUser(t *testing.T) {
	router := mux.NewRouter()

	//success
	// request, _ := http.NewRequest("GET", "/get-user?email=jeddi@gmail.com", nil)

	//no content
	request, _ := http.NewRequest("GET", "/get-user", nil)

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

func TestSearchHistories(t *testing.T) {
	router := mux.NewRouter()

	//success
	// request, _ := http.NewRequest("GET", "/search-histories?email=jedediah@gmail.com", nil)

	//no content
	request, _ := http.NewRequest("GET", "/search-histories", nil)

	response := httptest.NewRecorder()
	router.HandleFunc("/search-histories", c.GetAllHistories).Methods("GET").GetHandler().ServeHTTP(response, request)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var historyResponse models.HistoryResponse
	json.Unmarshal(body, &historyResponse)
	if historyResponse.Status == 204 {
		t.Error(historyResponse.Message)
	}
}

func TestGetFilmById(t *testing.T) {
	router := mux.NewRouter()

	//success
	// request, _ := http.NewRequest("GET", "/search-film-by-id/1", nil)

	//no content
	request, _ := http.NewRequest("GET", "/search-film-by-id/10", nil)

	response := httptest.NewRecorder()
	router.HandleFunc("/search-film-by-id", c.GetFilmById).Methods("GET").GetHandler().ServeHTTP(response, request)
	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var filmResponse models.FilmResponse
	// json.Unmarshal(body, &filmResponse)
	// if filmResponse.Status == 204 {
	// 	t.Error(filmResponse.Message)
	// }
	t.Error(response.Body)
}
