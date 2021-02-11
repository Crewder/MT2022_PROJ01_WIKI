package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

// Fetch all User
func UsersIndex(write http.ResponseWriter, request *http.Request) {

	write.Header().Set("Content-type", "application/json;charset=UTF-8")
	write.WriteHeader(http.StatusOK)
	//TODO function allUser()
	// json.NewEncoder(write).Encode(Users.AllUser())
}

// Create a User
func CreateUsers(write http.ResponseWriter, request *http.Request) {

	write.Header().Set("Content-type", "application/json;charset=UTF-8")
	write.WriteHeader(http.StatusOK)
}

// Struct to encode JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//Create JWT Key
var jwtKey = []byte("Ceci est un lapin et non un secret")

// Create a struct for the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Seed
var Users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Auth the user
// Return Cookie with JWT string
func AuthUsers(write http.ResponseWriter, request *http.Request) {

	fmt.Println("test")

	var creds Credentials
	err := json.NewDecoder(request.Body).Decode(&creds)

	// Verify the structure of the body
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verify Password
	expectedPassword, ok := Users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		write.WriteHeader(http.StatusUnauthorized)
	}

	// Expiration Time Token
	expirationTime := time.Now().Add(5 * time.Minute)

	//Create the JWT Claims
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// declare the JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		write.WriteHeader(http.StatusInternalServerError)
		return
	}

	// define the cookie
	http.SetCookie(write, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
