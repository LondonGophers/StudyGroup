package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, r *http.Request) {
	var palette = []color.Color{
		color.Black,
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff},
	}

	defaults := map[string]float64{
		"cycles":  5,     // number of complete x oscillator revolutions
		"res":     0.001, // angular resolution
		"size":    100,   // image canvas covers [-size..+size]
		"nframes": 64,    // number of animation frames
		"delay":   8,     // delay between frames in 10ms units
	}

	for key, value := range defaults {
		param := chi.URLParam(r, key)
		fParam, errParse := strconv.ParseFloat(param, 64)
		if errParse != nil {
			log.Fatalf("Error parsing float from parameter '%s': %v", param, errParse)
		}
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: int(defaults["nframes"])}
	phase := 0.0 // phase difference
	for i := 0; i < int(defaults["nframes"]); i++ {
		rect := image.Rect(0, 0, 2*defaults["size"]+1, 2*defaults["size"]+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(defaults["size"]+int(x*defaults["size"]+0.5), defaults["size"]+int(y*defaults["size"]+0.5),
				uint8((i/8)%(len(palette)-1))+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	_ = gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
