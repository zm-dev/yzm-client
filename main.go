package main

import (
	"github.com/3tnet/yzm-client/handler"
	"net/http"
	"log"

)

func main() {

	r := handler.CreateHTTPAPIHandler()
	log.Fatal(http.ListenAndServe(":8000", r))
}
