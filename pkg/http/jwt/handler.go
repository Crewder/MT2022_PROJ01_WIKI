package jwt

import (
	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv"
	"net/http"
	"time"
)

type Claims struct {
	CSRF string `json:"CSRF"`
	jwt.StandardClaims
}

// Struct for the request body
type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateNewTokens() (authTokenString, refreshTokenString, csrfSecret string, err error) {
	csrfSecret = CSRFKey
	authTokenString, err = CreateAuthTokenString(csrfSecret)
	refreshTokenString, err = createRefreshTokenString(csrfSecret)

	//todo ajout du refresh token dans la bdd

	if err != nil {
		return
	}
	return
}

func SetCookies(w http.ResponseWriter, authTokenString string, refreshTokenString string) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    authTokenString,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &authCookie)

	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshTokenString,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &refreshCookie)
}

func CreateAuthTokenString(csrfSecret string) (authTokenString string, err error) {
	expirationTime := time.Now().Add(20 * time.Minute).Unix()

	authClaims := &Claims{
		CSRF: csrfSecret,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: expirationTime,
			IssuedAt:  time.Now().Unix(),
		},
	}
	authJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)
	authTokenString, err = authJwt.SignedString(JwtKey)
	return
}

func createRefreshTokenString(csrfString string) (refreshTokenString string, err error) {
	expirationTime := time.Now().Add(72 * time.Hour).Unix()

	refreshClaims := &Claims{
		CSRF: csrfString,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: expirationTime,
			IssuedAt:  time.Now().Unix(),
		},
	}
	refreshJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err = refreshJwt.SignedString(JwtKey)
	return
}

func ClearSession(w http.ResponseWriter) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &authCookie)

	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &refreshCookie)
}

func getandrefreshtokens() {

	//si le cookie est valide
	//Si authToken valide
	//refresh l'expiration du authtoken
	//Sinon
	//si le refreshtoken est valid
	//si le authtoken est expiré
	//Update de l'expiration de l'authtoken
	//Update de l'expiration du refreshToken
	//Update du CSRF dans le refreshToken
	//sinon
	//OK
	//sinon Unauthorized
	//sinon Unauthorized

}

func UpdateAuthToken() {

	// return JWT + signature
}
func UpdateRefreshToken() {

	// return JWT + signature
}
