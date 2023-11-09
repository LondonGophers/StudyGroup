// Take an existing CPU-bound sequential program, such as the Mandelbrot program of
// Section 3.3 of the 3D surface computation of Section 3.2, and execute its main
// loop in parallel using channels for communication. How much faster does it run
// on a multiprocessor machine? What is the optimal number of goroutines to use?
//
// Based on andr.io/ch8/ex3_8
package mandelbrot_test

import (
	"io/ioutil"
	"runtime"
	"testing"

	"andr.io/ch8/ex8_5/mandelbrot"
)

// -- Benchmarks --

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrot.Generate(ioutil.Discard)
	}
}

func BenchmarkGenerateConcurrentWithOneWorker(b *testing.B) {
	if runtime.GOMAXPROCS(-1) >= 1 {
		for i := 0; i < b.N; i++ {
			mandelbrot.GenerateConcurrent(ioutil.Discard, 1)
		}
	} else {
		b.Skip()
	}
}

func BenchmarkGenerateConcurrentWithTwoWorkers(b *testing.B) {
	if runtime.GOMAXPROCS(-1) >= 2 {
		for i := 0; i < b.N; i++ {
			mandelbrot.GenerateConcurrent(ioutil.Discard, 2)
		}
	} else {
		b.Skip()
	}
}

func BenchmarkGenerateConcurrentWithThreeWorkers(b *testing.B) {
	if runtime.GOMAXPROCS(-1) >= 3 {
		for i := 0; i < b.N; i++ {
			mandelbrot.GenerateConcurrent(ioutil.Discard, 3)
		}
	} else {
		b.Skip()
	}
}

func BenchmarkGenerateConcurrentWithFourWorkers(b *testing.B) {
	if runtime.GOMAXPROCS(-1) >= 4 {
		for i := 0; i < b.N; i++ {
			mandelbrot.GenerateConcurrent(ioutil.Discard, 4)
		}
	} else {
		b.Skip()
	}
}

func BenchmarkGenerateConcurrentWithFiveWorkers(b *testing.B) {
	if runtime.GOMAXPROCS(-1) >= 5 {
		for i := 0; i < b.N; i++ {
			mandelbrot.GenerateConcurrent(ioutil.Discard, 5)
		}
	} else {
		b.Skip()
	}
}

func BenchmarkGenerateConcurrentWithSixWorkers(b *testing.B) {
	if runtime.GOMAXPROCS(-1) >= 6 {
		for i := 0; i < b.N; i++ {
			mandelbrot.GenerateConcurrent(ioutil.Discard, 6)
		}
	} else {
		b.Skip()
	}
}

func BenchmarkGenerateConcurrentWithSevenWorkers(b *testing.B) {
	if runtime.GOMAXPROCS(-1) >= 7 {
		for i := 0; i < b.N; i++ {
			mandelbrot.GenerateConcurrent(ioutil.Discard, 7)
		}
	} else {
		b.Skip()
	}
}

func BenchmarkGenerateConcurrentWithEightWorkers(b *testing.B) {
	if runtime.GOMAXPROCS(-1) >= 8 {
		for i := 0; i < b.N; i++ {
			mandelbrot.GenerateConcurrent(ioutil.Discard, 8)
		}
	} else {
		b.Skip()
	}
}

func BenchmarkGenerateConcurrentWithNineWorkers(b *testing.B) {
	if runtime.GOMAXPROCS(-1) >= 9 {
		for i := 0; i < b.N; i++ {
			mandelbrot.GenerateConcurrent(ioutil.Discard, 9)
		}
	} else {
		b.Skip()
	}
}

func BenchmarkGenerateConcurrentWithTenWorkers(b *testing.B) {
	if runtime.GOMAXPROCS(-1) >= 10 {
		for i := 0; i < b.N; i++ {
			mandelbrot.GenerateConcurrent(ioutil.Discard, 10)
		}
	} else {
		b.Skip()
	}
}

// $ go test -bench=. ./andr.io/ch3/ex3_8/mandelbrot
//
// goos: darwin
// goarch: amd64
// pkg: andr.io/ch8/ex8_5/mandelbrot
// BenchmarkGenerate-8             	       4	 257366328 ns/op
// BenchmarkGenerateConcurrent-8   	      12	  92844830 ns/op
// PASS
// ok  	andr.io/ch8/ex8_5/mandelbrot	3.292s
