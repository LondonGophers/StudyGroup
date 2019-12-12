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

	params := map[string]float64{
		"cycles":  5,     // number of complete x oscillator revolutions
		"res":     0.001, // angular resolution
		"size":    100,   // image canvas covers [-size..+size]
		"nframes": 64,    // number of animation frames
		"delay":   8,     // delay between frames in 10ms units
	}

	for key, value := range r.URL.Query() {
		// If the key does not already exist in our params map, skip this iteration of the loop and don't process it.
		if _, exists := params[key]; !exists {
			continue
		}

		// Only consider the last value, if multiple are passed in through the URL query parameter.
		fParam, errParse := strconv.ParseFloat(value[len(value)-1], 64)
		if errParse != nil {
			log.Printf("Error parsing float64 from parameter '%s': %v", value, errParse)
			continue // If it doesn't parse correctly, bin it.
		}

		params[key] = fParam
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: int(params["nframes"])}
	phase := 0.0 // phase difference

	for i := 0; i < int(params["nframes"]); i++ {
		rect := image.Rect(0, 0, int(2*params["size"]+1), int(2*params["size"]+1))
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < params["cycles"]*2*math.Pi; t += params["res"] {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				int(params["size"]+x*params["size"]+0.5),
				int(params["size"]+y*params["size"]+0.5),
				uint8((i/8)%(len(palette)-1))+1,
			)
		}

		phase += 0.1

		anim.Delay = append(anim.Delay, int(params["delay"]))
		anim.Image = append(anim.Image, img)
	}

	_ = gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
