package popcount_test

import (
	"math/bits"
	"testing"

	"github.com/matryer/is"

	popcount "github.com/LondonGophers/StudyGroup/workspaces/jlucktay/ch2/ex2_5"
)

// -- Alternative implementations --
func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x -= ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x += (x >> 8)
	x += (x >> 16)
	x += (x >> 32)

	return int(x & 0x7f)
}

func PopCountByClearing(x uint64) int {
	n := 0

	for x != 0 {
		x &= x - 1 // clear rightmost non-zero bit
		n++
	}

	return n
}

func PopCountByShifting(x uint64) int {
	n := 0

	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}

	return n
}

// Test implementations for accuracy
func TestImplementations(t *testing.T) {
	testCases := map[string]struct {
		impl func(uint64) int
	}{
		"popcount.PopCount": {
			impl: popcount.PopCount,
		},
		"popcount.Loop": {
			impl: popcount.Loop,
		},
		"popcount.Shift": {
			impl: popcount.Shift,
		},
		"popcount.Rightmost": {
			impl: popcount.Rightmost,
		},
		"bits.OnesCount64": {
			impl: bits.OnesCount64,
		},
		"BitCount": {
			impl: BitCount,
		},
		"PopCountByClearing": {
			impl: PopCountByClearing,
		},
		"PopCountByShifting": {
			impl: PopCountByShifting,
		},
	}
	for desc, tC := range testCases {
		tC := tC

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)
			is.Equal(tC.impl(1), 1)  // 00000001
			is.Equal(tC.impl(2), 1)  // 00000010
			is.Equal(tC.impl(3), 2)  // 00000011
			is.Equal(tC.impl(4), 1)  // 00000100
			is.Equal(tC.impl(5), 2)  // 00000101
			is.Equal(tC.impl(6), 2)  // 00000110
			is.Equal(tC.impl(7), 3)  // 00000111
			is.Equal(tC.impl(8), 1)  // 00001000
			is.Equal(tC.impl(9), 2)  // 00001001
			is.Equal(tC.impl(10), 2) // 00001010
			is.Equal(tC.impl(11), 3) // 00001011
			is.Equal(tC.impl(12), 2) // 00001100
			is.Equal(tC.impl(13), 3) // 00001101
			is.Equal(tC.impl(14), 3) // 00001110
			is.Equal(tC.impl(15), 4) // 00001111
			is.Equal(tC.impl(16), 1) // 00010000
		})
	}
}

// -- Benchmarks --
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.Loop(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.Shift(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountClearRightmost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.Rightmost(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountStdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bits.OnesCount64(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}
