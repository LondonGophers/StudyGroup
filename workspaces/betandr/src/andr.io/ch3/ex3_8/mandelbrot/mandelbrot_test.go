// Rendering fractals at high zoom levels demands great arithmetic precision.
// Implement the same fractal using four different representations of numbers:
// `complex64`, `complex128`, `big.Float`, and `big.Rat`. (The latter two types
// are found in the `math/big` package. `Float` uses arbitrary but
// bounded-precision floating-point; `Rat` uses unbounded-precision rational
// numbers.) How do they compare in performance and memory usage? At what zoom
// levels do rendering artefacts become visible?
package mandelbrot_test

import (
	"io/ioutil"
	"testing"

	"andr.io/ch3/ex3_8/mandelbrot"
)

// -- Benchmarks --

func BenchmarkGenerate64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrot.Generate64(ioutil.Discard)
	}
}

func BenchmarkGenerate128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrot.Generate128(ioutil.Discard)
	}
}

// $ go test -cpu=2 -bench=. andr.io/ch3/ex3_8/mandelbrot
// goos: darwin
// goarch: amd64
// pkg: andr.io/ch3/ex3_8/mandelbrot
// BenchmarkGenerate64-2    	       3	 365219698 ns/op
// BenchmarkGenerate128-2   	       5	 290191841 ns/op
// PASS
// ok  	andr.io/ch3/ex3_8/mandelbrot	4.815s
