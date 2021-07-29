package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gowiki-api/wiki/models"
	"github.com/gowiki-api/wiki/tools"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	comment := &models.Comment{}
	_ = json.NewDecoder(r.Body).Decode(comment)

	claims, error := tools.ExtractDataToken(w, r)
	if error {
		CoreResponse(w, http.StatusBadRequest, nil)
	}

	Uintdata := claims["Uintdata"].(map[string]interface{})
	comment.UserId = uint(Uintdata["Id"].(float64))

	_, err := models.GetArticleById(int64(*comment.ArticleId))

	if !err {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	}

	if !models.NewComment(comment) {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	}
	CoreResponse(w, http.StatusCreated, nil)
}

func GetCommentsByArticle(w http.ResponseWriter, r *http.Request) {
	articleId := chi.URLParam(r, "id")
	comments, result := models.GetAllCommentsByArticle(articleId)
	if result || len(comments) == 0 {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	}
	CoreResponse(w, http.StatusOK, comments)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	var err error

	id := chi.URLParam(r, "id")
	comment, result := models.GetComment(id)
	if result {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &comment)
	if !models.UpdateComment(comment) {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	}

	CoreResponse(w, http.StatusNoContent, nil)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	comment, result := models.GetComment(id)
	if result {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	}
	if !models.DeleteComment(comment) {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	}
	CoreResponse(w, http.StatusNoContent, nil)
}
