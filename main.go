package main

import (
	"github.com/gowiki-api/router"
	"net/http"
)

func main() {
	router := router.Router()
	http.ListenAndServe(":8080", router)
}
