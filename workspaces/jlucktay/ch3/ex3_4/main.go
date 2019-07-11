package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		poly(w, r)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
