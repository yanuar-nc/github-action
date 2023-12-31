package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello world\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/headers", headers)

	log.Println("server is running")
	http.ListenAndServe(":8090", nil)
}
