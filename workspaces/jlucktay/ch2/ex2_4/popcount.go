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

// PopCountLoop returns the population count (number of set bits) of x. The internal implementation uses a loop.
func PopCountLoop(x uint64) (sum int) {
	for index := uint(0); index < 8; index++ {
		sum += int(pc[byte(x>>(index*8))])
	}

	return
}

// PopCountShift returns the population count (number of non-zero bits) of x, by shifting its argument through 64 bit
// positions, testing the rightmost bit each time.
func PopCountShift(x uint64) (sum int) {
	for index := 0; index < 64; index++ {
		sum += int(byte(x) & 1)
		x = x >> 1
	}

	return
}
