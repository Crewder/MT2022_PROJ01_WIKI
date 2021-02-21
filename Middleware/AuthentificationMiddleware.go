package Middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gowiki-api/Services"
	"net/http"
)

func AuthentificationMiddleware(next http.Handler) http.Handler {

	fn := http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {

		cookie, err := request.Cookie("AuthToken")

		if err != nil {
			if err == http.ErrNoCookie {
				write.WriteHeader(http.StatusUnauthorized)
				return
			}
			write.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get the JWT string from the cookie
		tokenString := cookie.Value
		claims := &Services.Claims{}

		// Parse the JWT string and store the result in claims
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return Services.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				write.WriteHeader(http.StatusUnauthorized)
				return
			}
			write.WriteHeader(http.StatusBadRequest)
			return
		}

		if !token.Valid {
			write.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Call the next handler
		next.ServeHTTP(write, request)
	})
	return fn
}
