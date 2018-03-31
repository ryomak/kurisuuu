package router

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/ryomak/kurisuuu/controller"
	"github.com/urfave/negroni"
	"log"
	"net/http"
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

func (r *Router) Route(addr string) {
	r.HandleFunc("/", controller.Index)
	//r.HandleFunc("/movies",controller.MovieHandler)
	a := r.PathPrefix("/api").Subrouter()
	a.Path("/qiita").HandlerFunc(controller.GetQiita).Methods("GET")
	a.Path("/github").HandlerFunc(controller.GetGithub).Methods("GET")
	a.Path("/movie").HandlerFunc(controller.GetMovie).Methods("GET")
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(r)

	log.Println("server start :" + addr)
	http.ListenAndServe(":"+addr, n)
}
