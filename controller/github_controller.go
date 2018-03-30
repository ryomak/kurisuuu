package controller

import (
"net/http"
"github.com/ryomak/kurisuuu/core"
"encoding/json"
"github.com/ryomak/kurisuuu/interactor"
	"log"
)


func GetGithub(w http.ResponseWriter,r *http.Request){
	res,err := interactor.ResGithubRepository()
	if err != nil {
		log.Println(err)
		w = core.SetHeader(w,http.StatusBadRequest)
		return
	}
	j ,_:= json.Marshal(res)
	w = core.SetHeader(w,http.StatusOK)
	w.Write(j)
}
