package main

import (
	"net/http"
)

func main() {
	router := Router()
	_ = http.ListenAndServe(":8080", router)
}
