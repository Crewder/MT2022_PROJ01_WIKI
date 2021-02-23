package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gowiki-api/Tools"
	"net/http"
	"time"
)

//todo a set dans le .env
var JwtKey = Tools.GenerateARandomString()

// Struct to encode JWT
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Struct for the request body
type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateToken(write http.ResponseWriter, creds Credentials) (http.ResponseWriter, error) {
	var err error
	expirationTime := time.Now().Add(1 * time.Minute)

	claims := &Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	// Generate the JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		write.WriteHeader(http.StatusInternalServerError)
		return nil, err
	}

	http.SetCookie(write, &http.Cookie{
		Name:     "AuthToken",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	})

	return write, err
}

func RefreshToken(write http.ResponseWriter, request *http.Request) {

	ExtractCookieAndVerifyToken(write, request)

	claims := &Claims{}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		write.WriteHeader(http.StatusBadRequest)
		return
	}

	// set the new detail token
	//TODO Recuperation de l'email
	// claims.email := Cookie.EMAIL
	expirationTime := time.Now().Add(30 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().Unix()

	// Generate the JWT Token
	NewToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := NewToken.SignedString(JwtKey)

	if err != nil {
		write.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token
	http.SetCookie(write, &http.Cookie{
		Name:     "AuthToken",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	})
}

func ExtractCookieAndVerifyToken(write http.ResponseWriter, request *http.Request) (*jwt.Token, error) {
	c, err := request.Cookie("AuthToken")
	if err != nil {
		if err == http.ErrNoCookie {
			write.WriteHeader(http.StatusUnauthorized)
			return nil, nil
		}
		write.WriteHeader(http.StatusBadRequest)
		return nil, nil
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			write.WriteHeader(http.StatusUnauthorized)
			return nil, nil
		}
		write.WriteHeader(http.StatusBadRequest)
		return nil, nil
	}
	if !tkn.Valid {
		write.WriteHeader(http.StatusUnauthorized)
		return nil, nil
	}
	return tkn, err
}

func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "AuthToken",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
