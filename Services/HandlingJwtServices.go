package Services

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

//todo a set dans le .env
var JwtKey = []byte("Ceci est un lapin et non un secret")

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
	expirationTime := time.Now().Add(5 * time.Minute)

	//Claims = TokenDetails
	claims := &Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			//TODO Check createuser (Validation or not)
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

	//Todo recuperation du claims du middleware
	claims := &Claims{}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		write.WriteHeader(http.StatusBadRequest)
		return
	}

	//todo refacto Detailtoken ( refresh / create )

	// set the new detail token
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().Unix()

	// Generate the JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

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
