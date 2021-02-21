package controllers

import (
	"encoding/json"
	"github.com/gowiki-api/Services"
	"github.com/gowiki-api/Tools"
	"github.com/gowiki-api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

var db *gorm.DB

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

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Fatal(err)
	}
	var article models.Article

	err = json.Unmarshal(body, &article)

	models.NewUser(user)
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

	PasswordIsOk := Tools.ComparePasswords(Users.Password, []byte(creds.Password))

	if !PasswordIsOk {
		write.WriteHeader(http.StatusUnauthorized)
		return
	}

	Services.CreateToken(write, Services.Credentials(creds))

	write.WriteHeader(http.StatusOK)
}
