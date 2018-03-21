package adapter

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"context"
	"github.com/ryomak/kurisuuu/adapter/controller"
)

type Router struct {
	*mux.Router
	ctx *context.Context
}

func New() *Router {
	return &Router{
		mux.NewRouter(),
		nil,
	}
}

func (r *Router) Route() {
	r.HandleFunc("/", controller.Index)
	r.HandleFunc("/movies",controller.MovieHandler)
	r.HandleFunc("/qiita",controller.QiitaHandler)
}

func (r *Router) Run(addr string) {
	log.Println("server start PORT :", addr)
	http.ListenAndServe(":"+addr,r)
}
