package Middleware

import (
	"github.com/gorilla/csrf"
	"net/http"
)

func CSRFMiddleware(next http.Handler) http.Handler {
	fn := http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		csrf.Protect(
			[]byte("Ceci est un secret et non un lapins"),
			csrf.RequestHeader("Authenticity-Token"),
			csrf.FieldName("authenticity_token"),
			csrf.CookieName("X-CSRF-TOKEN"),
		)
		// Call the next handler
		next.ServeHTTP(write, request)

	})
	return fn
}
