package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gowiki-api/pkg/models"
	"github.com/gowiki-api/pkg/storage"
	_ "github.com/joho/godotenv"
	"net/http"
	"time"
)

var JwtKey = []byte(storage.GoDotEnvVariable("JWTKey"))
var CSRFKey = storage.GoDotEnvVariable("CSRFKey")

func CreateNewTokens(user *models.User) (authTokenString, csrfSecret string, err error) {
	csrfSecret = CSRFKey
	authTokenString, err = CreateAuthTokenString(csrfSecret, user)

	if err != nil {
		return
	}
	return
}

func SetCookies(w http.ResponseWriter, authTokenString string) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    authTokenString,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &authCookie)
}

func CreateAuthTokenString(csrfSecret string, user *models.User) (authTokenString string, err error) {
	expirationTime := time.Now().Add(20 * time.Minute).Unix()

	authClaims := &jwt.MapClaims{
		"Stringdata": map[string]string{
			"Role": user.Role,
			"CSRF": csrfSecret,
		},
		"Uintdata": map[string]uint{
			"Id": user.ID,
		},

		"NotBefore": time.Now().Unix(),
		"ExpiresAt": expirationTime,
		"IssuedAt":  time.Now().Unix(),
	}
	authJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)
	authTokenString, err = authJwt.SignedString(JwtKey)
	return
}

func ClearSession(w http.ResponseWriter) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &authCookie)
}
