package main

import (
	"github.com/gowiki-api/Router"
	"net/http"
)

func main() {
	router := Router.Router()
	http.ListenAndServe(":8080", router)
}
