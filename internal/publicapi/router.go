package publicapi

import (
	"fmt"
	"net/http"
	"url_shortener"
)

type Router struct {
	app url_shortener.PublicAPIService
}

func NewRouter(s url_shortener.PublicAPIService) *Router {
	return &Router{
		app: s,
	}
}

func (ro *Router) createUrlHandler() http.Handler {
	return http.HandlerFunc(createUrl)
}

func (ro *Router) getByCodeHandler() http.Handler {
	return http.HandlerFunc(getByCode)
}

func (ro *Router) getByIdHandler() http.Handler {
	return http.HandlerFunc(getById)
}

func createUrl(w http.ResponseWriter, r *http.Request) {
	s := NewApiService()
	err := s.CreateUrl(r.Context())
	if err != nil {
		fmt.Println(err)
	}
}

func getByCode(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
