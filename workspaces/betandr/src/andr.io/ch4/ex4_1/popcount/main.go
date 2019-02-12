// Write a function that counts the number of bits that are different in two
// SHA256 hashes. (see Popcount from Section 2.6.2)
package popcount

// Count returns the population count (number of set bits) of x by clearing.
func Count(x uint64) int {
	result := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		result++
	}
	return result
}

//!-
