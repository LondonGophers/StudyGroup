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

// This fails the tests. I dont understand the code clearly so not sure what the bug is.
// Rotate rotates a slice of integers `s` at the position `pos` with a single pass
// through the array, swapping from the last position to the first position
// until all items are in order, running in O(n) time and using O(n) space.
//func betandr(pos int, s []int) {
//	k := len(s) - pos
//	stop := 0
//	if len(s)%2 == 0 {
//		stop = 1
//	}
//
//	for i, j := len(s)-1, pos; i > stop; i, j = i-1, j-1 {
//		if j <= 0 {
//			j = pos
//			k = k - pos
//		}
//		l := i - k
//		if i <= 2 {
//			l = 0
//		}
//		s[l], s[i] = s[i], s[l]
//	}
//}
//$ go test
//--- FAIL: TestBetandr (0.00s)
//    main_test.go:74: want []int{2, 3, 4, 5, 6, 0, 1}, got []int{4, 2, 3, 5, 6, 0, 1}
//    main_test.go:74: want []int{3, 4, 5, 6, 7, 0, 1, 2}, got []int{4, 3, 5, 6, 7, 0, 1, 2}
//FAIL
//
//exit status 1
//FAIL	_/Users/ilanpillemer/git/study-group/workspaces/ilanpillemer/solutions/ch4/ex4.4	0.009s
//go test: exit status 1
//$
