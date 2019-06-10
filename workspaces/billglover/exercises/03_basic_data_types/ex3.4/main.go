package main

import (
	"net/http"
	"path"

	"github.com/go-london-user-group/study-group/workspaces/billglover/exercises/03_basic_data_types/ex3.4/surface"
)

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "image/svg+xml")

	var shape surface.Surface

	switch path.Base(r.URL.Path) {
	case "eggbox":
		shape = surface.EggBox
	case "saddle":
		shape = surface.Saddle
	case "monkeysaddle":
		shape = surface.MonkeySaddle
	case "paper":
		shape = surface.Paper
	default:
		shape = surface.Original
	}

	surface.Render(w, shape)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8085", nil)
}
