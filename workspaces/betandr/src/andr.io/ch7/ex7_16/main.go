// Write a web-based calculator program.
package main

import (
	"io"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		renderCalculator(w)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func renderCalculator(out io.Writer) {
	filePrefix, _ := filepath.Abs("./src/andr.io/ch7/ex7_16/templates/")
	calculator := template.Must(template.ParseFiles(filePrefix + "/index.html"))

	var data struct{}

	if err := calculator.Execute(out, data); err != nil {
		log.Fatal(err)
	}
}
