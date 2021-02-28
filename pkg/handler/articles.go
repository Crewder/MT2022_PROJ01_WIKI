package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/gowiki-api/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var article models.Article
	err = json.Unmarshal(body, &article)
	models.NewArticle(&article)
	CoreResponse(w, http.StatusCreated, nil)
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	articles := models.GetAllArticles()
	CoreResponse(w, http.StatusOK, articles)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	articleDetails := models.GetArticleBySlug(slug)

	CoreResponse(w, http.StatusOK, articleDetails)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
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

	CoreResponse(w, http.StatusOK, newArticle)
}
