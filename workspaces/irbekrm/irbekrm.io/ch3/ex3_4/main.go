package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"irbekrm.io/ch3/ex3_4/image"
)

const (
	defaultColor  = "white"
	defaultHeight = "320"
	defaultWidth  = "600"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	u := r.URL
	q := u.Query()
	color := q.Get("color")
	if color == "" {
		color = defaultColor
	}
	h := q.Get("height")
	if h == "" {
		h = defaultHeight
	}
	f_h, err := strconv.ParseFloat(h, 64)
	if err != nil {
		log.Fatalf("Cannot parse parameter 'height', error: %v\n", err)
	}
	width := q.Get("width")
	if width == "" {
		width = defaultWidth
	}
	f_width, err := strconv.ParseFloat(width, 64)
	if err != nil {
		log.Fatalf("Cannot parse parameter 'width', error: $v\n", err)
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	c := image.New(f_h, f_width, color)
	svg := c.BuildImage()
	fmt.Fprint(w, svg)
}
