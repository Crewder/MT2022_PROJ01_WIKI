package handler

import (
	"encoding/json"
	"net/http"
)

func coreResponse(w http.ResponseWriter, status int, array interface{}) {
	response, _ := json.Marshal(array)
	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	w.Write(response)
}
