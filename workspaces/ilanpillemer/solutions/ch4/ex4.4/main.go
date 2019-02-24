package main

// rotate operates in a single pass
// in order to do that the values of the slice are stored
// elsewhere so that values that are overwritten are not lost
// in the shuffle.
func rotateRight(input []int, rot int) []int {
	result := make([]int, len(input), len(input))

	for i := range input {
		// using the remainder gives the exact position of a rotation
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

