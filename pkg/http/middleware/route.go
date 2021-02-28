package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/cors"
	key "github.com/gowiki-api/pkg/http/jwt"
	"log"
	"net/http"
)

func AuthentificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {

		AuthCookie, authErr := request.Cookie("AuthToken")
		if authErr != nil {
			if authErr == http.ErrNoCookie {
				write.WriteHeader(http.StatusUnauthorized)
				return
			}
			write.WriteHeader(http.StatusBadRequest)
			return
		} else {
			jwtToken := AuthCookie.Value
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, authErr
				}
				return key.JwtKey, nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(request.Context(), "props", claims)
				next.ServeHTTP(write, request.WithContext(ctx))
			} else {
				write.WriteHeader(http.StatusUnauthorized)
				log.Fatal(err)
			}
		}
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		corshandler := cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
		})

		ctx := context.WithValue(request.Context(), "cors", corshandler)
		next.ServeHTTP(write, request.WithContext(ctx))
	})
}
