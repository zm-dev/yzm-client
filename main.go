package main

import (
	"github.com/zm-dev/yzm-client/src/backend/handler"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	handler.CreateWebHandler(r)
	handler.CreateHTTPAPIHandler(r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
