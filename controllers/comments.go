package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/gowiki-api/helpers"
	"github.com/gowiki-api/models"
	"github.com/gowiki-api/tools"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	comment := &models.Comment{}
	_ = json.NewDecoder(r.Body).Decode(comment)

	claims, err := tools.ExtractDataToken(r)
	helpers.HandleError(http.StatusBadRequest, err)

	Uintdata := claims["Uintdata"].(map[string]interface{})
	comment.UserId = uint(Uintdata["Id"].(float64))

	_, err = models.GetArticleById(int64(*comment.ArticleId))
	helpers.HandleError(http.StatusBadRequest, err)

	_, err = models.NewComment(comment)
	helpers.HandleError(http.StatusBadRequest, err)

	CoreResponse(w, http.StatusCreated, nil)
}

func GetCommentsByArticle(w http.ResponseWriter, r *http.Request) {
	articleId := chi.URLParam(r, "id")
	comments, err := models.GetAllCommentsByArticle(articleId)

	if err != nil || len(comments) == 0 {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	}

	CoreResponse(w, http.StatusOK, comments)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	var err error

	id := chi.URLParam(r, "id")
	comment, err := models.GetComment(id)

	helpers.HandleError(http.StatusBadRequest, err)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &comment)

	helpers.HandleError(http.StatusBadRequest, err)

	CoreResponse(w, http.StatusNoContent, nil)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	comment, err := models.GetComment(id)
	helpers.HandleError(http.StatusBadRequest, err)

	_, err = models.DeleteComment(comment)
	helpers.HandleError(http.StatusBadRequest, err)

	CoreResponse(w, http.StatusNoContent, nil)
}
