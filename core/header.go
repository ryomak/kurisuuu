package core

import (
	"net/http"
)

func SetHeader(w http.ResponseWriter, status int) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return w
}
