// Take an existing CPU-bound sequential program, such as the Mandelbrot program of
// Section 3.3 of the 3D surface computation of Section 3.2, and execute its main
// loop in parallel using channels for communication. How much faster does it run
// on a multiprocessor machine? What is the optimal number of goroutines to use?
//
// Based on andr.io/ch8/ex3_8
package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"math/cmplx"
	"sync"
)

// GenerateConcurrent generates a new mandelbrot set png to stdout using goroutines and channels
func GenerateConcurrent(w io.Writer, numWorkers int) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	wg := sync.WaitGroup{}

	// channel of rows to process
	rows := make(chan int, height)
	for row := 0; row < height; row++ {
		rows <- row
	}
	close(rows)

	for worker := 0; worker < numWorkers; worker++ {
		wg.Add(1)
		go func() {
			for py := range rows {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					img.Set(px, py, mandelbrot(z))
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	png.Encode(w, img)
}

// Generate generates a new mandelbrot set png to stdout
func Generate(w io.Writer) {
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
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50:
				return color.RGBA{0, 255, 0, 255}
			default:
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}
