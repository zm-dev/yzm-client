package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

func CreateWebHandler(r *mux.Router) {
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
}
