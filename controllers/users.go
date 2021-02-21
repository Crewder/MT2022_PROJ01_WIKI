package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gowiki-api/models"
	"net/http"
	"time"
)

// Struct to encode JWT
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var jwtKey = []byte("Ceci est un lapin et non un secret")

// Struct for the request body
type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Return Cookie with JWT string
func AuthUsers(write http.ResponseWriter, request *http.Request) {
	var creds Credentials
	err := json.NewDecoder(request.Body).Decode(&creds)
	Users, db := models.GetUserByEmail(creds.Email)

	if err != nil || db.RowsAffected != 1 {
		write.WriteHeader(http.StatusBadRequest)
		return
	}

	if Users.Password != creds.Password {
		write.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(15 * time.Minute)

	//Create the JWT Claims
	//Claims will be include in the payload token
	claims := &Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			//TODO Check createuser (Validation or not)
			NotBefore: time.Now().Unix(),
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Generate the JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		write.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(write, &http.Cookie{
		Name:     "AuthToken",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	})
}
