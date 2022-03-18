package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// var jwtKey = []byte("7233345")
var jwtKey = []byte("bebasapasaja")
var tokenName = "token"

type Claims struct {
	Email    string `json:"email"`
	UserType int    `json:"user_type"`
	jwt.StandardClaims
}

func generateToken(w http.ResponseWriter, email string, userType int) {
	tokenExpiryTime := time.Now().Add(20 * time.Minute)

	claims := &Claims{
		Email:    email,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: true,
	})
}

func resetUserToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	})
}

func Authenticate(next http.HandlerFunc, accessType int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidToken := validateUserToken(w, r, accessType)
		if !isValidToken {
			sendUnauthorizedResponse(w)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func validateUserToken(w http.ResponseWriter, r *http.Request, accessType int) bool {
	isAccessTokenValid, _, userType := validateTokenFromCookies(r)

	if isAccessTokenValid {
		isUserValid := userType == accessType
		if isUserValid {
			return true
		}
	}
	return false
}

func validateTokenFromCookies(r *http.Request) (bool, string, int) {
	if cookie, err := r.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.Email, accessClaims.UserType
		}
	}
	return false, "", -1
}

func GetEmailFromToken(r *http.Request) string {
	if cookie, err := r.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return accessClaims.Email
		}
	}
	return ""
}
