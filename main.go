package main

import (
	"yzm-client/pkg/distinguish_service"
	"io/ioutil"
	"fmt"
)

func main() {
	s, _ := distinguish_service.GetGRPCService("127.0.0.1:8080")
	b, _ := ioutil.ReadFile("/Users/taoyu/Desktop/yzm/data-1/train/0001.jpg")
	fmt.Println(s.DistinguishData1(b))
}
