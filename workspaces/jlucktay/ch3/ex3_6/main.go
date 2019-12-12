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
	height, width           = 500, 500
	xMin, yMin, xMax, yMax  = -2, -2, +2, +2
	superheight, superwidth = height * 2, width * 2
)

var results = map[uint8]uint{}

func main() {
	super := [superwidth][superheight]color.Color{}

	for py := 0; py < superheight; py++ {
		y := float64(py)/superheight*(yMax-yMin) + yMin

		for px := 0; px < superwidth; px++ {
			x := float64(px)/superwidth*(xMax-xMin) + xMin
			z := complex(x, y)
			super[px][py] = mandelbrot(z)
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r1, g1, b1, a1 := super[2*x][2*y].RGBA()
			r2, g2, b2, a2 := super[2*x+1][2*y].RGBA()
			r3, g3, b3, a3 := super[2*x][2*y+1].RGBA()
			r4, g4, b4, a4 := super[2*x+1][2*y+1].RGBA()

			avg := color.RGBA{
				R: uint8((r1 + r2 + r3 + r4) / (4 * 256)),
				G: uint8((g1 + g2 + g3 + g4) / (4 * 256)),
				B: uint8((b1 + b2 + b3 + b4) / (4 * 256)),
				A: uint8((a1 + a2 + a3 + a4) / (4 * 256)),
			}

			img.Set(x, y, avg)
		}
	}

	if errEncode := png.Encode(os.Stdout, img); errEncode != nil {
		log.Fatalf("error encoding PNG: %v", errEncode)
	}

	writeNumbersToTextFile()
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

func writeNumbersToTextFile() {
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
