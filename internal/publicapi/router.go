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
	r.HandleFunc("/urls", ro.createURLHandler).Methods("POST")
	r.HandleFunc("/urls/{id:[0-9]+}", ro.getByIDHandler).Methods("GET")
	r.HandleFunc("/urls/{code:[A-Za-z0-9]+}", ro.getByCodeHandler).Methods("GET")

	return r
}

func (ro *Router) createURLHandler(w http.ResponseWriter, r *http.Request) {
	var u url_shortener.URL

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println(err)
	}

	actualU, err := ro.app.CreateURL(r.Context(), u)
	if err != nil {
		_ = encodeErrorResp(err, w)
		return
	}

	err = encodeJSONResponse(w, actualU)
	if err != nil {
		_ = encodeErrorResp(err, w)
		return
	}
}

func (ro *Router) getByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		_ = encodeErrorResp(err, w)
		return
	}

	u, err := ro.app.GetByID(r.Context(), id)
	if err != nil {
		_ = encodeErrorResp(err, w)
		return
	}

	err = encodeJSONResponse(w, u)
	if err != nil {
		_ = encodeErrorResp(err, w)
		return
	}
}

func (ro *Router) getByCodeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	u, err := ro.app.GetByCode(r.Context(), code)
	if err != nil {
		_ = encodeErrorResp(err, w)
		return
	}

	err = encodeJSONResponse(w, u)
	if err != nil {
		_ = encodeErrorResp(err, w)
		return
	}
}

func encodeErrorResp(err error, w http.ResponseWriter) error {
	type errStruct struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	errResp := struct {
		errStruct `json:"error"`
	}{
		errStruct{
			Code:    "error number one:)",
			Message: err.Error(),
		},
	}

	w.WriteHeader(http.StatusBadRequest)

	return json.NewEncoder(w).Encode(errResp)
}

func encodeJSONResponse(w http.ResponseWriter, resp interface{}) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(resp)
}
