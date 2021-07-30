package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/gowiki-api/models"
	"github.com/gowiki-api/storage"
	_ "github.com/joho/godotenv"
	"net/http"
	"time"
)

var Key = []byte(storage.GoDotEnvVariable("JWTKey"))
var CSRFKey = storage.GoDotEnvVariable("CSRFKey")

// CreateNewTokens
// Generate auth Token with csrf key
func CreateNewTokens(user *models.User) (authTokenString, csrfSecret string, err error) {
	csrfSecret = CSRFKey
	authTokenString, err = CreateAuthTokenString(csrfSecret, user)

	if err != nil {
		return
	}
	return
}

// SetCookies
// return a new cookie with auth Token inside
func SetCookies(w http.ResponseWriter, authTokenString string) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    authTokenString,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &authCookie)
}

// CreateAuthTokenString
// Create the auth token string with Jwt key and Csrf
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
	authTokenString, err = authJwt.SignedString(Key)
	return
}

// ClearSession
//Define the expires time for deleting the cookie
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
