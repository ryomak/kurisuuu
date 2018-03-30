package controller

import (
	"net/http"
	"github.com/ryomak/kurisuuu/core"
	"encoding/json"
	"github.com/ryomak/kurisuuu/interactor"
)



func GetQiita(w http.ResponseWriter,r *http.Request){
	res,err := interactor.ResQiitaArticles()
	if err != nil {
		w = core.SetHeader(w,http.StatusBadRequest)
		return
	}
	j ,_:= json.Marshal(res)
	w = core.SetHeader(w,http.StatusOK)
	w.Write(j)
}