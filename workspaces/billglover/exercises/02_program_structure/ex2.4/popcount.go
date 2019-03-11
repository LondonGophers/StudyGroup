package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func popCountByteShift(x uint64) int {
	c := 0
	for b := uint8(0); b < 64; b += 8 {
		c += int(pc[byte(x>>b)])
	}
	return c
}

func popCountBitShift(x uint64) int {
	c := 0
	for b := uint8(0); b < 64; b++ {
		c += int(x & 1)
		x = x >> 1
	}
	return c
}
