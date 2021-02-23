package middleware

import (
	"github.com/gorilla/csrf"
	"github.com/gowiki-api/Tools"
	"net/http"
)

func CSRFMiddleware(next http.Handler) http.Handler {
	fn := http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		csrf.Protect(
			Tools.GenerateARandomString(),
			csrf.RequestHeader("Authenticity-Token"),
			csrf.FieldName("authenticity_token"),
			csrf.CookieName("X-CSRF-Token"),
		)
		// Call the next handler
		next.ServeHTTP(write, request)

	})
	return fn
}
