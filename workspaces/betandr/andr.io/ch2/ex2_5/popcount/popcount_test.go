// The expression `x&(x-1)` clears the rightmost non-zero bit of x. Write a version
// of `PopCount` that counts bits by using this fact, and assess its performance.
package popcount_test

import (
	"testing"

	"andr.io/ch2/ex2_5/popcount"
)

// -- Benchmarks --

func BenchmarkPopCountByLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.ByLookup(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.ByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.ByShifting(0x1234567890ABCDEF)
	}
}

// $ go test -cpu=2 -bench=. andr.io/ch2/ex2_5/popcount
// goos: darwin
// goarch: amd64
// pkg: andr.io/ch2/ex2_5/popcount
// BenchmarkPopCountByLookup-2     	100000000	        16.7 ns/op
// BenchmarkPopCountByClearing-2   	100000000	        15.4 ns/op
// BenchmarkPopCountByShifting-2   	20000000	        82.6 ns/op
// PASS
// ok  	andr.io/ch2/ex2_5/popcount	4.987s
