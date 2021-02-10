package controllers

import (
	"fmt"
	"net/http"
)

func ArticleCreate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json;json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello world", r.URL.Path[1:])

}

func Main(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world", r.URL.Path[1:])
}
