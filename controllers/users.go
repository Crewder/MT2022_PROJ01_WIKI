package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/gowiki-api/helpers"
	"github.com/gowiki-api/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	_ = json.NewDecoder(r.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	helpers.HandleError(http.StatusBadRequest, err)

	if len(user.Password) <= 4 {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	} else {
		user.Password = string(pass)
	}

	_, err = models.NewUser(user)
	helpers.HandleError(http.StatusBadRequest, err)

	CoreResponse(w, http.StatusCreated, nil)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()
	helpers.HandleError(http.StatusBadRequest, err)

	CoreResponse(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := chi.URLParam(r, "id")
	ID, err := strconv.ParseInt(userId, 0, 0)
	helpers.HandleError(http.StatusInternalServerError, err)

	user, err := models.GetUserById(ID)
	helpers.HandleError(http.StatusInternalServerError, err)

	CoreResponse(w, http.StatusOK, user)
}
