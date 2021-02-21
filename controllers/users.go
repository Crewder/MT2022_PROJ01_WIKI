package controllers

import (
	"encoding/json"
	"github.com/gowiki-api/Services"
	"github.com/gowiki-api/models"
	"net/http"
)

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

	Services.CreateToken(write, Services.Credentials(creds))

	write.WriteHeader(http.StatusOK)
}
