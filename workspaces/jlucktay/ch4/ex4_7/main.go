// Package reverse contains a func to reverse a slice.
package reverse

// Reverse reverses a slice of ints in place.
func Reverse(a *[]int) {
	for i, j := 0, len(*a)-1; i < j; i, j = i+1, j-1 {
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	}
}

// Modify `reverse` to reverse the characters of a `[]byte` slice that represents a UTF-8-encoded string, in place.
// Can you do it without allocating new memory?
