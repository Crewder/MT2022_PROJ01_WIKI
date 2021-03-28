package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gowiki-api/pkg/storage"
	_ "github.com/joho/godotenv"
	"net/http"
	"time"
)

var JwtKey = []byte(storage.GoDotEnvVariable("JWTKey"))
var CSRFKey = storage.GoDotEnvVariable("CSRFKey")

func CreateNewTokens(role string) (authTokenString, csrfSecret string, err error) {
	csrfSecret = CSRFKey
	authTokenString, err = CreateAuthTokenString(csrfSecret, role)

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

func CreateAuthTokenString(csrfSecret string, role string) (authTokenString string, err error) {
	expirationTime := time.Now().Add(20 * time.Minute).Unix()

	authClaims := &jwt.MapClaims{
		"data": map[string]string{
			"Role": role,
			"CSRF": csrfSecret,
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
