package controller

import "net/http"

func SetHeader(w http.ResponseWriter)http.ResponseWriter{
	//w.Header().Set("Authorization","Bearer c711e3310348accefd098bd7cc6d8104c961ae08")
	w.Header().Set("Content-Type","application/json")
	return w
}

func QiitaHandler(w http.ResponseWriter, r *http.Request) {

}

