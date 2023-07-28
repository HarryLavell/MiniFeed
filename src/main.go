package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!")
	}

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
