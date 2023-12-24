package main

import (
	"fmt"
	"net/http"
)


func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "texto retornado\n")
	fmt.Printf("Getting http request from %v\n", req.Header )
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, c := range headers {
			fmt.Fprintf(w, "name: %v, header: %v\n", name, c)
		}
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)
}	