package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

var v = flag.Bool("v", false, "print verbose logs for debugging")
var c = flag.String("c", "websafe", "show in colours")

type Grey = color.Gray

var iters = make(map[int]int)
var roots = make(map[complex128]int)

var key = -1
var colours = [4]color.RGBA{
	color.RGBA{255, 0, 0, 255},
	color.RGBA{0, 255, 0, 255},
	color.RGBA{0, 0, 255, 255},
	color.RGBA{0, 255, 255, 255},
}

func main() {
	const (
		yRangeMin, yRangeMax = -2, +2
		xRangeMin, xRangeMax = -2, +2
		width, height        = 1024, 1024
	)
	flag.Parse()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		// project py onto domain using y
		y := float64(py) / height       // ie go from 0 to 1
		y = y * (yRangeMax - yRangeMin) // expand range from 0 to 4
		y = y + yRangeMin               // translate range to -2 to +2
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xRangeMax-xRangeMin) + xRangeMin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}

	verbosef("# of roots: %d\n", len(roots))
	if *v {
		for k := range roots {
			verbosef("root is: %v\n", k)
		}
	}
	png.Encode(os.Stdout, img)
}

// next guess is z - ( f(x) / 'f(x) )
func newton(z complex128) color.Color {

	const (
		iter     = 50
		contrast = 20 // 50 * 5 is 250, which is good enough for me
		dz       = 0.000001
		rnd      = 6
	)

	// next guess is z - ( f(x) / 'f(x) )
	for i := 0; i < iter; i++ {
		nextGuess := (cmplx.Pow(z, 4) - 1) / (4 * cmplx.Pow(z, 3))
		z = z - nextGuess

		if cmplx.Abs(cmplx.Pow(z, 4)-1) < 0+dz {
			if _, ok := roots[round(z, rnd)]; !ok {
				key++
				roots[round(z, rnd)] = key
			}

			iters[i] = iters[i] + 1
			switch *c {
			case "grey":
				return Grey{255 - uint8(i)*contrast}
			case "websafe":
				return palette.WebSafe[(255-uint(i)*contrast)%216]
			case "plan9":
				return palette.Plan9[(255-uint(i)*contrast)%255]
			case "primary":
				return shaded(colours[roots[round(z, rnd)]], i, contrast)
			default:
				return palette.WebSafe[(255-uint(i)*contrast)%216]
			}
		}
	}
	return color.Black

}

func shaded(c color.Color, i int, contrast int) color.Color {
	r, g, b, _ := c.RGBA()
	if r != 0 {
		r = uint32(255-(i*contrast)) % 255
	}
	if g != 0 {
		g = uint32(255-(i*contrast)) % 255
	}
	if b != 0 {
		b = uint32(255-(i*contrast)) % 255
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

func round(z complex128, r int) complex128 {
	rl := real(z)
	mg := imag(z)

	rl = math.Round(rl * math.Pow10(r))
	mg = math.Round(mg * math.Pow10(r))

	rl = rl / math.Pow10(r)
	mg = mg / math.Pow10(r)

	return complex(rl, mg)
}

func verbosef(format string, args ...interface{}) {
	if *v {
		fmt.Fprintf(os.Stderr, format, args...)
	}
}
