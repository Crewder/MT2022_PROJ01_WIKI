package jwt

import (
	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv"
	"net/http"
	"time"
)

type Claims struct {
	CSRF string `json:"CSRF"`
	jwt.StandardClaims
}

// Struct for the request body
type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateNewTokens() (authTokenString, csrfSecret string, err error) {
	csrfSecret = CSRFKey
	authTokenString, err = CreateAuthTokenString(csrfSecret)

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

func CreateAuthTokenString(csrfSecret string) (authTokenString string, err error) {
	expirationTime := time.Now().Add(20 * time.Minute).Unix()

	authClaims := &Claims{
		CSRF: csrfSecret,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: expirationTime,
			IssuedAt:  time.Now().Unix(),
		},
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
