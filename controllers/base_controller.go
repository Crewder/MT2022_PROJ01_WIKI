package controllers

import (
	"encoding/json"
	"net/http"
)

func CoreResponse(w http.ResponseWriter, status int, array interface{}) {
	response, _ := json.Marshal(array)
	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	w.Write(response)
}
