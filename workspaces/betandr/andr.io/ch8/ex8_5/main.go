// Take an existing CPU-bound sequential program, such as the Mandelbrot program of
// Section 3.3 of the 3D surface computation of Section 3.2, and execute its main
// loop in parallel using channels for communication. How much faster does it run
// on a multiprocessor machine? What is the optimal number of goroutines to use?
package main

import (
	"os"

	"andr.io/andr.io/ch8/ex8_5/mandelbrot"
)

func main() {
	mandelbrot.GenerateConcurrent(os.Stdout)
}
