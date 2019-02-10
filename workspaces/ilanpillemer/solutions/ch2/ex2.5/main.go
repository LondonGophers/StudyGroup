// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) (count int) {
	var i uint
	for i = 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return
}

func PopCountShift(x uint64) (count int) {
	//shift the bit through all 64 positions and use a bitwise and
	//to check if that bit has a value, if it does increment pop count
	var i uint64
	for i = 0; i < 64; i++ {
		if (x & (1 << i)) != 0 {
			count++
		}
	}
	return
}

func PopCountClear(x uint64) (count int) {
	// since x & x ( x - 1) clears the rightmost bit each time
	// the number of times it takes to clear the bit is the pop count
	// so in the case used in benchmarking this should loop 32 times,
	// ... instead of the 64 times in the shifting
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return
}