package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "texto retornado\n")
	fmt.Printf("Getting http request from %v\n", req.Header )
}

func main() {
	http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)
}	