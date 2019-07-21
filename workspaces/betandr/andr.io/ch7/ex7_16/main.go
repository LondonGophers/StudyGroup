// Write a web-based calculator program.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"

	"andr.io/ch7/ex7_15/eval"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		renderCalculator(w)
	}

	evaluateHandler := func(w http.ResponseWriter, r *http.Request) {
		e := r.URL.Query().Get("e")

		expr, err := eval.Parse(e)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
		} else {
			vars := make(map[eval.Var]bool)
			err = expr.Check(vars)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				env := make(map[eval.Var]float64)
				result := strconv.FormatFloat(expr.Eval(env), 'f', -1, 64)
				fmt.Fprintf(w, result)
			}
		}
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/eval", evaluateHandler)
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
