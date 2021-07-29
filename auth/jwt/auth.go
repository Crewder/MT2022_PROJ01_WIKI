package jwt

import (
	"encoding/json"
	"net/http"

	"github.com/gowiki-api/wiki/controller"
	"github.com/gowiki-api/wiki/models"
)

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func AuthUsers(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	Users := models.GetUserByEmail(creds.Email)

	if err != nil {
		controllers.CoreResponse(w, http.StatusBadRequest, nil)
		return
	}

	PasswordIsOk := models.PasswordIsValid(Users.Password, []byte(creds.Password))

	if !PasswordIsOk {
		controllers.CoreResponse(w, http.StatusUnauthorized, nil)
		return
	}

	authTokenString, csrfSecret, err := CreateNewTokens(Users)

	SetCookies(w, authTokenString)
	if err != nil {
		controllers.CoreResponse(w, http.StatusInternalServerError, nil)
		return
	}

	w.Header().Set("X-CSRF-Token", csrfSecret)
	controllers.CoreResponse(w, http.StatusOK, nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	ClearSession(w)
	controllers.CoreResponse(w, http.StatusOK, nil)
}
