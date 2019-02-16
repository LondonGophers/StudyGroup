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
		higher, offset         = 16, 16
		width, height          = 1024*higher + offset, 1024*higher + offset
	)

	var pre [width + offset][height + offset]color.Color

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			pre[px][py] = mandelbrot(z)
		}
	}

	//super sample from the higher resolution
	superImg := image.NewRGBA(image.Rect(0, 0, 1024, 1024))

	for sy := 0; sy < 1024; sy++ {
		for sx := 0; sx < 1024; sx++ {

			colorUp := pre[sx*higher+offset][sy*higher-1+offset].(color.RGBA)
			colorDown := pre[sx*higher][sy*higher+1].(color.RGBA)
			colorLeft := pre[sx*higher-1+offset][sy*higher+offset].(color.RGBA)
			colorRight := pre[sx*higher+1][sy*higher].(color.RGBA)

			sampleRed := (uint64(colorUp.R) + uint64(colorDown.R) + uint64(colorLeft.R) + uint64(colorRight.R)) / 4
			sampleGreen := (uint64(colorUp.G) + uint64(colorDown.G) + uint64(colorLeft.G) + uint64(colorRight.G)) / 4
			sampleBlue := (uint64(colorUp.B) + uint64(colorDown.B) + uint64(colorLeft.B) + uint64(colorRight.B)) / 4
			sampleAlpha := (uint64(colorUp.A) + uint64(colorDown.A) + uint64(colorLeft.A) + uint64(colorRight.A)) / 4

			sampleColor := color.RGBA{uint8(sampleRed), uint8(sampleGreen), uint8(sampleBlue), uint8(sampleAlpha)}

			superImg.Set(sx, sy, sampleColor)
		}
	}
	png.Encode(os.Stdout, superImg)
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

	return color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)}
}
