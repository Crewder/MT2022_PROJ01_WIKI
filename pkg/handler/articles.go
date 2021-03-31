package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gowiki-api/pkg/models"
	"github.com/gowiki-api/pkg/tools"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {

	article := &models.Article{}
	_ = json.NewDecoder(r.Body).Decode(article)

	claims, error := tools.ExtractDataToken(w, r)
	if error {
		CoreResponse(w, http.StatusBadRequest, nil)
	}

	Uintdata := claims["Uintdata"].(map[string]interface{})
	article.UserId = uint(Uintdata["Id"].(float64))
	if !models.NewArticle(article) {
		CoreResponse(w, http.StatusBadRequest, nil)
	}
	CoreResponse(w, http.StatusCreated, nil)
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, error := models.GetAllArticles()
	if error {
		CoreResponse(w, http.StatusBadRequest, nil)
	}
	CoreResponse(w, http.StatusOK, articles)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	articleDetails, error := models.GetArticleBySlug(slug)
	if error || articleDetails.Title == "" || articleDetails.Content == "" {
		CoreResponse(w, http.StatusBadRequest, nil)
	}
	CoreResponse(w, http.StatusOK, articleDetails)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var err error

	slug := chi.URLParam(r, "slug")

	article, error := models.GetArticleBySlug(slug)
	if error {
		CoreResponse(w, http.StatusBadRequest, nil)
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
		CoreResponse(w, http.StatusBadRequest, nil)
	}

	newArticle, error := models.GetArticleBySlug(slug)
	if error {
		CoreResponse(w, http.StatusBadRequest, nil)
	}

	CoreResponse(w, http.StatusOK, newArticle)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	article, error := models.GetArticleBySlug(slug)
	if error {
		CoreResponse(w, http.StatusBadRequest, nil)
	}
	models.DeleteArticle(article)

	CoreResponse(w, http.StatusNoContent, article)
}
