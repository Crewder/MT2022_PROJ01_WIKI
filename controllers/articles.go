package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/gowiki-api/models"
	"io/ioutil"
	"log"
	"net/http"
)

func ArticleCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var article models.Article
	err = json.Unmarshal(body, &article)
	models.NewArticle(&article)
	coreResponse(w, http.StatusCreated, nil)
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	articles := models.GetAllArticles()
	coreResponse(w, http.StatusOK, articles)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	var err error

	slug := chi.URLParam(r, "slug")

	if err != nil {
		log.Fatal(err)
	}
	articleDetails := models.GetArticleBySlug(slug)

	coreResponse(w, http.StatusOK, articleDetails)
}

func ArticleUpdate(w http.ResponseWriter, r *http.Request) {
	var err error

	slug := chi.URLParam(r, "slug")

	article := models.GetArticleBySlug(slug)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &article)

	models.UpdateArticle(article)

	newArticle := models.GetArticleBySlug(slug)

	coreResponse(w, http.StatusOK, newArticle)
}
