// Package adjacent contains the Eliminate() func.
package adjacent

// Eliminate will remove adjacent duplicates in-place in a `[]string` slice.
// The slice argument must be given as a pointer so that changes can be effected upon the slice itself.
func Eliminate(s *[]string) {
	// Start from the second element, as the comparison loop looks at previous vs current index
	index := 1

	for index < len(*s) {
		if (*s)[index-1] == (*s)[index] {
			// Delete duplicate at 'index'
			copy((*s)[index:], (*s)[index+1:])

			// Shrink slice length by one
			*s = (*s)[:len(*s)-1]
		} else {
			index++
		}
	}
}
