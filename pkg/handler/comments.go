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
		CoreResponse(w, http.StatusNotAcceptable, nil)
	}
	CoreResponse(w, http.StatusCreated, nil)
}

func GetCommentsByArticle(w http.ResponseWriter, r *http.Request) {
	articleId := chi.URLParam(r, "id")
	comments := models.GetAllCommentsByArticle(articleId)
	CoreResponse(w, http.StatusOK, comments)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	comment := models.GetComment(id)
	models.DeleteComment(comment)

	CoreResponse(w, http.StatusNoContent, comment)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	var err error

	id := chi.URLParam(r, "id")

	comment := models.GetComment(id)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &comment)

	models.UpdateComment(comment)

	newComment := models.GetComment(id)

	CoreResponse(w, http.StatusOK, newComment)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	comment := models.GetComment(id)
	models.DeleteComment(comment)

	CoreResponse(w, http.StatusNoContent, comment)
}
