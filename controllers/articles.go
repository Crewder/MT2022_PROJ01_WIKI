package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gowiki-api/models"
)

func ArticleCreate(w http.ResponseWriter, r *http.Request) {

	// TODO: Récuperation ID User pour l'insertion dans la base
	// A CHECK changement au niveau des call a la bases

	w.Header().Set("content-type", "application/json;json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var article models.Article

	err = json.Unmarshal(body, &article)

	fmt.Fprint(w, &article)

	models.NewArticle(&article)

}
