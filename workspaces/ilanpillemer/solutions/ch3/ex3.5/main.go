package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 80
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if c := cmplx.Abs(v); c > 2 {
			switch {
			case c > 4.3:
				return color.RGBA{255 - contrast*n, 0, 0, uint8(255)}
			case c > 2.3:
				return color.RGBA{0, 0, 255 - contrast*n, uint8(255)}
			default:
				return color.RGBA{0, 255 - contrast*n, 0, uint8(255)}
			}

		}
	}
	return color.Black
}
