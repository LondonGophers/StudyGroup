package popcount

import (
	"math"
	"testing"
)

var cases = []struct {
	i uint64
	o int
}{
	{i: 0, o: 0},
	{i: 1, o: 1},
	{i: 2, o: 1},
	{i: 3, o: 2},
	{i: math.MaxUint64, o: 64},
}

func TestPopCount(t *testing.T) {
	for _, tc := range cases {
		if got, want := popCount(tc.i), tc.o; got != want {
			t.Errorf("test: %d, got: %d, want: %d", tc.i, got, want)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount(math.MaxUint64)
	}
}

func TestPopCountByteShift(t *testing.T) {
	for _, tc := range cases {
		if got, want := popCountByteShift(tc.i), tc.o; got != want {
			t.Errorf("test: %d, got: %d, want: %d", tc.i, got, want)
		}
	}
}

func BenchmarkPopCountByteShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCountByteShift(math.MaxUint64)
	}
}

func TestPopCountBitShift(t *testing.T) {
	for _, tc := range cases {
		if got, want := popCountBitShift(tc.i), tc.o; got != want {
			t.Errorf("test: %d, got: %d, want: %d", tc.i, got, want)
		}
	}
}

func BenchmarkPopCountBitShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCountBitShift(math.MaxUint64)
	}
}
