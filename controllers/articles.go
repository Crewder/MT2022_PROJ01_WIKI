package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gowiki-api/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

	vars := mux.Vars(r)
	articleId := vars["id"]
	ID, err := strconv.ParseInt(articleId, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	articleDetails := models.GetArticleById(ID)

	coreResponse(w, http.StatusOK, articleDetails)
}
