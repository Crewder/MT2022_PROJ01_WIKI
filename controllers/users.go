package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/gowiki-api/wiki/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	_ = json.NewDecoder(r.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		CoreResponse(w, http.StatusInternalServerError, nil)
		return
	}

	if len(user.Password) <= 4 {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	} else {
		user.Password = string(pass)
	}

	if err != nil {
		CoreResponse(w, http.StatusInternalServerError, nil)
		return
	}

	if !models.NewUser(user) {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	}

	CoreResponse(w, http.StatusCreated, nil)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, result := models.GetAllUsers()

	if result {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	}
	CoreResponse(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := chi.URLParam(r, "id")
	ID, err := strconv.ParseInt(userId, 0, 0)

	if err != nil {
		CoreResponse(w, http.StatusInternalServerError, nil)
		return
	}

	user, result := models.GetUserById(ID)
	if result {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	}
	CoreResponse(w, http.StatusOK, user)
}
