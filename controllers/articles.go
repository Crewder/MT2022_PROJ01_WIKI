package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gowiki-api/models"
)

func ArticleCreate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json;json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var article models.Article

	err = json.Unmarshal(body, &article)

	models.NewArticle(&article)

}
