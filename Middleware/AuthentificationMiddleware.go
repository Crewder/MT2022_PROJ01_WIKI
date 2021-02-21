package Middleware

import (
	"github.com/gowiki-api/Services"
	"net/http"
)

func AuthentificationMiddleware(next http.Handler) http.Handler {

	fn := http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {

		Services.ExtractCookieAndVerifyToken(write, request)

		// Call the next handler

		next.ServeHTTP(write, request)
	})
	return fn
}
