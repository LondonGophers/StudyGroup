// Rendering fractals at high zoom levels demands great arithmetic precision.
// Implement the same fractal using four different representations of numbers:
// `complex64`, `complex128`, `big.Float`, and `big.Rat`. (The latter two types
// are found in the `math/big` package. `Float` uses arbitrary but
// bounded-precision floating-point; `Rat` uses unbounded-precision rational
// numbers.) How do they compare in performance and memory usage? At what zoom
// levels do rendering artefacts become visible?
package main

import (
	"os"

	"andr.io/ch3/ex3_8/mandelbrot"
)

func main() {
	mandelbrot.Generate128(os.Stdout)
}
