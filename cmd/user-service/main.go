package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Simple Hello world

	hello := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello World!\n")
	}

	http.HandleFunc("/hello", hello)
	log.Println("Listing requests at http://localhost:8080/hello")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
