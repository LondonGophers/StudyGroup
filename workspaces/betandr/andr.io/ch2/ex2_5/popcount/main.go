// The expression `x&(x-1)` clears the rightmost non-zero bit of x. Write a version
// of `PopCount` that counts bits by using this fact, and assess its performance.
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

// ByClearing returns the population count (number of set bits) of x by clearing.
func ByClearing(x uint64) int {
	result := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		result++
	}
	return result
}
