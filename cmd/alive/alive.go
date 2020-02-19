package main

import (
	"bytes"
	"fmt"

	"github.com/valyala/fasthttp"
)

var OK = []byte("OK")

func main() {
	status, body, err := fasthttp.Get(nil, "http://localhost:8081/alive")
	if err != nil {
		panic(err)
	}
	if status != 200 {
		panic(fmt.Sprintf("Unexpected status %d", status))
	}
	if !bytes.Equal(body, OK) {
		panic("Unexpected response")
	}
	println("OK")
}
