package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/gowiki-api/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	_ = json.NewDecoder(r.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		CoreResponse(w, http.StatusInternalServerError, nil)
	}

	user.Password = string(pass)
	if err != nil {
		log.Fatal(err)
	}

	models.NewUser(user)
	CoreResponse(w, http.StatusCreated, nil)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	CoreResponse(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := chi.URLParam(r, "id")
	ID, err := strconv.ParseInt(userId, 0, 0)

	if err != nil {
		log.Fatal(err)
	}

	user := models.GetUserById(ID)
	CoreResponse(w, http.StatusOK, user)
}
