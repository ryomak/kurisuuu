package controller

import (
	"github.com/ryomak/kurisuuu/interactor"
	"github.com/ryomak/kurisuuu/core"
	"encoding/json"
	"net/http"
	"log"
)

func GetMovie(w http.ResponseWriter,r *http.Request){
	res,err := interactor.ResMovies()
	if err != nil {
		log.Println(err)
		w = core.SetHeader(w,http.StatusBadRequest)
		return
	}
	j ,_:= json.Marshal(res)
	w = core.SetHeader(w,http.StatusOK)
	w.Write(j)
}