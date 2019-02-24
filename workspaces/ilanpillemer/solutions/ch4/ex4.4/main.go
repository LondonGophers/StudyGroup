package main

//rotates left
func rotate(input []int, rot int) []int {
	return rotateRight(input, len(input)-rot) // moving left equivalent in terms of moving right
}

func rotate2(input []int, rot int) {
	rotateInPlaceRight(input, len(input)-rot) // moving left equivalent
}

// the first solution does in one pass but returns a new slice.
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

// rotateInPlaceRight does it in one pass replacing the values in the slice.
// This also uses a temporary storage area for the values that were replaced.
func rotateInPlaceRight(input []int, rot int) {
	var remembered = make(map[int]int)
	for i := range input { // one pass
		dest := (i + rot) % len(input)
		remembered[dest] = input[dest]
		value, ok := remembered[i]
		if ok {
			input[dest] = value
			continue
		}
		input[dest] = input[i]
	}
}
