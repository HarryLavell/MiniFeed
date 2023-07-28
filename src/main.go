package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("index.html"))
		template.Execute(w, nil)
	}

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
