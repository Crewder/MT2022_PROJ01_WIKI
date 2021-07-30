package jwt

import (
	"encoding/json"
	"github.com/gowiki-api/controllers"
	"github.com/gowiki-api/helpers"
	"github.com/gowiki-api/models"
	"net/http"
)

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func AuthUsers(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	helpers.HandleError(http.StatusBadRequest, err)

	Users, err := models.GetUserByEmail(creds.Email)
	helpers.HandleError(http.StatusBadRequest, err)

	PasswordIsOk := models.PasswordIsValid(Users.Password, []byte(creds.Password))

	if !PasswordIsOk {
		controllers.CoreResponse(w, http.StatusUnauthorized, nil)
		return
	}

	authTokenString, csrfSecret, err := CreateNewTokens(Users)

	SetCookies(w, authTokenString)

	w.Header().Set("X-CSRF-Token", csrfSecret)
	controllers.CoreResponse(w, http.StatusOK, nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	ClearSession(w)
	controllers.CoreResponse(w, http.StatusOK, nil)
}
