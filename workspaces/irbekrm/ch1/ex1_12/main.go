package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	backgroundColorIndex = iota
	patternColorIndex
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		cycles  = 5
		nframes = 64
		delay   = 8
	)
	if err := r.ParseForm(); err != nil {
		fmt.Printf("Recieved error: %v, continuing with the default params", err)
	} else {
		// try to set cycles, nframes, delay from http request params
		cycles, nframes, delay = parseParams(r.Form, cycles, nframes, delay)
	}
	lissajous(w, cycles, nframes, delay)
}

// writes a gif using the given metrics to the io.Writer passes as arg
func lissajous(out io.Writer, cycles, nframes, delay int) {
	const (
		res  = 0.001 // angular resolution
		size = 100   // image canvas covers [-size..+size]
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		palette := makePalette()
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), patternColorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

// takes url.Values and default param values
// if finds param values in url.Values, return those, else return defaults
func parseParams(form url.Values, cycles, nframes, delay int) (int, int, int) {
	for k, v := range form {
		if c, set := checkForParam(k, v, "cycles"); set {
			cycles = c
			break
		}
		if nf, set := checkForParam(k, v, "nframes"); set {
			nframes = nf
			break
		}
		if d, set := checkForParam(k, v, "delay"); set {
			delay = d
			break
		}
	}
	return cycles, nframes, delay
}

func checkForParam(key string, value []string, param string) (int, bool) {
	if key == param {
		if newVal, err := strconv.Atoi(value[0]); err == nil {
			return newVal, true
		}
	}
	return 0, false
}
func makePalette() []color.Color {
	c := randomColor()
	palette := []color.Color{color.Black, c}
	return palette
}

func randomColor() color.RGBA {
	tempColors := []uint8{}
	for i := 0; i < 4; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randInt := uint8(r.Intn(255))
		tempColors = append(tempColors, randInt)
	}
	return color.RGBA{tempColors[0], tempColors[1], tempColors[2], tempColors[3]}
}
