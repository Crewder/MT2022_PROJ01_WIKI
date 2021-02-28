package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/gowiki-api/pkg/http/jwt"
	"github.com/gowiki-api/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

// Struct for the request body
type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := chi.URLParam(r, "id")

	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	userDetails := models.GetUserById(ID)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func AddUser(write http.ResponseWriter, request *http.Request) {
	user := &models.User{}
	json.NewDecoder(request.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		write.WriteHeader(http.StatusInternalServerError)
	}
	user.Password = string(pass)
	if err != nil {
		log.Fatal(err)
	}
	models.NewUser(user)

	write.WriteHeader(http.StatusCreated)
}

// Return Cookie with JWT string
func AuthUsers(write http.ResponseWriter, request *http.Request) {
	var creds Credentials
	err := json.NewDecoder(request.Body).Decode(&creds)
	Users := models.GetUserByEmail(creds.Email)

	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		return
	}

	PasswordIsOk := models.ComparePasswords(Users.Password, []byte(creds.Password))

	if !PasswordIsOk {
		write.WriteHeader(http.StatusUnauthorized)
		return
	}

	authTokenString, refreshTokenString, csrfSecret, err := jwt.CreateNewTokens()

	jwt.SetCookies(write, authTokenString, refreshTokenString)
	if err != nil {
		write.WriteHeader(http.StatusInternalServerError)
		return
	}

	write.Header().Set("X-CSRF-Token", csrfSecret)
	write.WriteHeader(http.StatusOK)
}

func Logout(write http.ResponseWriter, request *http.Request) {
	jwt.ClearSession(write)
	write.WriteHeader(http.StatusOK)
}
