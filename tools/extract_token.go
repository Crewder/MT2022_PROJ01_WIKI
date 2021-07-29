package tools

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gowiki-api/wiki/storage"
	_ "github.com/joho/godotenv"
)

var JwtKey = []byte(storage.GoDotEnvVariable("JWTKey"))

func ExtractDataToken(w http.ResponseWriter, r *http.Request) (jwt.MapClaims, bool) {

	AuthCookie, authErr := r.Cookie("AuthToken")
	jwtToken := AuthCookie.Value
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, authErr
		}
		return JwtKey, nil
	})

	if err != nil {
		return nil, true
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims, false
}
