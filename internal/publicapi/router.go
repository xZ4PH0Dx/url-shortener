package publicapi

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xZ4PH0Dx/url_shortener"
	"net/http"
	"strconv"
)

type Router struct {
	app url_shortener.PublicAPIService
}

func NewRouter(s url_shortener.PublicAPIService) *Router {
	return &Router{
		app: s,
	}
}

func (ro *Router) Handler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/urls", ro.createUrlHandler).Methods("POST")
	r.HandleFunc("/urls/{id:[0-9]+}", ro.getByIdHandler).Methods("GET")
	//ro.router.HandleFunc("/urls/{code:[A-Za-z0-9]+", ro.getByCode).Methods("GET")
	return r
}

//func (ro *Router) InitializeRoutes() {
//}

func (ro *Router) createUrlHandler(w http.ResponseWriter, r *http.Request) {
	var u url_shortener.Url
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println(err)
	}
	actualU, err := ro.app.CreateUrl(r.Context(), u)
	if err != nil {
		_ = encodeErrorResp(err, w)
	}
	err = encodeJSONResponse(w, actualU)
	if err != nil {
		_ = encodeErrorResp(err, w)
	}
}

func (ro *Router) getByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		_ = encodeErrorResp(err, w)
	}
	u, err := ro.app.GetById(r.Context(), id)
	if err != nil {
		_ = encodeErrorResp(err, w)
	}
	//mu, err := json.Marshal(u)
	//if err != nil {
	//	encodeErrorResp(err, w)
	//}
	//w.WriteHeader(http.StatusOK)
	//w.Write(mu)
	err = encodeJSONResponse(w, u)
	if err != nil {
		_, _ = w.Write([]byte(encodeErrorResp(err, w).Error()))
	}
}

//func (ro *Router) getByCode(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	code, _ := vars["code"]
//	u, err := ro.app.GetByCode(r.Context(), code)
//	if err != nil {
//	encodeErrorResp(err, w)
//	}
//	fmt.Println("Here is URL:", u)
//
//}

func encodeErrorResp(err error, w http.ResponseWriter) error {
	w.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(w).Encode(err)
}

func encodeJSONResponse(w http.ResponseWriter, resp interface{}) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(resp)
}
