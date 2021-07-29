package main

import (
	"github.com/gowiki-api/pkg/http/rest"
	"net/http"
)

func main() {
	router := Router()
	_ = http.ListenAndServe(":8080", router)
}
