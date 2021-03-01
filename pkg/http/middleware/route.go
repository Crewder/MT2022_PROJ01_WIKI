package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/cors"
	"github.com/gowiki-api/pkg/handler"
	key "github.com/gowiki-api/pkg/http/jwt"
	"log"
	"net/http"
)

func AuthentificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AuthCookie, authErr := r.Cookie("AuthToken")
		if authErr != nil {
			if authErr == http.ErrNoCookie {
				handler.CoreResponse(w, http.StatusUnauthorized, nil)
				return
			}
			handler.CoreResponse(w, http.StatusBadRequest, nil)
			return
		} else {
			jwtToken := AuthCookie.Value
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, authErr
				}
				return key.JwtKey, nil
			})

			// CSRF Verification
			actualCSRF := GetCsrfFromReq(r)
			expectedCSRF := key.CSRFKey

			if actualCSRF != expectedCSRF {
				handler.CoreResponse(w, http.StatusForbidden, nil)
			} else {
				//TODO verification Refresh Token
				//handler.getandrefreshtokens

				//Jwt Validity verification
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					ctx := context.WithValue(r.Context(), "props", claims)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					handler.CoreResponse(w, http.StatusUnauthorized, nil)
					log.Fatal(err)
				}
			}
		}
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		corsHandler := cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
		})

		ctx := context.WithValue(r.Context(), "cors", corsHandler)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetCsrfFromReq(r *http.Request) string {
	csrfFromForm := r.FormValue("X-CSRF-Token")
	if csrfFromForm != "" {
		return csrfFromForm
	} else {
		return r.Header.Get("X-CSRF-Token")
	}
}
