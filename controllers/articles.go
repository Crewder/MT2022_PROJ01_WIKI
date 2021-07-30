package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/gowiki-api/helpers"
	"github.com/gowiki-api/models"
	"github.com/gowiki-api/tools"
	"io/ioutil"
	"net/http"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {

	article := &models.Article{}
	_ = json.NewDecoder(r.Body).Decode(article)

	claims, err := tools.ExtractDataToken(r)
	helpers.HandleError(http.StatusBadRequest, err)

	Uintdata := claims["Uintdata"].(map[string]interface{})
	article.UserId = uint(Uintdata["Id"].(float64))

	_, err = models.NewArticle(article)
	helpers.HandleError(http.StatusBadRequest, err)

	CoreResponse(w, http.StatusCreated, nil)
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := models.GetAllArticles()
	helpers.HandleError(http.StatusBadRequest, err)
	CoreResponse(w, http.StatusOK, articles)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	articleDetails, err := models.GetArticleBySlug(slug)
	helpers.HandleError(http.StatusBadRequest, err)

	CoreResponse(w, http.StatusOK, articleDetails)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var err error
	slug := chi.URLParam(r, "slug")

	article, err := models.GetArticleBySlug(slug)
	helpers.HandleError(http.StatusBadRequest, err)

	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleError(http.StatusBadRequest, err)

	err = json.Unmarshal(body, &article)
	helpers.HandleError(http.StatusBadRequest, err)

	_, err = models.UpdateArticle(article)
	helpers.HandleError(http.StatusBadRequest, err)

	CoreResponse(w, http.StatusNoContent, nil)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	article, err := models.GetArticleBySlug(slug)
	helpers.HandleError(http.StatusBadRequest, err)

	_, err = models.DeleteArticle(article)
	helpers.HandleError(http.StatusBadRequest, err)

	CoreResponse(w, http.StatusNoContent, nil)
}
