package tools

import (
	"github.com/golang-jwt/jwt"
	"github.com/gowiki-api/helpers"
	"github.com/gowiki-api/storage"
	_ "github.com/joho/godotenv"
	"net/http"
)

var JwtKey = []byte(storage.GoDotEnvVariable("JWTKey"))

// ExtractDataToken
// Function that return the data inside the jwt cookie
func ExtractDataToken(r *http.Request) (jwt.MapClaims, error) {

	AuthCookie, err := r.Cookie("AuthToken")
	helpers.HandleError(http.StatusInternalServerError, err)

	jwtToken := AuthCookie.Value
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return JwtKey, err
	})

	helpers.HandleError(http.StatusInternalServerError, err)
	claims := token.Claims.(jwt.MapClaims)

	return claims, err
}
