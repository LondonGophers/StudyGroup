package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

const (
	height, width          = 500, 500
	xMin, yMin, xMax, yMax = -2, -2, +2, +2
)

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
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z

		if cmplx.Abs(v) > 2 {
			return palette.Plan9[255-contrast*n]

			// rand.Seed(time.Now().UnixNano())
			// switch rand.Intn(3) {
			// case 0:
			// 	return color.RGBA{R: 255 - contrast*n, G: 0, B: 0, A: 255}
			// case 1:
			// 	return color.RGBA{R: 0, G: 255 - contrast*n, B: 0, A: 255}
			// case 2:
			// 	return color.RGBA{R: 0, G: 0, B: 255 - contrast*n, A: 255}
			// }
		}
	}

	return color.Black
}
