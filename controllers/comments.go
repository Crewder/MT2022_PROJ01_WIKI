package controllers

import (
	"encoding/json"
	"github.com/gowiki-api/models"
	"io/ioutil"
	"log"
	"net/http"
)

func CommentCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var comment models.Comment
	err = json.Unmarshal(body, &comment)
	models.NewComment(&comment)
	coreResponse(w, http.StatusCreated, nil)
}
