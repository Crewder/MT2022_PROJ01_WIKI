package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/gowiki-api/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var comment models.Comment
	err = json.Unmarshal(body, &comment)
	if !models.NewComment(&comment) {
		CoreResponse(w, http.StatusBadRequest, nil)
	}
	CoreResponse(w, http.StatusCreated, nil)
}

func GetCommentsByArticle(w http.ResponseWriter, r *http.Request) {
	articleId := chi.URLParam(r, "id")
	comments, result := models.GetAllCommentsByArticle(articleId)
	if result {
		CoreResponse(w, http.StatusBadRequest, nil)
	}
	CoreResponse(w, http.StatusOK, comments)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	var err error

	id := chi.URLParam(r, "id")
	comment, result := models.GetComment(id)
	if result {
		CoreResponse(w, http.StatusBadRequest, nil)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &comment)
	if !models.UpdateComment(comment) {
		CoreResponse(w, http.StatusBadRequest, nil)
	}
	newComment, result := models.GetComment(id)
	if result {
		CoreResponse(w, http.StatusBadRequest, nil)
	}
	CoreResponse(w, http.StatusOK, newComment)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	comment, result := models.GetComment(id)
	if result {
		CoreResponse(w, http.StatusBadRequest, nil)
	}
	if !models.DeleteComment(comment) {
		CoreResponse(w, http.StatusBadRequest, nil)
	}
	CoreResponse(w, http.StatusNoContent, comment)
}
