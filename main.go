package main

import (
	"github.com/zm-dev/yzm-client/handler"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	handler.CreateHTTPAPIHandler(r)
	handler.CreateWebHandler(r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
