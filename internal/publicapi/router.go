package publicapi

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"url_shortener"
)

type Router struct {
	app    url_shortener.PublicAPIService
	router *mux.Router
}

func NewRouter(s url_shortener.PublicAPIService) *Router {
	r := mux.NewRouter()
	return &Router{
		app:    s,
		router: r,
	}
}

func (ro *Router) InitializeRoutes() {
	ro.router.HandleFunc("/create", ro.createUrl).Methods("POST")
	ro.router.HandleFunc("/url/{id:[0-9]+}", ro.getById).Methods("GET")
	//ro.router.HandleFunc("/url/{code:[A-Za-z0-9]+", ro.getByCode).Methods("GET")
}

func (ro *Router) Run(addr string) {
	if addr == "" {
		addr = ":8000"
	}
	log.Fatal(http.ListenAndServe(addr, ro.router))
}

func (ro *Router) createUrl(w http.ResponseWriter, r *http.Request) {
	var u url_shortener.Url
	err := json.NewDecoder(r.Body).Decode(&u)
	id, err := ro.app.CreateUrl(r.Context(), u)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(id)))
}

func (ro *Router) getById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	u, err := ro.app.GetById(r.Context(), id)
	if err != nil {
		fmt.Println(err)
	}
	mu, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(mu)
}

//func (ro *Router) getByCode(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	code, _ := vars["code"]
//	u, err := ro.app.GetByCode(r.Context(), code)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("Here is URL:", u)
//
//}
