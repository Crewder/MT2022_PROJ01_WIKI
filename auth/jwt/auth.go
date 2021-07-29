package jwt

import (
	"encoding/json"
	"github.com/gowiki-api/controllers"
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
