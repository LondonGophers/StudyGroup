// Another simple fractal uses Newton's method to find complex solutions to a
// function such as `z^4-1 = 0`. Shade each starting point by the number of
// iterations required to get close to one of the four roots. Color each point by
// the root it approaches.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

var colors = []color.RGBA{
	{255, 0, 0, 255},
	{0, 255, 0, 255},
	{0, 0, 255, 255},
	{255, 255, 0, 255},
}

var colorMap = map[complex128]color.RGBA{}

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
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7

	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			root := complex(round(real(z), 4), round(imag(z), 4))
			c, ok := colorMap[root]
			if !ok {
				c = colors[0]
				colors = colors[1:]
				colorMap[root] = c
			}
			// Convert to YCbCr to make producing different shades easier.
			y, cb, cr := color.RGBToYCbCr(c.R, c.G, c.B)
			scale := math.Log(float64(i)) / math.Log(iterations)
			y -= uint8(float64(y) * scale)
			return color.YCbCr{y, cb, cr}
		}
	}
	return color.Black
}

func round(f float64, digits int) float64 {
	if math.Abs(f) < 0.5 {
		return 0
	}
	pow := math.Pow10(digits)
	return math.Trunc(f*pow+math.Copysign(0.5, f)) / pow
}
