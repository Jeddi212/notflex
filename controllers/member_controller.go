package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/notflex/models"
)

// UpdateMember | Member Function - Update Logged-in Member Profile Data
func UpdateMember(w http.ResponseWriter, r *http.Request) {
	// Connect to database
	db := Connect()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// Get email from logged member
	email := GetEmailFromToken(r)

	// Set object data to hold user data
	var user models.User

	// Get user data first before updates
	db.Where("email=?", email).First(&user)

	// Only can set Name - BirthDate - Gender
	if name := r.Form.Get("name"); name != "" {
		user.Name = name
	}
	if birthdate := r.Form.Get("birthdate"); birthdate != "" {
		user.BirthDate = birthdate
	}
	if gender := r.Form.Get("gender"); gender != "" {
		user.Gender = gender
	}

	// Update user data to database
	result := db.Save(&user)

	// Set response
	var response models.BasicResponse
	if result.Error == nil {
		response.Status = 200
		response.Message = "Success Update User Data"
	} else {
		response.Status = 400
		response.Message = "Update Failed"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Subscribe | Member Function - Set Member Subscribe Status (Basic / Premium)
// User must add credit card to complete the process
func Subscribe(w http.ResponseWriter, r *http.Request) {
	// Connect to database
	db := Connect()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// Get email from logged member
	email := GetEmailFromToken(r)

	// Get current time
	now := time.Now()

	// Get subscribe type & credit card information from request
	subscribe := r.Form.Get("subscribe")
	cardNumber := r.Form.Get("cardNumber")
	exp := r.Form.Get("exp")
	cvc := r.Form.Get("cvc")

	// Set response
	var response models.SubscribeResponse

	// If credit card information not provided by member | return error response
	if cardNumber == "" || exp == "" || cvc == "" {
		response.Status = 400
		response.Message = "Subscribe Failed | Form Data Invalid"
	} else {
		// Subscribe type submitted by member must ( Basic || Premium )
		if subscribe == "Premium" || subscribe == "Basic" {
			// Update subscribe status to database
			result := db.Model(&models.User{}).Where("email = ?", email).Updates(map[string]interface{}{"subscribe": subscribe, "sub_date": now})

			// Set object model to hold credit card data
			var creditCard models.Credit

			// Set credit card data to object
			creditCard.CardNumber = cardNumber
			creditCard.Exp = exp
			creditCard.Cvc = cvc
			creditCard.UserID = email

			if result.Error == nil {
				// Add new credit card - Bind it to current user
				err := db.Save(&creditCard)
				if err.Error == nil {
					response.Status = 200
					response.Message = "Subscribe success"
					response.Type = subscribe
				} else {
					response.Status = 400
					response.Message = "Subscribe Failed | " + result.Error.Error()
				}
			} else {
				response.Status = 400
				response.Message = "Subscribe Failed | " + result.Error.Error()
			}
		} else {
			response.Status = 400
			response.Message = "Subscribe Failed | Invalid Subscribe Type (Basic | Premium)"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Unsubscribe | Member Function - Update Member Subscribe Status to NULL
func Unsubscribe(w http.ResponseWriter, r *http.Request) {
	// Connect to database
	db := Connect()

	// Get email from logged in user
	email := GetEmailFromToken(r)

	// Set object data to hold user data
	var user models.User

	// Set response
	var response models.UnsubscribeResponse

	// If user found in database | continue
	if err := db.Where("email = ?", email).First(&user).Error; err == nil {
		// If user has subscribed before | continue
		if user.Subscribe == "Basic" || user.Subscribe == "Premium" {
			// Update subscribe status to database
			result := db.Model(&models.User{}).Where("email = ?", email).
				Updates(map[string]interface{}{"subscribe": "", "sub_date": nil})

			if result.Error == nil {
				response.Status = 200
				response.Message = "We are so sad you unsubscribed :("
			} else {
				response.Status = 400
				response.Message = "Unsubscribe Failed!" + result.Error.Error()
			}
		} else {
			response.Status = 400
			response.Message = "You Haven't Subscribe or Have Unsubscribe before"
		}
	} else {
		response.Status = 400
		response.Message = "There is Something Wrong"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetAllHistories | Member Function - Show All Watch Histories From Logged-in Member
func GetAllHistories(w http.ResponseWriter, r *http.Request) {
	// Connect to database
	db := Connect()

	// Get email from logged member
	email := GetEmailFromToken(r)

	// Set object data to hold user data
	var result []models.Result

	// Get user watch histories from database
	db.Model(&models.History{}).
		Select("films.id, films.title, films.genre, films.year, films.director, films.actor, films.synopsis, histories.date").
		Joins("join films on films.id = histories.film_id").
		Joins("join users on histories.user_email = users.email").
		Where("histories.user_email = ?", email).
		Scan(&result)

	// Set response
	var response models.HistoryResponse
	if len(result) > 0 {
		response.Status = 200
		response.Message = "Success Get All Histories"
		response.History = result
	} else {
		response.Status = 204
		response.Message = "Get histories failed, data empty"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
