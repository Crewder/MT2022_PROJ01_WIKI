package jwt

import (
	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv"
	"net/http"
	"time"
)

// Struct to encode JWT
type Claims struct {
	Email string `json:"email"`
	CSRF  []byte `json:"X-CSRF-TOKEN"`
	jwt.StandardClaims
}

// Struct for the request body
type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateToken(write http.ResponseWriter, request *http.Request, creds Credentials) error {
	var err error
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Email: creds.Email,
		CSRF:  CSRFKey,
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
		print(err)
		write.WriteHeader(http.StatusInternalServerError)
		return err
	}

	http.SetCookie(write, &http.Cookie{
		Name:     "AuthToken",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		Path:     "/",
	})

	return err
}

func RefreshToken(write http.ResponseWriter, request *http.Request) {

	var creds Credentials
	claims := &Claims{}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		write.WriteHeader(http.StatusBadRequest)
		return
	}
	expirationTime := time.Now().Add(30 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	email := creds.Email
	claims.Email = email
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
		Path:     "/",
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
