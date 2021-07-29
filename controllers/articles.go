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

	claims, error := tools.ExtractDataToken(w, r)
	if error {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	}

	Uintdata := claims["Uintdata"].(map[string]interface{})
	article.UserId = uint(Uintdata["Id"].(float64))

	helpers.HandleError(http.StatusBadRequest, nil, !models.NewArticle(article))

	CoreResponse(w, http.StatusCreated, nil)
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, error := models.GetAllArticles()
	if error {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	}
	CoreResponse(w, http.StatusOK, articles)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	articleDetails, error := models.GetArticleBySlug(slug)
	if error {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	}
	CoreResponse(w, http.StatusOK, articleDetails)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var err error
	slug := chi.URLParam(r, "slug")

	article, error := models.GetArticleBySlug(slug)
	if error {
		CoreResponse(w, http.StatusBadRequest, nil)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleError(http.StatusBadRequest, err, false)

	err = json.Unmarshal(body, &article)
	helpers.HandleError(http.StatusBadRequest, err, false)
	helpers.HandleError(http.StatusBadRequest, nil, !models.UpdateArticle(article))

	CoreResponse(w, http.StatusNoContent, nil)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	article, error := models.GetArticleBySlug(slug)
	if error {
		CoreResponse(w, http.StatusBadRequest, nil)
	}
	models.DeleteArticle(article)

	CoreResponse(w, http.StatusNoContent, nil)
}
