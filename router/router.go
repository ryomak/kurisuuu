package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"context"
	"github.com/ryomak/kurisuuu/controller"
	"github.com/urfave/negroni"
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
	http.ListenAndServe(":"+addr,n)
}
