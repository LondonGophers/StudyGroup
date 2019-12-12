package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"io/ioutil"
	"log"
	"math/cmplx"
	"os"
	"sort"

	"github.com/pkg/errors"
)

const (
	height, width          = 500, 500
	xMin, yMin, xMax, yMax = -2, -2, +2, +2
)

var results = map[uint8]uint{}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(yMax-yMin) + yMin

		for px := 0; px < height; px++ {
			x := float64(px)/width*(xMax-xMin) + xMin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	if errEncode := png.Encode(os.Stdout, img); errEncode != nil {
		log.Fatalf("error encoding PNG: %v", errEncode)
	}

	var writeMe string

	keys := []int{}

	for key := range results {
		keys = append(keys, int(key))
	}

	sort.Ints(keys)

	for _, key := range keys {
		writeMe += fmt.Sprintf("%7d=%7d\n", key, results[uint8(key)])
	}

	if errWrite := ioutil.WriteFile("text.log", []byte(writeMe), 0644); errWrite != nil {
		log.Fatal(errors.Wrap(errWrite, "goodbye!"))
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations, contrast = 200, 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z

		if cmplx.Abs(v) > 2 {
			result := 255 - contrast*n
			results[result]++

			return palette.Plan9[result]
		}
	}

	return color.Black
}
