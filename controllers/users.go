package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gowiki-api/models"
	"log"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var err error

	vars := mux.Vars(r)
	userId := vars["id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	userDetails := models.GetUserById(ID)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
