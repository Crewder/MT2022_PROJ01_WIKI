package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gowiki-api/pkg/models"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var article models.Article
	err = json.Unmarshal(body, &article)
	if !models.NewArticle(&article) {
		CoreResponse(w, http.StatusInternalServerError, nil)
	}
	CoreResponse(w, http.StatusCreated, nil)
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, result := models.GetAllArticles()
	if !result {
		CoreResponse(w, http.StatusInternalServerError, nil)
	}
	CoreResponse(w, http.StatusOK, articles)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	articleDetails, result := models.GetArticleBySlug(slug)
	if !result {
		CoreResponse(w, http.StatusInternalServerError, nil)
	}
	CoreResponse(w, http.StatusOK, articleDetails)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var err error

	slug := chi.URLParam(r, "slug")

	article, result := models.GetArticleBySlug(slug)
	if !result {
		CoreResponse(w, http.StatusInternalServerError, nil)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &article)
	if err != nil {
		log.Fatal(err)
	}

	if !models.UpdateArticle(article) {
		CoreResponse(w, http.StatusInternalServerError, nil)
	}

	newArticle, result := models.GetArticleBySlug(slug)
	if !result {
		CoreResponse(w, http.StatusInternalServerError, nil)
	}

	CoreResponse(w, http.StatusOK, newArticle)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	article, result := models.GetArticleBySlug(slug)

	if !result {
		CoreResponse(w, http.StatusBadRequest, nil)
	}

	models.DeleteArticle(article)

	CoreResponse(w, http.StatusNoContent, article)
}
