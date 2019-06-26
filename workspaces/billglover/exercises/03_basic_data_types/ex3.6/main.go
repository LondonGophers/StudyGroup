package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"math/cmplx"
	"os"
	//"fmt"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width*2, height*2))
	for py := 0; py < height*2; py++ {
		y := float64(py)/(height*2)*(ymax-ymin) + ymin
		for px := 0; px < width*2; px++ {
			x := float64(px)/(width*2)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	img_sub := image.NewRGBA(image.Rect(0, 0, width, height))

	for i := 0; i<height;i++ {
		for j := 0; j<width; j++ {
			c1 := img.RGBAAt(i*2, j*2)
			c2 := img.RGBAAt((i*2)+1, j*2)
			c3 := img.RGBAAt(i*2, (j*2)+1)
			c4 := img.RGBAAt((i*2)+1, (j*2)+1)

			r1, g1, b1, a1 := c1.RGBA()
			r2, g2, b2, a2 := c2.RGBA()
			r3, g3, b3, a3 := c3.RGBA()
			r4, g4, b4, a4 := c4.RGBA()

			c := color.RGBA{
				uint8((r1 + r2 + r3 + r4) / 1028),
				uint8((g1 + g2 + g3 + g4) / 1028),
				uint8((b1 + b2 + b3 + b4) / 1028),
				uint8((a1 + a2 + a3 + a4) / 1028),
			}

			img_sub.Set(i, j, c)
		}
	}

	png.Encode(os.Stdout, img_sub) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	
	// hack to avoid the ugly circular effect at Abs(z) = 2
	if cmplx.Abs(z) >= 2 {
		return palette.WebSafe[1]
	}

	var v complex128
	for n:=uint8(0); n<iterations; n++{
		v = v*v + z
		if cmplx.Abs(v)>2 {
			return palette.WebSafe[n]
		}	
	}

	return color.Black
}
