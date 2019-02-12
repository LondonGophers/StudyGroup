// Write a version of `rotate` that operates in a single pass

package main

import "fmt"

// Rotate rotates a slice of integers `s` at the position `pos` with a single pass
// through the array, swapping from the last position to the first position
// until all items are in order, running in O(n) time and using O(n) space.
func Rotate(pos int, s []int) {
	k := len(s) - pos
	stop := 0
	if len(s)%2 == 0 {
		stop = 1
	}

	for i, j := len(s)-1, pos; i > stop; i, j = i-1, j-1 {
		if j <= 0 {
			j = pos
			k = k - pos
		}
		l := i - k
		if i <= 2 {
			l = 0
		}
		s[l], s[i] = s[i], s[l]
	}
}

func main() {
	even := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Rotate(4, even[:])
	fmt.Println(even)

	odd := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Rotate(4, odd[:])
	fmt.Println(odd)
}
