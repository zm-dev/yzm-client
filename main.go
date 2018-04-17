package main

import (
	"fmt"
)

type a struct {
	a int
}

func (a a) test() {
	fmt.Println("test")
}

func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	//r := handler.CreateHTTPAPIHandler()
	//log.Fatal(http.ListenAndServe(":8000", r))
}
