package controllers

import (
	"encoding/json"
	"github.com/gowiki-api/Tools"
	"github.com/gowiki-api/models"
	"github.com/gowiki-api/services"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// Struct for the request body
type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateUser(write http.ResponseWriter, request *http.Request) {

	user := &models.User{}
	json.NewDecoder(request.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		write.WriteHeader(http.StatusInternalServerError)
	}
	user.Password = string(pass)
	if err != nil {
		log.Fatal(err)
	}
	models.NewUser(user)

	write.WriteHeader(http.StatusCreated)
}

// Return Cookie with JWT string
func AuthUsers(write http.ResponseWriter, request *http.Request) {
	var creds Credentials
	err := json.NewDecoder(request.Body).Decode(&creds)
	Users := models.GetUserByEmail(creds.Email)

	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		return
	}

	PasswordIsOk := Tools.ComparePasswords(Users.Password, []byte(creds.Password))

	if !PasswordIsOk {
		write.WriteHeader(http.StatusUnauthorized)
		return
	}

	services.CreateToken(write, services.Credentials(creds))

	write.WriteHeader(http.StatusOK)
}

func Logout(write http.ResponseWriter, request *http.Request) {
	services.ClearSession(write)
	write.WriteHeader(http.StatusOK)
}
