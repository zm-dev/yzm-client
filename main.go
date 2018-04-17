package main

import (
	"github.com/zm-dev/yzm-client/handler"
	"log"
	"net/http"
)

func main() {

	r := handler.CreateHTTPAPIHandler()
	log.Fatal(http.ListenAndServe(":8000", r))
}
