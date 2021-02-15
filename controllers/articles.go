package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gowiki-api/models"
	"log"
	"net/http"
	"strconv"
)

func Main(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world", r.URL.Path[1:])
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
