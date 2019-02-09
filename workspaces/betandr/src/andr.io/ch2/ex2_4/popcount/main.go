// Write a version of `PopCount` that counts bits by shifting its argument through
// 64 bit positions, testing the rightmost bit each time. Compare its performance
// to the table-lookup version.
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// ByLookup returns the population count (number of set bits) of x.
func ByLookup(x uint64) int {
	result := 0
	for i := uint(0); i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}

	return result
}

// ByShifting returns the population count (number of set bits) of x by shifting.
func ByShifting(x uint64) int {
	result := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			result++
		}
	}
	return result
}
