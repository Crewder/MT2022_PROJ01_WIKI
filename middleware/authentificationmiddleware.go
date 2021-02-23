package middleware

import (
	"github.com/gowiki-api/services"
	"net/http"
)

func AuthentificationMiddleware(next http.Handler) http.Handler {

	fn := http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {

		services.ExtractCookieAndVerifyToken(write, request)

		// Call the next handler
		next.ServeHTTP(write, request)
	})
	return fn
}
