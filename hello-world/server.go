package main

import (
	"io"
	"log"
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func (res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello, you've requested: %s\n", req.URL.Path)
	})

	helloHandler := func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "Hello, world!\n")
	}

	http.HandleFunc("/hello/", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}