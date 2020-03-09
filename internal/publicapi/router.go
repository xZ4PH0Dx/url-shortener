package publicapi

import (
	"net/http"
	"url_shortener"
)

type Router struct {
	app url_shortener.PublicAPIServer
}

func NewRouter(s url_shortener.PublicAPIServer) *Router {
	return &Router{app: s}
}

func (ro *Router) Handler() http.Handler {
	//r := mux.NewRouter()
	//q := r.HandleFunc("/url", ById).Methods("POST")
	return ById
}

func ById(w http.ResponseWriter, r *http.Request) {
}
