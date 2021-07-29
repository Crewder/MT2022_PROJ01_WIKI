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
	helpers.HandleError(http.StatusBadRequest, nil, err)

	Uintdata := claims["Uintdata"].(map[string]interface{})
	comment.UserId = uint(Uintdata["Id"].(float64))

	_, err = models.GetArticleById(int64(*comment.ArticleId))

	helpers.HandleError(http.StatusBadRequest, nil, !err)
	helpers.HandleError(http.StatusBadRequest, nil, !models.NewComment(comment))

	CoreResponse(w, http.StatusCreated, nil)
}

func GetCommentsByArticle(w http.ResponseWriter, r *http.Request) {
	articleId := chi.URLParam(r, "id")
	comments, result := models.GetAllCommentsByArticle(articleId)

	helpers.HandleError(http.StatusBadRequest, nil, result)
	helpers.HandleError(http.StatusBadRequest, nil, len(comments) == 0)

	CoreResponse(w, http.StatusOK, comments)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	var err error

	id := chi.URLParam(r, "id")
	comment, result := models.GetComment(id)

	helpers.HandleError(http.StatusBadRequest, nil, result)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &comment)

	helpers.HandleError(http.StatusBadRequest, err, false)

	CoreResponse(w, http.StatusNoContent, nil)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	comment, result := models.GetComment(id)

	helpers.HandleError(http.StatusBadRequest, nil, result)
	helpers.HandleError(http.StatusBadRequest, nil, !models.DeleteComment(comment))

	CoreResponse(w, http.StatusNoContent, nil)
}
