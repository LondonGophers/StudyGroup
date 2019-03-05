package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-london-user-group/study-group/workspaces/billglover/exercises/01_tutorial/ex1.12/lissajous"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "server: %v\n", err)
		os.Exit(1)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles, err := strconv.Atoi(r.URL.Query().Get("cycles"))
	if err != nil || cycles <= 0 {

		// if a cycles value was set then return an error
		if r.URL.Query().Get("cycles") != "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "unable to parse 'cycles', not a positive integer")
			return
		}

		// otherwise use a default value
		cycles = 5
	}

	freq, err := strconv.ParseFloat(r.URL.Query().Get("freq"), 64)
	if err != nil {

		// if a freq value was set then return an error
		if r.URL.Query().Get("freq") != "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "unable to parse 'freq', not a number")
			return
		}

		// otherwise use a default value
		freq = 3.0
	}

	lissajous.Lissajous(w, float64(cycles), freq)
}
