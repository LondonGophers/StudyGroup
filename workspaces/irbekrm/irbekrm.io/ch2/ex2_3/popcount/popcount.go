// Package popcount has functionality to count the number of set bits in an integer
package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Count returns the population count (number of set bits) of x.
func Count(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// CountLoop returns the population count (number of set bits) of x. Calculates result using a for-loop
func CountLoop(x uint64) int {
	var result int
	var i uint
	for i = 0; i < 9; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}
