package main

// rotate operate in a single pass
// in order to do that the values of the slice are stored
// elsewhere so that values that are overwritten are not lost
// in the shuffle.
func rotateRight(input []int, rot int) []int {
	result := make([]int, len(input), len(input))

	for i := range input {
		// using the remainder give the exact position of a rotation
		// when moving right
		dest := (i + rot) % len(input)
		result[dest] = input[i]
	}

	return result
}

//rotates left
func rotate(input []int, rot int) []int {
	return rotateRight(input, len(input)-rot) // moving left equivalent in terms of moving right
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
