package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)

type router struct {
	router *mux.Router
}

type Router interface {
	Router() *mux.Router
	RegisterRoute(pattern string, cb handler, method string)
}

func New() Router {
	a := &router{}

	r := mux.NewRouter()
	a.router = r
	return a
}

func (r *router) Router() *mux.Router {
	return r.router
}

func Logger(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request from <= " + r.RemoteAddr + " to => " + r.URL.String() + " Method: " + r.Method)
		f(w, r)
	}
}

func (r *router) RegisterRoute(pattern string, cb handler, method string) {
	r.router.HandleFunc(pattern, Logger(cb)).Methods(method)
}
