package middleware

import (
	"context"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/cors"
	jwt2 "github.com/gowiki-api/pkg/auth/jwt"
	"github.com/gowiki-api/pkg/handler"
	"github.com/gowiki-api/pkg/models"
	"log"
	"net/http"
	"strings"
)

// Verify JWT Token validity and the CSRF Inside the JWt Token
// will return 401 if CSRF OR JWT is no valid
// Then Verify the Policy
// will return 403 if Policy
func AuthentificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//get the Cookie value
		AuthCookie, authErr := r.Cookie("AuthToken")
		if authErr != nil {
			if authErr == http.ErrNoCookie {
				handler.CoreResponse(w, http.StatusUnauthorized, nil)
				return
			}
			handler.CoreResponse(w, http.StatusBadRequest, nil)
			return
		} else {
			// Parse The token value and return the JWT key if everything is valid
			jwtToken := AuthCookie.Value
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, authErr
				}
				return jwt2.JwtKey, nil
			})

			//Fetch the data inside the token
			claims := token.Claims.(jwt.MapClaims)
			Stringdata := claims["Stringdata"].(map[string]interface{})
			Uintdata := claims["Uintdata"].(map[string]interface{})

			// And verify if the CSRF token on header is equals to CSRF inside the JWT
			actualCSRF := GetCsrfFromReq(r)
			expectedCSRF := Stringdata["CSRF"].(string)

			if actualCSRF != fmt.Sprintf(expectedCSRF) {
				handler.CoreResponse(w, http.StatusUnauthorized, nil)
			} else {
				//Jwt Validity verification
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					ctx := context.WithValue(r.Context(), "props", claims)
					method := r.Method
					path := r.URL.Path
					keys, ok := r.URL.Query()["key"]

					// fetching current role on JWT
					role := Stringdata["Role"].(string)
					userid := Uintdata["Id"].(int)

					if role == "" {
						role = "anonymous"
					}

					if ok {

						if role != "admin" {
							if method == "DELETE" || method == "PUT" {
								if strings.Contains(path, "Article ") {
									Article := models.GetArticleBySlug(keys[0])
									if Article.UserId != userid {
										handler.CoreResponse(w, http.StatusForbidden, nil)
									}
								}
								if strings.Contains(path, "Comment ") {
									Comment := models.GetComment(keys[0])

									if Comment.UserId != userid {
										handler.CoreResponse(w, http.StatusForbidden, nil)
									}
								}
							}
						}
					}

					//Create an enforcer with path for the policy in csv file and the model
					// We will verify with this enforcer if the action is allowed for this role
					e := casbin.NewEnforcer("pkg/auth/roles/auth_model.conf", "pkg/auth/roles/auth_policy.csv")

					if e.Enforce(role, path, method) {
						next.ServeHTTP(w, r.WithContext(ctx))
					} else {
						handler.CoreResponse(w, http.StatusForbidden, nil)
					}
				} else {
					handler.CoreResponse(w, http.StatusUnauthorized, nil)
					log.Fatal(err)
				}
			}
		}
	})
}

// Configure the CORS with default value
// Will Allow request from all Source
// return and serve a  context with the CorsOptionHandler
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		corsOptionsHandler := cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
		})
		ctx := context.WithValue(r.Context(), "cors", corsOptionsHandler)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// verify the  CSRF token on header of the request
func GetCsrfFromReq(r *http.Request) string {
	csrfFromForm := r.FormValue("X-CSRF-Token")
	if csrfFromForm != "" {
		return csrfFromForm
	} else {
		return r.Header.Get("X-CSRF-Token")
	}
}
