package main

import (
	"github.com/gowiki-api/pkg/http/rest"
	"net/http"
)

func main() {
	router := rest.Router()
	_ = http.ListenAndServe(":8080", router)
}
