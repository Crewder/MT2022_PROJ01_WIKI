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
	refreshTokenString, err = createRefreshTokenString(csrfSecret)
	authTokenString, err = createAuthTokenString(csrfSecret)

	if err != nil {
		return
	}
	return
}

func SetCookies(write http.ResponseWriter, authTokenString string, refreshTokenString string) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    authTokenString,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(write, &authCookie)

	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshTokenString,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(write, &refreshCookie)
}

func createAuthTokenString(csrfSecret string) (authTokenString string, err error) {
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

func ClearSession(write http.ResponseWriter) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(write, &authCookie)

	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(write, &refreshCookie)
}

func getandrefreshtokens() {
	// Verifier la validité de mon cookie
	// si c'est valid : Return
	//Si pas valid
	// recupere le refresh dans ma bdd
	// si c'est valide je refresh le AuthToken
	// Sinon return 401
}

func GetCsrfFromReq(r *http.Request) string {
	csrfFromForm := r.FormValue("X-CSRF-Token")
	if csrfFromForm != "" {
		return csrfFromForm
	} else {
		return r.Header.Get("X-CSRF-Token")
	}
}
