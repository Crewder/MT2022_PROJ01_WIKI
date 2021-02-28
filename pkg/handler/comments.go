package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/gowiki-api/pkg/models"
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

func GetCommentsByArticle(w http.ResponseWriter, r *http.Request) {
	articleId := chi.URLParam(r, "id")
	comments := models.GetAllCommentsByArticle(articleId)
	coreResponse(w, http.StatusOK, comments)
}
